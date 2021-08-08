package engine

import (
	"MaisrForAdvancedSystems/go-biller/middlewares"
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"runtime/debug"

	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbpool"
	irespo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/interfaces"
	respo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/repositories"
)

// Info(context.Context, *Empty) (*ServiceInfo, error)
// 	Calulate(context.Context, *CalculationRequest) (*BillResponce, error)
// 	Confirm(context.Context, *BillResponce) (*Empty, error)
// 	GetCustomerByCustkey(context.Context, *Key) (*Customer, error)
// 	GetLoockup(context.Context, *Entity) (*LookUpsResponce, error)

var VERSION string = "v2.4.0"
var SERVICE_NAME string = "go_biller"

var empty = &billing.Empty{}

type Engine struct {
	ChargeService  billing.BillingChargeServiceServer
	DataProvider   billing.BillingDataProviderServer
	DataConsumer   billing.BillingDataCousumerServer
	TariffProvider billing.BillingTariffProviderServer
	billing.UnimplementedEngineServer
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
func (e *Engine) GetBillsByCustkey(ctx context.Context, rq *billing.GetBillRequest) (resp *billing.BillResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", string(debug.Stack())))
		}
	}()
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
func (e *Engine) GetBillsByFormNo(ctx context.Context, rq *billing.GetBillRequest) (resp *billing.BillResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", string(debug.Stack())))
		}
	}()
	if rq.FormNo == nil {
		return nil, errors.New("لا يوجد طلب")
	}
	formno, err := strconv.ParseInt(*rq.FormNo, 10, 64)
	if err != nil {
		return nil, errors.New("قم بادخال رقم للطلب")
	}
	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	} else {
		conn.Debug = true
	}
	var cancelbill irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	req, err := cancelbill.GetByFormNo(formno)
	if err != nil {
		return nil, err
	}
	if len(req) == 0 {
		return nil, errors.New("لا يوجد طلب بهذا الرقم")
	}
	if req[0].CLOSED != nil && *req[0].CLOSED {
		return nil, errors.New("الطلب تم اغلاقه")
	}
	var lucancelledBillsState irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	if req[0].STATE != nil {
		reqState, err := lucancelledBillsState.GetByID(*req[0].STATE)
		if err != nil {
			return nil, err
		}
		if len(reqState) > 0 {
			if reqState[0].RECAL_READY != nil && !*reqState[0].RECAL_READY {
				return nil, errors.New("الطلب لابد ان يكون في مرحله للاحتساب")
			}
		} else {
			return nil, errors.New("لابد ان يكون للطلب حاله")
		}
	} else {
		return nil, errors.New("لابد ان يكون للطلب حاله")
	}

	response, err := e.DataProvider.GetBillsByFormNo(ctx, rq)
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
func (e *Engine) Login(ctx context.Context, rq *billing.LoginRequest) (resp *billing.LoginResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at Login %v", string(debug.Stack())))
		}
	}()
	resp, err = e.DataProvider.Login(ctx, rq)
	if err != nil {
		return nil, err
	}
	if rq.Username == nil {
		return resp, err
	}
	tkn, err := middlewares.CreateToken(*rq.Username)
	if err != nil {
		return nil, err
	}
	resp.Token = &tkn
	return resp, err
}
func (e *Engine) Info(ctx context.Context, rq *billing.Empty) (resp *billing.ServiceInfo, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
	return &billing.ServiceInfo{Version: &VERSION, Name: &SERVICE_NAME}, nil
}
func (e *Engine) GetCustomerByCustkey(ctx context.Context, rq *billing.Key) (resp *billing.Customer, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
	return e.DataProvider.GetCustomerByCustkey(ctx, rq)
}
func (e *Engine) GetLoockup(ctx context.Context, rq *billing.Entity) (resp *billing.LookUpsResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
	return e.DataProvider.GetLoockup(ctx, rq)
}
func (e *Engine) GetCtgs(ctx context.Context, rq *billing.Empty) (resp *billing.CtgsResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
	return e.DataProvider.GetCtgs(ctx, rq)
}
func (e *Engine) Calulate(ctx context.Context, rq *billing.ChargeRequest) (resp *billing.BillResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at Calulate %v", string(debug.Stack())))
		}
	}()
	log.Println("calculate")
	cst := rq.Customer
	if cst == nil {
		return nil, errors.New("Customer not found:")
	}

	bs, err := e.ChargeService.Charge(ctx, &billing.ChargeRequest{
		Customer:         cst,
		ServicesReadings: rq.ServicesReadings,
		Setting:          rq.Setting,
		Services:         rq.Services,
		OldBill:          rq.OldBill,
	})
	if err != nil {
		return nil, err
	}
	return bs, nil
}
func (e *Engine) Setup() (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
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
func (e *Engine) HandleRequest(cont context.Context, key string, setting *billing.ChargeSetting, readings []*billing.ServiceReading) (resp *billing.BillResponce, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at GetBillsByCustkey %v", er))
		}
	}()
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
func (e *Engine) Post(cont context.Context, msg *billing.PostMessage) (resp *billing.Empty, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprintf("panic at Confirm %v", string(debug.Stack())))
		}
	}()
	_, err = e.DataConsumer.WriteFinantialData(cont, msg)
	if err != nil {
		log.Println("Engine:Confirm Error ", err.Error())
	}
	return &billing.Empty{}, err
}
