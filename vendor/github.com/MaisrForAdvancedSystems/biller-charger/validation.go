package charge_service

import (
	"context"
	"errors"
	"github.com/MaisrForAdvancedSystems/biller-charger/tools"
	"strings"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
)

/////////////////////vlidations/////////////////
// ValidateCustomer validate
// all ctypes inclides in ctgs in the engin
// all ctype customer service tarrifs founded in engin tarrifes
func (s *BillingChargeService) validateCustomer(con context.Context, cust *Customer) (*bool, error) {
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if cust == nil {
		return nil, errors.New("invalied customer data")
	}
	if cust.Property == nil {
		return nil, errors.New("missing customer properties")
	}
	p := cust.Property
	if p == nil {
		return nil, errors.New("invalied customer properties")
	}
	if p.Services == nil || len(p.Services) == 0 {
		return nil, errors.New("missing propertie services")
	}
	ok:=true
	for sdx := range p.Services {
		srv := p.Services[sdx]
		if srv == nil {
			continue
		}
		okk,err:=s.validateConn(con,srv.Connection,srv.GetServiceType())
		if err!=nil{
			return nil,err
		}
		if okk!=nil{
			ok=ok&&*okk
		}
	}
	return &ok, nil
}

func (s *BillingChargeService) validateConn(con context.Context, conn *Connection,serviceType SERVICE_TYPE) (*bool, error) {
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if conn == nil {
		return nil, errors.New("invalied connection data")
	}
	if conn.SubConnections==nil{
		conn.SubConnections=make([]*SubConnection,0)
	}
	if len(conn.SubConnections)==1{
		return nil,errors.New("الانشطة المختلطة غير صحيحة:الحد الادني لعدد الانشطة 2")
	}
	isHaveMain:=false
	var estim207 float64=0
	for sid:=range conn.SubConnections{
		sb:=conn.SubConnections[sid]
		if sb.CType==nil || sb.CType.CType==nil{
			return nil,errors.New("نشاط فرعي غير معرف")
		}
		if conn.CType!=nil && conn.CType.CType!=nil{
			if *sb.CType.CType==*conn.CType.CType{
				isHaveMain=true
			}
		}
		ok, err := s.isCtgFound(sb.CType.CType, serviceType)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("missing tarrif for " + *sb.CType.CType)
		}
		if sb.EstimateConsumption!=nil && *sb.EstimateConsumption>0{
			estim207=estim207+*sb.EstimateConsumption
		}
		//round
		sb.ConsumptionPercentage=tools.RoundFloatP(sb.ConsumptionPercentage)
		sb.EstimateConsumption=tools.RoundFloatP(sb.EstimateConsumption)
		if sb.NoUnits==nil || *sb.NoUnits<1{
			var on int64=1
			sb.NoUnits=&on
		}
	}
	if len(conn.SubConnections)>1{
		//update main ctype
		if !isHaveMain{
			conn.CType=conn.SubConnections[0].CType
		}
		if estim207>0{
			conn.EstimCons=&estim207
		}

	}else{
		if conn.CType==nil{
			return nil,errors.New("النشاط الرئيسي غير معرف")
		}
		ok, err := s.isCtgFound(conn.CType.CType, serviceType)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("missing tarrif for " + *conn.CType.CType)
		}
	}
	ok := true
	return &ok, nil
}

// IsCtgFound check if the ctype is included into the ctgs
// also check if the tarrif for the service is founded in engin tarrifs
func (s *BillingChargeService) isCtgFound(ctg *string, serviceType SERVICE_TYPE) (bool, error) {
	if s.Ctgs == nil || ctg == nil {
		return false, nil
	}
	_ctg := strings.TrimSpace(*ctg)
	cg, ok := s.Ctgs[_ctg]
	if ok {
		if cg.Tariffs != nil {
			found := false
			for cgIdx := range cg.Tariffs {
				tar := cg.Tariffs[cgIdx]
				if *tar.ServiceType == serviceType {
					found = true
					if !s.isTariffFound(tar.TariffCode) {
						return false, errors.New("missing tarrif for " + *cg.CType + " " + serviceType.String())
					}
				}
			}
			if !found {
				return false, errors.New("missing tarrif for " + *cg.CType)
			}
		}
	}
	return ok, nil
}

// IsTariffFound check if tarrif found in engin tarrifes map
func (s *BillingChargeService) isTariffFound(tar *string) bool {
	if s.Tariffs == nil || tar == nil {
		return false
	}
	_tar := strings.TrimSpace(*tar)
	_, ok := s.Tariffs[_tar]
	return ok
}
