package sv_tst

import (
	"context"
	"errors"

	"github.com/MaisrForAdvancedSystems/biller-charger/sample"
	pr "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type JsonTest struct {
	sample.JsonTestService
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
