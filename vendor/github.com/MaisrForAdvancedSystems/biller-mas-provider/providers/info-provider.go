package providers

import (
	"context"

	. "github.com/MaisrForAdvancedSystems/biller-mas-provider/tools"
	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type InfoProvider struct {
}

const version = "v1.129.0"

func (s *InfoProvider) Info(cn context.Context, empty *Empty) (*ServiceInfo, error) {
	return &ServiceInfo{
		Name:    ToStringPointer("MasDataService"),
		Version: ToStringPointer(version),
	}, nil
}
