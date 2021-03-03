package charge

import billing "MaisrForAdvancedSystems/go-biller/proto"

type EntityTypeGetter billing.ENTITY_TYPE
func (t *EntityTypeGetter) getValue(c *billing.Customer) interface{}{
	var ty billing.ENTITY_TYPE=billing.ENTITY_TYPE(*t)
	return GetEntityTypeValue(c,ty)
}

func GetEntityTypeValue(c *billing.Customer,typ billing.ENTITY_TYPE) interface{} {
	if typ == billing.ENTITY_TYPE_CUSTOMER_TYPE {
		return c.CustType
	}
	if typ == billing.ENTITY_TYPE_CUSTOMER_FLAG1 {
		return c.InfoFlag1
	}
	if typ == billing.ENTITY_TYPE_CUSTOMER_FLAG2 {
		return c.InfoFlag2
	}
	if typ == billing.ENTITY_TYPE_CUSTOMER_FLAG3 {
		return c.InfoFlag3
	}
	if typ == billing.ENTITY_TYPE_CUSTOMER_FLAG4 {
		return c.InfoFlag4
	}
	if typ == billing.ENTITY_TYPE_CUSTOMER_FLAG5 {
		return c.InfoFlag5
	}
	if c.Property != nil {
		if typ == billing.ENTITY_TYPE_PROPERTY_VACATED {
			return c.Property.IsVacated
		}
		if typ == billing.ENTITY_TYPE_PROPERTY_FLAG1 {
			return c.Property.InfoFlag1
		}
		if typ == billing.ENTITY_TYPE_PROPERTY_FLAG2 {
			return c.Property.InfoFlag2
		}
		if typ == billing.ENTITY_TYPE_PROPERTY_FLAG3 {
			return c.Property.InfoFlag3
		}
		if typ == billing.ENTITY_TYPE_PROPERTY_FLAG4 {
			return c.Property.InfoFlag4
		}
		if typ == billing.ENTITY_TYPE_PROPERTY_FLAG5 {
			return c.Property.InfoFlag5
		}
		services := c.Property.Services
		if services != nil && len(services) > 0 {
			if typ == billing.ENTITY_TYPE_SERVICE {
				return services
			}
			for _, srv := range services {
				if srv.Connection != nil {
					if typ == billing.ENTITY_TYPE_CONNECTION_DIAMETER {
						return srv.Connection.ConnDiameter
					}
					if typ == billing.ENTITY_TYPE_CONNECTION_STATUS {
						return srv.Connection.ConnectionStatus
					}
					if typ == billing.ENTITY_TYPE_CONNECTION_ISBULK_METER {
						return srv.Connection.IsBulkMeter
					}
					if typ == billing.ENTITY_TYPE_CTG_TYPE {
						ctypes := make([]*string, 0)
						if srv.Connection.SubConnections != nil && len(srv.Connection.SubConnections) > 0 {
							for idx := range srv.Connection.SubConnections {
								ctypes = append(ctypes, srv.Connection.SubConnections[idx].CType)
							}
						} else {
							ctypes = append(ctypes, srv.Connection.CType)
						}
						return ctypes
					}
					if srv.Connection.Meter != nil {
						if typ == billing.ENTITY_TYPE_METER_DIAMETER {
							return srv.Connection.Meter.Diameter
						}
					}
				}
			}
		}
	}
	return nil
}
