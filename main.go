package main

import (
	"MaisrForAdvancedSystems/go-biller/engine"
	stest "MaisrForAdvancedSystems/go-biller/test"
	"context"
	"time"

	"log"

	chrg "github.com/MaisrForAdvancedSystems/biller-charger"
	sample "github.com/MaisrForAdvancedSystems/biller-charger/sample"
	prov "github.com/MaisrForAdvancedSystems/biller-mas-provider"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var empty *billing.Empty = &billing.Empty{}

func main() {
	charger := &chrg.BillingChargeService{IsTrace: true}
	masProvider := &prov.MasProvider{}
	sample := &sample.JsonTestService{}
	sample.Init("test_pattern.json")
	eng, err := engine.NewEngine(masProvider, charger, masProvider, sample)
	if err != nil {
		log.Println(err)
		return
	}
	custkey := "120252964"
	var cycleLength int64 = 1
	bilngDate := time.Date(2021, 1, 31, 0, 0, 0, 0, time.UTC)
	stmapBilngDate := timestamppb.New(bilngDate)
	setting := billing.ChargeSetting{
		CycleLength:      &cycleLength,
		BilingDate:       stmapBilngDate,
		IgnoreTimeEffect: new(bool),
	}
	water := billing.SERVICE_TYPE_WATER
	sewer := billing.SERVICE_TYPE_SEWER
	var consump float64 = 25
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
	if rs.FTransactions == nil {
		log.Println("invalied responce transcation")
	}
	if len(rs.FTransactions) == 0 {
		log.Println("invalied responce empty transcation")
	}
	log.Println("succssed")
	for _, t := range rs.FTransactions {
		log.Println(t.GetCode(), t.GetAmount())
	}
}

func TestJsonService() {
	errF := log.Println
	logF := log.Println
	errff := log.Printf
	logff := log.Printf
	stest.TestService(logF, errF, errff, logff, "test_patter.json")
}
