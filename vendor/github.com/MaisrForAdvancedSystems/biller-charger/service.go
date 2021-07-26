package charge_service

import (
	"context"
	errors "errors"
	"fmt"
	"log"
	"strings"
	"time"

	rg "github.com/MaisrForAdvancedSystems/biller-charger/charge/regular_charge"
	tr "github.com/MaisrForAdvancedSystems/biller-charger/charge/tariff_calc"
	"github.com/MaisrForAdvancedSystems/biller-charger/consumptions"
	"github.com/MaisrForAdvancedSystems/biller-charger/tools"

	//dbmodels "github.com/MaisrForAdvancedSystems/mas-db-models"

	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"

	"math"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var empty Empty = Empty{}

type BillingChargeService struct {
	Tariffs        map[string]map[time.Time]*Tariff
	Ctgs           map[string]*Ctg
	RegularCharges []*RegularCharge
	IsTrace        bool
	TransCodes     map[string]*TransCode
}
type RegularChargeResponceItem struct {
	Reg       *RegularCharge
	Amount    float64
	TaxAmount *float64
}

// Charge calculate charge for all services
func (s *BillingChargeService) Charge(c context.Context, r *ChargeRequest) (*BillResponce, error) {
	log.Println("Charge Now ....")
	if s.Ctgs == nil {
		return nil, errors.New("Missing Consumtion types lookup")
	}
	if s.Tariffs == nil {
		return nil, errors.New("Missing Traiff lookup")
	}
	if r == nil {
		return nil, errors.New("invalied request")
	}
	if r.Customer == nil {
		return nil, errors.New("invalied request:missing customer")
	}
	if r.Setting == nil || r.Setting.BilingDate == nil {
		return nil, errors.New("invalied bilng date")
	}
	yr := r.GetSetting().GetBilingDate().AsTime().Year()
	if yr > 2100 || yr < 1900 {
		return nil, errors.New(fmt.Sprintf("invalied bilng date %v", r.Setting.GetBilingDate().AsTime().String()))
	}
	isCustValied, err := s.validateCustomer(c, r.Customer)
	if err != nil {
		return nil, err
	}
	if isCustValied == nil || !*isCustValied {
		s.trace("valid result ", isCustValied)
		return nil, errors.New("customer data is not valied ")
	}
	if r.Customer.Property == nil {
		return nil, nil
	}
	cust_services := r.Customer.Property.Services
	if cust_services == nil || len(cust_services) == 0 {
		return nil, errors.New("customer have no services")
	}
	chargedServicesType := r.Services
	if chargedServicesType == nil || len(chargedServicesType) == 0 {
		return nil, errors.New("no services definded for charge")
	}
	services := make([]*Service, 0)
	for _, chrg_type := range chargedServicesType {
		added := false
		for idx := range cust_services {
			_srv := cust_services[idx]
			if _srv.GetServiceType() == chrg_type {
				services = append(services, _srv)
				added = true
				break
			}
		}
		if !added {
			//copy the main service with new service type
			srv := *cust_services[0]
			srv.ServiceType = &chrg_type
			services = append(services, &srv)
		}
	}
	if len(services) == 0 {
		return nil, errors.New("no services definded match customer services")
	}
	//validate readings
	serviceRdgs := make(map[SERVICE_TYPE]*Reading)
	if r.ServicesReadings != nil && len(r.ServicesReadings) > 0 {
		for idx := range services {
			srv := services[idx]
			for sx := range r.ServicesReadings {
				rdg := r.ServicesReadings[sx]
				if rdg.ServiceType == nil {
					continue
				}
				if rdg.Reading != nil {
					if rdg.Reading.Consump == nil {
						return nil, errors.New("الاستهلاك غير معرف")
					}
					if rdg.Reading.ReadType == nil {
						return nil, errors.New("نوع الاستهلاك غير معرف")
					}
					if *rdg.Reading.ReadType == READING_TYPE_ACTUAL {
						if rdg.Reading.CrReading == nil {
							return nil, errors.New("القراءة الحالية غير معرفة")
						}
						if rdg.Reading.PrReading == nil {
							return nil, errors.New("القراءة السابقة غير معرفة")
						}
					}
				}
				if *rdg.ServiceType == srv.GetServiceType() {
					serviceRdgs[srv.GetServiceType()] = rdg.Reading
					break
				}
			}
		}
	}

	bill := &Bill{
		PaymentNo:        r.Setting.PaymentNo,
		BilngDate:        r.Setting.BilingDate,
		FTransactions:    []*FinantialTransaction{},
		ServicesReadings: r.ServicesReadings,
		Customer:         r.Customer,
	}
	resp := BillResponce{Bills: []*Bill{bill}}
	bill.FTransactions = make([]*FinantialTransaction, 0)
	isConnected := true
	for idx := range services {
		log.Println(idx)
		srv := services[idx]
		rdg, _ := serviceRdgs[srv.GetServiceType()]
		s.trace(srv.GetServiceType())
		if srv.Connection != nil && srv.Connection.ConnectionStatus != nil {
			if *srv.Connection.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER || *srv.Connection.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER {
				isConnected = false
			}
		}
		sTrans, err := s.calcForService(r.Setting, srv, rdg, r.Customer)
		if err != nil {
			s.trace(err)
			return nil, err
		}
		s.trace(sTrans)
		if sTrans != nil {
			bill.FTransactions = append(bill.FTransactions, sTrans...)
		}
	}
	if isConnected {
		cycle_length := r.Setting.CycleLength
		regsTrans, err := s.regCharge(r.Customer, r.Setting.BilingDate.AsTime(), cycle_length)
		if err != nil {
			return nil, err
		}
		bill.FTransactions = append(bill.FTransactions, regsTrans...)
	}

	var zero float64 = 0
	ser1 := services[0]

	// Adding CreditCode - DebtCode - InstallCode From OldTransactions & Set All Transaction Calculated Edit = false
	if r.OldBill != nil {
		var mainCtype Ctg = *ser1.Connection.CType
		mainCtg, _ := s.Ctgs[*mainCtype.CType]
		if r.OldBill.FTransactions != nil {
			s.ExtraCreditDebtInst(ser1, bill, r, zero, mainCtg, ser1, ser1.GetConnection().NoUnits)
		}
		bill.PaymentNo = r.OldBill.PaymentNo
		//bill.RecalcFormDate = r.OldBill.RecalcFormDate
		//bill.RecalcFormNo = r.OldBill.RecalcFormNo
	}

	var taxAmt float64 = 0
	var taxNoUnits *int64
	var taxCtype *Ctg
	var totalAmt float64 = 0
	for _, t := range bill.FTransactions {
		if t.Amount != nil {
			totalAmt = totalAmt + *t.Amount
		}
		t.Amount = roundTo(t.Amount, 0.001)
		t.TaxAmount = roundTo(t.TaxAmount, 0.001)
		if t.TaxAmount != nil && *t.TaxAmount > 0 {
			taxAmt = taxAmt + *t.TaxAmount
			taxNoUnits = t.NoUnits
			taxCtype = t.Ctype
			t.TaxAmount = &zero
		}
		t.Description = s.getTransCodeDescription(t.Code)

	}
	taxCode, _ := s.TransCodes["TAX_AMT"]
	roundCode, _ := s.TransCodes["ROUND_AMT"]

	if taxAmt > 0 {
		totalAmt = totalAmt + taxAmt
		rndTax := roundTo(&taxAmt, 0.001)
		bill.FTransactions = append(bill.FTransactions, &FinantialTransaction{
			Code:           taxCode.Code,
			EffDate:        r.Setting.BilingDate,
			BilngDate:      r.Setting.BilingDate,
			Amount:         rndTax,
			TaxAmount:      &zero,
			DiscountAmount: &zero,
			Ctype:          taxCtype,
			PropRef:        r.Customer.Property.PropRef,
			ServiceType:    ser1.ServiceType,
			NoUnits:        taxNoUnits,
			Description:    taxCode.Description,
		})
	}
	//round
	pRound := roundTo(&totalAmt, 0.5)
	roundAmt := *pRound
	log.Printf("total amount:%v roundAmount:%v", totalAmt, roundAmt)
	if totalAmt > roundAmt {
		roundAmt = roundAmt + float64(0.5)
	}
	_dif := roundAmt - totalAmt
	roundDif := roundTo(&_dif, 0.001)
	log.Printf("total amount:%v roundAmount:%v dif:%v", totalAmt, roundAmt, roundDif)
	if *roundDif >= 0.001 {
		bill.FTransactions = append(bill.FTransactions, &FinantialTransaction{
			Code:           roundCode.Code,
			EffDate:        r.Setting.BilingDate,
			BilngDate:      r.Setting.BilingDate,
			Amount:         roundDif,
			TaxAmount:      &zero,
			DiscountAmount: &zero,
			Ctype:          taxCtype,
			PropRef:        r.Customer.Property.PropRef,
			ServiceType:    ser1.ServiceType,
			NoUnits:        ser1.Connection.NoUnits,
			Description:    roundCode.Description,
		})
	}
	return &resp, nil
}

func (s *BillingChargeService) ExtraCreditDebtInst(service *Service, bill *Bill, r *ChargeRequest, zero float64, mainctype *Ctg, ser1 *Service, taxNoUnits *int64) {
	for idx := range bill.FTransactions {
		newTrans := bill.FTransactions[idx]
		newTrans.Editable = tools.ToBoolPointer(false)
	}
	CreditCode, ok := s.TransCodes["CRDT_AMT"]
	if !ok {
		return
	}
	DebtCode := s.TransCodes["DBT_AMT"]
	InstallCode := s.TransCodes["INSTALLS_AMT"]
	//multi charges
	ctgs := make(map[*Ctg]float64) //ctgs with consumption ratio
	if len(service.Connection.SubConnections) > 1 {
		for id := range service.Connection.SubConnections {
			subCon := service.Connection.SubConnections[id]
			if subCon.GetConsumptionPercentage() > 0 {
				ctgs[subCon.CType] = math.Round(10*subCon.GetConsumptionPercentage()) / 1000
			}
		}
	} else {
		ctgs[service.Connection.CType] = 1
	}
	for idx := range r.OldBill.FTransactions {
		oldTrans := r.OldBill.FTransactions[idx]
		if oldTrans != nil && oldTrans.Amount != nil && *oldTrans.Amount != 0 && (*oldTrans.Code == *InstallCode.Code || *oldTrans.Code == *DebtCode.Code || *oldTrans.Code == *CreditCode.Code) {
			oldTrans.Editable = tools.ToBoolPointer(true)
			for c, ratio := range ctgs {
				if c.CType != nil {
					ctype := *c.CType
					desc := ""
					if c.Description != nil {
						desc = *c.Description
					}
					//bill.FTransactions = append(bill.FTransactions, oldTrans)
					tr := *oldTrans
					nwAmt := *oldTrans.Amount * ratio
					tr.Amount = &nwAmt
					tr.Ctype = &Ctg{CType: &ctype, Description: &desc}
					bill.FTransactions = append(bill.FTransactions, &tr)
				}
			}
		}
	}
}

//calc for service and may be not used
func (s *BillingChargeService) calcForService(setting *ChargeSetting, service *Service, reading *Reading, cust *Customer) ([]*FinantialTransaction, error) {
	var conn = service.Connection
	var empty = make([]*FinantialTransaction, 0)
	if conn == nil {
		s.tsraceF("no service connection %v", service.GetServiceType())
		return empty, nil
	}
	if *conn.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER {
		return empty, nil
	}
	if *conn.ConnectionStatus == CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER {
		return empty, nil
	}
	cons, err := consumptions.GetConnectionsConsumption(s.Ctgs, service.Connection, setting.BilingDate.AsTime(), reading, s.IsTrace)
	if err != nil {
		return nil, err
	}
	trans := make([]*FinantialTransaction, 0)
	var crReading *float64 = nil
	var prReading *float64 = nil
	var MeterType *string = nil
	var MeterRef *string = nil
	serviceType := service.GetServiceType()
	bilngDate := timestamppb.New(setting.BilingDate.AsTime())
	if reading != nil {
		crReading = reading.CrReading
		prReading = reading.PrReading
	}
	if conn.Meter != nil {
		MeterType = conn.Meter.MeterType
		MeterRef = conn.Meter.MeterRef
	}
	propref := ""
	if cust != nil && cust.Property != nil && cust.Property.PropRef != nil {
		propref = *cust.Property.PropRef
	}
	for idc := range cons {
		c := cons[idc]
		ctg, ok := s.Ctgs[c.Ctype]
		if !ok {
			return nil, errors.New("missing ctype :" + c.Ctype)
		}
		transCode := "MISSING_CODE"
		var taxPercentage float64 = 0
		var discountPercentage float64 = 0
		if ctg.Tariffs == nil || len(ctg.Tariffs) == 0 {
			return nil, errors.New("missing tariff for ctype :" + c.Ctype)
		}

		if ctg.Tariffs != nil {
			tarifCode := ""
			isZeroTariff := false
			for _, srvTar := range ctg.Tariffs {
				if *srvTar.ServiceType == *service.ServiceType {
					if srvTar.TariffCode == nil {
						return nil, errors.New("missing tariff code:")
					}
					tarifCode = *srvTar.TariffCode
					if srvTar.TransCode != nil {
						transCode = *srvTar.TransCode
					} else {
						return nil, errors.New("missing trans code tariff :" + c.Ctype)
					}
					if srvTar.IsZeroTarif != nil {
						isZeroTariff = *srvTar.IsZeroTarif
					}
					if isZeroTariff {
						break
					} else {
						if srvTar.TariffCode == nil {
							return nil, errors.New("missing tariff :" + c.Ctype)
						}
					}

					if srvTar.TaxPercentage != nil && *srvTar.TaxPercentage > 0 {
						taxPercentage = *srvTar.TaxPercentage
					}
					if srvTar.DiscountPercentage != nil && *srvTar.DiscountPercentage > 0 {
						discountPercentage = *srvTar.DiscountPercentage
					}
					break
				}
			}

			if taxPercentage > 100 {
				taxPercentage = 100
			}
			if taxPercentage < 0 {
				taxPercentage = 0
			}
			if discountPercentage > 100 {
				discountPercentage = 100
			}
			if discountPercentage < 0 {
				discountPercentage = 0
			}
			if tarifCode == "" && !isZeroTariff {
				return nil, errors.New("missing tariff for ctype :" + c.Ctype)
			}
			var tariff *Tariff = nil
			if !isZeroTariff {
				tariff, err = s.findBestTarrifMatch(tarifCode, setting.GetBilingDate().AsTime())
				if err != nil {
					return nil, err
				}
			}
			chargAmt, err := tr.Calc(service, setting, c.NoUnits, c.Consump, c.ReadType, &c.CrDate, &c.PrDate, tariff, isZeroTariff)
			if err != nil {
				return nil, err
			}
			if chargAmt != nil && *chargAmt >= 0 {
				taxAmount := taxPercentage * (*chargAmt) / float64(100)
				disAmount := discountPercentage * (*chargAmt) / float64(100)
				amount := *chargAmt - disAmount
				trans = append(trans, &FinantialTransaction{
					Code:           &transCode,
					EffDate:        bilngDate,
					BilngDate:      bilngDate,
					Amount:         &amount,
					TaxAmount:      &taxAmount,
					DiscountAmount: &disAmount,
					Ctype:          ctg,
					PropRef:        &propref,
					ServiceType:    &serviceType,
					NoUnits:        &c.NoUnits,
					MTransaction: &MeasuredTransaction{
						CrReading: crReading,
						PrReading: prReading,
						Consump:   &c.Consump,
						ReadType:  &c.ReadType,
						MeterType: MeterType,
						MeterRef:  MeterRef,
					},
				})
			}
			if tariff != nil && tariff.ExtraFees != nil {
				for eid := range tariff.ExtraFees {
					extra := tariff.ExtraFees[eid]
					if extra != nil {
						if extra.TransCode == nil {
							return nil, errors.New("missing transcode for extra tariff :" + c.Ctype)
						}
						var amountPerMeter float64 = 0
						var fixed float64 = 0
						if extra.AmountPerMeter != nil {
							amountPerMeter = *extra.AmountPerMeter
						}
						if extra.FixedAmount != nil {
							fixed = *extra.FixedAmount
						}
						amount := fixed + (amountPerMeter * c.Consump)
						var zero float64 = 0
						taxAmount := zero
						if extra.TaxPercentage != nil {
							if *extra.TaxPercentage > 0 && *extra.TaxPercentage <= 100 {
								taxAmount = amount * (*extra.TaxPercentage) / float64(100)
							}
						}
						//if extra.t
						trans = append(trans, &FinantialTransaction{
							Code:           extra.TransCode,
							EffDate:        bilngDate,
							BilngDate:      bilngDate,
							Amount:         &amount,
							TaxAmount:      &taxAmount,
							DiscountAmount: &zero,
							Ctype:          ctg,
							PropRef:        &propref,
							ServiceType:    &serviceType,
							NoUnits:        &c.NoUnits,
							MTransaction: &MeasuredTransaction{
								CrReading: crReading,
								PrReading: prReading,
								Consump:   &c.Consump,
								ReadType:  &c.ReadType,
								MeterType: MeterType,
								MeterRef:  MeterRef,
							},
						})
					}
				}
			}
		}
	}

	return trans, nil
}

func (s *BillingChargeService) getTransCodeDescription(code *string) *string {
	str := "UnkownTransCode"
	if code == nil {
		return &str
	}
	desc, ok := s.TransCodes[*code]
	if !ok || desc.Description == nil {
		return code
	}
	return desc.Description
}

func (s *BillingChargeService) findBestTarrifMatch(tarId string, bilngDate time.Time) (*Tariff, error) {
	log.Println("find best tariff for " + tarId + " @" + bilngDate.String())
	if s.Tariffs == nil || len(s.Tariffs) == 0 {
		return nil, errors.New("no tariffes defined")
	}
	tarifs, ok := s.Tariffs[tarId]
	if !ok {
		s.trace(s.Tariffs)
		return nil, errors.New("missing tariff " + tarId)
	}
	var bestDown *time.Time = nil
	for dt := range tarifs {
		if dt.Before(bilngDate) || dt == bilngDate {
			if bestDown != nil {
				if bestDown.Before(dt) {
					iDatee := dt
					bestDown = &iDatee
				}
			} else {
				iDatee := dt
				bestDown = &iDatee
			}
		}
	}
	log.Println("bestDown:", bestDown)
	if bestDown == nil {
		return nil, errors.New("miisng tariff for " + tarId + " @" + bilngDate.String())
	}
	return tarifs[*bestDown], nil
}

func (s *BillingChargeService) regCharge(customer *Customer, bilngDate time.Time, cycle_length *int64) ([]*FinantialTransaction, error) {
	if s.RegularCharges == nil || len(s.RegularCharges) == 0 {
		return nil, nil
	}
	rs := make([]*FinantialTransaction, 0)
	for id := range s.RegularCharges {
		regCh := s.RegularCharges[id]
		//log.Println("service recharge ", regCh.ChargeType.String(), *regCh.RegularChargeId)
		if regCh == nil {
			return nil, errors.New("invalied regular charge setup")
		}
		chTrans, err := rg.CalcCharge(regCh, customer, bilngDate, cycle_length, s.Ctgs)
		if err != nil {
			return nil, err
		}
		if chTrans != nil {
			rs = append(rs, chTrans...)
		}
	}
	return rs, nil
}

func (s *BillingChargeService) printInfo(cst *Customer) {
	if !s.IsTrace {
		return
	}
	if cst == nil {
		log.Println("no customer")
		return
	}
	if cst.Property == nil {
		log.Printf("customer has no properties %v", cst.Custkey)
		return
	}
	prop := cst.Property
	if prop.Services == nil {
		log.Printf("invalied customer services %v", cst.Custkey)
		return
	}

	if len(prop.Services) == 0 {
		log.Printf("customer has no services %v", cst.Custkey)
		return
	}

	for _, s := range prop.Services {
		log.Printf("Service:%v", s.GetServiceType())
		if s.Connection == nil {
			log.Printf("service:%v has no connection %v", s.GetServiceType(), cst.GetCustkey())
			continue
		}
		conn := s.Connection
		log.Printf("status %v ", conn.GetConnectionStatus())
		log.Printf("estim cons %v ", conn.GetEstimCons())
		log.Printf("main ctype %v", conn.GetCType())
		log.Printf("main ctype group %v", conn.CType.CTypeGroupid)
		if conn.Meter == nil {
			log.Printf("%v connection has no meter", s.GetServiceType())
			continue
		}
		meter := conn.Meter
		log.Printf("meter diam:%v", meter.Diameter)
		log.Printf("meter status:%v", meter.OpStatus)
	}
}

func (s *BillingChargeService) getCtype(ct string) (*Ctg, error) {
	ct = strings.TrimSpace(ct)
	cg, ok := s.Ctgs[ct]
	if !ok {
		return nil, errors.New("missing ctype:" + ct)
	}
	return cg, nil
}

func (s *BillingChargeService) getAllCtype(cust *Customer, serviceType SERVICE_TYPE) ([]*Ctg, error) {
	if cust == nil || cust.Property == nil || cust.Property.Services == nil {
		return []*Ctg{}, nil
	}
	if len(cust.Property.Services) == 0 {
		return []*Ctg{}, nil
	}
	for _, sv := range cust.Property.Services {
		if sv != nil && sv.GetServiceType() == serviceType && sv.Connection != nil {
			if sv.Connection.SubConnections == nil || len(sv.Connection.SubConnections) <= 1 {
				cg, ok := s.Ctgs[*sv.Connection.CType.CType]
				if !ok {
					return nil, errors.New("missing ctype:" + *sv.Connection.CType.CType)
				}
				return []*Ctg{cg}, nil
			} else {
				cgs := make([]*Ctg, 0)
				for _, sb := range sv.Connection.SubConnections {
					cg, ok := s.Ctgs[*sb.CType.CType]
					if !ok {
						return nil, errors.New("missing ctype:" + *sb.CType.CType)
					}
					cgs = append(cgs, cg)
				}
				return cgs, nil
			}
		}
	}
	return []*Ctg{}, nil
}

func getNoUnits(cust *Customer, serviceType SERVICE_TYPE) *int64 {
	if cust == nil || cust.Property == nil || cust.Property.Services == nil {
		return nil
	}
	if len(cust.Property.Services) == 0 {
		return nil
	}
	for _, sv := range cust.Property.Services {
		if sv != nil && sv.GetServiceType() == serviceType && sv.Connection != nil {
			return sv.Connection.NoUnits
		}
	}
	return nil
}

func getMainCtype(cust *Customer, serviceType SERVICE_TYPE) *string {
	if cust == nil || cust.Property == nil || cust.Property.Services == nil {
		return nil
	}
	if len(cust.Property.Services) == 0 {
		return nil
	}
	for _, sv := range cust.Property.Services {
		if sv != nil && sv.GetServiceType() == serviceType && sv.Connection != nil {
			return sv.Connection.CType.CType
		}
	}
	return nil
}

func roundTo(f *float64, _to float64) *float64 {
	if f == nil {
		return nil
	}
	if math.IsNaN(*f) {
		return nil
	}
	if math.IsInf(*f, 0) {
		return nil
	}
	if _to <= 0 {
		rVal := math.Round(*f)
		return &rVal
	}
	to := 1 / _to
	rVal := math.Round((*f)*float64(to)) / float64(to)
	log.Printf("round :%v %v %v", *f, rVal, to)
	return &rVal
}
