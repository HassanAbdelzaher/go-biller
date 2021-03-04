package charge

import (	. "MaisrForAdvancedSystems/go-biller/proto"
	. "MaisrForAdvancedSystems/go-biller/tools"
	"testing")

func TestCustTypeIsChargeEnable(t *testing.T) {
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
	cr:=getCustTypeChargeRegularSample(entityType ,MappedValues)
	cust:=getNoramlCustomer(1,false,"00/01",10)
	isEnable,err:=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if isEnable{
		t.Errorf("%s:must be disabled while found enabled",entityType)
	}
	///
	cust=getNoramlCustomer(2,false,"00/01",10)
	isEnable,err=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if !isEnable{
		t.Errorf("%s:must be enabled while found diabled",entityType)
	}

	cust=getNoramlCustomer(3,false,"00/01",10)
	isEnable,err=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if isEnable{
		t.Error("must be disabled while found enabled")
	}
}

func TestPropertyVacatedChargeEnable(t *testing.T) {
	entityType:=ENTITY_TYPE_PROPERTY_VACATED
	MappedValues:=make([]*EntityEnableMappedValue,0)
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("0"),
		Value:ToBoolPointer(false),
	})
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("1"),
		Value:ToBoolPointer(true),
	})
	cr:=getCustTypeChargeRegularSample(entityType ,MappedValues)
	cust:=getNoramlCustomer(1,false,"00/01",10)
	isEnable,err:=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if isEnable{
		t.Errorf("%s:must be disabled while found enabled",entityType)
	}
	///
	cust=getNoramlCustomer(2,true,"00/01",10)
	isEnable,err=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if !isEnable{
		t.Errorf("%s:must be enabled while found diabled",entityType)
	}
}



func TestPropertyServiceChargeEnable(t *testing.T) {
	entityType:=ENTITY_TYPE_SERVICE
	MappedValues:=make([]*EntityEnableMappedValue,0)
	water:=int64(SERVICE_TYPE_WATER)
	sewer:=int64(SERVICE_TYPE_SEWER)
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:Int64ToString(&water),
		Value:ToBoolPointer(false),
	})
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:Int64ToString(&sewer),
		Value:ToBoolPointer(true),
	})
	cr:=getCustTypeChargeRegularSample(entityType ,MappedValues)
	cust:=getNoramlCustomer(1,true,"00/01",10)
	isEnable,err:=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if isEnable{
		t.Errorf("%s:must be disabled while found enabled",entityType)
	}
}


func TestMultiCTypeChargeEnable(t *testing.T) {
	entityType:=ENTITY_TYPE_CTYPE
	MappedValues:=make([]*EntityEnableMappedValue,0)
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("00/01"),
		Value:ToBoolPointer(true),
	})
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("00/02"),
		Value:ToBoolPointer(false),
	})
	cr:=getCustTypeChargeRegularSample(entityType ,MappedValues)
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
	cust:=getMultiConnectionCustomer(1,true,30,"00/01",conns)
	isEnable,err:=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if !isEnable{
		t.Errorf("%s:must be enable while found diabled",entityType)
	}
}



func TestCTypeChargeEnable(t *testing.T) {
	entityType:=ENTITY_TYPE_CTYPE
	MappedValues:=make([]*EntityEnableMappedValue,0)
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("00/01"),
		Value:ToBoolPointer(true),
	})
	MappedValues=append(MappedValues,&EntityEnableMappedValue{
		LuKey:ToStringPointer("00/02"),
		Value:ToBoolPointer(false),
	})
	cr:=getCustTypeChargeRegularSample(entityType ,MappedValues)
	cust:=getNoramlCustomer(1,true,"00/01",10)
	isEnable,err:=IsChargeEnable(cr,cust)
	if err!=nil{
		t.Error(err)
	}
	if !isEnable{
		t.Errorf("%s:must be enabled while found diabled",entityType)
	}
}

