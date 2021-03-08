package regular_charge

import (. "MaisrForAdvancedSystems/go-biller/proto"
	"errors"
	"fmt"
	"log"
	"time"
)
type RegularChargeAmount struct {
	Amount float64
	TaxAmount *float64
}
func CalcCharge(fee *RegularCharge,cust *Customer,bilngDate time.Time,lastCharge *time.Time) (*RegularChargeAmount,error){
	log.Println("calc reg charge:"+*fee.Code)
	var amount float64=0
	var taxAmount float64=0
	if fee==nil || cust==nil{
		return nil,errors.New("Invalied request")
	}
	if fee.IsChargable==nil ||*fee.IsChargable==false {
		return nil,nil
	}
	if fee.EffectDate==nil{
		return nil,errors.New("Missing Effect Date for charge regular")
	}
	var effDate time.Time=fee.EffectDate.AsTime()
	if effDate.After(bilngDate){
		log.Println(effDate.String(),bilngDate.String())
		log.Println("charge skipped effect date")
		return nil, nil
	}
	if fee.Code==nil{
		return nil,errors.New("Missing TransCode Date for charge regular")
	}
	if fee.ChargeCalcPeriod==nil{
		return nil,errors.New("Missing Calc Period Date for charge regular")
	}
	bypass:=false
	if fee.Bypass!=nil{
		bypass=*fee.Bypass
	}
	if !bypass{
		if fee.RelationEnableEntity==nil{
			return nil,errors.New("missing enabled entity for charge regular")
		}
		ree:=fee.RelationEnableEntity
		if ree.EntityType==nil{
			return nil,errors.New("missing enabled entity type for charge regular")
		}
		isEnabled,err:=check(fee,cust,bilngDate,nil)
		if err!=nil{
			return nil,err
		}
		if !isEnabled{
			return nil,nil
		}
	}
	// calculate charge for fixed type
	if fee.ChargeType!=nil || *fee.ChargeType==ChargeType_FIXED{
		if fee.FixedCharge==nil{
			return nil,errors.New(fmt.Sprintf("missing fixed value for charge regular %v",fee.Code))
		}
		amount=*fee.FixedCharge
		var noUnits int64=0
		if cust.Property!=nil && cust.Property.Services!=nil {
			for _,sv:=range cust.Property.Services{
				if sv!=nil && sv.Connection!=nil && sv.Connection.NoUnits!=nil{
					noUnits=noUnits+*sv.Connection.NoUnits
				}
			}
		}
		if fee.PerUnit!=nil && *fee.PerUnit{
			if fee.GetVatPercentage()>0{
				taxAmount=amount*fee.GetVatPercentage()/float64(100)
			}
			return &RegularChargeAmount{
				Amount:amount,
				TaxAmount:&taxAmount,
			},nil
		}
	}
	/////////////////CALC//////////////////////
	if fee.RelationChargeEntity.EntityType==nil{
	}
	custValues:=customerValues(*fee.RelationChargeEntity.EntityType,cust)
	log.Println(custValues)
	///validation of entity
	if fee.GetVatPercentage()>0{
		taxAmount=amount*fee.GetVatPercentage()/float64(100)
	}
	return &RegularChargeAmount{
		Amount:amount,
		TaxAmount:&taxAmount,
	},nil
}
