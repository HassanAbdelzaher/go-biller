package charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"log"
)

func IsChargeEnable(fee *RegularCharge,c *Customer) (bool,error){
	if fee.Bypass!=nil{
		if *fee.Bypass{
			return true,nil
		}
	}
	if fee.RelationEnableEntity==nil{
		return false,errors.New("missing enabled entity for charge regular")
	}
	ree:=fee.RelationEnableEntity
	if ree.EntityType==nil{
		return false,errors.New("missing enabled entity type for charge regular")
	}
	typ:=*ree.EntityType
	var mappedValues=ree.MappedValues
	if mappedValues==nil || len(mappedValues)==0{
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_TYPE {
		val:=tools.Int64ToString(c.CustType)
		for _,m:=range mappedValues{
			log.Println(*m.LuKey,*val)
			if tools.StringComparePointer(m.LuKey,val){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG1 {
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,c.InfoFlag1){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG2 {
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,c.InfoFlag2){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG3 {
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,c.InfoFlag3){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG4 {
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,c.InfoFlag4){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG5 {
		for _,m:=range mappedValues{
			if tools.StringComparePointer(m.LuKey,c.InfoFlag5){
				return m.GetValue(),nil
				break;
			}
		}
		return false,nil
	}
	if c.Property != nil {
		if typ == ENTITY_TYPE_PROPERTY_VACATED {
			val:=tools.BoolToString(c.Property.IsVacated)
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,val){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG1 {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.InfoFlag1){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG2 {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.InfoFlag2){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG3 {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.InfoFlag3){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG4 {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.InfoFlag4){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG5 {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.InfoFlag5){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		if typ == ENTITY_TYPE_TOWINSHIP {
			for _,m:=range mappedValues{
				if tools.StringComparePointer(m.LuKey,c.Property.TOWINSHIP){
					return m.GetValue(),nil
					break;
				}
			}
			return false,nil
		}
		services := c.Property.Services
		if services != nil && len(services) > 0 {
			if typ == ENTITY_TYPE_SERVICE {
				for _, srv := range services {
					if srv.ServiceType != nil {
						srvType := int64(*srv.ServiceType)
						val := tools.Int64ToString(&srvType)
						for _, m := range mappedValues {
							if tools.StringComparePointer(m.LuKey, val) {
								return m.GetValue(), nil
								break;
							}
						}
					}
				}
				return false, nil
			}
			for _, srv := range services {
				if srv.Connection != nil {
					conn:=srv.Connection
					if typ == ENTITY_TYPE_CONNECTION_DIAMETER {
						val:=tools.Int64ToString(conn.ConnDiameter)
						for _,m:=range mappedValues{
							if tools.StringComparePointer(m.LuKey,val){
								return m.GetValue(),nil
								break;
							}
						}
						return false,nil
					}
					if typ == ENTITY_TYPE_CONNECTION_STATUS {
						if conn.ConnectionStatus!=nil{
							var cSttaus int32=int32(*conn.ConnectionStatus)
							val:=tools.Int32ToString(&cSttaus)
							for _,m:=range mappedValues{
								if tools.StringComparePointer(m.LuKey,val){
									return m.GetValue(),nil
									break;
								}
							}
						}
						return false,nil
					}
					if typ == ENTITY_TYPE_CONNECTION_ISBULK_METER {
						val:=tools.BoolToString(conn.IsBulkMeter)
						for _,m:=range mappedValues{
							if tools.StringComparePointer(m.LuKey,val){
								return m.GetValue(),nil
								break;
							}
						}
						return false,nil
					}
					if typ == ENTITY_TYPE_CTYPE {
						ctypes := make([]*string, 0)
						if srv.Connection.SubConnections != nil && len(srv.Connection.SubConnections) > 0 {
							for idx := range srv.Connection.SubConnections {
								ctypes = append(ctypes, srv.Connection.SubConnections[idx].CType)
							}
						} else {
							ctypes = append(ctypes, srv.Connection.CType)
						}
						for _,ctype:=range ctypes{
							for _,m:=range mappedValues{
								if tools.StringComparePointer(m.LuKey,ctype){
									return m.GetValue(),nil
									break;
								}
							}
						}
					}
					if typ == ENTITY_TYPE_CTYPE_GROUP {
						ctypes_groups := make([]*string, 0)
						if srv.Connection.SubConnections != nil && len(srv.Connection.SubConnections) > 0 {
							for idx := range srv.Connection.SubConnections {
								ctypes_groups = append(ctypes_groups, srv.Connection.SubConnections[idx].CTYPE_GROUP)
							}
						} else {
							ctypes_groups = append(ctypes_groups, srv.Connection.CTYPE_GROUP)
						}
						for _,ctype:=range ctypes_groups{
							for _,m:=range mappedValues{
								if tools.StringComparePointer(m.LuKey,ctype){
									return m.GetValue(),nil
									break;
								}
							}
						}
					}
					if srv.Connection.Meter != nil {
						if typ == ENTITY_TYPE_METER_DIAMETER {
							val:=tools.Int64ToString(conn.Meter.Diameter)
							for _,m:=range mappedValues{
								if tools.StringComparePointer(m.LuKey,val){
									return m.GetValue(),nil
									break;
								}
							}
						}
					}
				}
			}
		}
	}
	return false,nil
}

