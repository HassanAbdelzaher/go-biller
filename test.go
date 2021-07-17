package main

import (
	"MaisrForAdvancedSystems/go-biller/engine"
	sv_tst "MaisrForAdvancedSystems/go-biller/test"
	"context"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
	chrg "github.com/MaisrForAdvancedSystems/biller-charger"
	prov "github.com/MaisrForAdvancedSystems/biller-mas-provider"
)

func runTest() {
	charger := &chrg.BillingChargeService{IsTrace: true}
	masProvider := &prov.MasProvider{}
	sample := &sv_tst.JsonTest{}
	sample.Init("test_pattern.json")
	eng, err := engine.NewEngine(masProvider, charger, masProvider, sample)
	if err != nil {
		log.Println(err)
		return
	}
	custkey := "100000992"
	var cycleLength int64 = 1
	bilngDate := time.Date(2021, 2, 28, 0, 0, 0, 0, time.UTC)
	stmapBilngDate := timestamppb.New(bilngDate)
	setting := billing.ChargeSetting{
		CycleLength:      &cycleLength,
		BilingDate:       stmapBilngDate,
		IgnoreTimeEffect: new(bool),
	}
	water := billing.SERVICE_TYPE_WATER
	sewer := billing.SERVICE_TYPE_SEWER
	var consump float64 = 155
	var zero float64 = 0
	readings := []*billing.ServiceReading{{
		ServiceType: &water,
		Reading: &billing.Reading{
			Consump:   &consump,
			PrReading: &zero,
			CrReading: &consump,
			PrDate:    timestamppb.New(time.Now().AddDate(-1, 0, 0)),
			CrDate:    timestamppb.Now(),
		},
	},
		{
			ServiceType: &sewer,
			Reading: &billing.Reading{
				Consump:   &consump,
				PrReading: &zero,
				CrReading: &consump,
				PrDate:    timestamppb.New(time.Now().AddDate(-1, 0, 0)),
				CrDate:    timestamppb.Now(),
			},
		}}
	rs, err := eng.HandleRequest(context.Background(), custkey, &setting, readings)
	if err != nil {
		log.Println(err)
		return
	}
	if rs == nil {
		log.Println("invalied responce")
	}
	if rs.Bills == nil {
		log.Println("invalied responce:null Bill")
	}
	if rs.Bills[0].FTransactions == nil {
		log.Println("invalied responce transcation")
	}
	if len(rs.Bills[0].FTransactions) == 0 {
		log.Println("invalied responce empty transcation")
	}
	log.Println("succssed")
	for _, t := range rs.Bills[0].FTransactions {
		log.Println(t.GetCode(), t.GetAmount())
	}
}

func TestJsonService() {
	errF := log.Println
	logF := log.Println
	errff := log.Printf
	logff := log.Printf
	sv_tst.TestService(logF, errF, errff, logff, "test_patter.json")
}
