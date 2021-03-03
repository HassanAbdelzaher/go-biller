package charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	errors "errors"
)
//Calc portable function for claulate charge for certain c_type
func Calc(no_units int64, consump float64, tariff *Tariff) (*float64, error) {
	var amt float64 = 0
	uCons := consump / float64(no_units)
	bands := tariff.Bands
	if bands == nil || len(bands) == 0 {
		return nil, errors.New("missing tarrif bands")
	}
	bCons := make(map[*TariffBand]float64)
	reminder := uCons
	for bx := range bands {
		//log.Println(reminder)
		band := bands[bx]
		if uCons < *band.From || reminder <= 0 {
			bCons[band] = 0
			continue
		}
		if uCons >= *band.To {
			cc := *band.To - *band.From
			bCons[band] = cc
			reminder = reminder - cc
			continue
		}
		cc := reminder
		bCons[band] = cc
		reminder = 0
	}
	//log.Println("==================")
	amt = 0
	for band, cons := range bCons {
		if cons > 0 {
			charge := *band.Constant + *band.Factor*cons
			amt = amt + charge
		}
		//log.Println(amt)
	}
	amt = float64(no_units) * amt
	return &amt, nil
}