package main

import (
	engine2 "MaisrForAdvancedSystems/go-biller/engine"
	"MaisrForAdvancedSystems/go-biller/service"
	"context"
	"log"

	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

var empty *billing.Empty = &billing.Empty{}

func main() {
	TestJsonService()
}

func TestJsonService() {

	srv := &JsonTestService{}
	srv.Init()
	chargeService := &service.BillingService{IsTrace: true}
	engine, err := engine2.NewEngine(chargeService, srv, srv, srv)
	if err != nil {
		log.Println(err)
		return
	}
	for id := range srv.file.Cases {
		cas := srv.file.Cases[id]
		trans, err := engine.HandleRequest(context.Background(), *cas.Customer.Custkey, srv.file.Setting, cas.Readings)
		if err != nil {
			log.Println(err)
			return
		}
		if trans == nil {
			log.Println("invalied responce")
			return
		}
		if trans.FTransactions == nil {
			log.Println("invalied responce FTransactions")
			return
		}
		var totalAmount float64 = 0
		for _, r := range trans.FTransactions {
			totalAmount = totalAmount + r.GetAmount()
		}
		if totalAmount != cas.TotalExpectedValue {
			for _, r := range trans.FTransactions {
				log.Printf("code:%v service:%v ctype:%v amount:%v no_units:%v", r.GetCode(), r.GetServiceType(), r.GetCtype(), r.GetAmount(), r.GetNoUnits())
			}
			log.Printf("expected value %v while found %v", cas.TotalExpectedValue, totalAmount)
		}
	}
}
