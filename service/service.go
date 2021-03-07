package service

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/charge/tariff_calc"
	"context"
	errors "errors"
	"fmt"
	"log"
	"strings"
)
var empty Empty=Empty{}
type BillingService struct {
	Tariffs map[string]*Tariff
	Ctgs    map[string]*Ctg
	RegularCharges []*RegularCharge
	IsTrace bool
}
func (s *BillingService) Trace(v ...interface{}){
	if s.IsTrace{
		log.Println(v...)
	}
}

func (s *BillingService) TraceF(str string,v ...interface{}){
	if s.IsTrace{
		log.Printf(str,v...)
	}
}
//setup for the engine with all settings
func (s *BillingService) Setup(c context.Context, r *SetupRequest)(*Empty, error){
	_,err:=s.SetCtg(c,r.Ctgs)
	if err!=nil{
		return nil,err
	}
	_,err=s.SetTariff(c,r.Tariffs)
	if err!=nil{
		return nil,err
	}
	s.RegularCharges=r.RegularCharge
	return &empty,nil
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
	isCustValied,err:=s.ValidateCustomer(c,r.Customer)
	if err!=nil{
		return nil,err
	}
	if isCustValied==nil || !*isCustValied{
		return nil, errors.New("customer data is not valied ")
	}
	if r.Customer.Property == nil {
		return nil, nil
	}
	services:=r.Customer.Property.Services
	if services==nil || len(services)==0{
		return nil,nil
	}
	rdgs:=make(map[*Service]*ServiceReading)
	if r.ServicesReadings != nil && len(r.ServicesReadings) >0 {
		for idx:=range services{
			srv:=services[idx]
			for sx:=range r.ServicesReadings{
				if *r.ServicesReadings[sx].ServiceType==*srv.ServiceType{
					rdgs[srv]=r.ServicesReadings[sx]
					break
				}
			}
		}
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
	return nil, nil
}



// SetCtg setup the ctg for the engin
func (s *BillingService) SetCtg(c context.Context, Ctgs []*Ctg) (*Empty, error) {
	if Ctgs == nil ||len(Ctgs)==0 {
		return nil, errors.New("invalied request")
	}
	s.Ctgs = make(map[string]*Ctg)
	for idx := range Ctgs {
		if Ctgs[idx].CType == nil {
			return nil, errors.New("invalied ctg data")
		}
		if Ctgs[idx].Tariffs == nil || len(Ctgs[idx].Tariffs) == 0 {
			return nil, errors.New("invalied ctg tarrif data")
		}
		id := strings.TrimSpace(*Ctgs[idx].CType)
		s.Ctgs[id] = Ctgs[idx]
	}
	return nil, nil
}

// SetTariff setup the tariffes for the engin
func (s *BillingService) SetTariff(c context.Context,tariffs []*Tariff) (*Empty, error) {
	if tariffs == nil || len(tariffs) == 0 {
		return nil, errors.New("invalied request")
	}
	s.Tariffs = make(map[string]*Tariff)
	for idx := range tariffs {
		if tariffs[idx].TariffId == nil {
			return nil, errors.New("invalied tariff data")
		}
		if tariffs[idx].Bands == nil || len(tariffs[idx].Bands) == 0 {
			return nil, errors.New("invalied  tarrif bands")
		}
		id := strings.TrimSpace(*tariffs[idx].TariffId)
		s.Tariffs[id] = tariffs[idx]
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
//calc for service and may be not used
func (s *BillingService) CalcForService(setting *ChargeSetting,service *Service,reading *ServiceReading) (*float64, error) {
	var conn=service.Connection
	if conn==nil{
		return nil,nil
	}
	mainCtype:=conn.CType;
	if conn.SubConnections==nil || len (conn.SubConnections)==0{
		return s.CalcForConnection(setting,service.ServiceType,mainCtype,conn.NoUnits,conn.EstimCons,reading);
	}
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



