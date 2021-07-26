package providers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/HassanAbdelzaher/lama"
	"github.com/MaisrForAdvancedSystems/biller-mas-provider/tools"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	dbmodels "github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DataConsumer struct {
	InfoProvider
}

type BILL_CTYPES_TRANS struct {
	Trans *billing.FinantialTransaction
	Item  *dbmodels.BILL_CTYPES
}

func gethand(custkey string, billdate *timestamppb.Timestamp, tx *lama.LamaTx) (*dbmodels.HAND_MH_ST, bool, error) {
	hand := dbmodels.HAND_MH_ST{}
	bill_date := billdate.AsTime()
	err := tx.Where(&dbmodels.HAND_MH_ST{
		CUSTKEY:    custkey,
		BILNG_DATE: &bill_date,
	}).First(&hand)
	forArc := false
	if err == sql.ErrNoRows {
		err = tx.Table("ARC_HAND_MH_ST").Where(&dbmodels.HAND_MH_ST{
			CUSTKEY:    custkey,
			BILNG_DATE: &bill_date,
		}).First(&hand)
		if err == nil {
			forArc = true
		}
	}
	if err != nil {
		return nil, forArc, err
	}
	return &hand, forArc, nil
}

func getBill_Item(custkey string, cycleId int32, tx *lama.LamaTx) (*dbmodels.BILL_ITEM, error) {
	billItem := dbmodels.BILL_ITEM{}
	err := tx.Where(&dbmodels.BILL_ITEM{
		CUSTKEY:  custkey,
		CYCLE_ID: &cycleId,
	}).First(&billItem)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return &billItem, err
	}
	return &billItem, nil
}

func ctgBillCtypes(data []*billing.FinantialTransaction, ctype string) (*dbmodels.BILL_CTYPES, *billing.FinantialTransaction, error) {
	if data == nil || len(data) == 0 {
		return nil, nil, nil
	}
	values := make(map[string]float64)
	var trans *billing.FinantialTransaction
	for id := range data {
		bi := data[id]
		if bi == nil {
			continue
		}
		if bi.Code == nil && bi.Amount != nil && *bi.Amount != 0 {
			return nil, nil, errors.New("missing code for transaction")
		}
		if bi.Code == nil || bi.Amount == nil {
			continue
		}
		if bi.Ctype == nil || bi.Ctype.CType == nil {
			continue
		}
		if bi.Ctype.GetCType() != ctype {
			continue
		}
		if trans == nil || trans.MTransaction == nil {
			trans = bi
		}
		if _, ok := values[*bi.Code]; ok {
			values[*bi.Code] = values[*bi.Code] + *bi.Amount
		} else {
			values[*bi.Code] = *bi.Amount
		}
	}
	bitm := &dbmodels.BILL_CTYPES{BILL_ITEMS: dbmodels.BILL_ITEMS{}, C_TYPE: ctype}
	err := tools.SetBillItemValues(&bitm.BILL_ITEMS, values)
	return bitm, trans, err
}

func ctgBillCtypesAll(data []*billing.FinantialTransaction) ([]*BILL_CTYPES_TRANS, error) {
	ctypes := make(map[string]string)
	for _, d := range data {
		if d.Ctype != nil && d.Ctype.CType != nil {
			ctypes[*d.Ctype.CType] = *d.Ctype.CType
		}
	}
	items := make([]*BILL_CTYPES_TRANS, 0)
	for ctype := range ctypes {
		bl, trans, err := ctgBillCtypes(data, ctype)
		if err != nil {
			return nil, err
		}
		if bl != nil {
			items = append(items, &BILL_CTYPES_TRANS{
				Trans: trans,
				Item:  bl,
			})
		}
	}
	return items, nil
}

func createCancelledBillAction(tx *lama.LamaTx, recalcid *int64, userid *int32) error {
	if recalcid == nil {
		return nil
	}
	if *recalcid <= 0 {
		return nil
	}
	var req dbmodels.CANCELLED_REQUEST
	err := tx.Where(dbmodels.CANCELLED_REQUEST{FORM_NO: *recalcid}).First(&req)
	if err == sql.ErrNoRows {
		er := fmt.Sprintf(" %v طلب الالغاء غير موجود ", *recalcid)
		return errors.New(er)
	}
	if err != nil {
		return err
	}
	if req.CLOSED != nil && *req.CLOSED {
		er := fmt.Sprintf(" %v طلب الالغاء مغلق ", *recalcid)
		return errors.New(er)
	}
	//validate state
	if req.STATE != nil {
		var cr_state dbmodels.LU_CANCELLED_BILL_STATE
		err = tx.Where(dbmodels.LU_CANCELLED_BILL_STATE{
			ID: *req.STATE,
		}).First(&cr_state)
		if err != nil {
			return err
		}
		if cr_state.RECAL_READY == nil || *cr_state.RECAL_READY == false {
			er := fmt.Sprintf(" %v حالة الطلب لا تسمح باعادة الاحتساب للفواتير ", *recalcid)
			return errors.New(er)
		}
	}

	//
	var actions []*dbmodels.LU_CANCELLED_BILLS_ACTION
	tr := true
	err = tx.Model(dbmodels.LU_CANCELLED_BILLS_ACTION{}).Where(dbmodels.LU_CANCELLED_BILLS_ACTION{RECALC_DONE_ACTION: &tr}).Find(&actions)
	if err == sql.ErrNoRows {
		er := fmt.Sprintf(" %v لا يوجد تعريف لاجراء اعادة الاحتساب ", *recalcid)
		return errors.New(er)
	}
	if err != nil {
		return err
	}
	if len(actions) == 0 {
		return nil
	}
	action := actions[0]
	if action.CURRENT_STATE == nil {
		if err == sql.ErrNoRows {
			er := fmt.Sprintf(" %v الاجراء غير معرف لع الحالة الحالية ", action.ID)
			return errors.New(er)
		}
	}
	if req.STATE == nil {
		er := fmt.Sprintf(" %v حالة الطب لا غير صحيحة ", *recalcid)
		return errors.New(er)
	}
	if *req.STATE != *action.CURRENT_STATE {
		er := fmt.Sprintf(" %v حالة الطب لا تسمح بالاجراء ", *recalcid)
		return errors.New(er)
	}
	comment := "تم اعادة احتساب الفاتورة"
	now := time.Now()
	canceAction := dbmodels.CANCELLED_BILLS_ACTION{
		FORM_NO:     req.FORM_NO,
		CUSTKEY:     req.CUSTKEY,
		DOCUMENT_NO: req.DOCUMENT_NO,
		ACTION_ID:   action.ID,
		STAMP_DATE:  &now,
		COMMENT:     &comment,
		STAMP_USER:  nil,
		USER_ID:     userid,
	}
	err = tx.Add(canceAction)
	if err != nil {
		return err
	}
	if action.NEXT_STATE != nil {
		req.STATE = action.NEXT_STATE
		var lustate dbmodels.LU_CANCELLED_BILL_STATE
		err = tx.Where(dbmodels.LU_CANCELLED_BILL_STATE{
			ID: *action.NEXT_STATE,
		}).First(&lustate)
		if err != nil {
			return err
		}
		req.STATUS = lustate.DESCRIPTION
		err = tx.Save(req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DataConsumer) UpdateHandMhSt(bl *billing.Bill, m_trans *billing.MeasuredTransaction, tx *lama.LamaTx, msg *billing.PostMessage, user *dbmodels.USERS, ustation *dbmodels.STATIONS, formNo *int64) (*dbmodels.HAND_MH_ST, error) {
	cst := bl.Customer
	billdate := bl.GetBilngDate()
	var cl_blnce float64 = 0
	if cst == nil {
		return nil, errors.New("missing customer data")
	}
	if cst.Property == nil {
		return nil, errors.New("missing Property data")
	}
	if cst.Property.Services == nil || len(cst.Property.Services) == 0 {
		return nil, errors.New("missing services data")
	}
	if bl.FTransactions != nil {
		for _, tr := range bl.FTransactions {
			if tr.Amount != nil {
				cl_blnce = cl_blnce + *tr.Amount
			}
		}
		//return nil, errors.New("invalied request:missing bill finantinal transactions")
	}
	conn := cst.Property.Services[0].Connection
	if conn != nil {
		log.Println("Conn Found")
		if conn.ConnectionStatus != nil {
			log.Println("Conn Status", conn.ConnectionStatus.String())
		}
	}
	custkey := cst.GetCustkey()
	// oldhand save
	hand, _, err := gethand(custkey, billdate, tx)
	if err != nil {
		return nil, err
	}
	if ustation != nil && hand.STATION_NO != nil && !IsHq(ustation) {
		if *hand.STATION_NO != ustation.STATION_NO {
			return nil, errors.New("الفاتورة تخص فرع اخر")
		}
	}
	log.Println("Consumer: Success Get Hand DB", custkey)
	hsthand := &dbmodels.HST_HAND_MH_ST{
		RECALC_ID:  formNo,
		HAND_MH_ST: *hand,
	}
	err = tx.AddIfNotExists(hsthand)
	if err != nil {
		return nil, err
	} //never save to keep first view of record
	if cst.Property != nil {
		serv := cst.Property.Services
		if serv != nil && len(serv) > 0 {

		}
	}
	if bl.Customer.Property != nil {
		hand.Prop_ref = bl.Customer.Property.PropRef
	}
	/*if hand.INSTALMENT!=nil{
		cl_blnce=cl_blnce+*hand.INSTALMENT
	}*/
	hand.BILLGROUP = bl.Customer.Billgroup
	now := time.Now()
	hand.STAMP_DATE = &now
	hand.STAMP_USER = user.USER_NAME
	hand.NOTE_R = tools.ToStringPointer("اعادة احتساب")
	hand.RECALC_ID = msg.CancelledRequestFormNo
	hand.No_units = conn.NoUnits
	hand.Cl_blnce = &cl_blnce
	hand.MODIFIED_AVRG_CONSUMP = conn.EstimCons
	if conn.CType != nil {
		hand.C_type = conn.CType.CType
		hand.Ctypegrp_id = conn.CType.CTypeGroupid
	}
	//upgrade ctype for multiconn
	if conn.SubConnections != nil && len(conn.SubConnections) > 1 {
		frst := conn.SubConnections[0]
		if frst != nil && frst.CType != nil {
			hand.C_type = frst.CType.CType
			hand.Ctypegrp_id = frst.CType.CTypeGroupid
		}
	}
	if conn.Meter != nil {
		hand.Meter_ref = conn.Meter.MeterRef
		hand.Meter_type = conn.Meter.MeterType
	}
	connStatus := int32(*conn.ConnectionStatus)
	if m_trans == nil && conn.ConnectionStatus != nil && (*conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER || *conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER) {
		log.Println("Conn Not Connected")
		sZero := float64(0)
		hand.S_CONSUMP = &sZero
		hand.Cl_blnce = &sZero
	}
	hand.CONN_STATUS = &connStatus
	if m_trans != nil {
		t := m_trans
		if t.ReadType != nil {
			rty := int64(*t.ReadType)
			hand.READ_TYPE = &rty
		}
		hand.S_CR_READING = t.CrReading
		hand.S_PR_READING = t.PrReading
		if t.CrReading != nil {
			cr := *t.CrReading
			var pr float64 = 0
			if t.PrReading != nil {
				pr = *t.PrReading
			}
			cons := cr - pr
			hand.S_CONSUMP = &cons
		} else {
			if t.Consump != nil {
				hand.S_CONSUMP = conn.EstimCons
			}
		}
	}
	log.Println("Consumer: Init Save Hand", custkey)
	err = tx.Upsert(hand)
	if err != nil {
		return nil, err
	}
	_, err = tx.DB.Exec(fmt.Sprintf("DELETE FROM ARC_HAND_MH_ST where custkey='%v' and cycle_id=%v", hand.CUSTKEY, *hand.CYCLE_ID))
	if err != nil {
		return nil, err
	}
	if hand != nil {
		log.Println("Hand After Upsert:", hand.CUSTKEY)
	}
	if hand.Payment_no != nil {
		log.Println("Hand After Upsert:", *hand.Payment_no)
	}
	if hand.S_CONSUMP != nil {
		log.Println("Hand S_Consump:", *hand.S_CONSUMP)
	}
	if hand.Consump != nil {
		log.Println("Hand Consump:", *hand.Consump)
	}
	return hand, nil
}

func (s *DataConsumer) UpdateBillItems(bl *billing.Bill, cycle_id int32, stationNo *int32, tx *lama.LamaTx, msg *billing.PostMessage, user *dbmodels.USERS, ustation *dbmodels.STATIONS, formNo *int64) error {
	cst := bl.Customer
	if cst == nil {
		return errors.New("missing customer data")
	}
	if cst.Property == nil {
		return errors.New("missing Property data")
	}
	if cst.Property.Services == nil || len(cst.Property.Services) == 0 {
		return errors.New("missing services data")
	}
	conn := cst.Property.Services[0].Connection
	if conn.ConnectionStatus != nil && (*conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER || *conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER) {
		if bl.FTransactions == nil {
			bl.FTransactions = make([]*billing.FinantialTransaction, 0)
			//return nil, errors.New("invalied request:missing bill finantinal transactions")
		}
	} else {
		if len(bl.FTransactions) == 0 {
			return errors.New("invalied request:empty finantinal transactions slice")
		}
	}
	if bl.FTransactions == nil {
		return errors.New("invalied request:missing bill finantinal transactions")
	}
	custkey := cst.GetCustkey()
	paymentNo := bl.GetPaymentNo()
	// oldbillitem save
	oldBillItem, err := getBill_Item(custkey, cycle_id, tx)
	if err != nil {
		return err
	}
	//restrive old bill items
	//important note oldBillItem_Temp maybe null without error
	//hstBillItem := getOldBill(msg.CancelledRequestFormNo,hand.STATION_NO,custkey, hand.CYCLE_ID,oldBillItem)
	var hstBillItem *dbmodels.HST_BILL_ITEM = nil
	if oldBillItem != nil {
		hstBillItem = &dbmodels.HST_BILL_ITEM{
			RECALC_ID: formNo,
			BILL_ITEM: *oldBillItem,
		}
	}
	if hstBillItem != nil {
		err = tx.AddIfNotExists(hstBillItem)
		if err != nil {
			return err
		}
	}
	log.Println("Consumer: Temp Bill In Trans Save", custkey, paymentNo)
	if conn.ConnectionStatus != nil && (*conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER || *conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER) {
		log.Println("conn Is DisConnected")
		tx.Delete(oldBillItem)
		newBillItems := dbmodels.BILL_ITEM{CUSTKEY: custkey, CYCLE_ID: tools.ToInt32Pointer(int32(cycle_id)), STATION_NO: stationNo}
		err = tx.Upsert(newBillItems)
		if err != nil {
			return err
		}
	} else {
		newBillItems, _, err := tools.CreateBillItems(custkey, int64(cycle_id), stationNo, bl.FTransactions)
		if err != nil {
			return err
		}
		err = tx.Upsert(newBillItems)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DataConsumer) UpdateBillCtypes(bl *billing.Bill, cycle_id int32, payment_no *string, tx *lama.LamaTx, msg *billing.PostMessage, user *dbmodels.USERS, ustation *dbmodels.STATIONS, formNo *int64) error {
	cst := bl.Customer
	if cst == nil {
		return errors.New("missing customer data")
	}
	if cst.Property == nil {
		return errors.New("missing Property data")
	}
	if cst.Property.Services == nil || len(cst.Property.Services) == 0 {
		return errors.New("missing services data")
	}
	custkey := cst.GetCustkey()
	log.Println("Consumer: ", custkey)

	//save previous hst values
	cnt, err := tx.Model(dbmodels.HST_BILL_CTYPES{}).Where(dbmodels.BILL_CTYPES{
		CUSTKEY:   custkey,
		CYCLE_ID:  cycle_id,
		RECALC_ID: formNo,
	}).Count()
	if err != nil {
		return errors.New("missing Current Ctypes")
	}
	currentCtypes := []dbmodels.BILL_CTYPES{}
	if *cnt == 0 {
		//adding Current BILL_CTYPES To HST_BILL_CTYPES Then Delete In BILL_CTYPES Then Add New In BILL_CTYPES
		err := tx.Model(dbmodels.BILL_CTYPES{}).Where(dbmodels.BILL_CTYPES{
			CUSTKEY:  custkey,
			CYCLE_ID: cycle_id,
		}).Find(&currentCtypes)
		if err != nil {
			return errors.New("missing Current Ctypes")
		}
		for idx := range currentCtypes {
			currentCtype := currentCtypes[idx]
			Hnew := &dbmodels.HST_BILL_CTYPES{
				RECALC_ID:   formNo,
				BILL_CTYPES: currentCtype,
			}
			err = db.AddIfNotExists(Hnew)
			if err != nil {
				return err
			}
		}
	} else {
		err := tx.Model(dbmodels.HST_BILL_CTYPES{}).Where(dbmodels.BILL_CTYPES{
			CUSTKEY:  custkey,
			CYCLE_ID: cycle_id,
		}).Find(&currentCtypes)
		if err != nil {
			return errors.New("missing Current Ctypes")
		}
	}

	if cst.Property != nil && cst.Property.Services != nil && len(cst.Property.Services) > 0 {
		conn := bl.Customer.Property.Services[0].Connection
		//stm := fmt.Sprintf(`delete from BILL_CTYPES where custkey='%v' and cycle_id=%d`, custkey, cycle_id)
		//_, err := db.DB.Exec(stm)
		err = db.Model(&dbmodels.BILL_CTYPES{}).Where(&dbmodels.BILL_CTYPES{
			CUSTKEY:  custkey,
			CYCLE_ID: cycle_id,
		}).DeleteAll()
		if err != nil {
			return err
		}
		ctype_bill_items, err := ctgBillCtypesAll(bl.FTransactions)
		if err != nil {
			return err
		}
		if len(ctype_bill_items) == 0 {
			log.Println("No ctype_bill_items")
			zerov := float64(0)
			log.Println("currentCtypes", len(currentCtypes))
			for idx := range currentCtypes {
				currentCtype := currentCtypes[idx]
				nw := &dbmodels.BILL_CTYPES{
					CUSTKEY:      custkey,
					CYCLE_ID:     cycle_id,
					C_TYPE:       currentCtype.C_TYPE,
					PAYMENT_NO:   payment_no,
					CONSUMP_PERC: currentCtype.CONSUMP_PERC,
					CONSUMP:      &zerov,
					NO_UNITS:     currentCtype.NO_UNITS,
					BILL_ITEMS:   dbmodels.BILL_ITEMS{},
					RECALC_ID:    formNo,
				}
				err = db.Upsert(nw)
				if err != nil {
					return err
				}
			}
		}
		for cid := range ctype_bill_items {
			sb := ctype_bill_items[cid]
			trans := sb.Trans
			itm := sb.Item
			var cp *float64
			if trans == nil || trans.Ctype == nil || trans.Ctype.CType == nil {
				return errors.New("invalied ctype for sub connection")
			}
			if len(conn.SubConnections) > 0 {
				for id := range conn.SubConnections {
					if conn.SubConnections[id].CType != nil {
						if conn.SubConnections[id].CType.GetCType() == *trans.Ctype.CType {
							cp = conn.SubConnections[id].ConsumptionPercentage
							break
						}
					}
				}
			}
			var consump *float64
			if trans.MTransaction != nil {
				consump = trans.MTransaction.Consump
			}
			nw := &dbmodels.BILL_CTYPES{}
			if conn.ConnectionStatus != nil && (*conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER || *conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER) {
				nw = &dbmodels.BILL_CTYPES{
					CUSTKEY:      custkey,
					CYCLE_ID:     cycle_id,
					C_TYPE:       *trans.Ctype.CType,
					PAYMENT_NO:   payment_no,
					CONSUMP_PERC: cp,
					CONSUMP:      consump,
					NO_UNITS:     trans.NoUnits,
					BILL_ITEMS:   dbmodels.BILL_ITEMS{},
					RECALC_ID:    formNo,
				}
			} else {
				nw = &dbmodels.BILL_CTYPES{
					CUSTKEY:      custkey,
					CYCLE_ID:     cycle_id,
					C_TYPE:       *trans.Ctype.CType,
					PAYMENT_NO:   payment_no,
					CONSUMP_PERC: cp,
					CONSUMP:      consump,
					NO_UNITS:     trans.NoUnits,
					BILL_ITEMS:   itm.BILL_ITEMS,
					RECALC_ID:    formNo,
				}
			}

			err = db.Upsert(nw)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (s *DataConsumer) WriteFinantialData(cn context.Context, msg *billing.PostMessage) (*billing.Empty, error) {
	log.Printf("write data by :%v", cn.Value("username"))
	username, ok := cn.Value("username").(string)
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

	//dateNow := time.Now()
	//return nil, errors.New("NotImplemented")

	if msg == nil {
		return nil, errors.New("invalied request")
	}
	data := msg.Data
	if data == nil {
		return nil, errors.New("invalied request")
	}
	if data.Bills == nil {
		return nil, errors.New("invalied request:missing bill")
	}
	if len(data.Bills) == 0 {
		return nil, errors.New("invalied request:empty request")
	}
	db.Debug = true
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	var formNo int64 = -1
	if msg.IsCancelledRequest != nil && *msg.IsCancelledRequest {
		if msg.CancelledRequestFormNo == nil {
			return nil, errors.New("رقم الطلب غير معرف ")
		}
		formNo = *msg.CancelledRequestFormNo
		var req dbmodels.CANCELLED_REQUEST
		err := tx.Where(dbmodels.CANCELLED_REQUEST{FORM_NO: formNo}).First(&req)
		if err == sql.ErrNoRows {
			er := fmt.Sprintf(" %v طلب الالغاء غير موجود ", formNo)
			return nil, errors.New(er)
		}
		if err != nil {
			return nil, err
		}
		createCancelledBillAction(tx, &formNo, &user.ID)
		cnt, err := tx.Model(dbmodels.CANCELLED_BILL{}).Where(dbmodels.CANCELLED_BILL{
			FORM_NO: formNo,
		}).Count()
		if err != nil {
			return nil, err
		}
		if int64(len(data.Bills)) != *cnt {
			return nil, errors.New(fmt.Sprintf("متوقع عدد فواتير %v بينما وجد %v", len(data.Bills), *cnt))
		}
	} else {
		formNo = time.Now().UnixNano()
	}
	/*if len(data.Bills)>1 {
		return nil, errors.New("currently we support only on bill for each request")
	}*/
	for id := range data.Bills {
		bl := data.Bills[id]
		cst := bl.Customer
		if cst == nil {
			return nil, errors.New("missing customer data")
		}
		if cst.Property == nil {
			return nil, errors.New("missing Property data")
		}
		if cst.Property.Services == nil || len(cst.Property.Services) == 0 {
			return nil, errors.New("missing services data")
		}
		conn := cst.Property.Services[0].Connection
		if conn.SubConnections == nil {
			conn.SubConnections = make([]*billing.SubConnection, 0)
		}
		if len(conn.SubConnections) == 1 {
			return nil, errors.New("الانشطة المختلطة غير صحيحة:الحد الادني لعدد الانشطة 2")
		}
		isHaveMain := false
		var estim207 float64 = 0
		for _, sb := range conn.SubConnections {
			if sb.CType == nil || sb.CType.CType == nil {
				return nil, errors.New("نشاط فرعي غير معرف")
			}
			if conn.CType != nil && conn.CType.CType != nil {
				if *sb.CType.CType == *conn.CType.CType {
					isHaveMain = true
				}
			}
			if sb.EstimateConsumption != nil && *sb.EstimateConsumption > 0 {
				estim207 = estim207 + *sb.EstimateConsumption
			}
		}
		if len(conn.SubConnections) > 1 {
			//update main ctype
			if !isHaveMain {
				conn.CType = conn.SubConnections[0].CType
			}
			if estim207 > 0 {
				conn.EstimCons = &estim207
			}

		} else {
			if conn.CType == nil {
				return nil, errors.New("النشاط الرئيسي غير معرف")
			}
		}
		if conn.ConnectionStatus != nil && (*conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITH_METER || *conn.ConnectionStatus == billing.CONNECTION_STATUS_TYPE_DISCONNECTED_WITHOUT_METER) {
			if bl.FTransactions == nil {
				bl.FTransactions = make([]*billing.FinantialTransaction, 0)
				//return nil, errors.New("invalied request:missing bill finantinal transactions")
			}
		} else {
			if len(bl.FTransactions) == 0 {
				return nil, errors.New("invalied request:empty finantinal transactions slice")
			}
		}
		if bl.FTransactions == nil {
			return nil, errors.New("invalied request:missing bill finantinal transactions")
		}
		var mTrans *billing.MeasuredTransaction
		for id, ft := range bl.FTransactions {
			if ft.MTransaction != nil && ft.MTransaction.Consump != nil {
				mTrans = bl.FTransactions[id].MTransaction
			}
		}
		hand, err := s.UpdateHandMhSt(bl, mTrans, tx, msg, user, ustation, &formNo)
		if err != nil {
			return nil, err
		}
		if hand == nil {
			return nil, errors.New("internal error:hand is nil")
		}
		if hand.CYCLE_ID == nil {
			return nil, errors.New(fmt.Sprintf("missing cycle id:%v", hand.Payment_no))
		}
		err = s.UpdateBillItems(bl, *hand.CYCLE_ID, hand.STATION_NO, tx, msg, user, ustation, &formNo)
		if err != nil {
			return nil, err
		}
		err = s.UpdateBillCtypes(bl, *hand.CYCLE_ID, hand.Payment_no, tx, msg, user, ustation, &formNo)
		if err != nil {
			return nil, err
		}
		err = createCancelledBillAction(tx, &formNo, &user.ID)
		if err != nil {
			return nil, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &billing.Empty{}, nil
}
