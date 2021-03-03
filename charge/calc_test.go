package charge

import (
	"fmt"
	. "math"
	"testing"
)

const errMargin float64 = 0.000000001

func TestCalc(t *testing.T) {
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
	tar := getTariffSample()
	for consump, value := range consumps {
		amt, err := Calc(no_units, consump, tar)
		if err != nil {
			t.Error(err)
		}
		if amt == nil {
			t.Error("invalied amount")
		}
		if Abs(*amt-value) > errMargin {
			t.Error(fmt.Sprintf("expectd %f while found %f", *amt, value))
		}
	}

}

func TestCalc_no_units_3(t *testing.T) {
	var no_units int64 = 3
	var consumps map[float64]float64 = map[float64]float64{}
	consumps[20] = 13
	consumps[40] = 35.5
	consumps[60] = 67.5
	consumps[80] = 112.5
	consumps[100] = 275
	consumps[200] = 630
	tar := getTariffSample()
	for consump, value := range consumps {
		amt, err := Calc(no_units, consump, tar)
		if err != nil {
			t.Error(err)
		}
		if amt == nil {
			t.Error("invalied amount")
		}
		if Abs(*amt-value) > errMargin {
			t.Error(*amt - value)
			t.Error(fmt.Sprintf("expectd %f while found %f", *amt, value))
		}
	}

}
