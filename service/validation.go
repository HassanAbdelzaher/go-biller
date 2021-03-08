package service

import (
	"context"
	"errors"
	. "MaisrForAdvancedSystems/go-biller/proto"
	"strings"
)

/////////////////////vlidations/////////////////
// ValidateCustomer validate
// all ctypes inclides in ctgs in the engin
// all ctype customer service tarrifs founded in engin tarrifes
func (s *BillingService) ValidateCustomer(con context.Context, cust *Customer) (*bool, error) {
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if cust == nil {
		return nil, errors.New("invalied customer data")
	}
	if cust.Property == nil {
		return nil, errors.New("missing customer properties")
	}
	p := cust.Property
	if p == nil {
		return nil, errors.New("invalied customer properties")
	}
	if p.Services == nil || len(p.Services) == 0 {
		return nil, errors.New("missing propertie services")
	}
	for sdx := range p.Services {
		srv := p.Services[sdx]
		if srv == nil || srv.Connection == nil {
			continue
		}
		if srv.Connection.CType == nil {
			return nil, errors.New("missing ctype for connection ")
		}
		ok, err := s.IsCtgFound(srv.Connection.CType, srv)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("missing tarrif for " + *srv.Connection.CType)
		}
		if srv.Connection.SubConnections != nil {
			for ux := range srv.Connection.SubConnections {
				sb := srv.Connection.SubConnections[ux]
				if sb == nil {
					return nil, errors.New("missing sub connection data " + *srv.Connection.CType)
				}
				if sb.CType == nil {
					return nil, errors.New("missing ctype for sub connection ")
				}
				ok, err := s.IsCtgFound(sb.CType, srv)
				if err != nil {
					return nil, err
				}
				if !ok {
					return nil, errors.New("missing tarrif for " + *srv.Connection.CType)
				}
			}
		}
	}
	ok:=true
	return &ok, nil
}


// IsCtgFound check if the ctype is included into the ctgs
// also check if the tarrif for the service is founded in engin tarrifs
func (s *BillingService) IsCtgFound(ctg *string, service *Service) (bool, error) {
	if s.Ctgs == nil || ctg == nil {
		return false, nil
	}
	_ctg := strings.TrimSpace(*ctg)
	cg, ok := s.Ctgs[_ctg]
	if ok {
		if cg.Tariffs != nil && service != nil {
			found := false
			for cgIdx := range cg.Tariffs {
				tar := cg.Tariffs[cgIdx]
				if *tar.ServiceType == *service.ServiceType {
					found = true
					if !s.IsTariffFound(tar.TarifId) {
						return false, errors.New("missing tarrif for " + *cg.CType + " " + service.ServiceType.String())
					}
				}
			}
			if !found {
				return false, errors.New("missing tarrif for " + *cg.CType)
			}
		}
	}
	return ok, nil
}

// IsTariffFound check if tarrif found in engin tarrifes map
func (s *BillingService) IsTariffFound(tar *string) bool {
	if s.Tariffs == nil || tar == nil {
		return false
	}
	_tar := strings.TrimSpace(*tar)
	_, ok := s.Tariffs[_tar]
	return ok
}
