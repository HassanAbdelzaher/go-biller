package engine

import (
	"context"
	"errors"
	"log"
	"sort"

	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

// Info(context.Context, *Empty) (*ServiceInfo, error)
// 	Calulate(context.Context, *CalculationRequest) (*BillResponce, error)
// 	Confirm(context.Context, *BillResponce) (*Empty, error)
// 	GetCustomerByCustkey(context.Context, *Key) (*Customer, error)
// 	GetLoockup(context.Context, *Entity) (*LookUpsResponce, error)

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

// GetBillsByCustkey Imp
func (e *Engine) GetBillsByCustkey(ctx context.Context, rq *billing.GetBillRequest) (*billing.BillResponce, error) {
	response, err := e.DataProvider.GetBillsByCustkey(ctx, rq)
	if err != nil {
		return nil, err
	}
	if response != nil {
		if len(response.Bills) > 0 {
			bills := []*billing.Bill{}
			for idx := range response.Bills {
				bill := response.Bills[idx]
				if bill != nil && bill.BilngDate != nil {
					bills = append(bills, bill)
				}
			}
			sort.SliceStable(bills, func(i, j int) bool {
				return (*bills[j].BilngDate).AsTime().Before((*bills[i].BilngDate).AsTime())
			})
			response.Bills = bills
		}
	}
	return response, nil
}

func (e *Engine) GetBillsByFormNo(ctx context.Context, rq *billing.GetBillRequest) (*billing.BillResponce, error) {
	return e.DataProvider.GetBillsByFormNo(ctx, rq)
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
	log.Println("calculate")
	log.Println("Len,", len(rq.OldFTransactions))
	for idx := range rq.OldFTransactions {
		oldTrans := rq.OldFTransactions[idx]
		log.Println("Code", *oldTrans.Code, "Amount", *oldTrans.Amount)
	}
	cst := rq.Customer
	if cst == nil {
		return nil, errors.New("Customer not found:")
	}

	bs, err := e.ChargeService.Charge(ctx, &billing.ChargeRequest{
		Customer:         cst,
		ServicesReadings: rq.ServicesReadings,
		Setting:          rq.Setting,
		Services:         rq.Services,
		OldFTransactions: rq.OldFTransactions,
	})
	if err != nil {
		return nil, err
	}
	return bs, nil
}
func (e *Engine) Setup() error {
	if e.TariffProvider == nil || e.ChargeService == nil {
		return errors.New("Engine not created properly")
	}
	setups, err := e.TariffProvider.GetSetupData(context.Background(), empty)
	if err != nil {
		return err
	}
	_, err = e.ChargeService.Setup(context.Background(), &billing.SetupData{
		Tariffs:        setups.GetTariffs(),
		Ctgs:           setups.GetCtgs(),
		RegularCharges: setups.GetRegularCharges(),
		TransCodes:     setups.TransCodes,
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
	log.Println("Engine:Confirm Bills Count", len(reps.Bills))
	log.Println("Engine:Confirm First Bill")
	rep := billing.BillResponce{Bills: []*billing.Bill{}}
	for idx := range reps.Bills {
		bill := *reps.Bills[idx]
		log.Println("Engine:Confirm Bill ", *bill.PaymentNo, *bill.Customer.CustType)
		rep.Bills = append(rep.Bills, &bill)
	}
	_, err := e.DataConsumer.WriteFinantialData(cont, &rep)
	if err != nil {
		log.Println("Engine:Confirm Error ", err.Error())
	}
	return &billing.Empty{}, err
}
