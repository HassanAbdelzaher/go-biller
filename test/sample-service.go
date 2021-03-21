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

func (s *JsonTest) GetLoockup(cn context.Context, en *pr.Entity) (*pr.LookUpsResponce, error) {
	return nil, errors.New("json service : dos't support lookups")
}
