package samples

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	pt "github.com/golang/protobuf/ptypes"
	"time"
)

func GetCtgMap() map[string]*Ctg{
	ctype:="00/01"
	ctype2:="00/02"
	water:=SERVICE_TYPE_WATER
	c1:=Ctg{
		CType:                ToStringPointer(ctype),
		CTypeGroupid:         ToStringPointer("00"),
		Tariffs:              []*ServiceTariff{&ServiceTariff{
			ServiceType:          &water,
			TarifId:              ToStringPointer("0-1"),
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
		}},
		OP_ESTIM_CONS:        ToFloatPointer(float64(90)),
		NOOP_ESTIM_CONS:      ToFloatPointer(float64(150)),
	}
	return map[string]*Ctg{*c1.CType:&c1,*c2.CType:&c2}
}

func GetTariffSampleMap() map[string]*Tariff{
	tar:=GetTariffSample()
	return map[string]*Tariff{tar.GetTariffId():tar}
}
func GetTariffSample() *Tariff {
	tar := Tariff{}
	tar.TariffId=ToStringPointer("0-1")
	tar.Bands = make([]*TariffBand, 0)
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
	return &tar
}

func GetCustTypeChargeRegularSample(entityType ENTITY_TYPE,maped []*EntityEnableMappedValue) *RegularCharge{
	var cp=RegularChargePeriod_BILL
	var effDate=time.Date(2000,1,1,0,0,0,0,time.Local)
	var tsmp,_=pt.TimestampProto(effDate)
	rg:=RegularCharge{}
	srvType:=SERVICE_TYPE_WATER
	rg.ServiceType=&srvType
	crel:=ChargeType_RELATION
	rg.ChargeType=&crel;
	rg.TransCode=ToIntPointer(60);
	rg.TransSCode=ToIntPointer(50);
	rg.TransTitle=ToStringPointer("Water")
	rg.TransTitle=ToStringPointer("Estidama")
	rg.VatPercentage=ToFloatPointer(14)
	rg.ChargeCalcPeriod=&cp
	rg.EffectiveDate=tsmp
	rg.FixedCharge=ToFloatPointer(12)
	rg.IsChargable=ToBoolPointer(true)
	en:=RegularEnableEntity{}
	en.Code=ToStringPointer("98")
	en.MappedValues=make([]*EntityEnableMappedValue,0)
	en.EntityType=&entityType
	en.MappedValues=maped
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

