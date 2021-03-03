package charge

import (. "MaisrForAdvancedSystems/go-biller/proto"
	"errors"
	"time"
)

func CalcRegularCharge(fee *RegularCharge,cust *Customer,bilngDate time.Time,lastCharge *time.Time) (*float64,error){
	var amount float64=0
	if fee==nil || cust==nil{
		return nil,errors.New("Invalied request")
	}
	if fee.IsChargable==nil ||*fee.IsChargable==false {
		return nil,nil
	}
	if fee.EffectiveDate==nil{
		return nil,errors.New("Missing Effect Date for charge regular")
	}
	var effDate time.Time=fee.EffectiveDate.AsTime()
	if effDate.After(bilngDate){
		return nil, nil
	}
	if fee.TransCode==nil{
		return nil,errors.New("Missing TransCode Date for charge regular")
	}
	if fee.TransSCode==nil{
		return nil,errors.New("Missing TransSCode Date for charge regular")
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
		isEnabled,err:=IsChargeEnable(fee,cust)
		if err!=nil{
			return nil,err
		}
		if !isEnabled{
			return nil,nil
		}
	}
	///validation of entity
	return &amount,nil
}
