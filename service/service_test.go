package service

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	"fmt"
	"testing"
	. "math"
)

const errMargin float64 = 0.000000001

func getTariff() *Tariff {
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
	return &tar
}
func TestCalc(t *testing.T) {
	srv := BillingService{}
	var no_units int64 = 1
	var consumps map[float64]float64 = map[float64]float64{}
	consumps[5] = 3.25
	consumps[10] = 6.5
	consumps[15] = 14.5
	consumps[20] = 22.5
	consumps[25] = 33.75
	consumps[30] = 45
	consumps[35] = 96.25
	consumps[40] = 110
	consumps[45] = 141.75
	consumps[50] = 157.5
	consumps[55] = 173.25
	consumps[100] = 315
	tar := getTariff()
	for consump, value := range consumps {
		amt, err := srv.Calc(no_units, consump, tar)
		if err != nil {
			t.Error(err)
		}
		if amt == nil {
			t.Error("invalied amount")
		}
		if Abs(*amt - value) > errMargin {
			t.Error(fmt.Sprintf("expectd %f while found %f", *amt, value))
		}
	}

}

func TestCalc_no_units_3(t *testing.T) {
	srv := BillingService{}
	var no_units int64 = 3
	var consumps map[float64]float64 = map[float64]float64{}
	consumps[20] = 13
	consumps[40] = 35.5
	consumps[60] = 67.5
	consumps[80] = 112.5
	consumps[100] = 275
	consumps[200] = 630
	tar := getTariff()
	for consump, value := range consumps {
		amt, err := srv.Calc(no_units, consump, tar)
		if err != nil {
			t.Error(err)
		}
		if amt == nil {
			t.Error("invalied amount")
		}
		if Abs(*amt - value) > errMargin {
			t.Error(*amt - value)
			t.Error(fmt.Sprintf("expectd %f while found %f", *amt, value))
		}
	}

}

func toFloat(fl float64) *float64 {
	return &fl
}
