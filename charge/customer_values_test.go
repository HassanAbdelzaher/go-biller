package charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	"testing")
func TestCustomerValues(t *testing.T) {
	entityType := ENTITY_TYPE_CUSTOMER_TYPE
	cust := getNoramlCustomer(1, false, "00/01", 10)
	values:=CustomerValues(entityType,cust)
	if values==nil{
		t.Errorf("TestCustomerValues:%s invalied return -- null value",entityType)
	}

	if len(values)!=1{
		t.Errorf("TestCustomerValues:%s expected length 1 while found %d",entityType,len(values))
	}

	if values[0]==nil{
		t.Errorf("TestCustomerValues:%s expected length 1 while found null values %d",entityType,len(values))
	}
	cstType:="1"
	if *values[0]!=cstType{
		t.Errorf("TestCustomerValues:%s expected value %s while found %s",entityType,cstType,*values[0])
	}
}

func TestMultiCtypeCustomerValues(t *testing.T) {
	entityType := ENTITY_TYPE_CTYPE
	conns:=[]*SubConnection{
		&SubConnection{
			CType:                      ToStringPointer("00/01"),
			EstimateConsumptionPerUnit: ToFloatPointer(10),
			ConsumptionPercentage:      ToFloatPointer(10),
			NoUnits:                    ToIntPointer(1),
		},
		&SubConnection{
			CType:                      ToStringPointer("00/01"),
			EstimateConsumptionPerUnit: ToFloatPointer(20),
			ConsumptionPercentage:      ToFloatPointer(90),
			NoUnits:                    ToIntPointer(9),
		},
		&SubConnection{
			CType:                      ToStringPointer("00/03"),
			EstimateConsumptionPerUnit: ToFloatPointer(20),
			ConsumptionPercentage:      ToFloatPointer(90),
			NoUnits:                    ToIntPointer(9),
		},
	}
	cust:=getMultiConnectionCustomer(1,true,30,"00/01",conns)
	values:=CustomerValues(entityType,cust)
	if values==nil{
		t.Errorf("TestCustomerValues:%s invalied return -- null value",entityType)
	}

	if len(values)!=len(conns){
		t.Errorf("TestCustomerValues:%s expected length %d while found %d",entityType,len(conns),len(values))
	}

	disitnctCtypes:=make(map[string]int32)
	for _,c:=range conns{
		disitnctCtypes[*c.CType]=disitnctCtypes[*c.CType]+1
	}

	for ctyp,count:=range disitnctCtypes{
		var cnt int32=0
		for _,v:=range values{
			if ctyp==*v{
				cnt++
			}
		}
		if cnt!=count{
			t.Errorf("TestCustomerValues:%s  %s expected found %d while found %d",entityType,ctyp,count,cnt)
			break
		}
	}
}

