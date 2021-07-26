package consumptions

import (
	"errors"
	"log"
	"time"

	"github.com/MaisrForAdvancedSystems/biller-charger/tools"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

type EstimConsumption struct {
	*SubConnection
	IsMain         bool
	MainConnection *Connection
}

type ConnectionConsumption struct {
	Ctype    string
	NoUnits  int64
	Consump  float64
	PrDate   time.Time
	CrDate   time.Time
	ReadType READING_TYPE
}

func GetConnectionsConsumption(ctgs map[string]*Ctg, conn *Connection, bilngDate time.Time, _reading *Reading, trace bool) (map[string]ConnectionConsumption, error) {
	Trace := TraceFFunc(trace)
	if conn == nil {
		Trace("Invalied request connection is null")
		return nil, errors.New("Invalied Property Connection")
	}
	//raise error if connection status not defined
	if conn.ConnectionStatus == nil {
		Trace("Connection Status Not Defined")
		return nil, errors.New("حالة التوصيلة غير معرفة")
	}
	cons := make(map[string]ConnectionConsumption)
	//reject disconnected connections
	if *conn.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER || *conn.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER {
		Trace("Reject disconnected connections")
		return cons, nil
	}
	//working meters without readings always have no charge
	if *conn.ConnectionStatus == CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER {
		if conn.Meter != nil {
			opStatus := conn.Meter.GetOpStatus()
			log.Println("opStatus", opStatus)
			if opStatus == MeterOperationStatus_WORKING {
				return handleWorkingMeter(ctgs, conn, bilngDate, _reading, trace)
			} else {
				return handleNoneWorking(ctgs, conn, bilngDate, _reading, trace)
			}
		} else {
			return nil, errors.New("بيانات العداد غير موجودة")
		}
	} else {
		return handleNoneWorking(ctgs, conn, bilngDate, _reading, trace)
	}
}

//handleWorkingMeter handle working meters
func handleWorkingMeter(ctgs map[string]*Ctg, conn *Connection, bilngDate time.Time, _reading *Reading, trace bool) (map[string]ConnectionConsumption, error) {
	Trace := TraceFFunc(trace)
	//working meters without readings always have no charge
	if *conn.ConnectionStatus != CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER {
		return nil, errors.New("handling error:connection should be connected with meter")
	}
	cons := make(map[string]ConnectionConsumption)
	if conn.Meter == nil {
		return nil, errors.New("بيانات العداد غير موجودة")
	}
	opStatus := conn.Meter.GetOpStatus()
	if opStatus != MeterOperationStatus_WORKING {
		return handleNoneWorking(ctgs, conn, bilngDate, _reading, trace)
	}
	//handle working meters
	if _reading == nil || _reading.Consump == nil {
		Trace("Reject working meters without readings")
		return cons, nil
	}
	var reading = _reading
	var handred float64 = 100
	defaultPrDate := bilngDate.AddDate(0, -1, 0)
	//handle actual reading
	consump := reading.GetConsump()
	if consump < 0 {
		consump = 0
	}

	prDate := tools.DefaultTimeStamp(reading.PrDate, defaultPrDate)
	crDate := tools.DefaultTimeStamp(reading.CrDate, bilngDate)
	if conn.SubConnections == nil || len(conn.SubConnections) <= 1 {
		Trace("working with meters have actual readings (one connection)")
		_noUnits := tools.DefaultI(conn.NoUnits, int64(1))
		conn.NoUnits = &_noUnits
		cons[*conn.CType.CType] = ConnectionConsumption{
			Ctype:    *conn.CType.CType,
			NoUnits:  _noUnits,
			Consump:  tools.RoundFloat(&consump),
			PrDate:   prDate,
			CrDate:   crDate,
			ReadType: READING_TYPE_ACTUAL,
		}
		return cons, nil
	}
	Trace("working with meters have actual readings (multi connection)")
	for id := range conn.SubConnections {
		sb := conn.SubConnections[id]
		if sb != nil {
			if sb.GetConsumptionPercentage() >= 0 {
				_noUnits := tools.DefaultI(sb.NoUnits, int64(1))
				sb.NoUnits = &_noUnits
				ratio := tools.Divide(sb.ConsumptionPercentage, &handred)
				supConsump := tools.Multiply(&consump, ratio)
				cons[*sb.CType.CType] = ConnectionConsumption{
					Ctype:    *sb.CType.CType,
					NoUnits:  _noUnits,
					Consump:  tools.RoundFloat(supConsump),
					PrDate:   prDate,
					CrDate:   crDate,
					ReadType: READING_TYPE_ACTUAL,
				}
			}
		}
	}
	return cons, nil
}

// handleNoneWorking handling of non working meter or connection without meter
func handleNoneWorking(ctgs map[string]*Ctg, conn *Connection, bilngDate time.Time, _reading *Reading, trace bool) (map[string]ConnectionConsumption, error) {
	Trace := TraceFFunc(trace)
	cons := make(map[string]ConnectionConsumption)
	//estims := make([]*EstimConsumption, 0)
	isHaveCustom := false
	var handred float64 = 100
	defaultPrDate := bilngDate.AddDate(0, -1, 0)
	if conn.SubConnections == nil || len(conn.SubConnections) <= 1 {
		//handle connection without sub connections
		Trace("find estim consumptions for main connection")
		if conn.EstimCons != nil && *conn.EstimCons > 0 {
			Trace("missing estim for main connection")
			_noUnits := tools.DefaultI(conn.NoUnits, int64(1))
			cons[*conn.CType.CType]=ConnectionConsumption{
				Ctype:                 *conn.CType.CType,
				Consump:   tools.RoundFloat(conn.EstimCons),
				NoUnits:  _noUnits,
				PrDate:   defaultPrDate,
				CrDate:   bilngDate,
				ReadType: READING_TYPE_ESTIMATE,
			}
		}
		return cons,nil
	} else {
		//working with customers have custom estim for each sub connection 207
		Trace("find estim consumptions from SubConnection with custom estim (207)")
		for id := range conn.SubConnections {
			sb := conn.SubConnections[id]
			if sb.CType==nil{
				return nil,errors.New("نشاط فرعي غير معرف")
			}
			if sb.CType.CType==nil{
				return nil,errors.New("نشاط فرعي كود غير معرف")
			}
			if sb.EstimateConsumption==nil && sb.ConsumptionPercentage==nil{
				return nil,errors.New("لابد من تحديد الاستهلاك او نسبة الاستهلاك للمشاط الفرعي")
			}
			if sb != nil && sb.EstimateConsumption != nil && *sb.EstimateConsumption>0 {
				isHaveCustom = true
			}
		}
		if isHaveCustom{
			for id := range conn.SubConnections {
				sb := conn.SubConnections[id]
				if sb.EstimateConsumption == nil || *sb.EstimateConsumption<0 {
					return nil,errors.New("نشاط فرعي بدون استهلاك معرف "+*sb.CType.CType)
				}
			}
		}else{
			if conn.EstimCons==nil{
				return nil,errors.New("النشاط الرئيسي غير معرف الاستهلاك")
			}
			for id := range conn.SubConnections {
				//working with customers have custom estim for each sub connection 207
				sb := conn.SubConnections[id]
				if sb == nil {
					return nil, errors.New("missing sub connection data")
				}
				if sb.ConsumptionPercentage == nil {
					return nil, errors.New("النشاط الفرعي لابد ان يكون له نسبة استهلاك")
				}
				_noUnits := tools.DefaultI(sb.NoUnits, int64(1))
				sb.NoUnits=&_noUnits
				ratio := tools.Divide(sb.ConsumptionPercentage, &handred)
				sb.EstimateConsumption = tools.Multiply(conn.EstimCons, ratio)

			}
		}
		for id := range conn.SubConnections {
			sb := conn.SubConnections[id]
			_noUnits := tools.DefaultI(sb.NoUnits, int64(1))
			sb.NoUnits=&_noUnits
			cons[*sb.CType.CType] = ConnectionConsumption{
				Ctype:    *sb.CType.CType,
				NoUnits:  _noUnits,
				Consump:  tools.RoundFloat(sb.EstimateConsumption),
				PrDate:   defaultPrDate,
				CrDate:   bilngDate,
				ReadType: READING_TYPE_ESTIMATE,
			}
		}
		return cons,nil
	}
}

func TraceFunc(IsTrace bool) func(v ...interface{}) {
	return func(v ...interface{}) {
		if IsTrace {
			log.Println(v...)
		}
	}
}

func TraceFFunc(IsTrace bool) func(str string, v ...interface{}) {
	return func(str string, v ...interface{}) {
		if IsTrace {
			log.Printf(str, v...)
		}
	}
}
