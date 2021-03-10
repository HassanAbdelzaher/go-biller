package service

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"context"
	errors "errors"
	"log"
	"strings"
	"time"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

func (s *BillingService) Info(cn context.Context, empty *Empty) (*ServiceInfo, error) {
	return &ServiceInfo{
		Name:    tools.ToStringPointer("ChargeService"),
		Version: tools.ToStringPointer(ChargeServiceVersion),
	}, nil
}
func (s *BillingService) Trace(v ...interface{}) {
	if s.IsTrace {
		log.Println(v...)
	}
}
func (s *BillingService) TraceF(str string, v ...interface{}) {
	if s.IsTrace {
		log.Printf(str, v...)
	}
}

//setup for the engine with all settings
func (s *BillingService) Setup(c context.Context, r *SetupRequest) (*Empty, error) {
	s.TraceF("writing setup data")
	_, err := s.SetCtg(c, r.Ctgs)
	if err != nil {
		return nil, err
	}
	_, err = s.SetTariff(c, r.Tariffs)
	if err != nil {
		return nil, err
	}
	s.RegularCharges = r.RegularCharge
	return &empty, nil
}

// Charge calculate charge for all services

// SetCtg setup the ctg for the engin
func (s *BillingService) SetCtg(c context.Context, Ctgs []*Ctg) (*Empty, error) {
	if Ctgs == nil || len(Ctgs) == 0 {
		return nil, errors.New("invalied ctg request")
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
func (s *BillingService) SetTariff(c context.Context, tariffs []*Tariff) (*Empty, error) {
	if tariffs == nil || len(tariffs) == 0 {
		return nil, errors.New("invalied tarif request")
	}
	s.Tariffs = make(map[string]map[time.Time]*Tariff)
	for idx := range tariffs {
		inTarf := tariffs[idx]
		if inTarf.TariffId == nil {
			return nil, errors.New("invalied tariff id")
		}
		s.Trace(*inTarf.TariffId)
		if inTarf.Bands == nil || len(tariffs[idx].Bands) == 0 {
			return nil, errors.New("invalied  tarrif bands")
		}
		if inTarf.EffectDate == nil {
			return nil, errors.New("missing tarrif effect date")
		}
		tarifId := strings.TrimSpace(*inTarf.TariffId)
		//check tarif
		tar, ok := s.Tariffs[tarifId]
		if !ok || tar == nil {
			tar = make(map[time.Time]*Tariff)
			s.Tariffs[tarifId] = tar
		}
		effDate := inTarf.GetEffectDate().AsTime()
		_, ok = tar[effDate]
		if !ok {
			tar[effDate] = inTarf
		}
	}
	return nil, nil
}
