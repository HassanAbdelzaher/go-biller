package engine

import (
	"context"
	"errors"
	"log"

	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

var empty = &billing.Empty{}

type Engine struct {
	ChargeService  billing.BillingChargeServiceServer
	DataProvider   billing.BillingDataProviderServer
	DataConsumer   billing.BillingDataCousumerServer
	TariffProvider billing.BillingTariffProviderServer
}

func NewEngine(tariffProvider billing.BillingTariffProviderServer,
	chargeProvider billing.BillingChargeServiceServer,
	dataprovider billing.BillingDataProviderServer,
	dataConsumer billing.BillingDataCousumerServer) (*Engine, error) {
	eng := &Engine{
		ChargeService:  chargeProvider,
		DataProvider:   dataprovider,
		DataConsumer:   dataConsumer,
		TariffProvider: tariffProvider,
	}
	err := eng.Setup()
	if err != nil {
		return nil, err
	}
	return eng, nil
}
func (e *Engine) Setup() error {
	if e.TariffProvider == nil || e.ChargeService == nil {
		return errors.New("Engine not created properly")
	}
	setups, err := e.TariffProvider.GetSetupData(context.Background(), empty)
	if err != nil {
		return err
	}
	_, err = e.ChargeService.Setup(context.Background(), &billing.SetupRequest{
		Tariffs:       setups.GetTariffs(),
		Ctgs:          setups.GetCtgs(),
		RegularCharge: setups.GetRegularCharges(),
	})
	if err != nil {
		return err
	}
	return nil
}
func (e *Engine) HandleRequest(cont context.Context, key string, setting *billing.ChargeSetting, readings []*billing.ServiceReading) (*billing.BillResponce, error) {
	dataReq := billing.Key{
		Key:       &key,
		BilngDate: setting.BilingDate,
	}
	log.Println(dataReq.BilngDate.AsTime())
	cust, err := e.DataProvider.GetCustomerByCustkey(context.Background(), &dataReq)
	if err != nil {
		return nil, err
	}
	log.Println(cust)
	charges, err := e.ChargeService.Charge(context.Background(), &billing.ChargeRequest{
		Customer:         cust,
		ServicesReadings: readings,
		Setting:          setting,
	})
	if err != nil {
		return nil, err
	}
	return charges, nil
}

func (e *Engine) Confirm(cont context.Context, reps *billing.BillResponce) error {
	_, err := e.DataConsumer.WriteFinantialData(cont, reps)
	return err
}
