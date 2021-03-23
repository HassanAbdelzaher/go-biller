package sv_tst

import (
	engine2 "MaisrForAdvancedSystems/go-biller/engine"
	"context"

	ch "github.com/MaisrForAdvancedSystems/biller-charger"
)

type LogFun func(args ...interface{})
type LogfFun func(stm string, args ...interface{})

func TestService(lg LogFun, erF LogFun, erFF LogfFun, logFF LogfFun, fileName string) {
	sampleSrv := &JsonTest{}
	sampleSrv.Init(fileName)
	chaSrv := &ch.BillingChargeService{IsTrace: true}
	engine, err := engine2.NewEngine(sampleSrv, chaSrv, sampleSrv, sampleSrv)
	if err != nil {
		erF(err)
		return
	}
	casses := sampleSrv.GetCases()
	for id := range casses {
		cas := casses[id]
		trans, err := engine.HandleRequest(context.Background(), *cas.Customer.Custkey, sampleSrv.GetSettings(), cas.Readings)
		if err != nil {
			erF(err)
			return
		}
		if trans == nil {
			erF("invalied responce")
			return
		}
		if trans.Bill == nil {
			erF("invalied responce")
			return
		}

		if trans.Bill.FTransactions == nil {
			erF("invalied responce FTransactions")
			return
		}
		var totalAmount float64 = 0
		for _, r := range trans.Bill.FTransactions {
			totalAmount = totalAmount + r.GetAmount()
		}
		if totalAmount != cas.TotalExpectedValue {
			for _, r := range trans.Bill.FTransactions {
				logFF("code:%v service:%v ctype:%v amount:%v no_units:%v", r.GetCode(), r.GetServiceType(), r.GetCtype(), r.GetAmount(), r.GetNoUnits())
			}
			erFF("expected value %v while found %v", cas.TotalExpectedValue, totalAmount)
		}
	}
}
