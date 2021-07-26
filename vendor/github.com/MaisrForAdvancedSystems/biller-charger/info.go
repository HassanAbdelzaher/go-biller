package charge_service

import (
	"context"

	"github.com/MaisrForAdvancedSystems/biller-charger/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

var ChargeServiceVersion = "v1.83.0"

func (s *BillingChargeService) Info(cn context.Context, empty *Empty) (*ServiceInfo, error) {
	return &ServiceInfo{
		Name:    tools.ToStringPointer("DefaultChargeService"),
		Version: tools.ToStringPointer(ChargeServiceVersion),
	}, nil
}
