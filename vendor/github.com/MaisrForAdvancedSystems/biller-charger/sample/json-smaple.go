package sample

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	. "github.com/MaisrForAdvancedSystems/biller-charger/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type JsonCase struct {
	Customer           *Customer
	Readings           []*ServiceReading
	TotalExpectedValue float64
}

type JsonFile struct {
	Ctgs           []*Ctg
	Tariffs        []*Tariff
	RegularCharges []*RegularCharge
	Cases          []*JsonCase
	Setting        *ChargeSetting
}

type JsonTestService struct {
	file *JsonFile
}

func (s *JsonTestService) Init(fileName string) {
	data, err := ioutil.ReadFile("test_pattern.json")
	if err != nil {
		panic(err)
	}
	var jf JsonFile
	err = json.Unmarshal(data, &jf)
	if err != nil {
		panic(err)
	}
	s.file = &jf
}
func (s *JsonTestService) GetSettings() *ChargeSetting {
	if s.file == nil {
		return nil
	}
	return s.file.Setting
}
func (s *JsonTestService) GetCases() []*JsonCase {
	return s.file.Cases
}
func (s *JsonTestService) Info(cn context.Context, empty *Empty) (*ServiceInfo, error) {
	return &ServiceInfo{
		Name:    ToStringPointer("JsonTestService"),
		Version: ToStringPointer("v1.0.0"),
	}, nil
}
func (s *JsonTestService) GetCtgs(cn context.Context, empty *Empty) (*CtgsResponce, error) {
	return &CtgsResponce{
		Ctgs:                 s.file.Ctgs,
	}, nil
}
func (s *JsonTestService) GetSetupData(cn context.Context, empty *Empty) (*SetupData, error) {
	if s.file == nil {
		return nil, errors.New("empty json file")
	}
	log.Println("file tarrifes")
	log.Println(s.file.Tariffs)
	return &SetupData{
		Ctgs:           s.file.Ctgs,
		Tariffs:        s.file.Tariffs,
		RegularCharges: s.file.RegularCharges,
	}, nil
}
func (s *JsonTestService) GetCustomerByCustkey(cn context.Context, key *Key) (*Customer, error) {
	if s.file == nil {
		return nil, errors.New("empty json file")
	}
	if s.file.Cases == nil {
		return nil, errors.New("empty cases json file")
	}
	if len(s.file.Cases) == 0 {
		return nil, errors.New("empty  cases json file")
	}
	cst := s.file.Cases[0].Customer
	return cst, nil
}
func (s *JsonTestService) GetCustomersByBillgroup(cn context.Context, key *Key) (*CustomersList, error) {
	if s.file == nil {
		return nil, errors.New("empty json file")
	}
	if s.file.Cases == nil {
		return nil, errors.New("empty cases json file")
	}
	csts := make([]*Customer, 0)
	for id := range s.file.Cases {
		csts = append(csts, s.file.Cases[id].Customer)
	}
	return &CustomersList{
		Customers: csts,
	}, nil
}
func (s *JsonTestService) WriteFinantialData(cn context.Context, msg *PostMessage) (*Empty, error) {
	if msg == nil {
		log.Println("data is null nor responce data")
		return nil, nil
	}
	data := msg.Data
	if data == nil {
		log.Println("data is null nor responce data")
		return nil, nil
	}
	if data.Bills == nil || len(data.Bills) == 0 {
		log.Println("data is null nor responce data")
		return nil, nil
	}
	if data.Bills[0].FTransactions == nil {
		log.Println("data FTransactions is null")
		return nil, nil
	}
	if len(data.Bills[0].FTransactions) == 0 {
		log.Println("data length is zero")
		return nil, nil
	}
	for _, t := range data.Bills[0].FTransactions {
		if t == nil {
			log.Println("error:transction is null ")
			continue
		}
		stm := fmt.Sprintf("ctype:%v service:%v amount:%v taxAmount:%v discountAmount:%v", t.GetCtype(), t.GetServiceType(), t.GetAmount(), t.GetTaxAmount(), t.GetDiscountAmount())
		log.Println(stm)
	}
	return &Empty{}, nil
}
