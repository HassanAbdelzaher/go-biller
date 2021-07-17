package sv_tst

import (
	"context"
	"errors"

	"github.com/MaisrForAdvancedSystems/biller-charger/sample"
	pr "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type JsonTest struct {
	sample.JsonTestService
	pr.UnimplementedBillingDataCousumerServer
	pr.UnimplementedBillingDataProviderServer
	pr.UnimplementedBillingTariffProviderServer
	pr.UnimplementedBillingChargeServiceServer
	pr.UnimplementedEngineServer
}

func (e *JsonTest) Info(cn context.Context, empty *pr.Empty) (*pr.ServiceInfo, error) {
	return &pr.ServiceInfo{}, nil
}

func (e *JsonTest) GetSetupData(cn context.Context, empty *pr.Empty) (*pr.SetupData, error) {
	return &pr.SetupData{}, nil
}

func (e *JsonTest) GetCtgs(cn context.Context, eType *pr.Empty) (resp *pr.CtgsResponce, err error) {
	return &pr.CtgsResponce{}, nil
}

func (e *JsonTest) GetCustomerByCustkey(cn context.Context, key *pr.Key) (cst *pr.Customer, err error) {
	return &pr.Customer{}, nil
}

func (e *JsonTest) GetCustomersByBillgroup(cn context.Context, key *pr.Key) (*pr.CustomersList, error) {
	return &pr.CustomersList{}, nil
}

func (e *JsonTest) Login(ctx context.Context, rq *pr.LoginRequest) (*pr.LoginResponce, error) {
	return &pr.LoginResponce{}, nil
}

func (s *JsonTest) GetLoockup(cn context.Context, en *pr.Entity) (*pr.LookUpsResponce, error) {
	return nil, errors.New("json service : dos't support lookups")
}

func (s *JsonTest) GetBillsByCustkey(cn context.Context, en *pr.GetBillRequest) (*pr.BillResponce, error) {
	return nil, errors.New("json service : dos't support lookups")
}

func (s *JsonTest) GetBillsByFormNo(cn context.Context, en *pr.GetBillRequest) (*pr.BillResponce, error) {
	return nil, errors.New("json service : dos't support lookups")
}

func (s *JsonTest) WriteFinantialData(cn context.Context, data *pr.PostMessage) (*pr.Empty, error) {
	return s.JsonTestService.WriteFinantialData(cn, data)
}
