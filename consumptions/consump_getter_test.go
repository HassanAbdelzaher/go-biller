package consumptions
// this the most important part in testing
// because most of bissness belong to this operation
import (
	billing "MaisrForAdvancedSystems/go-biller/proto"
	samples "MaisrForAdvancedSystems/go-biller/samples"
	"MaisrForAdvancedSystems/go-biller/tools"
	"google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"testing"
	"time"
)

func TestConnectionActualReading(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_WORKING
	cust:=samples.GetNoramlCustomer(0,false,"00/01",12,2,meterWorking)
	conn:=cust.Property.Services[0].Connection
	crDate:=timestamppb.New(time.Now())
	prDate:=timestamppb.New(time.Now().AddDate(0,-1,0))
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),&billing.Reading{
		Consump:              tools.ToFloatPointer(30),
		PrReading:            tools.ToFloatPointer(0),
		CrReading:            tools.ToFloatPointer(90),
		PrDate:               prDate,
		CrDate:               crDate,
	},true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=1{
		t.Errorf("expect return with length %d while found %d ",1,len(data))
		return
	}

	if data[*cust.Property.Services[0].Connection.CType].Consump!=30{
		t.Errorf("expect consump %v while found %v ",30,data[*cust.Property.Services[0].Connection.CType].Consump)
		return
	}

	if data[*cust.Property.Services[0].Connection.CType].NoUnits!=2{
		t.Errorf("expect consump %v while found %v ",2,data[*cust.Property.Services[0].Connection.CType].NoUnits)
		return
	}

}
func TestConnectionWithoutActualReading(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_WORKING
	cust:=samples.GetNoramlCustomer(0,false,"00/01",12,2,meterWorking)
	conn:=cust.Property.Services[0].Connection
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),nil,true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=0{
		t.Errorf("expect return with length %d while found %d ",0,len(data))
		return
	}
}
func TestMultiConnectionWithActualReading(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_WORKING
	var consump float64=60
	connections:=[]*billing.SubConnection{
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/01"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(25),
			NoUnits:                    tools.ToIntPointer(10),
		},
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/02"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(75),
			NoUnits:                    tools.ToIntPointer(20),
		},
	}
	cust:=samples.GetMultiConnectionCustomer(0,false,"00/01",12,2,meterWorking,connections)
	conn:=cust.Property.Services[0].Connection
	crDate:=timestamppb.New(time.Now())
	prDate:=timestamppb.New(time.Now().AddDate(0,-1,0))
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),&billing.Reading{
		Consump:              tools.ToFloatPointer(60),
		PrReading:            tools.ToFloatPointer(0),
		CrReading:            tools.ToFloatPointer(60),
		PrDate:               prDate,
		CrDate:               crDate,
	},true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=2{
		t.Errorf("expect return with length %d while found %d ",2,len(data))
		return
	}

	for _,c:=range connections{
		errVal:=data[*c.CType].Consump-*c.ConsumptionPercentage*consump/100
		if math.Abs(errVal)>0.000001{
			t.Errorf("expect consump %v while found %v ",*c.ConsumptionPercentage*consump/100,data[*c.CType].Consump)
			return
		}
	}


}
func TestMultiConnectionOpWithoutActualReading(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_WORKING
	connections:=[]*billing.SubConnection{
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/01"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(25),
			NoUnits:                    tools.ToIntPointer(10),
		},
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/02"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(75),
			NoUnits:                    tools.ToIntPointer(20),
		},
	}
	cust:=samples.GetMultiConnectionCustomer(0,false,"00/01",12,2,meterWorking,connections)
	conn:=cust.Property.Services[0].Connection
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),nil,true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=0{
		t.Errorf("expect return with length %d while found %d ",0,len(data))
		return
	}
}
func TestMultiConnectionNoOpWithoutActualReading(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_NOT_WORKING
	connections:=[]*billing.SubConnection{
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/01"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(25),
			NoUnits:                    tools.ToIntPointer(10),
		},
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/02"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: nil,
			ConsumptionPercentage:      tools.ToFloatPointer(75),
			NoUnits:                    tools.ToIntPointer(20),
		},
	}
	cust:=samples.GetMultiConnectionCustomer(0,false,"00/01",12,2,meterWorking,connections)
	conn:=cust.Property.Services[0].Connection
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),nil,true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=2{
		t.Errorf("expect return with length %d while found %d ",2,len(data))
		return
	}
	ctgs:=samples.GetCtgMap()
	for _,c:=range connections{
		ctg,ok:=ctgs[*c.CType]
		if !ok{
			t.Errorf("missing ctype in test pattern ")
			return
		}
		if ctg.NOOP_ESTIM_CONS!=nil{
			if *ctg.NOOP_ESTIM_CONS!=data[*c.CType].Consump{
				t.Errorf("expect consump %v while found %v ",*ctg.NOOP_ESTIM_CONS,data[*c.CType].Consump)
				return
			}
		}
	}
}
func TestMultiConnectionNoOpWithoutActualReading207(t *testing.T){
	var meterWorking=billing.MeterOperationStatus_NOT_WORKING
	connections:=[]*billing.SubConnection{
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/01"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: tools.ToFloatPointer(10),
			ConsumptionPercentage:      tools.ToFloatPointer(25),
			NoUnits:                    tools.ToIntPointer(10),
		},
		&billing.SubConnection{
			CType:                      tools.ToStringPointer("00/02"),
			CTYPE_GROUP:                tools.ToStringPointer("00"),
			EstimateConsumptionPerUnit: tools.ToFloatPointer(30),
			ConsumptionPercentage:      tools.ToFloatPointer(75),
			NoUnits:                    tools.ToIntPointer(20),
		},
	}
	cust:=samples.GetMultiConnectionCustomer(0,false,"00/01",12,2,meterWorking,connections)
	conn:=cust.Property.Services[0].Connection
	data,err:=GetConnectionsConsumption(samples.GetCtgMap(),conn,time.Now(),nil,true)
	if err!=nil{
		t.Error(err)
		return
	}
	if data==nil{
		t.Error("Invalied return form GetConnectionsConsumption")
		return
	}
	for _,d:=range data{
		t.Log(d.Ctype,d.Consump,d.NoUnits,d.ReadType)
	}
	if len(data)!=2{
		t.Errorf("expect return with length %d while found %d ",2,len(data))
		return
	}

	for _,c:=range connections{
		noUn:=float64(tools.DefaultI(c.NoUnits,1))
		dataCons:=tools.Multiply(&noUn,c.EstimateConsumptionPerUnit)
		if dataCons==nil || *dataCons!=data[*c.CType].Consump{
			t.Errorf("expect consump %v while found %v ",*dataCons,data[*c.CType].Consump)
			return
		}
	}
}


