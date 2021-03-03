package charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	pt "github.com/golang/protobuf/ptypes"
	"time"
)

func getTariffSample() *Tariff {
	tar := Tariff{}
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

func getCustTypeChargeRegularSample(entityType ENTITY_TYPE,maped []*EntityEnableMappedValue) *RegularCharge{
	var cp=RegularChargePeriod_BILL
	var effDate=time.Date(2000,1,1,0,0,0,0,time.Local)
	var tsmp,_=pt.TimestampProto(effDate)
	rg:=RegularCharge{}
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
func getNoramlCustomer(typ int64,isVactated bool,ctype string,estimCons float64) *Customer{
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
	conn.NoUnits=ToIntPointer(1);
	conn.EstimCons=&estimCons
	meter:=Meter{}
	conn.Meter=&meter
	meter.Diameter=ToIntPointer(10)
	meter.MeterRef=ToStringPointer("1")
	meter.MeterType=ToStringPointer("2")
	return &cust
}


func getMultiConnectionCustomer(typ int64,isVactated bool,estimCons float64,mainCtype string,connections []*SubConnection) *Customer{
	if connections==nil || len(connections)==0{
		return getNoramlCustomer(typ,isVactated,mainCtype,estimCons)
	}
	cust:=getNoramlCustomer(typ,isVactated,mainCtype,estimCons)
	if cust.Property!=nil && len(cust.Property.Services)>0 {
		cust.Property.Services[0].Connection.SubConnections=connections
	}
	return cust
}

