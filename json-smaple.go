package main


import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

type JsonCase struct {
	Customer *Customer
	Readings []*ServiceReading
	TotalExpectedValue float64
}

type JsonFile struct {
	Ctgs []*Ctg
	Tariffs []*Tariff
	RegularCharges []*RegularCharge
	Cases []*JsonCase
	Setting *ChargeSetting
}

type JsonTestService struct {
	file *JsonFile
}
func (s *JsonTestService) Init(){
	data,err:=ioutil.ReadFile("test_pattern.json")
	if err!=nil{
		panic(err)
	}
	var jf JsonFile
	err=json.Unmarshal(data,&jf)
	if err!=nil{
		panic(err)
	}
	s.file=&jf
}
func (s *JsonTestService) Info(cn context.Context,empty *Empty) (*ServiceInfo,error){
	return &ServiceInfo{
		Name:                 ToStringPointer("JsonTestService"),
		Version:              ToStringPointer("v1.0.0"),
	},nil
}
func (s *JsonTestService) GetSetupData(cn context.Context,empty *Empty) (*ProviderSetupResponce,error){
	if s.file==nil{
		return nil,errors.New("empty json file")
	}
	log.Println("file tarrifes")
	log.Println(s.file.Tariffs)
	return &ProviderSetupResponce{
		Ctgs:s.file.Ctgs,
		Tariffs:s.file.Tariffs,
		RegularCharges:s.file.RegularCharges,
	},nil
}
func (s *JsonTestService) GetCustomerByCustkey(cn context.Context,key *Key) (*Customer,error){
	if s.file==nil{
		return nil,errors.New("empty json file")
	}
	if s.file.Cases==nil{
		return nil,errors.New("empty cases json file")
	}
	if len(s.file.Cases)==0{
		return nil,errors.New("empty  cases json file")
	}
	cst:=s.file.Cases[0].Customer
	return cst,nil
}
func (s *JsonTestService) GetCustomersByBillgroup(cn context.Context,key *Key) (*CustomersList,error){
	if s.file==nil{
		return nil,errors.New("empty json file")
	}
	if s.file.Cases==nil{
		return nil,errors.New("empty cases json file")
	}
	csts:=make([]*Customer,0)
	for id:=range s.file.Cases{
		csts=append(csts,s.file.Cases[id].Customer)
	}
	return &CustomersList{
		Customers:            csts,
	},nil
}
func (s *JsonTestService) WriteFinantialData(cn context.Context,data *BillResponce) (*Empty,error){
	if data==nil{
		log.Println("data is null nor responce data")
	}
	if data.FTransactions==nil{
		log.Println("data FTransactions is null")
	}
	if len(data.FTransactions)==0{
		log.Println("data length is zero")
	}
	for _,t:=range data.FTransactions{
		if t==nil{
			log.Println("error:transction is null ")
			continue
		}
		stm:=fmt.Sprintf("ctype:%v service:%v amount:%v taxAmount:%v discountAmount:%v",t.GetCtype(),t.GetServiceType(),t.GetAmount(),t.GetTaxAmount(),t.GetDiscountAmount())
		log.Println(stm)
	}
	return &Empty{},nil
}

