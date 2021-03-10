package regular_charge

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"fmt"
	"log"
	"time"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegularChargeAmount struct {
	Amount    float64
	TaxAmount *float64
}

func CalcCharge(fee *RegularCharge, cust *Customer, bilngDate time.Time, lastCharge *time.Time) ([]*FinantialTransaction, error) {
	log.Println("calc reg charge:" + *fee.TransCode)
	resp := make([]*FinantialTransaction, 0)
	stampBilnDate := timestamppb.New(bilngDate)
	if fee == nil || cust == nil {
		return nil, errors.New("Invalied request")
	}
	if fee.IsChargable == nil || *fee.IsChargable == false {
		return resp, nil
	}
	if fee.EffectDate == nil {
		return nil, errors.New("Missing Effect Date for charge regular")
	}
	if fee.TransCode == nil {
		return nil, errors.New("Missing TransCode Date for charge regular")
	}
	isEnabled, approvedValues, err := check(fee, cust, bilngDate, nil)
	if err != nil {
		return nil, err
	}
	if !isEnabled {
		return resp, nil
	}
	transEffectDate := stampBilnDate
	if fee.ChargeCalcPeriod != nil && *fee.ChargeCalcPeriod == RegularChargePeriod_MONTHLY {
		dy := 1
		if fee.ChargeMonthlyDay != nil {
			dy = int(*fee.ChargeMonthlyDay)
		}
		nwTransEffDate := time.Date(bilngDate.Year(), bilngDate.Month(), dy, 0, 0, 0, 0, time.Local)
		transEffectDate = timestamppb.New(nwTransEffDate)
	}
	var mainNoUnits int64 = 1
	var mainCtype = ""
	propRef := ""
	if cust.Property != nil && cust.Property.Services != nil {
		propRef = cust.Property.GetPropRef()
		for _, sv := range cust.Property.Services {
			if sv != nil && sv.Connection != nil && sv.GetServiceType() == fee.GetServiceType() {
				if sv.Connection.NoUnits != nil && *sv.Connection.NoUnits > mainNoUnits {
					mainNoUnits = *sv.Connection.NoUnits
				}
				if sv.Connection.CType != nil {
					mainCtype = *sv.Connection.CType
				}
			}
		}
	}
	if mainNoUnits < 1 {
		mainNoUnits = 1
	}
	// calculate charge for fixed type
	if fee.ChargeType == nil || *fee.ChargeType == ChargeType_FIXED {
		log.Println("charge fixed regular charge")
		var amount float64 = 0
		var taxAmount float64 = 0
		if fee.FixedCharge == nil {
			return nil, errors.New(fmt.Sprintf("missing fixed value for charge regular %v", fee.TransCode))
		}
		amount = *fee.FixedCharge
		if fee.PerUnit != nil && *fee.PerUnit && mainNoUnits > 1 {
			amount = amount * float64(mainNoUnits)
		}
		if fee.GetVatPercentage() > 0 {
			taxAmount = amount * fee.GetVatPercentage() / float64(100)
		}
		resp = append(resp, &FinantialTransaction{
			ServiceType: fee.ServiceType,
			Code:        fee.TransCode,
			Amount:      &amount,
			TaxAmount:   &taxAmount,
			BilngDate:   stampBilnDate,
			EffDate:     transEffectDate,
			NoUnits:     &mainNoUnits,
			Ctype:       &mainCtype,
			PropRef:     &propRef,
		})
		return resp, nil
	}
	/////////////////CALC//////////////////////
	log.Printf("charge type %v", *fee.ChargeType)
	if fee.RelationChargeEntity == nil {
		return nil, errors.New("missing charge entity for charge regular")
	}
	ree := fee.RelationChargeEntity
	if ree.EntityType == nil {
		return nil, errors.New("missing charge entity type for charge regular")
	}
	chargeEntitytype := *ree.EntityType
	var feeValues = ree.MappedValues
	var custvalues = customerValues(chargeEntitytype, cust, fee.ServiceType) //all disitnct values
	if custvalues == nil || len(custvalues) == 0 {
		return resp, nil
	}
	if feeValues == nil || len(feeValues) == 0 || len(custvalues) == 0 {
		return resp, nil
	}
	mappedValues := map[string]struct {
		Value   float64
		NoUnits *int64
		Ctype   *string
	}{}
	for cstValue, _mapedValue := range custvalues {
		if cstValue == "" {
			continue
		}
		found := false
		mv := *_mapedValue //copy data
		for _, m := range feeValues {
			if tools.StringComparePointer(m.LuKey, &cstValue) {
				found = true
				mappedValues[cstValue] = struct {
					Value   float64
					NoUnits *int64
					Ctype   *string
				}{Value: m.GetValue(), NoUnits: mv.noUnits, Ctype: mv.cType}
			}
		}
		if !found {
			return nil, errors.New("missing lookup for charge regular:" + fee.GetTransCode() + " " + cstValue)
		}
	}
	if len(mappedValues) == 0 {
		return resp, nil
	}
	if fee.CTypeCalcBase == nil || *fee.CTypeCalcBase == ChargeRegularCalcStrategy_EACH_ONE {
		for k := range mappedValues {
			v := mappedValues[k]
			var amt float64 = v.Value //copy value
			var tax float64 = 0
			var subNoUnits int64 = 1
			if fee.PerUnit != nil && *fee.PerUnit {
				subNoUnits = tools.DefaultI(v.NoUnits, int64(1))
			}
			amt = amt * float64(subNoUnits)
			if fee.GetVatPercentage() > 0 {
				tax = amt * fee.GetVatPercentage() / float64(100)
			}
			resp = append(resp, &FinantialTransaction{
				ServiceType: fee.ServiceType,
				Code:        fee.TransCode,
				Amount:      &amt,
				TaxAmount:   &tax,
				BilngDate:   stampBilnDate,
				EffDate:     transEffectDate,
				NoUnits:     &subNoUnits,
				Ctype:       v.Ctype,
				PropRef:     &propRef,
			})
		}
		return resp, nil
	}
	calcBase := *fee.CTypeCalcBase
	var singleAmount float64 = 0
	var maxAmount float64 = 0
	for k := range mappedValues {
		v := mappedValues[k]
		if v.Value > maxAmount {
			maxAmount = v.Value
		}
		if calcBase == ChargeRegularCalcStrategy_SUM_ALL {
			singleAmount = singleAmount + v.Value
		}
		if calcBase == ChargeRegularCalcStrategy_HIGHTEST_AMOUNT {
			if v.Value > singleAmount {
				singleAmount = v.Value
			}
		}
		if calcBase == ChargeRegularCalcStrategy_LOWEST_AMOUNT {
			if v.Value < singleAmount {
				singleAmount = v.Value
			}
		}
	}
	if calcBase == ChargeRegularCalcStrategy_HIGHTEST_AMOUNT ||
		calcBase == ChargeRegularCalcStrategy_LOWEST_AMOUNT ||
		calcBase == ChargeRegularCalcStrategy_SUM_ALL {
		var tax float64 = 0
		amt := singleAmount
		if fee.PerUnit != nil && *fee.PerUnit && mainNoUnits > 1 {
			amt = amt * float64(mainNoUnits)
		}
		if fee.GetVatPercentage() > 0 {
			tax = singleAmount * fee.GetVatPercentage() / float64(100)
		}
		resp = append(resp, &FinantialTransaction{
			ServiceType: fee.ServiceType,
			Code:        fee.TransCode,
			Amount:      &amt,
			TaxAmount:   &tax,
			BilngDate:   stampBilnDate,
			EffDate:     transEffectDate,
			NoUnits:     &mainNoUnits,
			Ctype:       &mainCtype,
			PropRef:     &propRef,
		})
		return resp, nil
	}
	//for giza style they want to repeate the value for each ctype
	//giza style work with ctype group
	if fee.RelationEnableEntity == nil {
		return nil, errors.New("can not define calculation strategy missing RelationEnableEntity")
	}

	if !(fee.RelationEnableEntity.GetEntityType() == ENTITY_TYPE_CTYPE_GROUP || fee.RelationEnableEntity.GetEntityType() == ENTITY_TYPE_CTYPE) {
		return nil, errors.New("giza calculation strategy must be apply only if enable entity is ctype or ctype group")
	}

	for k := range approvedValues {
		apVale := approvedValues[k]
		amt := maxAmount
		var tax float64 = 0
		var noUnits int64 = 1
		if fee.PerUnit != nil && *fee.PerUnit && apVale.noUnits != nil {
			noUnits = *apVale.noUnits
			if noUnits < 1 {
				noUnits = 1
			}
			amt = amt * float64(noUnits)
		}
		if fee.GetVatPercentage() > 0 {
			tax = amt * fee.GetVatPercentage() / float64(100)
		}
		resp = append(resp, &FinantialTransaction{
			ServiceType: fee.ServiceType,
			Code:        fee.TransCode,
			Amount:      &amt,
			TaxAmount:   &tax,
			BilngDate:   stampBilnDate,
			EffDate:     transEffectDate,
			NoUnits:     &noUnits,
			Ctype:       apVale.cType,
			PropRef:     &propRef,
		})
		return resp, nil
	}
	return resp, nil
}
