package regular_charge

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MaisrForAdvancedSystems/biller-charger/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

func check(fee *RegularCharge, c *Customer, bilngDate time.Time, lastChargeDate *time.Time) (bool, CustomerValues, error) {
	log.Println("Check..")
	var approvedValues CustomerValues = make(map[string]*MappedData)
	var custValues CustomerValues = make(map[string]*MappedData)
	if fee == nil {
		return false, approvedValues, nil
	}
	if fee.IsChargable == nil || !*fee.IsChargable {
		return false, approvedValues, nil
	}
	if fee.EffectDate == nil {
		return false, approvedValues, errors.New("Missing Effect Date for charge regular")
	}
	var effDate time.Time = fee.EffectDate.AsTime()
	if effDate.After(bilngDate) {
		return false, approvedValues, nil
	}
	if fee.EffectDateTo != nil && bilngDate.After(fee.EffectDateTo.AsTime()) {
		return false, approvedValues, nil
	}
	if fee.RelationEnableEntity != nil {
		log.Println("fee entity", fee.RelationEnableEntity.GetEntityType())
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
	for kk, vv := range custValues {
		log.Println("custValues ", "key:", kk)
		if vv != nil {
			log.Println("custValues Not Null")
			if vv.cType != nil {
				log.Println("custValues ", " ctype:", *vv.cType)
			}
			if vv.cTypeGroup != nil {
				log.Println("custValues ", " group:", *vv.cTypeGroup)
			}
			if vv.noUnits != nil {
				log.Println("custValues ", " unitsNo:", *vv.noUnits)
			}
		}
	}

	//typ:=*ree.EntityType
	var mappedValues = ree.MappedValues
	if mappedValues == nil || len(mappedValues) == 0 {
		return false, approvedValues, errors.New(fmt.Sprintf("missing regular charge values :%v @ %v", ree.EntityType.String(), fee.RegularChargeId))
	}
	if custValues == nil || len(custValues) == 0 {
		return false, approvedValues, errors.New(fmt.Sprintf("regular charge:missing customer values :%v @ %v", ree.EntityType.String(), fee.RegularChargeId))
	}
	enabled := false
	found := false
	for cstValue, _ := range custValues {
		for _, m := range mappedValues {
			if tools.StringComparePointer(m.LuKey, &cstValue) {
				found = true
				if m.GetValue() {
					enabled = true
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
	if !found {
		keys := []string{}
		for k, _ := range custValues {
			keys = append(keys, k)
		}
		missidValues := strings.Join(keys, " , ")
		return false, approvedValues, errors.New(fmt.Sprintf("regular charge:missing  values :%v @ %v values is :%s", ree.EntityType.String(), fee.RegularChargeId, missidValues))
	}
	log.Printf("approvedValues  %v", approvedValues)
	log.Printf("enabled  %v", enabled)
	return enabled, approvedValues, nil
}
