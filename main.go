package main

import (
	billing "MaisrForAdvancedSystems/go-biller/proto"
	"MaisrForAdvancedSystems/go-biller/samples"
	"MaisrForAdvancedSystems/go-biller/service"
	"MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strconv"
	"time"
)

func floatToString(f *float64) *string{
	if f==nil{
		return nil
	}
	str:=strconv.FormatFloat(*f,'f',-1,64)
	return &str
}

func main() {
	bs:=service.BillingService{
		Tariffs:        samples.GetTariffSampleMap(),
		Ctgs:           samples.GetCtgMap(),
		//RegularCharges: samples.GetCustTypeChargeRegularSample(),
		IsTrace:        true,
	}
	setting:=billing.ChargeSetting{
		CycleLength:          tools.ToIntPointer(1),
		BilingDate:           timestamppb.New(time.Now()),
		IgnoreTimeEffect:     nil,
	}
	cst:=samples.GetNoramlCustomer(0,false,"00/01",10,1,billing.MeterOperationStatus_WORKING)
	rdg:=billing.Reading{
		Consump:              tools.ToFloatPointer(30),
		PrReading:            nil,
		CrReading:            nil,
		PrDate:               nil,
		CrDate:               nil,
	}
	water:=billing.SERVICE_TYPE_WATER
	rq:=billing.BillRequest{
		Customer:             cst,
		ServicesReadings:     []*billing.ServiceReading{
			&billing.ServiceReading{
				ServiceType:          &water,
				Reading:              &rdg,
			},
		},
		Setting:              &setting,
	}
	data,err:=bs.Charge(context.Background(),&rq)
	if err!=nil || data==nil{
		log.Println(err)
		return
	}
	log.Println(data)
	for _,d:=range data.Charges{
		if d.Charges!=nil{
			for _,ch:=range d.Charges{
				log.Println(*ch.CType,*ch.Amount)
			}
		}
	}
}
