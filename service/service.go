package service

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/charge"
	"context"
	errors "errors"
	"strings"
)

type BillingService struct {
	Tariffs map[string]*Tariff
	Ctgs    map[string]*Ctg
}

// Charge calculate charge for all services
func (s *BillingService) Charge(c context.Context, r *BillRequest) (*BillResponce, error) {
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if s.Tariffs == nil {
		return nil, errors.New("Missing Traiff lookup")
	}
	if r == nil || r.Customer == nil {
		return nil, errors.New("invalied request")
	}
	if r.ServicesReadings == nil {
		return nil, errors.New("invalied request services")
	}
	if len(r.ServicesReadings) == 0 {
		return nil, errors.New("invalied request services")
	}

	return nil, nil
}

// SetCtg setup the ctg for the engin
func (s *BillingService) SetCtg(c context.Context, r *CtgRequest) (*Empty, error) {
	if r == nil || r.Ctgs == nil {
		return nil, errors.New("invalied request")
	}
	s.Ctgs = make(map[string]*Ctg)
	for idx := range r.Ctgs {
		if r.Ctgs[idx].CType == nil {
			return nil, errors.New("invalied ctg data")
		}
		if r.Ctgs[idx].Tariffs == nil || len(r.Ctgs[idx].Tariffs) == 0 {
			return nil, errors.New("invalied ctg tarrif data")
		}
		id := strings.TrimSpace(*r.Ctgs[idx].CType)
		s.Ctgs[id] = r.Ctgs[idx]
	}
	return nil, nil
}

// SetTariff setup the tariffes for the engin
func (s *BillingService) SetTariff(c context.Context, r *TariffRequest) (*Empty, error) {
	if r == nil || len(r.Tariffs) == 0 {
		return nil, errors.New("invalied request")
	}
	s.Tariffs = make(map[string]*Tariff)
	for idx := range r.Tariffs {
		if r.Tariffs[idx].TariffId == nil {
			return nil, errors.New("invalied tariff data")
		}
		if r.Tariffs[idx].Bands == nil || len(r.Tariffs[idx].Bands) == 0 {
			return nil, errors.New("invalied  tarrif bands")
		}
		id := strings.TrimSpace(*r.Tariffs[idx].TariffId)
		s.Tariffs[id] = r.Tariffs[idx]
	}
	return nil, nil
}

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
	if cust.Properties == nil || len(cust.Properties) == 0 {
		return nil, errors.New("missing customer properties")
	}
	for idx := range cust.Properties {
		p := cust.Properties[idx]
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
	}
	return nil, nil
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


func (s *BillingService) CalcForService(ctype string, no_units int64, consump float64, serv *SERVICE_TYPE) (*float64, error) {
	ctg, ok := s.Ctgs[ctype]
	if !ok || ctg.Tariffs == nil {
		return nil, errors.New("missing ctype in setup" + ctype)
	}
	var tariff *Tariff
	found := false
	for idx := range ctg.Tariffs {
		if *ctg.Tariffs[idx].ServiceType == *serv {
			found = true
			tarfid := ctg.Tariffs[idx].TarifId
			if !s.IsTariffFound(tarfid) {
				return nil, errors.New("missing tarrif" + ctype)
			}
			tariff = s.Tariffs[*tarfid]
			break
		}
	}
	if !found || tariff == nil {
		return nil, errors.New("missing tarrif" + ctype)
	}
	if no_units < 1 {
		no_units = 1
	}

	return Calc(no_units, consump, tariff)
}
