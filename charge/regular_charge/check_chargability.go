package regular_charge

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"log"
	"time"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

func check(fee *RegularCharge, c *Customer, bilngDate time.Time, lastChargeDate *time.Time) (bool, CustomerValues, error) {
	var approvedValues CustomerValues = make(map[string]*MappedData)
	var custValues CustomerValues = make(map[string]*MappedData)
	if fee == nil {
		return false, approvedValues, nil
	}
	if fee.IsChargable == nil || !*fee.IsChargable {
		return false, approvedValues, nil
	}
	if *fee.ChargeCalcPeriod == RegularChargePeriod_MONTHLY {
		if lastChargeDate != nil {
			var period int64 = 1 //on month
			if fee.ChargeInterval != nil && *fee.ChargeInterval > 1 {
				period = period
			}
			nextChargeDate := lastChargeDate.AddDate(0, int(period), 0)
			if nextChargeDate.After(time.Now()) {
				return false, approvedValues, nil
			}
		}
	} else {
		if fee.EffectDate == nil {
			return false, approvedValues, errors.New("Missing Effect Date for charge regular")
		}
		var effDate time.Time = fee.EffectDate.AsTime()
		if effDate.After(bilngDate) {
			return false, approvedValues, nil
		}
	}
	if fee.RelationEnableEntity != nil {
		custValues = customerValues(fee.RelationEnableEntity.GetEntityType(), c, fee.ServiceType)
	}
	if fee.Bypass != nil {
		if *fee.Bypass {
			return true, custValues, nil
		}
	}
	if fee.ServiceType == nil {
		log.Println("missing service type")
		return false, approvedValues, nil
	}
	if c.Property == nil {
		log.Println("missing property")
		return false, approvedValues, nil
	}
	if c.Property.Services == nil || len(c.Property.Services) == 0 {
		log.Println("missing property services")
		return false, approvedValues, nil
	}
	haveService := false
	for _, sv := range c.Property.Services {
		if sv.ServiceType != nil && *sv.ServiceType == *fee.ServiceType {
			haveService = true
			break
		}
	}
	if !haveService {
		log.Println("no customer services match")
		return false, approvedValues, nil
	}
	if fee.RelationEnableEntity == nil {
		return false, approvedValues, errors.New("missing enabled entity for charge regular")
	}
	ree := fee.RelationEnableEntity
	if ree.EntityType == nil {
		return false, approvedValues, errors.New("missing enabled entity type for charge regular")
	}
	log.Printf("MappedValues  %v", ree.MappedValues)
	log.Printf("custValues  %v", custValues)
	//typ:=*ree.EntityType
	var mappedValues = ree.MappedValues
	if mappedValues == nil || len(mappedValues) == 0 || custValues == nil || len(custValues) == 0 {
		log.Printf("no mapped")
		return false, approvedValues, nil
	}
	found := false
	for cstValue, _ := range custValues {
		for _, m := range mappedValues {
			if tools.StringComparePointer(m.LuKey, &cstValue) {
				if m.GetValue() {
					found = true
				}
				k := cstValue //copy
				approvedValues[k] = nil
				if custValues[cstValue] != nil {
					val := *custValues[cstValue]
					approvedValues[k] = &val
				}
			}
		}
	}
	log.Printf("found  %v", found)

	return found, approvedValues, nil
}
