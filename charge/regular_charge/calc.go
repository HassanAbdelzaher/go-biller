package regular_charge

import (. "MaisrForAdvancedSystems/go-biller/proto"
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)
type RegularChargeAmount struct {
	Amount float64
	TaxAmount *float64
}
func CalcCharge(fee *RegularCharge,cust *Customer,bilngDate time.Time,lastCharge *time.Time) ([]*FinantialTransaction,error){
	log.Println("calc reg charge:"+*fee.Code)
	var amount float64=0
	var taxAmount float64=0
	resp:=make([]*FinantialTransaction,0)
	stampBilnDate:=timestamppb.New(bilngDate)
	if fee==nil || cust==nil{
		return nil,errors.New("Invalied request")
	}
	if fee.IsChargable==nil ||*fee.IsChargable==false {
		return resp,nil
	}
	if fee.EffectDate==nil{
		return nil,errors.New("Missing Effect Date for charge regular")
	}
	if fee.Code==nil{
		return nil,errors.New("Missing TransCode Date for charge regular")
	}
	isEnabled,err:=check(fee,cust,bilngDate,nil)
	if err!=nil{
		return nil,err
	}
	if !isEnabled{
		return resp,nil
	}
	transEffectDate:=stampBilnDate
	if fee.ChargeCalcPeriod!=nil && *fee.ChargeCalcPeriod==RegularChargePeriod_MONTHLY{
		dy:=1
		if fee.ChargeMonthlyDay!=nil{
			dy=int(*fee.ChargeMonthlyDay)
		}
		nwTransEffDate:=time.Date(bilngDate.Year(),bilngDate.Month(),dy,0,0,0,0,time.Local)
		transEffectDate=timestamppb.New(nwTransEffDate)
	}
	var noUnits int64=1
	var mainCtype=""
	propRef:=""
	var conn *Connection
	if cust.Property!=nil && cust.Property.Services!=nil {
		propRef=cust.Property.GetPropRef()
		for _,sv:=range cust.Property.Services{
			if sv!=nil && sv.Connection!=nil && sv.GetServiceType()==fee.GetServiceType(){
				if sv.Connection.NoUnits!=nil && *sv.Connection.NoUnits>noUnits{
					noUnits=*sv.Connection.NoUnits
				}
				if sv.Connection.CType!=nil{
					mainCtype=*sv.Connection.CType
				}
				conn=sv.Connection
			}
		}
	}
	if noUnits<1{
		noUnits=1
	}
	// calculate charge for fixed type
	if fee.ChargeType!=nil || *fee.ChargeType==ChargeType_FIXED{
		if fee.FixedCharge==nil{
			return nil,errors.New(fmt.Sprintf("missing fixed value for charge regular %v",fee.Code))
		}
		amount=*fee.FixedCharge
		if fee.PerUnit!=nil && *fee.PerUnit &&noUnits>1 {
			amount=amount*float64(noUnits)
		}
		if fee.GetVatPercentage()>0{
			taxAmount=amount*fee.GetVatPercentage()/float64(100)
		}
		resp=append(resp,&FinantialTransaction{
			ServiceType:fee.ServiceType,
			Code:fee.Code,
			Amount:&amount,
			TaxAmount:&taxAmount,
			BilngDate:stampBilnDate,
			EffDate:transEffectDate,
			NoUnits:&noUnits,
			Ctype:&mainCtype,
			PropRef:&propRef,
		})
		return resp,nil
	}
	/////////////////CALC//////////////////////
	log.Printf("charge type %v",*fee.ChargeType)
	if fee.RelationChargeEntity==nil{
		return nil,errors.New("missing charge entity for charge regular")
	}
	ree:=fee.RelationChargeEntity
	if ree.EntityType==nil{
		return nil,errors.New("missing charge entity type for charge regular")
	}
	typ:=*ree.EntityType
	var feeValues=ree.MappedValues
	var customerValues=customerValues(typ,cust)
	if feeValues==nil || len(feeValues)==0 || customerValues==nil || len(customerValues)==0{
		return resp,nil
	}
	mappedValues:=map[string]float64{}
	for _,cstValue:=range customerValues{
		if cstValue==nil{
			continue
		}
		found:=false
		for _,m:=range feeValues{
			if tools.StringComparePointer(m.LuKey,cstValue){
				found=true
				mappedValues[*cstValue]=m.GetValue()
			}
			if found{
				return nil,errors.New("missing lookup for charge regular:"+fee.GetCode()+" "+*cstValue)
			}
		}
	}
	if len (mappedValues)==0{
		return resp,nil
	}
	var totalMappedValue float64=0
	for _,v:=range mappedValues{
		totalMappedValue=totalMappedValue+v
	}
	if len (mappedValues)==1 || fee.CTypeCalcBase==nil || *fee.CTypeCalcBase==ChargeRegularCTypeCalcStrategy_SUM_CTYPES || typ!=ENTITY_TYPE_CTYPE{
		amount=totalMappedValue
		if fee.PerUnit!=nil && *fee.PerUnit &&noUnits>1 {
			amount=amount*float64(noUnits)
		}
		if fee.GetVatPercentage()>0{
			taxAmount=amount*fee.GetVatPercentage()/float64(100)
		}
		resp=append(resp,&FinantialTransaction{
			ServiceType:fee.ServiceType,
			Code:fee.Code,
			Amount:&amount,
			TaxAmount:&taxAmount,
			BilngDate:stampBilnDate,
			EffDate:transEffectDate,
			NoUnits:&noUnits,
			Ctype:&mainCtype,
			PropRef:&propRef,
		})
		return resp,nil
	}
	calcBase:=*fee.CTypeCalcBase
	if calcBase==ChargeRegularCTypeCalcStrategy_EACH_CTYPE{
		for k,v:=range mappedValues{
			var amt float64=v
			var tax float64=0
			if fee.PerUnit!=nil && *fee.PerUnit {
				var subNoUnits=1;
				if conn!=nil && conn.SubConnections!=nil {
					for _,sbConn:=range conn.SubConnections{
						if sbConn!=nil && sbConn.NoUnits!=nil && sbConn.CType!=nil && *sbConn.CType==k{

						}
					}
				}
				amt=amt*float64(subNoUnits)
			}
			if fee.GetVatPercentage()>0{
				tax=amount*fee.GetVatPercentage()/float64(100)
			}
			resp=append(resp,&FinantialTransaction{
				ServiceType:fee.ServiceType,
				Code:fee.Code,
				Amount:&amt,
				TaxAmount:&tax,
				BilngDate:stampBilnDate,
				EffDate:transEffectDate,
				NoUnits:&noUnits,
				Ctype:&k,
				PropRef:&propRef,
			})
		}
		return resp,nil
	}
	if calcBase==ChargeRegularCTypeCalcStrategy_MAIN_CTYPE{
		var amt float64=0
		var tax float64=0
		for k,v:=range mappedValues{
			if k==mainCtype{
				amt=v
			}
		}
		if fee.PerUnit!=nil && *fee.PerUnit &&noUnits>1 {
			amt=amt*float64(noUnits)
		}
		if fee.GetVatPercentage()>0{
			tax=amount*fee.GetVatPercentage()/float64(100)
		}
		resp=append(resp,&FinantialTransaction{
			ServiceType:fee.ServiceType,
			Code:fee.Code,
			Amount:&amt,
			TaxAmount:&tax,
			BilngDate:stampBilnDate,
			EffDate:transEffectDate,
			NoUnits:&noUnits,
			Ctype:&mainCtype,
			PropRef:&propRef,
		})

		return resp,nil
	}
	return nil,errors.New("Unkown ctype calculation strategy")
}
