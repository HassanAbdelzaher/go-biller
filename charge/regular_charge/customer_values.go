package regular_charge

import (
	"MaisrForAdvancedSystems/go-biller/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type MappedData struct {
	cType      *string
	cTypeGroup *string
	noUnits    *int64
}

type CustomerValues map[string]*MappedData

func (c CustomerValues) append(key *string, v *MappedData) {
	if key != nil {
		if v == nil {
			v = &MappedData{
				cType:   nil,
				noUnits: nil,
			}
		}
		c[*key] = v
	}
}

func customerValues(entityType ENTITY_TYPE, c *Customer, certainCerviceType *SERVICE_TYPE) map[string]*MappedData {
	typ := entityType
	var values CustomerValues = make(map[string]*MappedData)
	var mainMappedData = &MappedData{
		cType:   nil,
		noUnits: nil,
	}
	if typ == ENTITY_TYPE_CUSTOMER_TYPE {
		val := tools.Int64ToString(c.CustType)
		values.append(val, mainMappedData)
		if val != nil {
			values[*val] = nil
		}
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG1 {
		values.append(c.InfoFlag1, mainMappedData)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG2 {
		values.append(c.InfoFlag2, mainMappedData)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG3 {
		values.append(c.InfoFlag3, mainMappedData)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG4 {
		values.append(c.InfoFlag4, mainMappedData)
	}
	if typ == ENTITY_TYPE_CUSTOMER_FLAG5 {
		values.append(c.InfoFlag5, mainMappedData)
	}
	if c.Property != nil {
		if typ == ENTITY_TYPE_PROPERTY_VACATED {
			val := tools.BoolToString(c.Property.IsVacated)
			values.append(val, mainMappedData)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG1 {
			values.append(c.Property.InfoFlag1, mainMappedData)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG2 {
			values.append(c.Property.InfoFlag2, mainMappedData)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG3 {
			values.append(c.Property.InfoFlag3, mainMappedData)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG4 {
			values.append(c.Property.InfoFlag4, mainMappedData)
		}
		if typ == ENTITY_TYPE_PROPERTY_FLAG5 {
			values.append(c.Property.InfoFlag5, mainMappedData)
		}
		if typ == ENTITY_TYPE_TOWINSHIP {
			values.append(c.Property.Township, mainMappedData)
		}
		services := c.Property.Services
		if services != nil && len(services) > 0 {
			if typ == ENTITY_TYPE_SERVICE {
				for _, srv := range services {
					if srv.ServiceType != nil {
						srvType := int64(*srv.ServiceType)
						val := tools.Int64ToString(&srvType)
						values.append(val, mainMappedData)
					}
				}
			}
			for _, srv := range services {
				if certainCerviceType != nil && *certainCerviceType != srv.GetServiceType() {
					continue
				}
				if srv.Connection != nil {
					conn := srv.Connection
					mainNoUnits := tools.DefaultI(conn.NoUnits, int64(1))
					mainMappedData.noUnits = &mainNoUnits
					mainMappedData.cType = conn.CType
					mainMappedData.cTypeGroup = conn.CTYPE_GROUP
					if typ == ENTITY_TYPE_CONNECTION_DIAMETER {
						val := tools.Int64ToString(conn.ConnDiameter)
						values.append(val, mainMappedData)
					}
					if typ == ENTITY_TYPE_CONNECTION_STATUS {
						if conn.ConnectionStatus != nil {
							var cSttaus int32 = int32(*conn.ConnectionStatus)
							val := tools.Int32ToString(&cSttaus)
							values.append(val, mainMappedData)
						}
					}
					if typ == ENTITY_TYPE_CONNECTION_ISBULK_METER {
						val := tools.BoolToString(conn.IsBulkMeter)
						values.append(val, mainMappedData)
					}
					if typ == ENTITY_TYPE_CTYPE {
						//ctypes := make([]*string, 0)
						if srv.Connection.SubConnections != nil && len(srv.Connection.SubConnections) > 0 {
							for idx := range srv.Connection.SubConnections {
								subMappedType := &MappedData{
									cType:      nil,
									cTypeGroup: nil,
									noUnits:    nil,
								}
								values.append(srv.Connection.SubConnections[idx].CType, subMappedType)
							}
						} else {
							values.append(srv.Connection.CType, mainMappedData)
						}
					}
					if typ == ENTITY_TYPE_CTYPE_GROUP {
						if srv.Connection.SubConnections != nil && len(srv.Connection.SubConnections) > 0 {
							for idx := range srv.Connection.SubConnections {
								subMappedType := &MappedData{
									cType:      nil,
									cTypeGroup: nil,
									noUnits:    nil,
								}
								values.append(srv.Connection.SubConnections[idx].CTYPE_GROUP, subMappedType)
							}
						} else {
							values.append(srv.Connection.CTYPE_GROUP, mainMappedData)
						}
					}
					if srv.Connection.Meter != nil {
						if typ == ENTITY_TYPE_METER_DIAMETER {
							val := tools.Int64ToString(conn.Meter.Diameter)
							values.append(val, mainMappedData)
						}
					}
				}
			}
		}
	}
	return values
}
