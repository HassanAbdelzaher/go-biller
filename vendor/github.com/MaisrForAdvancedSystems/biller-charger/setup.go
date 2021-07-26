package charge_service

import (
	"context"
	errors "errors"
	"log"
	"strings"
	"time"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

func (s *BillingChargeService) trace(v ...interface{}) {
	if s.IsTrace {
		log.Println(v...)
	}
}
func (s *BillingChargeService) tsraceF(str string, v ...interface{}) {
	if s.IsTrace {
		log.Printf(str, v...)
	}
}

//setup for the engine with all settings
func (s *BillingChargeService) Setup(c context.Context, r *SetupData) (*Empty, error) {
	s.trace("writing setup data")
	_, err := s.setCtg(c, r.Ctgs)
	if err != nil {
		return nil, err
	}
	_, err = s.setTariff(c, r.Tariffs)
	if err != nil {
		return nil, err
	}
	s.RegularCharges = r.RegularCharges
	if r.TransCodes != nil {
		_, err = s.setTransCodes(c, r.TransCodes)
	}
	return &empty, err
}

func (s *BillingChargeService) GetSetupDate() struct {
	Tariffs        map[string]map[time.Time]*Tariff
	Ctgs           map[string]*Ctg
	RegularCharges []*RegularCharge
} {
	return struct {
		Tariffs        map[string]map[time.Time]*Tariff
		Ctgs           map[string]*Ctg
		RegularCharges []*RegularCharge
	}{
		Tariffs:        s.Tariffs,
		Ctgs:           s.Ctgs,
		RegularCharges: s.RegularCharges,
	}
}

// Charge calculate charge for all services

// SetCtg setup the ctg for the engin
func (s *BillingChargeService) setCtg(c context.Context, Ctgs []*Ctg) (*Empty, error) {
	if Ctgs == nil || len(Ctgs) == 0 {
		return nil, errors.New("invalied ctg request")
	}
	s.Ctgs = make(map[string]*Ctg)
	for idx := range Ctgs {
		ctype := Ctgs[idx]
		if ctype == nil {
			return nil, errors.New("invalied ctg data")
		}
		if ctype.CType == nil {
			return nil, errors.New("invalied ctg ctype code")
		}
		if ctype.Tariffs == nil || len(ctype.Tariffs) == 0 {
			log.Println("missing tarrif for ctype " + *ctype.CType)
			if ctype.Tariffs == nil {
				ctype.Tariffs = make([]*ServiceTariff, 0)
			}
			//return nil, errors.New("invalied ctg tarrif data " + *ctype.CType)
		}
		id := strings.TrimSpace(*ctype.CType)
		s.Ctgs[id] = ctype
	}
	return nil, nil
}

// SetTariff setup the tariffes for the engin
func (s *BillingChargeService) setTariff(c context.Context, tariffs []*Tariff) (*Empty, error) {
	if tariffs == nil || len(tariffs) == 0 {
		return nil, errors.New("invalied tarif request")
	}
	s.Tariffs = make(map[string]map[time.Time]*Tariff)
	for idx := range tariffs {
		inTarf := tariffs[idx]
		if inTarf.TariffCode == nil {
			return nil, errors.New("invalied tariff id")
		}
		s.trace(*inTarf.TariffCode)
		if inTarf.Bands == nil || len(tariffs[idx].Bands) == 0 {
			return nil, errors.New("invalied  tarrif bands :"+*inTarf.TariffCode)
		}
		if inTarf.EffectDate == nil {
			return nil, errors.New("missing tarrif effect date:"+*inTarf.TariffCode)
		}
		tarifId := strings.TrimSpace(*inTarf.TariffCode)
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

func (s *BillingChargeService) setTransCodes(c context.Context, codes []*TransCode) (*Empty, error) {
	s.TransCodes = make(map[string]*TransCode)
	for id := range codes {
		cd := codes[id]
		s.TransCodes[cd.GetCode()] = cd
	}
	return &Empty{}, nil
}
