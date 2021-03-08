package main

import (
	engine2 "MaisrForAdvancedSystems/go-biller/engine"
	"MaisrForAdvancedSystems/go-biller/service"
	"context"
	"testing"
)

func TestService(t *testing.T){
	srv:=&JsonTestService{}
	srv.Init()
	chargeService:=&service.BillingService{IsTrace:true}
	engine,err:=engine2.NewEngine(chargeService,srv,srv,srv)
	if err!=nil{
		t.Error(err)
		return
	}
	for id:=range srv.file.Cases{
		cas:=srv.file.Cases[id]
		trans,err:=engine.HandleRequest(context.Background(),*cas.Customer.Custkey,srv.file.Setting,cas.Readings)
		if err!=nil{
			t.Error(err)
			return
		}
		if trans== nil{
			t.Error("invalied responce")
			return
		}
		if trans.FTransactions== nil{
			t.Error("invalied responce FTransactions")
			return
		}
		var totalAmount float64=0
		for _,r:=range trans.FTransactions{
			totalAmount=totalAmount+r.GetAmount()
		}
		if totalAmount!=cas.TotalExpectedValue{
			for _,r:=range trans.FTransactions{
				t.Logf("code:%v service:%v ctype:%v amount:%v no_units:%v",r.GetCode(),r.GetServiceType(),r.GetCtype(),r.GetAmount(),r.GetNoUnits())
			}
			t.Errorf("expected value %v while found %v",cas.TotalExpectedValue,totalAmount)
		}
	}
}

