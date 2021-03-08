package samples

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"fmt"
	pt "github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type TestDummyService struct {

}
func (s *TestDummyService) Info(cn context.Context,empty *Empty) (*ServiceInfo,error){
	return &ServiceInfo{
		Name:                 ToStringPointer("Dummy"),
		Version:              ToStringPointer("v1.0.0"),
	},nil
}
func (s *TestDummyService) GetSetupData(cn context.Context,empty *Empty) (*ProviderSetupResponce,error){
	return &ProviderSetupResponce{
		Ctgs:GetCtgs(),
		Tariffs:GetTariffSample(),
		RegularCharges:GetChargeRegulars(),
	},nil
}
func (s *TestDummyService) GetCustomerByCustkey(cn context.Context,key *Key) (*Customer,error){
	cst:=GetNoramlCustomer(0,false,"00/01",10,1,MeterOperationStatus_WORKING);
	return cst,nil
}
func (s *TestDummyService) GetCustomersByBillgroup(cn context.Context,key *Key) (*CustomersList,error){
	cst:=GetNoramlCustomer(0,false,"00/01",10,1,MeterOperationStatus_WORKING);
	conns:=[]*SubConnection{
		&SubConnection{
			CType:                      ToStringPointer("00/01"),
			EstimateConsumptionPerUnit: ToFloatPointer(10),
			ConsumptionPercentage:      ToFloatPointer(10),
			NoUnits:                    ToIntPointer(1),
		},
		&SubConnection{
			CType:                      ToStringPointer("00/02"),
			EstimateConsumptionPerUnit: ToFloatPointer(20),
			ConsumptionPercentage:      ToFloatPointer(90),
			NoUnits:                    ToIntPointer(9),
		},
	}
	cust2:=GetMultiConnectionCustomer(1,true,"00/01",30,1,MeterOperationStatus_WORKING,conns)
	cust3:=GetMultiConnectionCustomer(1,true,"00/01",60,5,MeterOperationStatus_NOT_WORKING,conns)
	csts:=[]*Customer{cst,cust2,cust3}
	lst:=CustomersList{
		Customers:            csts,
	}
	return &lst,nil
}
func (s *TestDummyService) WriteFinantialData(cn context.Context,data *BillResponce) (*Empty,error){
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
func GetCtgs() []*Ctg{
	ctype:="00/01"
	ctype2:="00/02"
	water:=SERVICE_TYPE_WATER
	c1:=Ctg{
		CType:                ToStringPointer(ctype),
		CTypeGroupid:         ToStringPointer("00"),
		Tariffs:              []*ServiceTariff{&ServiceTariff{
			ServiceType:          &water,
			TarifId:              ToStringPointer("0-1"),
			Code:ToStringPointer("WATER_AMT"),
			TaxPercentage:ToFloatPointer(14),
			DiscountPercentage:ToFloatPointer(10),
		}},
		OP_ESTIM_CONS:        ToFloatPointer(float64(20)),
		NOOP_ESTIM_CONS:      ToFloatPointer(float64(40)),
	}
	c2:=Ctg{
		CType:                ToStringPointer(ctype2),
		CTypeGroupid:         ToStringPointer("00"),
		Tariffs:              []*ServiceTariff{&ServiceTariff{
			ServiceType:          &water,
			TarifId:              ToStringPointer("0-2"),
			Code:ToStringPointer("WATER_AMT"),
		}},
		OP_ESTIM_CONS:        ToFloatPointer(float64(90)),
		NOOP_ESTIM_CONS:      ToFloatPointer(float64(150)),
	}
	return []*Ctg{&c1,&c2}
}
func GetCtgMap() map[string]*Ctg{
	ctgs:=GetCtgs()
	mp:=make(map[string]*Ctg);
	for id:=range ctgs{
		mp[*ctgs[id].CType]=ctgs[id]
	}
	return mp
}
func GetTariffSample() []*Tariff {
	tar := Tariff{}
	tar.TariffId=ToStringPointer("0-1")
	tar.Bands = make([]*TariffBand, 0)
	effDate:=time.Now().AddDate(-10,0,0)
	tar.EffectDate=timestamppb.New(effDate)
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     ToFloatPointer(0),
		To:       ToFloatPointer(10),
		Factor:   ToFloatPointer(0.65),
		Constant: ToFloatPointer(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     ToFloatPointer(10),
		To:       ToFloatPointer(20),
		Factor:   ToFloatPointer(1.6),
		Constant: ToFloatPointer(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     ToFloatPointer(20),
		To:       ToFloatPointer(30),
		Factor:   ToFloatPointer(2.25),
		Constant: ToFloatPointer(0),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     ToFloatPointer(30),
		To:       ToFloatPointer(40),
		Factor:   ToFloatPointer(2.75),
		Constant: ToFloatPointer(37.5),
	})
	tar.Bands = append(tar.Bands, &TariffBand{
		From:     ToFloatPointer(40),
		To:       ToFloatPointer(99999999),
		Factor:   ToFloatPointer(3.15),
		Constant: ToFloatPointer(16),
	})
	return []*Tariff{&tar}
}
func GetChargeRegulars()[]*RegularCharge{
	return []*RegularCharge{GetCustTypeChargeRegularSample()}
}
func GetCustTypeChargeRegularSample() *RegularCharge{
	entityType:=ENTITY_TYPE_CUSTOMER_TYPE
	MappedValues:=make([]*EntityEnableMappedValue,0)
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("1"),
		Value:ToBoolPointer(false),
	})
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("2"),
		Value:ToBoolPointer(true),
	})
	rg:=GetChargeRegularSample(entityType,MappedValues)
	return rg
}
func GetChargeRegularSample(entityType ENTITY_TYPE,MappedValues []*EntityEnableMappedValue) *RegularCharge{
	var cp=RegularChargePeriod_BILL
	var effDate=time.Date(2000,1,1,0,0,0,0,time.Local)
	var tsmp,_=pt.TimestampProto(effDate)
	rg:=RegularCharge{}
	srvType:=SERVICE_TYPE_WATER
	rg.ServiceType=&srvType
	crel:=ChargeType_RELATION
	rg.ChargeType=&crel;
	rg.Code=ToStringPointer("BASIC_AMT");
	rg.Title=ToStringPointer("Water")
	rg.VatPercentage=ToFloatPointer(14)
	rg.ChargeCalcPeriod=&cp
	rg.EffectiveDate=tsmp
	rg.FixedCharge=ToFloatPointer(12)
	rg.IsChargable=ToBoolPointer(true)
	en:=RegularEnableEntity{}
	en.Code=ToStringPointer("98")
	en.MappedValues=make([]*EntityEnableMappedValue,0)
	en.EntityType=&entityType
	en.MappedValues=MappedValues
	rg.RelationEnableEntity=&en
	return &rg
}
func GetNoramlCustomer(typ int64,isVactated bool,ctype string,estimCons float64,noUnits int64,meteOpStatus MeterOperationStatus) *Customer{
	cust:=Customer{}
	cust.CustType=&typ
	cust.Custkey=ToStringPointer("1234")
	cust.InfoFlag1=ToStringPointer("1")
	prop:=Property{}
	cust.Property=&prop
	prop.InfoFlag1=ToStringPointer("1")
	prop.IsVacated=&isVactated
	prop.NoRooms=ToIntPointer(12);
	prop.PropRef=ToStringPointer("123444")
	srv:=Service{}
	prop.Services=[]*Service{&srv}
	water:=SERVICE_TYPE_WATER
	srv.ServiceType=&water
	conn:=Connection{}
	srv.Connection=&conn
	conn.CType=ToStringPointer(ctype)
	conn.ConnDiameter=ToIntPointer(10)
	cStatus:=CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER
	conn.ConnectionStatus=&cStatus
	conn.IsBulkMeter=ToBoolPointer(false)
	conn.NoUnits=ToIntPointer(noUnits);
	conn.EstimCons=&estimCons
	meter:=Meter{}
	conn.Meter=&meter
	meter.Diameter=ToIntPointer(10)
	meter.MeterRef=ToStringPointer("1")
	meter.MeterType=ToStringPointer("2")
	meter.OpStatus=&meteOpStatus
	return &cust
}
func GetMultiConnectionCustomer(typ int64,isVactated bool,mainCtype string,estimCons float64,noUnits int64,meteOpStatus MeterOperationStatus,connections []*SubConnection) *Customer{
	if connections==nil || len(connections)==0{
		return GetNoramlCustomer(typ,isVactated,mainCtype,estimCons,noUnits,meteOpStatus)
	}
	cust:=GetNoramlCustomer(typ,isVactated,mainCtype,estimCons,noUnits,meteOpStatus)
	if cust.Property!=nil && len(cust.Property.Services)>0 {
		cust.Property.Services[0].Connection.SubConnections=connections
	}
	return cust
}

