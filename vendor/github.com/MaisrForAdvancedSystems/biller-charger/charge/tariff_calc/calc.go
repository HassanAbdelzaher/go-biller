package tariff_calc

import (
	errors "errors"
	"log"
	"time"

	"github.com/MaisrForAdvancedSystems/biller-charger/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

//Calc portable function for claulate charge for certain c_type
func Calc(service *Service, setting *ChargeSetting, no_units int64, consump float64, read_type READING_TYPE, cr_date *time.Time, pr_date *time.Time, tariff *Tariff, isZeroTarif bool) (*float64, error) {
	var amt float64 = 0
	if isZeroTarif {
		return &amt, nil
	}
	if no_units < 1 {
		no_units = 1
	}
	var cycleLength int64 = 1
	ignoreTimeEffect := false
	if setting != nil {
		if setting.IgnoreTimeEffect != nil {
			ignoreTimeEffect = *setting.IgnoreTimeEffect
		}
		if setting.CycleLength != nil {
			cycleLength = *setting.CycleLength
		}
	}
	if cycleLength < 1 {
		cycleLength = 1
	}
	no_units = no_units * cycleLength
	var noMonths int64 = 1
	if !ignoreTimeEffect && pr_date != nil && cr_date != nil && read_type == READING_TYPE_ACTUAL {
		if cr_date.After(*pr_date) {
			hours := cr_date.Sub(*pr_date).Hours()
			dayes := hours / (24)
			noMonths = int64(dayes) / 30
		}
	}
	if noMonths < 1 {
		noMonths = 1
	}
	no_units = no_units * noMonths
	if tariff == nil {
		return nil, errors.New("missing tarrif:")
	}
	uCons := consump / float64(no_units)
	log.Println("ucons:", uCons, " comsump", consump, " nounits:", no_units)
	bands := tariff.Bands
	if bands == nil || len(bands) == 0 {
		return nil, errors.New("missing tarrif bands")
	}
	bCons := make(map[*TariffBand]float64)
	reminder := uCons
	for bx := range bands {
		//log.Println(reminder)
		band := bands[bx]
		log.Println("Bands Now .. ", *band.From, " - ", *band.To, " - ", *band.Charge)
		if uCons < *band.From || reminder <= 0 {
			bCons[band] = 0
			continue
		}
		if uCons >= *band.To {
			cc := *band.To - *band.From
			bCons[band] = cc
			reminder = reminder - cc
			log.Println("Reminder : ", reminder)
			continue
		}
		log.Println("Last Reminder : ", reminder)
		cc := reminder
		bCons[band] = cc
		reminder = 0
	}
	//log.Println("==================")
	amt = 0
	for band, cons := range bCons {
		cot := tools.DefaultF(band.Constant, float64(0))
		chrg := tools.DefaultF(band.Charge, float64(0))
		if cons > 0 {
			charge := cot + chrg*cons
			amt = amt + charge
		}
		//log.Println(amt)
	}
	amt = float64(no_units) * amt
	return &amt, nil
}
