package providers

import (
	"context"
	"database/sql"
	"encoding/base64"
	"math"
	"strconv"

	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/HassanAbdelzaher/lama"
	"github.com/MaisrForAdvancedSystems/biller-mas-provider/dbcontext"
	"github.com/MaisrForAdvancedSystems/biller-mas-provider/tools"
	. "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	dbmodels "github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
 rpc Info(Empty) returns (ServiceInfo){}
  rpc GetCustomerByCustkey (Key) returns (Customer) {}
  rpc GetCustomersByBillgroup (Key) returns (CustomersList) {}
  rpc GetLoockup (Entity) returns (LookUpsResponce) {}
  rpc GetBillByCustkey (GetBillRequest) returns (BillResponce) {}*/

type DataProvider struct {
	InfoProvider
}

var db *lama.Lama = dbcontext.DbConnPool

func GetUser(_username *string) (*dbmodels.USERS, error) {
	if _username == nil {
		return nil, errors.New("اسم الدخول غير صحيح")
	}
	username := *_username
	var users []*dbmodels.USERS
	err := db.Model(dbmodels.USERS{}).Where(dbmodels.USERS{USER_NAME: &username}).Find(&users)
	if err != nil {
		return nil, err
	}
	if users == nil || len(users) == 0 {
		return nil, errors.New("اسم الدخول او كلمة المرور غير صحيحة")
	}
	if len(users) > 1 {
		return nil, errors.New(username + "تكرار باسم الدخول بقواعد بيانات المستخدمين ")
	}
	return users[0], nil
}

func GetStation(_stationNo *int64) (*dbmodels.STATIONS, error) {
	if _stationNo == nil {
		return nil, errors.New("رقم الفرع غير صحيح")
	}
	stationNo := int32(*_stationNo)
	var stations []*dbmodels.STATIONS
	//station no may be zero so we must use map for filter
	err := db.Model(dbmodels.STATIONS{}).Where(map[string]interface{}{"STATION_NO": stationNo}).Find(&stations)
	if err != nil {
		return nil, err
	}
	if stations == nil || len(stations) == 0 {
		return nil, errors.New("رقم الفرع غير صحيح")
	}
	if len(stations) > 1 {
		return nil, errors.New(fmt.Sprintf("%v  تكرار برقم الفرع ", stationNo))
	}
	return stations[0], nil
}

func IsHq(st *dbmodels.STATIONS) bool {
	if st == nil || st.IS_HEADQUARTERS == nil {
		return false
	}
	if *st.IS_HEADQUARTERS == int64(1) {
		return true
	}
	return false
}

func (s *DataProvider) Login(ctx context.Context, rq *LoginRequest) (resp *LoginResponce, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	user, err := GetUser(rq.Username)
	if err != nil {
		return nil, err
	}
	pssword := user.PASSWORD
	if pssword == "" {
		return nil, errors.New("كلمة المرور غير صحيحة")
	}
	_decodePassword, err := base64.StdEncoding.Strict().DecodeString(pssword)
	if err != nil {
		return nil, err
	}
	if string(_decodePassword) != *rq.Password {
		return nil, errors.New("كلمة المرور غير صحيحة")
	}
	if user.EDAMS_RECALC_NEW == nil || !*user.EDAMS_RECALC_NEW {
		return nil, errors.New("المستخدم لا يمتلك صلاحية كافية")
	}
	tru := true
	return &LoginResponce{
		Succssed: &tru,
	}, nil
}

func (s *DataProvider) GetBillsByCustkey(ctx context.Context, rq *GetBillRequest) (resp *BillResponce, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	log.Printf("write data by :%v", ctx.Value("username"))
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	user, err := GetUser(&username)
	if err != nil {
		return nil, err
	}
	ustation, err := GetStation(user.STATION_NO)
	if err != nil {
		return nil, err
	}
	if rq.Custkey == nil {
		return nil, errors.New("invalied request missing custkey")
	}
	query := fmt.Sprintf(`select CUSTKEY,CYCLE_ID from dbo.HAND_MH_ST where custkey ='%s' and BILNG_DATE>=@startdate and BILNG_DATE <=@enddate order by BILNG_DATE`, *rq.Custkey)
	arc_query := fmt.Sprintf(`select CUSTKEY,CYCLE_ID from dbo.ARC_HAND_MH_ST where custkey ='%s' and BILNG_DATE>=@startdate and BILNG_DATE <=@enddate order by BILNG_DATE`, *rq.Custkey)
	var data []*dbmodels.HAND_MH_ST
	var arc_data []*dbmodels.HAND_MH_ST
	from := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC)
	if rq.BilngDateFrom != nil {
		from = rq.BilngDateFrom.AsTime().Add(-2 * time.Hour)
		if from.Year() > 2100 || from.Year() < 1900 {
			return nil, errors.New(fmt.Sprintf("invalied date:%v", from.String()))
		}
	}
	if rq.BilngDateTo != nil {
		to = rq.BilngDateTo.AsTime().Add(2 * time.Hour)
		if to.Year() > 2100 || to.Year() < 1900 {
			return nil, errors.New(fmt.Sprintf("invalied date:%v", to.String()))
		}
	}
	err = db.DB.Unsafe().Select(&data, query, sql.Named("startdate", from), sql.Named("enddate", to))
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	err = db.DB.Unsafe().Select(&arc_data, arc_query, sql.Named("startdate", from), sql.Named("enddate", to))
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	resp = &BillResponce{Bills: make([]*Bill, 0)}
	if arc_data != nil {
		for id := range arc_data {
			d := arc_data[id]
			if ustation != nil && d.STATION_NO != nil && !IsHq(ustation) {
				if *d.STATION_NO != ustation.STATION_NO {
					return nil, errors.New("الفاتورة تخص فرع اخر")
				}
			}
			if d.CYCLE_ID == nil {
				return nil, errors.New(fmt.Sprintf("missing cycle id:%s", d.CUSTKEY))
			}
			bl, err := s.getBill(ctx, *rq.Custkey, *d.CYCLE_ID)
			if err != nil {
				return nil, err
			}
			resp.Bills = append(resp.Bills, bl)
		}
	}
	for id := range data {
		d := data[id]
		if ustation != nil && d.STATION_NO != nil && !IsHq(ustation) {
			if *d.STATION_NO != ustation.STATION_NO {
				return nil, errors.New("الفاتورة تخص فرع اخر")
			}
		}
		if d.CYCLE_ID == nil {
			return nil, errors.New(fmt.Sprintf("missing cycle id:%s", d.CUSTKEY))
		}
		bl, err := s.getBill(ctx, *rq.Custkey, *d.CYCLE_ID)
		if err != nil {
			return nil, err
		}
		resp.Bills = append(resp.Bills, bl)
	}

	return resp, nil
}
func (s *DataProvider) GetBillsByFormNo(ctx context.Context, rq *GetBillRequest) (resp *BillResponce, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	log.Printf("write data by :%v", ctx.Value("username"))
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	if rq.FormNo == nil {
		return nil, errors.New("invalied request missing form_no")
	}
	formNo, err := strconv.ParseInt(*rq.FormNo, 0, 64)
	if err != nil {
		return nil, errors.New("invalied form no")
	}
	user, err := GetUser(&username)
	if err != nil {
		return nil, err
	}
	ustation, err := GetStation(user.STATION_NO)
	if err != nil {
		return nil, err
	}
	var can_req dbmodels.CANCELLED_REQUEST
	err = db.Where(dbmodels.CANCELLED_REQUEST{FORM_NO: formNo}).First(&can_req)
	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf(" %v لا يوجد طلب بالرقم ", formNo))
	}
	if err != nil {
		return nil, err
	}
	if can_req.COMMENT != nil {
		log.Print("comment  CANCELLED_REQUEST", *can_req.COMMENT)
	}
	log.Print("CANCELLED_REQUEST", can_req)
	querycan_reqs := fmt.Sprintf(`select * from dbo.CANCELLED_REQUESTS where FORM_NO=%v `, formNo)
	var can_reqs []*dbmodels.CANCELLED_REQUEST
	err = db.DB.Unsafe().Select(&can_reqs, querycan_reqs)
	if err != nil {
		return nil, err
	}
	if len(can_reqs) > 0 {
		if can_reqs[0].COMMENT != nil {
			log.Print("comment  CANCELLED_REQUEST", *can_reqs[0].COMMENT)
		}
		log.Print("CANCELLED_REQUEST", can_reqs[0])
	}

	// cancelled bills
	var can_bills []*dbmodels.CANCELLED_BILL
	err = db.Where(dbmodels.CANCELLED_BILL{FORM_NO: formNo}).Find(&can_bills)
	if err != nil {
		return nil, err
	}
	log.Print("Count CANCELLED_BILL ", len(can_bills))

	query := fmt.Sprintf(`select h.* from dbo.CANCELLED_BILLS c,HAND_MH_ST h where h.CUSTKEY=c.CUSTKEY and h.PAYMENT_NO=c.PAYMENT_NO and FORM_NO=%v order by h.BILNG_DATE `, formNo)
	arc_query := fmt.Sprintf(`select h.* from dbo.CANCELLED_BILLS c,ARC_HAND_MH_ST h where h.CUSTKEY=c.CUSTKEY and h.PAYMENT_NO=c.PAYMENT_NO and FORM_NO=%v order by h.BILNG_DATE`, formNo)
	var data []*dbmodels.HAND_MH_ST
	var arc_data []*dbmodels.HAND_MH_ST
	err = db.DB.Unsafe().Select(&data, query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	err = db.DB.Unsafe().Select(&arc_data, arc_query)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	resp = &BillResponce{
		Bills: make([]*Bill, 0),
		RecalcForm: &CancelledRequest{
			FormNo:   &can_req.FORM_NO,
			FormDate: create_timestamp(can_req.REQUEST_DATE),
			Status:   can_req.STATUS,
			Comment:  can_req.COMMENT,
		},
	}

	if arc_data != nil {
		for id := range arc_data {
			d := arc_data[id]
			if ustation != nil && d.STATION_NO != nil && !IsHq(ustation) {
				if *d.STATION_NO != ustation.STATION_NO {
					return nil, errors.New("الفاتورة تخص فرع اخر")
				}
			}
			log.Println("arc", d.CUSTKEY, d.CYCLE_ID)
			if d.CYCLE_ID == nil {
				return nil, errors.New(fmt.Sprintf("missing cycle id:%s", d.CUSTKEY))
			}
			bl, err := s.getBill(ctx, d.CUSTKEY, *d.CYCLE_ID)
			if err != nil {
				return nil, err
			}
			for idx := range can_bills {
				can_bill := can_bills[idx]
				if can_bill.PAYMENT_NO == *bl.PaymentNo {
					bl.Comment = can_bill.COMMENT
					break
				}
			}
			resp.Bills = append(resp.Bills, bl)
		}
	}

	for id := range data {
		d := data[id]
		if ustation != nil && d.STATION_NO != nil && !IsHq(ustation) {
			if *d.STATION_NO != ustation.STATION_NO {
				return nil, errors.New("الفاتورة تخص فرع اخر")
			}
		}
		log.Println(d.CUSTKEY, d.CYCLE_ID)
		if d.CYCLE_ID == nil {
			return nil, errors.New(fmt.Sprintf("missing cycle id:%s", d.CUSTKEY))
		}
		bl, err := s.getBill(ctx, d.CUSTKEY, *d.CYCLE_ID)
		if err != nil {
			return nil, err
		}
		for idx := range can_bills {
			can_bill := can_bills[idx]
			if can_bill.PAYMENT_NO == *bl.PaymentNo {
				bl.Comment = can_bill.COMMENT
				break
			}
		}
		resp.Bills = append(resp.Bills, bl)
	}
	return resp, nil
}

func (s *DataProvider) GetLoockup(cn context.Context, eType *Entity) (resp *LookUpsResponce, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	lookups := make([]*LookUp, 0)
	resp = &LookUpsResponce{LookUps: lookups}
	if eType.GetEntityType() == ENTITY_TYPE_METER_DIAMETER || eType.GetEntityType() == ENTITY_TYPE_CONNECTION_DIAMETER {
		for i := 2; i < 30; i++ {
			dm := 5 * i
			if i > 20 {
				dm = 100 + 50*(i-20)
			}
			code := fmt.Sprintf("%v", dm)
			describe := fmt.Sprintf("قطر %v مم", dm)
			lk := &LookUp{Code: &code, Description: &describe}
			lookups = append(lookups, lk)
		}
		return &LookUpsResponce{LookUps: lookups}, nil
	}
	if eType.GetEntityType() == ENTITY_TYPE_CTYPE_GROUP {
		var ctgGrops []*dbmodels.CTG_CONSUMPTIONTYPEGRPS
		err := db.Model(dbmodels.CTG_CONSUMPTIONTYPEGRPS{}).Find(&ctgGrops)
		if err != nil {
			return nil, err
		}
		if ctgGrops == nil {
			return nil, errors.New("no data found")
		}
		for id := range ctgGrops {
			ctg := ctgGrops[id]
			if ctg.DESCRIPTION == nil {
				continue
			}
			code := ctg.CTYPEGRP_ID
			describe := *ctg.DESCRIPTION
			lk := &LookUp{Code: &code, Description: &describe}
			lookups = append(lookups, lk)
		}
		resp.LookUps = lookups

		return resp, nil
	}
	if eType.GetEntityType() == ENTITY_TYPE_CTYPE {
		var ctgGrops []*dbmodels.CTG_CONSUMPTIONTYPES
		query := `select * from  CTG_CONSUMPTIONTYPES`
		err := db.DB.Unsafe().Select(&ctgGrops, query)
		if err != nil {
			return nil, err
		}
		if ctgGrops == nil {
			return nil, errors.New("no data found")
		}
		for id := range ctgGrops {
			ctg := ctgGrops[id]
			if ctg.DESCRIPTION == nil {
				continue
			}
			code := ctg.CTYPE_ID
			describe := *ctg.DESCRIPTION
			lk := &LookUp{Code: &code, Description: &describe}
			lookups = append(lookups, lk)
		}
		resp.LookUps = lookups
		return resp, nil
	}
	return nil, errors.New("Not Supported")
}

func (s *DataProvider) GetCtgs(cn context.Context, eType *Empty) (resp *CtgsResponce, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	lookups := make([]*Ctg, 0)
	resp = &CtgsResponce{Ctgs: lookups}
	var ctgGrops []*dbmodels.CTG_CONSUMPTIONTYPES
	query := `select * from  CTG_CONSUMPTIONTYPES`
	err = db.DB.Unsafe().Select(&ctgGrops, query)
	if err != nil {
		return nil, err
	}
	if ctgGrops == nil {
		return nil, errors.New("no data found")
	}
	for id := range ctgGrops {
		ctg := ctgGrops[id]
		if ctg.CTYPE_ID == "" {
			continue
		}
		if ctg.DESCRIPTION == nil {
			ctg.DESCRIPTION = &ctg.CTYPE_ID
		}
		lk := &Ctg{
			CType:            &ctg.CTYPE_ID,
			CTypeGroupid:     &ctg.CTYPEGRP_ID,
			Tariffs:          nil,
			OP_ESTIM_CONS:    nil,
			NOOP_ESTIM_CONS:  nil,
			Description:      ctg.DESCRIPTION,
			GroupDescription: ctg.DESCRIPTION,
			Weigth:           ctg.WEIGHT,
		}
		lookups = append(lookups, lk)
	}
	resp.Ctgs = lookups
	return resp, nil
	//return nil, errors.New("Not Supported")
}

func (s *DataProvider) GetCustomersByBillgroup(cn context.Context, key *Key) (*CustomersList, error) {
	return nil, errors.New("Not Immplemented")
}

func (s *DataProvider) GetCustomerByCustkey(cn context.Context, key *Key) (cst *Customer, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	query := `select BILLGROUP,CUSTKEY,CYCLE_ID,BILNG_DATE,payment_no from dbo.HAND_MH_ST where custkey=@ck and bilng_date=@bd`
	var data []*dbmodels.HAND_MH_ST
	err = db.DB.Unsafe().Select(&data, query, sql.Named("ck", key.GetKey()), sql.Named("bd", key.BilngDate.AsTime()))

	if err == sql.ErrNoRows {
		return nil, errors.New(fmt.Sprintf("No data fount for %v @ %v ", key.GetKey(), key.GetBilngDate().AsTime().String()))
	}
	if err != nil {
		return nil, err
	}

	if data == nil || len(data) == 0 {
		return nil, errors.New(fmt.Sprintf("No data fount for %v @ %v ", key.GetKey(), key.GetBilngDate().AsTime().String()))
	}
	var hand *dbmodels.HAND_MH_ST = data[0]
	log.Println(hand)
	if hand == nil || hand.CYCLE_ID == nil {
		return nil, errors.New("invalied record cycle id not defined")
	}
	reps, err := s.getBill(cn, hand.CUSTKEY, *hand.CYCLE_ID)
	return reps.Customer, err
}

func (s *DataProvider) getBill(cn context.Context, custkey string, cycle_id int32) (bill *Bill, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	log.Println("get bill")
	query := fmt.Sprintf(`select * from dbo.HAND_MH_ST where custkey ='%s' and cycle_id=@cycle_id`, custkey)
	var data []*dbmodels.HAND_MH_ST
	err = db.DB.Unsafe().Select(&data, query, sql.Named("cycle_id", cycle_id))
	if err != nil {
		return nil, err
	}
	if data == nil || len(data) == 0 {
		query = fmt.Sprintf(`select * from dbo.ARC_HAND_MH_ST where custkey ='%s' and cycle_id=@cycle_id`, custkey)
		err = db.DB.Unsafe().Select(&data, query, sql.Named("cycle_id", cycle_id))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("no data found for %v at %v", custkey, cycle_id))
		}
	}
	if len(data) > 1 {
		return nil, errors.New("Conflict more then one record match")
	}
	if len(data) == 0 {
		return nil, errors.New(fmt.Sprintf("رقم الفاتورة غير موجودة %v @%v", custkey, cycle_id))
	}
	var hand *dbmodels.HAND_MH_ST = data[0]
	if hand == nil {
		return nil, errors.New("No Data Found")
	}
	if hand.CYCLE_ID == nil {
		return nil, errors.New("invalied cycle id:" + hand.CUSTKEY)
	}
	if hand.BILNG_DATE == nil {
		return nil, errors.New("invalied bilng date:" + hand.CUSTKEY)
	}
	var billCtypes []*dbmodels.BILL_CTYPES
	err = db.Where(&dbmodels.BILL_CTYPES{CUSTKEY: hand.CUSTKEY, CYCLE_ID: *hand.CYCLE_ID}).Find(&billCtypes)
	if err != nil {
		return nil, err
	}
	var prevConnTariffAloc []*dbmodels.CONN_DTL_TARIFF_ALLOC
	//found all tariff allocations for previous
	var prevCycleId *int32 //inital value
	query = fmt.Sprintf(`select max(cycle_id) from dbo.CONN_DTL_TARIFF_ALLOC where custkey ='%s' and bilng_date<@bdate`, custkey)
	err = db.DB.Unsafe().Get(&prevCycleId, query, sql.Named("bdate", *hand.BILNG_DATE))
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if prevCycleId == nil || *prevCycleId == 0 {
		prevCycleId = hand.CYCLE_ID
	}
	query = fmt.Sprintf(`select * from dbo.CONN_DTL_TARIFF_ALLOC where custkey ='%s' and cycle_id=%v order by cycle_id`, custkey, *prevCycleId)
	err = db.DB.Unsafe().Select(&prevConnTariffAloc, query)
	if err != nil {
		return nil, err
	}
	/*if len(connTariffAloc) > 1 {
		return nil, errors.New("الانشظة المختلطة معطلة موقتا")
	}*/
	cStatus := CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER
	opStatus := MeterOperationStatus_WORKING
	if hand.Meter_ref != nil || (hand.READ_TYPE != nil && *hand.READ_TYPE == 0) {
		cStatus = CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER
		if hand.Op_status != nil && *hand.Op_status == 1 {
			opStatus = MeterOperationStatus_NOT_WORKING
		} else {
			opStatus = MeterOperationStatus_WORKING
		}
	} else {
		cStatus = CONNECTION_STATUS_TYPE_CONNECTED_WITHOUT_METER
		opStatus = MeterOperationStatus_NOT_WORKING
	}
	if hand.CONN_STATUS != nil {
		if *hand.CONN_STATUS == 2 {
			cStatus = CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER
			opStatus = MeterOperationStatus_NOT_WORKING
		}

		if *hand.CONN_STATUS == 3 {
			cStatus = CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER
			opStatus = MeterOperationStatus_NOT_WORKING
		}
	}

	water := SERVICE_TYPE_WATER
	sewer := SERVICE_TYPE_SEWER
	var one int64 = 1
	var mainCtype dbmodels.CTG_CONSUMPTIONTYPES
	var minaCtypeGroup dbmodels.CTG_CONSUMPTIONTYPEGRPS
	err = db.Where(dbmodels.CTG_CONSUMPTIONTYPES{CTYPE_ID: *hand.C_type}).First(&mainCtype)
	if err != nil {
		return nil, err
	}
	err = db.Where(dbmodels.CTG_CONSUMPTIONTYPEGRPS{CTYPEGRP_ID: mainCtype.CTYPEGRP_ID}).First(&minaCtypeGroup)
	if err != nil {
		return nil, err
	}
	mainCtg := &Ctg{
		CType:            &mainCtype.CTYPE_ID,
		CTypeGroupid:     &mainCtype.CTYPEGRP_ID,
		Description:      mainCtype.DESCRIPTION,
		GroupDescription: minaCtypeGroup.DESCRIPTION,
	}
	wconn := Connection{
		CType:            mainCtg,
		NoUnits:          hand.No_units,
		IsBulkMeter:      new(bool),
		ConnDiameter:     new(int64),
		EstimCons:        hand.CONN_AVRG_CONSUMP,
		ConnectionStatus: &cStatus,
		//SubConnections:   []*SubConnection{},
	}
	if cStatus == CONNECTION_STATUS_TYPE_CONNECTED_WITH_METER {
		mtype := ""
		mref := ""
		if hand.Meter_type != nil {
			mtype = *hand.Meter_type
		}
		if hand.Meter_ref != nil {
			mref = *hand.Meter_ref
		}
		wconn.Meter = &Meter{
			MeterType:       &mtype,
			MeterRef:        &mref,
			Diameter:        new(int64),
			ConverterFactor: &one,
			OpStatus:        &opStatus,
		}
	}
	sconn := wconn //copy data
	wconn.SubConnections = make([]*SubConnection, 0)
	sconn.SubConnections = make([]*SubConnection, 0)
	if billCtypes != nil && len(billCtypes) > 1 {
		for id := range billCtypes {
			xt := billCtypes[id]
			if xt != nil {
				var ctype dbmodels.CTG_CONSUMPTIONTYPES
				err = db.Where(dbmodels.CTG_CONSUMPTIONTYPES{CTYPE_ID: xt.C_TYPE}).First(&ctype)
				if err != nil {
					return nil, err
				}
				var ctypeGroup dbmodels.CTG_CONSUMPTIONTYPEGRPS
				err = db.Where(dbmodels.CTG_CONSUMPTIONTYPEGRPS{CTYPEGRP_ID: ctype.CTYPEGRP_ID}).First(&ctypeGroup)
				if err != nil {
					return nil, err
				}
				if xt.CONSUMP_PERC == nil && prevConnTariffAloc != nil {
					for _, pcn := range prevConnTariffAloc {
						if pcn.C_TYPE == xt.C_TYPE {
							xt.CONSUMP_PERC = pcn.ALLOC_PERC
						}
					}
				}
				wconn.SubConnections = append(wconn.SubConnections, &SubConnection{
					CType: &Ctg{
						CType:            &xt.C_TYPE,
						CTypeGroupid:     &ctype.CTYPEGRP_ID,
						Description:      ctype.DESCRIPTION,
						GroupDescription: ctypeGroup.DESCRIPTION,
					},
					EstimateConsumption:   xt.CONSUMP,
					ConsumptionPercentage: xt.CONSUMP_PERC,
					NoUnits:               xt.NO_UNITS,
				})
				sconn.SubConnections = append(sconn.SubConnections, &SubConnection{
					CType: &Ctg{
						CType:            &xt.C_TYPE,
						CTypeGroupid:     &ctype.CTYPEGRP_ID,
						Description:      ctype.DESCRIPTION,
						GroupDescription: ctypeGroup.DESCRIPTION,
					},
					EstimateConsumption:   xt.CONSUMP,
					ConsumptionPercentage: xt.CONSUMP_PERC,
					NoUnits:               xt.NO_UNITS,
				})
			}
		}
	}
	cust := &Customer{
		Custkey:   &hand.CUSTKEY,
		CustType:  new(int64),
		IsCompany: new(bool),
		InfoFlag1: new(string),
		InfoFlag2: new(string),
		InfoFlag3: new(string),
		InfoFlag4: new(string),
		InfoFlag5: new(string),
		Property: &Property{
			PropRef:   hand.Prop_ref,
			InfoFlag1: new(string),
			InfoFlag2: new(string),
			InfoFlag3: new(string),
			InfoFlag4: new(string),
			InfoFlag5: new(string),
			NoRooms:   new(int64),
			Services: []*Service{
				{
					ServiceType: &water,
					Connection:  &wconn,
				},
				{
					ServiceType: &sewer,
					Connection:  &sconn,
				},
			},
			IsVacated: new(bool),
			Township:  new(string),
		},
		Billgroup: hand.BILLGROUP,
	}
	var balance *float32
	if hand.Cl_blnce != nil {
		bl := float32(*hand.Cl_blnce)
		bl = float32(math.Round(1000*float64(bl)) / 1000)
		balance = tools.ToFloat32Pointer(bl)
	}
	bRs := Bill{PaymentNo: hand.Payment_no, ClBalance: balance}
	bRs.Customer = cust
	if hand.BILNG_DATE != nil {
		bRs.BilngDate = timestamppb.New(*hand.BILNG_DATE)
	}
	bRs.ServicesReadings = make([]*ServiceReading, 0)
	rdg := Reading{
		Consump:   hand.S_CONSUMP,
		CrReading: hand.S_CR_READING,
		PrReading: hand.S_PR_READING,
	}
	var readType READING_TYPE
	actual := READING_TYPE_ACTUAL
	estim := READING_TYPE_ESTIMATE
	if hand.READ_TYPE == nil {
		if (hand.Cr_reading != nil && hand.Pr_read1 != nil && *hand.Pr_read1 > 0 && *hand.Cr_reading > 0) || (hand.Op_status != nil && *hand.Op_status == 0 && hand.Meter_ref != nil) {
			readType = actual
		} else {
			readType = estim
		}
	} else {
		readType = READING_TYPE(int32(*hand.READ_TYPE))
	}
	rdg.ReadType = &readType
	bRs.ServicesReadings = append(bRs.ServicesReadings, &ServiceReading{
		ServiceType: &water,
		Reading:     &rdg,
	})
	bRs.ServicesReadings = append(bRs.ServicesReadings, &ServiceReading{
		ServiceType: &sewer,
		Reading:     &rdg,
	})

	if billCtypes != nil && len(billCtypes) > 1 {
		bRs.FTransactions, err = getBillItemsCtypesTransactions(hand, billCtypes)
		if err != nil {
			return nil, err
		}
	} else {
		bRs.FTransactions, err = getBillItemsTransactions(hand, mainCtg)
		if err != nil {
			return nil, err
		}
	}
	return &bRs, err
}
func getBillItemsCtypesTransactions(hand *dbmodels.HAND_MH_ST, billctypes []*dbmodels.BILL_CTYPES) ([]*FinantialTransaction, error) {
	bdate := timestamppb.New(*hand.BILNG_DATE)
	readType := READING_TYPE_ACTUAL
	if hand.READ_TYPE != nil {
		if int32(*hand.READ_TYPE) == int32(READING_TYPE_AVERAGE) {
			readType = READING_TYPE_AVERAGE
		}
		if int32(*hand.READ_TYPE) == int32(READING_TYPE_ESTIMATE) {
			readType = READING_TYPE_ESTIMATE
		}
	}
	finantialCtypeTrans := []*FinantialTransaction{}
	for idx := range billctypes {
		ctype := billctypes[idx]
		if ctype == nil {
			return nil, errors.New("نشاط غير معرف")
		}
		var ctype_ctg dbmodels.CTG_CONSUMPTIONTYPES
		err := db.Where(dbmodels.CTG_CONSUMPTIONTYPES{CTYPE_ID: ctype.C_TYPE}).First(&ctype_ctg)
		if err != nil {
			return nil, err
		}
		var ctypeGroup dbmodels.CTG_CONSUMPTIONTYPEGRPS
		err = db.Where(dbmodels.CTG_CONSUMPTIONTYPEGRPS{CTYPEGRP_ID: ctype_ctg.CTYPEGRP_ID}).First(&ctypeGroup)
		if err != nil {
			return nil, err
		}
		var measure *MeasuredTransaction
		if ctype.CONSUMP != nil {
			measure = &MeasuredTransaction{
				Consump:   ctype.CONSUMP,
				CrReading: hand.S_CR_READING, // hand
				PrReading: hand.S_PR_READING, // hand
				ReadType:  &readType,         // hand
				MeterType: hand.Meter_type,   // hand
				MeterRef:  hand.Meter_ref,    // hand
			}
		}

		values := make(map[string]float64)
		water := SERVICE_TYPE_WATER
		sewer := SERVICE_TYPE_SEWER
		elect := SERVICE_TYPE_ELECTRICITY
		fire := SERVICE_TYPE_FIRE
		hydrant := SERVICE_TYPE_HYDRANT
		_values, err := tools.GetBillItemValues(&ctype.BILL_ITEMS)
		if err != nil {
			return nil, err
		}
		for ii := range _values {
			values[ii] = _values[ii]
		}
		for k, v := range values {
			itm := k   //copy
			value := v //copy
			if v == 0 {
				continue
			}
			var service SERVICE_TYPE
			if strings.Contains(strings.ToUpper(itm), "SEWER") {
				service = sewer
			} else if strings.Contains(strings.ToUpper(itm), "WATER") {
				service = water
			} else if strings.Contains(strings.ToUpper(itm), "ELECTRICITY") {
				service = elect
			} else if strings.Contains(strings.ToUpper(itm), "FIRE") {
				service = fire
			} else if strings.Contains(strings.ToUpper(itm), "HYDRANT") {
				service = hydrant
			} else {
				continue
			}

			descr, ok := dbmodels.FinancialTransCodes[dbmodels.TransCode(itm)]
			if !ok {
				descr = itm
			}
			finantialCtypeTrans = append(finantialCtypeTrans, &FinantialTransaction{
				ServiceType:    &service,
				Code:           tools.ToStringPointer(itm),
				Description:    &descr,
				BilngDate:      bdate, // hand
				EffDate:        bdate, // hand
				Amount:         &value,
				TaxAmount:      new(float64),
				DiscountAmount: new(float64),
				Ctype: &Ctg{
					CType:            &ctype.C_TYPE,
					CTypeGroupid:     &ctype_ctg.CTYPEGRP_ID,
					Description:      ctype_ctg.DESCRIPTION,
					GroupDescription: ctypeGroup.DESCRIPTION,
				},
				NoUnits:      ctype.NO_UNITS,
				PropRef:      hand.Prop_ref, // hand
				MTransaction: measure,
			})
		}
	}
	return finantialCtypeTrans, nil
}

func getBillItemsTransactions(hand *dbmodels.HAND_MH_ST, ctg *Ctg) ([]*FinantialTransaction, error) {
	var bi dbmodels.BILL_ITEM
	err := db.Where(&dbmodels.BILL_ITEM{CUSTKEY: hand.CUSTKEY, CYCLE_ID: hand.CYCLE_ID}).First(&bi)
	if err != nil && err == sql.ErrNoRows {
		return []*FinantialTransaction{}, nil
	}
	if err != nil {
		return nil, err
	}
	var billCtypes []*dbmodels.BILL_CTYPES
	err = db.Where(&dbmodels.BILL_CTYPES{CUSTKEY: hand.CUSTKEY, CYCLE_ID: *hand.CYCLE_ID}).Find(&billCtypes)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	values := make(map[string]float64)
	water := SERVICE_TYPE_WATER
	sewer := SERVICE_TYPE_SEWER
	fts := make([]*FinantialTransaction, 0)
	if billCtypes == nil || len(billCtypes) < 2 {
		values, err = tools.GetBillItemValues(&bi.BILL_ITEMS)
		if err != nil {
			return nil, err
		}
	} else {
		for bid := range billCtypes {
			btc := billCtypes[bid]
			_values, err := tools.GetBillItemValues(&btc.BILL_ITEMS)
			if err != nil {
				return nil, err
			}
			for ii := range _values {
				values[ii] = _values[ii]
			}
		}
	}

	for k, v := range values {
		itm := k   //copy
		value := v //copy
		if v == 0 {
			continue
		}
		if strings.Contains(strings.ToUpper(itm), "SEWER") {
			fts = append(fts, getFtrans(hand, ctg, itm, &value, sewer))
		} else {
			fts = append(fts, getFtrans(hand, ctg, itm, &value, water))
		}
	}
	return fts, nil
}

func getFtrans(hand *dbmodels.HAND_MH_ST, ctg *Ctg, code string, amount *float64, service SERVICE_TYPE) *FinantialTransaction {
	bdate := timestamppb.New(*hand.BILNG_DATE)
	readType := READING_TYPE_ACTUAL
	if hand.READ_TYPE != nil {
		if int32(*hand.READ_TYPE) == int32(READING_TYPE_AVERAGE) {
			readType = READING_TYPE_AVERAGE
		}
		if int32(*hand.READ_TYPE) == int32(READING_TYPE_ESTIMATE) {
			readType = READING_TYPE_ESTIMATE
		}
	}
	descr, ok := dbmodels.FinancialTransCodes[dbmodels.TransCode(code)]
	if !ok {
		descr = code
	}
	var measure *MeasuredTransaction
	if hand.S_CONSUMP != nil {
		measure = &MeasuredTransaction{
			Consump:   hand.S_CONSUMP,
			CrReading: hand.S_CR_READING,
			PrReading: hand.S_PR_READING,
			ReadType:  &readType,
			MeterType: hand.Meter_type,
			MeterRef:  hand.Meter_ref,
		}
	}
	return &FinantialTransaction{
		ServiceType:    &service,
		Code:           tools.ToStringPointer(code),
		Description:    &descr,
		BilngDate:      bdate,
		EffDate:        bdate,
		Amount:         rountToThPointer(amount),
		TaxAmount:      new(float64),
		DiscountAmount: new(float64),
		Ctype:          ctg,
		NoUnits:        hand.No_units,
		PropRef:        hand.Prop_ref,
		MTransaction:   measure,
	}
}

func create_timestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}
func rountToTh(val float64) float64 {
	return math.Round(1000*val) / 1000
}
func rountToThPointer(val *float64) *float64 {
	if val == nil {
		return nil
	}
	rv := math.Round(1000*(*val)) / 1000
	return &rv
}
