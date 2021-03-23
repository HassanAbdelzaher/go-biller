package engine

import (
	"context"
	"errors"
	"log"

	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

/*
Info(context.Context, *Empty) (*ServiceInfo, error)
	Calulate(context.Context, *CalculationRequest) (*BillResponce, error)
	Confirm(context.Context, *BillResponce) (*Empty, error)
	GetCustomerByCustkey(context.Context, *Key) (*Customer, error)
	GetLoockup(context.Context, *Entity) (*LookUpsResponce, error)
*/
var VERSION string = "v1.0.1"
var SERVICE_NAME string = "v1.0.1"

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
func (e *Engine) GetBillByCustkey(ctx context.Context, rq *billing.GetBillRequest) (*billing.BillResponce, error) {
	return e.DataProvider.GetBillByCustkey(ctx, rq)
}
func (e *Engine) Info(ctx context.Context, rq *billing.Empty) (*billing.ServiceInfo, error) {
	return &billing.ServiceInfo{Version: &VERSION, Name: &SERVICE_NAME}, nil
}
func (e *Engine) GetCustomerByCustkey(ctx context.Context, rq *billing.Key) (*billing.Customer, error) {
	return e.DataProvider.GetCustomerByCustkey(ctx, rq)
}
func (e *Engine) GetLoockup(ctx context.Context, rq *billing.Entity) (*billing.LookUpsResponce, error) {
	return e.DataProvider.GetLoockup(ctx, rq)
}
func (e *Engine) Calulate(ctx context.Context, rq *billing.ChargeRequest) (*billing.BillResponce, error) {
	cst := rq.Customer
	if cst == nil {
		return nil, errors.New("Customer not found:")
	}
	return e.ChargeService.Charge(ctx, &billing.ChargeRequest{
		Customer:         cst,
		ServicesReadings: rq.ServicesReadings,
		Setting:          rq.Setting,
	})
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

func (e *Engine) Confirm(cont context.Context, reps *billing.BillResponce) (*billing.Empty, error) {
	_, err := e.DataConsumer.WriteFinantialData(cont, reps)
	return &billing.Empty{}, err
}
