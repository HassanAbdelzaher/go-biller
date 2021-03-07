package consumptions

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	. "MaisrForAdvancedSystems/go-biller/proto"
	"log"
	"time"
)
type EstimConsumption struct{
	*SubConnection
	IsMain bool
	MainConnection *Connection
}

type ConnectionConsumption struct{
	Ctype string
	NoUnits int64
	Consump float64
	PrDate time.Time
	CrDate time.Time
	ReadType READING_TYPE
}
func GetConnectionsConsumption(ctgs map[string]*Ctg,conn *Connection,bilngDate time.Time,reading *Reading,trace bool) (map[string]ConnectionConsumption,error){
	Trace:=TraceFFunc(trace)
	if conn==nil{
		Trace("Invalied request connection is null")
		return nil,errors.New("Invalied Property Connection")
	}
	//raise error if connection status not defined
	if conn.ConnectionStatus==nil{
		Trace("Connection Status Not Defined")
		return nil,errors.New("حالة التوصيلة غير معرفة")
	}
	cons:=make(map[string]ConnectionConsumption)
	//reject disconnected connections
	if *conn.ConnectionStatus==CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER||*conn.ConnectionStatus==CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER{
		Trace("Reject disconnected connections")
		return cons,nil
	}
	//working meters without readings always have no charge
	if conn.Meter!=nil && (reading==nil || reading.Consump==nil){
		opStatus:=conn.Meter.GetOpStatus()
		//handle working meters
		if opStatus==MeterOperationStatus_WORKING{
			Trace("Reject working meters without readings")
			return cons,nil
		}
	}
	estims:=make([]*EstimConsumption,0)
	totalUnits:=tools.DefaultI(conn.NoUnits,int64(1))
	totalUnitsF:=float64(totalUnits)
	isHaveCustom:=false
	var handred float64=100
	defaultPrDate:=bilngDate.AddDate(0,-1,0)

	// Handle No Meter connection or non working meters without readings
	if *conn.ConnectionStatus==CONNECTION_STATUS_TYPE_CONNECTED_WITHOUT_METER || conn.Meter==nil || reading==nil || reading.Consump==nil{
		//calculating estim consumptions
		//make slice for connection estim consumtions
		if conn.SubConnections!=nil && len(conn.SubConnections)>0{
			//working with customers have custom estim for each sub connection 207
			Trace("find estim consumptions from SubConnection with custom estim (207)")
			for id:=range conn.SubConnections{
				sb:=conn.SubConnections[id]
				if sb!=nil && sb.EstimateConsumptionPerUnit!=nil && *sb.EstimateConsumptionPerUnit>0{
					_noUnits:=tools.DefaultI(sb.NoUnits,int64(1))
					sb.NoUnits=&_noUnits
					estims=append(estims,&EstimConsumption{
						SubConnection: sb ,
						IsMain:         false,
						MainConnection: conn,
					})
					isHaveCustom=true
				}
			}
			//working with normal estim which divide estim cosumptions per each ctype
			if !isHaveCustom && conn.EstimCons!=nil && *conn.EstimCons>0{
				Trace("find estim consumptions for SubConnection without custom estim")
				for id:=range conn.SubConnections{
					//working with customers have custom estim for each sub connection 207
					sb:=conn.SubConnections[id]
					if sb!=nil && sb.EstimateConsumptionPerUnit!=nil && *sb.EstimateConsumptionPerUnit>0{
						if sb.GetConsumptionPercentage()>0{
							_noUnits:=tools.DefaultI(sb.NoUnits,int64(1))
							_noUnitsF:=float64(_noUnits)
							sb.NoUnits=&_noUnits
							ratio:=tools.Divide(sb.ConsumptionPercentage,&handred)
							sb.EstimateConsumptionPerUnit=tools.Divide(conn.EstimCons,&_noUnitsF)
							sb.EstimateConsumptionPerUnit=tools.Multiply(sb.EstimateConsumptionPerUnit,ratio)
							estims=append(estims,&EstimConsumption{
								SubConnection: sb ,
								IsMain:         false,
								MainConnection: conn,
							})
						}
					}
				}
			}
		}else {
			//handle connection without sub connections
			Trace("find estim consumptions for main connection")
			if conn.EstimCons!=nil && *conn.EstimCons>0{
				estims=append(estims,&EstimConsumption{
					SubConnection: &SubConnection{
						CType:                      conn.CType,
						CTYPE_GROUP:                conn.CTYPE_GROUP,
						EstimateConsumptionPerUnit: tools.Divide(conn.EstimCons,&totalUnitsF),
						ConsumptionPercentage:      tools.ToFloatPointer(100),
						NoUnits:                    conn.NoUnits,
					} ,
					IsMain:         false,
					MainConnection: conn,
				})
			}
		}
		//get estim consumptions from the customer ctype
		if len(estims)==0{
			Trace("find estim consumptions from ctype")
			for d:=range conn.SubConnections {
				if conn.SubConnections[d].CType==nil{
					return nil,errors.New("missing subconnection ctype:")
				}
				ctg,ok:=ctgs[*conn.SubConnections[d].CType ]
				if !ok{
					return nil,errors.New("missing ctype in config:"+*conn.SubConnections[d].CType )
				}
				estims=append(estims,&EstimConsumption{
					SubConnection: &SubConnection{
						CType:                      ctg.CType,
						CTYPE_GROUP:                ctg.CTypeGroupid,
						EstimateConsumptionPerUnit: tools.Divide(ctg.NOOP_ESTIM_CONS,&totalUnitsF),
						ConsumptionPercentage:      tools.ToFloatPointer(100),
						NoUnits:                    conn.NoUnits,
					} ,
					IsMain:         false,
					MainConnection: conn,
				})

			}
		}
		//return data for noop or no meter or no reading
		Trace("Handle NO_METER and None working meters without actual reading")
		for id:=range estims{
			estm:=estims[id]
			if estm.CType==nil{
				return nil,errors.New("missing ctype for the sub connection")
			}
			_noUnits:=tools.DefaultI(estm.NoUnits,1)
			_noUnitsF:=float64(_noUnits)
			sbCons:=tools.Multiply(&_noUnitsF,estm.EstimateConsumptionPerUnit)
			cons[*estims[id].CType]=ConnectionConsumption{
				Ctype:   *estm.CType,
				NoUnits: _noUnits,
				Consump: *sbCons,
				PrDate: defaultPrDate ,
				CrDate:  bilngDate,
				ReadType:READING_TYPE_ESTIMATE,
			}
		}
		return cons,nil
	}
	//handle actual reading
	consump:=reading.GetConsump()
	if consump<0{
		consump=0
	}
	prDate:=tools.DefaultTimeStamp(reading.PrDate,defaultPrDate)
	crDate:=tools.DefaultTimeStamp(reading.CrDate,bilngDate)
	if conn.SubConnections==nil || len(conn.SubConnections)==0{
		Trace("working with meters have actual readings (one connection)")
		_noUnits:=tools.DefaultI(conn.NoUnits,int64(1))
		conn.NoUnits=&_noUnits
		cons[*conn.CType]=ConnectionConsumption{
			Ctype:   *conn.CType,
			NoUnits: _noUnits,
			Consump: consump,
			PrDate:prDate ,
			CrDate:  crDate,
			ReadType:READING_TYPE_ACTUAL,
		}
		return cons,nil
	}
	Trace("working with meters have actual readings (multi connection)")
	for id:=range conn.SubConnections{
		sb:=conn.SubConnections[id]
		if sb!=nil{
			if sb.GetConsumptionPercentage()>=0{
				_noUnits:=tools.DefaultI(sb.NoUnits,int64(1))
				sb.NoUnits=&_noUnits
				ratio:=tools.Divide(sb.ConsumptionPercentage,&handred)
				supConsump:=tools.Multiply(&consump,ratio)
				cons[*sb.CType]=ConnectionConsumption{
					Ctype:   *sb.CType,
					NoUnits: _noUnits,
					Consump: tools.DefaultF(supConsump,0),
					PrDate:prDate ,
					CrDate:  crDate,
					ReadType:READING_TYPE_ACTUAL,
				}
			}
		}
	}
	return cons,nil
}

func TraceFunc(IsTrace bool) func(v ...interface{}) {
	return func(v ...interface{}) {
		if IsTrace {
			log.Println(v...)
		}
	}
}


func TraceFFunc(IsTrace bool) func(str string, v ...interface{}) {
	return func(str string,v ...interface{}) {
		if IsTrace {
			log.Printf(str,v...)
		}
	}
}
