package regular_charge

import (
	. "MaisrForAdvancedSystems/go-biller/proto"
	"MaisrForAdvancedSystems/go-biller/tools"
)

func CustomerValues(entityType ENTITY_TYPE,c *Customer) ([]*string){
	typ:=entityType
	var values=make([]*string,0)
	if typ == ENTITY_TYPE_CUSTOMER_TYPE {
		val:=tools.Int64ToString(c.CustType)
		values=append(values,val)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG1 {
		values=append(values,c.InfoFlag1)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG2 {
		values=append(values,c.InfoFlag2)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG3 {
		values=append(values,c.InfoFlag3)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG4 {
		values=append(values,c.InfoFlag4)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG5 {
		values=append(values,c.InfoFlag5)
	}
	if c.Property != nil {
		if typ == ENTITY_TYPE_PROPERTY_VACATED {
			val:=tools.BoolToString(c.Property.IsVacated)
			values=append(values,val)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG1 {
			values=append(values,c.Property.InfoFlag1)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG2 {
			values=append(values,c.Property.InfoFlag2)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG3 {
			values=append(values,c.Property.InfoFlag3)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG4 {
			values=append(values,c.Property.InfoFlag4)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG5 {
			values=append(values,c.Property.InfoFlag5)
		}
		if typ == ENTITY_TYPE_TOWINSHIP {
			values=append(values,c.Property.TOWINSHIP)
		}
		services := c.Property.Services
		if services != nil && len(services) > 0 {
			if typ == ENTITY_TYPE_SERVICE {
				for _, srv := range services {
					if srv.ServiceType != nil {
						srvType := int64(*srv.ServiceType)
						val := tools.Int64ToString(&srvType)
						values=append(values,val)
					}
				}
			}
			for _, srv := range services {
				if srv.Connection != nil {
					conn:=srv.Connection
					if typ == ENTITY_TYPE_CONNECTION_DIAMETER {
						val:=tools.Int64ToString(conn.ConnDiameter)
						values=append(values,val)
					}
					if typ == ENTITY_TYPE_CONNECTION_STATUS {
						if conn.ConnectionStatus!=nil{
							var cSttaus int32=int32(*conn.ConnectionStatus)
							val:=tools.Int32ToString(&cSttaus)
							values=append(values,val)
						}
					}
					if typ == ENTITY_TYPE_CONNECTION_ISBULK_METER {
						val:=tools.BoolToString(conn.IsBulkMeter)
						values=append(values,val)
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
						values=append(values,ctypes...)
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
						values=append(values,ctypes_groups...)
					}
					if srv.Connection.Meter != nil {
						if typ == ENTITY_TYPE_METER_DIAMETER {
							val:=tools.Int64ToString(conn.Meter.Diameter)
							values=append(values,val)
						}
					}
				}
			}
		}
	}
	return values
}

