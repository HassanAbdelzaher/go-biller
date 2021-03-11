package main

import (
	engine2 "MaisrForAdvancedSystems/go-biller/engine"
	"context"
	"testing"

	ch "github.com/MaisrForAdvancedSystems/biller-charger"
	"github.com/MaisrForAdvancedSystems/biller-charger/sample"
)

func TestService(t *testing.T) {
	sampleSrv := &sample.JsonTestService{}
	sampleSrv.Init("test_patter.json")
	chaSrv := &ch.BillingChargeService{IsTrace: true}
	engine, err := engine2.NewEngine(chaSrv, sampleSrv, sampleSrv, sampleSrv)
	if err != nil {
		t.Error(err)
		return
	}
	casses := sampleSrv.GetCases()
	for id := range casses {
		cas := casses[id]
		trans, err := engine.HandleRequest(context.Background(), *cas.Customer.Custkey, sampleSrv.GetSettings(), cas.Readings)
		if err != nil {
			t.Error(err)
			return
		}
		if trans == nil {
			t.Error("invalied responce")
			return
		}
		if trans.FTransactions == nil {
			t.Error("invalied responce FTransactions")
			return
		}
		var totalAmount float64 = 0
		for _, r := range trans.FTransactions {
			totalAmount = totalAmount + r.GetAmount()
		}
		if totalAmount != cas.TotalExpectedValue {
			for _, r := range trans.FTransactions {
				t.Logf("code:%v service:%v ctype:%v amount:%v no_units:%v", r.GetCode(), r.GetServiceType(), r.GetCtype(), r.GetAmount(), r.GetNoUnits())
			}
			t.Errorf("expected value %v while found %v", cas.TotalExpectedValue, totalAmount)
		}
	}
}
