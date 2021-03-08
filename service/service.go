package service

import (
	tr "MaisrForAdvancedSystems/go-biller/charge/tariff_calc"
	rg "MaisrForAdvancedSystems/go-biller/charge/regular_charge"
	"MaisrForAdvancedSystems/go-biller/consumptions"
	. "MaisrForAdvancedSystems/go-biller/proto"
	"context"
	errors "errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)
var empty Empty=Empty{}
var ChargeServiceVersion="v1.0.0"
type BillingService struct {
	Tariffs map[string]map[time.Time]*Tariff
	Ctgs    map[string]*Ctg
	RegularCharges []*RegularCharge
	IsTrace bool
}
type RegularChargeResponceItem struct {
	Reg *RegularCharge
	Amount float64
	TaxAmount *float64
}
// Charge calculate charge for all services
func (s *BillingService) Charge(c context.Context, r *ChargeRequest) (*BillResponce, error) {
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if s.Tariffs == nil {
		return nil, errors.New("Missing Traiff lookup")
	}
	if r == nil {
		return nil, errors.New("invalied request")
	}
	if r.Customer == nil {
		return nil, errors.New("invalied request:missing customer")
	}
	log.Println("bilngDate",r.Setting.BilingDate.String())
	isCustValied,err:=s.ValidateCustomer(c,r.Customer)
	if err!=nil{
		return nil,err
	}
	if isCustValied==nil || !*isCustValied{
		s.Trace("valid result ",isCustValied)
		return nil, errors.New("customer data is not valied ")
	}
	if r.Customer.Property == nil {
		return nil, nil
	}
	services:=r.Customer.Property.Services
	if services==nil || len(services)==0{
		return nil,nil
	}
	serviceRdgs:=make(map[*Service]*ServiceReading)
	if r.ServicesReadings != nil && len(r.ServicesReadings) >0 {
		for idx:=range services{
			srv:=services[idx]
			for sx:=range r.ServicesReadings{
				if *r.ServicesReadings[sx].ServiceType==*srv.ServiceType{
					serviceRdgs[srv]=r.ServicesReadings[sx]
					break
				}
			}
		}
	}
	resp:=BillResponce{}
	resp.FTransactions=make([]*FinantialTransaction,0)
	for k,v:=range serviceRdgs{
		s.Trace(k.ServiceType.String())
		sTrans,err:=s.CalcForService(r.Setting,k,v)
		if err!=nil{
			s.Trace(err)
			return nil,err
		}
		s.Trace(sTrans)
		for id:=range sTrans{
			t:=sTrans[id]
			resp.FTransactions=append(resp.FTransactions,t)
		}
	}
	regsTrans,err:=s.RegCharge(r.Customer,r.Setting.BilingDate.AsTime(),nil)
	if err!=nil{
		return nil,err
	}
	log.Println("regsTrans",len(regsTrans))
	for id:=range regsTrans{
		fee:=regsTrans[id]
		feesTrans:=&FinantialTransaction{
			ServiceType:          fee.Reg.ServiceType,
			Code:                 fee.Reg.Code,
			BilngDate:            r.Setting.BilingDate,
			EffDate:              r.Setting.BilingDate,
			Amount:               &fee.Amount,
			TaxAmount:           fee.TaxAmount,
			DiscountAmount:       nil,
			Ctype:                getCtype(r.Customer,fee.Reg.GetServiceType()),
			NoUnits:              getNoUnits(r.Customer,fee.Reg.GetServiceType()),
			PropRef:              nil,
		}
		resp.FTransactions=append(resp.FTransactions,feesTrans)
	}
	return &resp, nil
}

func (s *BillingService) RegCharge(customer * Customer,bilngDate time.Time,lastCharge *time.Time) ([]*RegularChargeResponceItem,error){
	if s.RegularCharges==nil || len (s.RegularCharges)==0{
		return nil,nil
	}
	rs:=make([]*RegularChargeResponceItem,0)
	for id:=range s.RegularCharges{
		regCh:=s.RegularCharges[id]
		if regCh==nil{
			return nil,errors.New("invalied regular charge setup")
		}
		chAmt,err:=rg.CalcCharge(regCh,customer,bilngDate,lastCharge)
		if err!=nil{
			return nil,err
		}
		if chAmt!=nil{
			rs=append(rs,&RegularChargeResponceItem{
				Reg:regCh,
				Amount:chAmt.Amount,
				TaxAmount:chAmt.TaxAmount,
			})
		}
	}
	return rs,nil
}

//calc for service and may be not used
func (s *BillingService) CalcForService(setting *ChargeSetting,service *Service,reading *ServiceReading) ([]*FinantialTransaction, error) {
	var conn=service.Connection
	if conn==nil{
		return nil,nil
	}
	cons,err:=consumptions.GetConnectionsConsumption(s.Ctgs,service.Connection,setting.BilingDate.AsTime(),reading.Reading,s.IsTrace)
	if err!=nil{
		return nil,err
	}
	trans:=make([]*FinantialTransaction,0)
	var crReading *float64=nil
	var prReading *float64=nil
	var MeterType *string =nil
	var MeterRef *string =nil
	serviceType:=service.GetServiceType()
	bilngDate:=timestamppb.New(setting.BilingDate.AsTime())
	if reading.Reading!=nil{
		crReading=reading.Reading.CrReading
		prReading=reading.Reading.PrReading
	}
	if conn.Meter!=nil{
		MeterType=conn.Meter.MeterType
		MeterRef=conn.Meter.MeterRef
	}
	for _,c:=range cons{
		ctg,ok:=s.Ctgs[c.Ctype]
		if !ok{
			return nil,errors.New("missing ctype :"+c.Ctype)
		}
		tarifId:="??"
		isZeroTariff:=false
		transCode :="??"
		var taxPercentage float64=0
		var discountPercentage float64=0

		if ctg.Tariffs!=nil{
			for _,srvTar:=range ctg.Tariffs{
				if *srvTar.ServiceType==*service.ServiceType{
					if srvTar.IsZeroTarif!=nil{
						isZeroTariff=*srvTar.IsZeroTarif
					}
					if srvTar.TarifId==nil && !isZeroTariff{
						return nil,errors.New("missing tariff :"+c.Ctype)
					}else{
						tarifId=*srvTar.TarifId
						if srvTar.Code!=nil{
							transCode=*srvTar.Code
						}else {
							return nil,errors.New("missing trans code tariff :"+c.Ctype)
						}
						if srvTar.TaxPercentage!=nil && *srvTar.TaxPercentage>0{
							taxPercentage=*srvTar.TaxPercentage
						}
						if srvTar.DiscountPercentage!=nil && *srvTar.DiscountPercentage>0{
							discountPercentage=*srvTar.DiscountPercentage
						}
					}
				}
			}
			if taxPercentage>100{
				taxPercentage=100
			}
			if taxPercentage<0{
				taxPercentage=0
			}
			if discountPercentage>100{
				discountPercentage=100
			}
			if discountPercentage<0{
				discountPercentage=0
			}
			tariff,err:=s.FindBestTarrifMatch(tarifId,setting.GetBilingDate().AsTime())
			if err!=nil{
				return nil,err
			}
			crgAmt,err:=tr.Calc(c.NoUnits,c.Consump,tariff,isZeroTariff)
			if err!=nil{
				return nil,err
			}
			if crgAmt!=nil && *crgAmt>=0{
				taxAmount:=taxPercentage*(*crgAmt)/float64(100)
				disAmount:=discountPercentage*(*crgAmt)/float64(100)
				amount:=*crgAmt-disAmount
				trans=append(trans,&FinantialTransaction{
					Code:           &transCode,
					EffDate:        bilngDate,
					BilngDate:      bilngDate,
					Amount:         &amount,
					TaxAmount:      &taxAmount,
					DiscountAmount: &disAmount,
					Ctype:          &c.Ctype,
					PropRef:        nil,
					ServiceType:    &serviceType,
					NoUnits:        &c.NoUnits,
					MTransaction:&MeasuredTransaction{
						CrReading: crReading,
						PrReading: prReading,
						Consump:   &c.Consump,
						ReadType:  &c.ReadType,
						MeterType: MeterType,
						MeterRef:  MeterRef,
					},
				})
			}
		}
	}
	return trans,nil
}

func (s *BillingService) FindBestTarrifMatch(tarId string, bilngDate time.Time) (*Tariff,error){
	if s.Tariffs==nil || len(s.Tariffs)==0 {
		return nil,errors.New("no tariffes defined")
	}
	tarifs,ok:=s.Tariffs[tarId]
	if !ok{
		s.Trace(s.Tariffs)
		return nil,errors.New("missing tariff "+tarId)
	}
	var bestUp *time.Time=nil
	var bestDown *time.Time=nil
	for dt,_:=range tarifs{
		iDate:=dt // copy of date value
		if bestUp==nil{
			bestUp=&iDate
		}
		if bestDown==nil{
			bestDown=&iDate
		}
		if iDate.After(bilngDate) && iDate.Before(*bestUp){
			bestUp=&iDate
		}
		if iDate.Before(bilngDate) && iDate.After(*bestDown){
			bestDown=&iDate
		}
	}
	if bestDown.After(*bestUp){
		bestDown=&(*bestUp)
	}
	return tarifs[*bestDown],nil
}

func getCtype(cust * Customer,serviceType SERVICE_TYPE) *string{
	if cust==nil || cust.Property==nil || cust.Property.Services==nil{
		return nil
	}
	if len(cust.Property.Services)==0{
		return nil
	}
	for _,sv:=range cust.Property.Services{
		if sv!=nil && sv.GetServiceType()==serviceType && sv.Connection!=nil{
			return sv.Connection.CType
		}
	}
	return nil
}


func getNoUnits(cust * Customer,serviceType SERVICE_TYPE) *int64{
	if cust==nil || cust.Property==nil || cust.Property.Services==nil{
		return nil
	}
	if len(cust.Property.Services)==0{
		return nil
	}
	for _,sv:=range cust.Property.Services{
		if sv!=nil && sv.GetServiceType()==serviceType && sv.Connection!=nil{
			return sv.Connection.NoUnits
		}
	}
	return nil
}




