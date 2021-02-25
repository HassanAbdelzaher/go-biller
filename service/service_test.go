package service

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	"testing"
)

func TestCalc(t *testing.T) {
	srv := BillingService{}
	var no_units int64 = 1
	var consump float64 = 10
	tar := Tariff{}
	tar.Bands = make([]*TariffBand, 0)
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     toFloat(0),
		To:       toFloat(10),
		Factor:   toFloat(0.65),
		Constant: toFloat(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     toFloat(10),
		To:       toFloat(20),
		Factor:   toFloat(1.6),
		Constant: toFloat(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     toFloat(20),
		To:       toFloat(30),
		Factor:   toFloat(2.25),
		Constant: toFloat(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     toFloat(30),
		To:       toFloat(40),
		Factor:   toFloat(2.75),
		Constant: toFloat(37.5),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     toFloat(40),
		To:       toFloat(99999999),
		Factor:   toFloat(3.15),
		Constant: toFloat(16),
	})
	amt, err := srv.Calc(no_units, consump, &tar)
	t.Error(*amt, "    ", err)
}

func toFloat(fl float64) *float64 {
	return &fl
}
