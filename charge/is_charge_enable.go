package charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"time"
)


func IsChargeEnable(fee *RegularCharge,c *Customer,bilngDate time.Time,lastChargeDate *time.Time) (bool,error){
	if fee==nil{
		return false,nil
	}
	if fee.IsChargable==nil || !*fee.IsChargable{
		return false,nil
	}
	if fee.ChargeCalcPeriod!=nil || *fee.ChargeCalcPeriod==RegularChargePeriod_MONTHLY{
		if fee.EffectiveDate==nil{
			return false,errors.New("Missing Effect Date for charge regular")
		}
		var effDate time.Time=fee.EffectiveDate.AsTime()
		if effDate.After(bilngDate){
			return false, nil
		}
	}else {
		if fee.EffectiveDate==nil{
			return false,errors.New("Missing Effect Date for charge regular")
		}
		var effDate time.Time=fee.EffectiveDate.AsTime()
		if effDate.After(bilngDate){
			return false, nil
		}
	}

	if fee.Bypass!=nil{
		if *fee.Bypass{
			return true,nil
		}
	}
	if fee.ServiceType==nil{
		return false,nil
	}
	if c.Property==nil{
		return false,nil
	}
	if c.Property.Services==nil || len(c.Property.Services)==0{
		return false,nil
	}
	found:=false
	for _,sv:=range c.Property.Services{
		if sv.ServiceType!=nil && *sv.ServiceType==*fee.ServiceType{
			found=true
			break
		}
	}
	if !found{
		return false,nil
	}
	if fee.ChargeType==nil && *fee.ChargeType==ChargeType_FIXED{
		return true,nil
	}
	if fee.RelationEnableEntity==nil{
		return false,errors.New("missing enabled entity for charge regular")
	}
	ree:=fee.RelationEnableEntity
	if ree.EntityType==nil{
		return false,errors.New("missing enabled entity type for charge regular")
	}
	typ:=*ree.EntityType
	var mappedValues=ree.MappedValues
	var customerValues=CustomerValues(typ,c)
	if mappedValues==nil || len(mappedValues)==0 || customerValues==nil || len(customerValues)==0{
		return false,nil
	}
	for _,cstValue:=range customerValues{
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,cstValue){
				return m.GetValue(),nil
				break;
			}
		}
	}
	return false,nil
}

