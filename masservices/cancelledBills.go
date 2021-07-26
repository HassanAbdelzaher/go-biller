package masservices

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"errors"
	"log"
	"sort"
	"strings"
	"time"

	pbdbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/dbmessages"
	pbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/messages"
	"github.com/MaisrForAdvancedSystems/go-biller-proto/go/serverhostmessages"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbpool"
	irespo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/interfaces"
	respo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/repositories"

	"google.golang.org/grpc/codes"
)

func cancelledBillListPP(ctx *context.Context, in *pbMessages.CancelledBillListRequest, perfectreq *bool) (rsp *pbMessages.CancelledBillListResponse, err error) {
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	if user.CANCEL_BILL == nil || !*user.CANCEL_BILL {
		return nil, errors.New("المستخدم لا يمتلك الصلاحية الكافية")
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	var cancelledBills irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledBillsData []*dbmodels.CANCELLED_REQUEST
	if in.State != nil && (station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0) {
		cancelledBillsData, err = cancelledBills.GetByClosedStatusStation(false, *in.State, station.STATION_NO)
	} else if in.State != nil {
		cancelledBillsData, err = cancelledBills.GetByClosedStatus(false, *in.State)
	} else if station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0 {
		cancelledBillsData, err = cancelledBills.GetByClosedStation(false, station.STATION_NO)
	} else {
		cancelledBillsData, err = cancelledBills.GetByClosed(false)
	}

	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.CancelledBillListResponse{}
	stringEmpty := ""
	for idx := range cancelledBillsData {
		usj := &pbdbMessages.CANCELLED_REQUEST{}
		cancelledBillsUse := cancelledBillsData[idx]
		//log.Println("IDX: ", idx, " CUSTKEY", cancelledBillsUse.CUSTKEY, "Comment ", cancelledBillsUse.COMMENT)
		usj.CLOSED = cancelledBillsUse.CLOSED
		if cancelledBillsUse.COMMENT == nil {
			usj.COMMENT = &stringEmpty
		} else {
			usj.COMMENT = cancelledBillsUse.COMMENT
		}
		usj.CUSTKEY = &cancelledBillsUse.CUSTKEY
		usj.DOCUMENT_NO = &cancelledBillsUse.DOCUMENT_NO
		usj.FORM_NO = &cancelledBillsUse.FORM_NO
		usj.REQUEST_BY = cancelledBillsUse.REQUEST_BY
		usj.REQUEST_DATE = create_timestamp(cancelledBillsUse.REQUEST_DATE)
		usj.STAMP_DATE = create_timestamp(cancelledBillsUse.STAMP_DATE)
		usj.STATE = cancelledBillsUse.STATE
		usj.STATION_NO = tools.StringToInt32(&cancelledBillsUse.STATION_NO)
		usj.STATUS = cancelledBillsUse.STATUS
		usj.SURNAME = cancelledBillsUse.SURNAME

		DataJ.CancelledBillList = append(DataJ.CancelledBillList, usj)
	}
	log.Println("End cancelledBillList..")
	*perfectreq = true
	return DataJ, nil
}
func getCustomerPaymentsPP(ctx *context.Context, in *pbMessages.GetCustomerPaymentsRequest, perfectreq *bool) (rsp *pbMessages.GetCustomerPaymentsResponse, err error) {
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	if in.Custkey == nil {
		return nil, errors.New("رقم الحساب غير صحيح")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var hand irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var ctgConTypeGr irespo.ICtgConsumptionTypeGroupsRepository = &respo.CtgConsumptionTypeGroupsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	ctgData, err := ctgConTypeGr.GetAll()
	if err != nil {
		return nil, err
	}
	handData, err := hand.GetPaymentsByCustKey(*in.Custkey)
	if err != nil {
		return nil, err
	}
	log.Println("handData Done ...")
	if len(handData) == 0 {
		return nil, errors.New("لا توجد اي فواتير للعميل " + *in.Custkey)
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	var stationNo *int32 = nil
	if !(station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 1) {
		stationNo = &station.STATION_NO
	}
	DataJ := &pbMessages.GetCustomerPaymentsResponse{}
	for idx := range handData {
		handDataUse := handData[idx]
		usj, err := getPayment(handDataUse.Payment_no, &handDataUse.CUSTKEY, tools.ToBoolPointer(true), nil, nil, stationNo, &hand, user.USER_NAME, ctgData)
		if err != nil {
			return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
		}
		DataJ.Items = append(DataJ.Items, usj)
	}
	log.Println("end ..")
	*perfectreq = true
	return DataJ, nil
}
func getPayment(paymentNo *string, custKey *string, skipBracodTrim *bool, forQuery *bool, cycle_id *int32, stationNo *int32, hand *irespo.IHandMhStRepository, userName *string, ctg []*dbmodels.CTG_CONSUMPTIONTYPEGRPS) (rsp *serverhostmessages.CollectionDestributionItem, err error) {
	if forQuery == nil {
		forQuery = tools.ToBoolPointer(false)
	}
	if skipBracodTrim == nil {
		skipBracodTrim = tools.ToBoolPointer(false)
	}
	if paymentNo == nil && custKey == nil {
		return nil, errors.New("برجاء تحديد كود الفاتورة أو رقم الحساب")
	}
	if paymentNo != nil {
		paymentNo = tools.ToStringPointer(strings.TrimSpace(*paymentNo))
		if !*skipBracodTrim {
			if int32(len(*paymentNo)) > BARCODE_LENGTH && BARCODE_LENGTH > 1 {
				paymentNo = tools.ToStringPointer((*paymentNo)[0:BARCODE_LENGTH])
			}
		}
	}
	pay, err := (*hand).GetPayment(paymentNo, stationNo, custKey, cycle_id)
	if err != nil {
		return nil, err
	}
	if len(pay) == 0 {
		return nil, errors.New("لا يوجد فاتورة بهذا الرقم")
	}
	if (custKey != nil && *custKey != "") && len(pay) > 1 {
		sort.SliceStable(pay, func(i, j int) bool {
			if pay[i].BILNG_DATE == nil && pay[j].BILNG_DATE == nil {
				return false
			} else if pay[i].BILNG_DATE == nil {
				return false
			} else if pay[j].BILNG_DATE == nil {
				return true
			}
			return (*pay[i].BILNG_DATE).After(*pay[j].BILNG_DATE)
		})
	}
	if len(pay) > 1 {
		if pay[0].BILNG_DATE != nil && pay[1].BILNG_DATE != nil && (*pay[0].BILNG_DATE).Equal(*pay[1].BILNG_DATE) {
			return nil, errors.New("يوجد اكثر من فاتورة بهذا الرقم")
		}
	}
	var cancelreq irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	cancelData, err := cancelreq.GetByCustKey(*custKey)
	if err != nil {
		return nil, err
	}
	for idx := range cancelData {
		cancelDataUse := cancelData[idx]
		cancelBillData, err := cancelreq.GetByFormNoPaymentNo(cancelDataUse.FORM_NO, *pay[0].Payment_no)
		if err != nil {
			return nil, err
		}
		if len(cancelBillData) > 0 {
			//return nil, errors.New("يوجد طلب مفتوح على الفاتورة قيد المراجعة")
			return nil, errors.New("Exist Opened Issue")
		}
	}
	var customerbook irespo.ICustomerBooksRepository = &respo.CustomerBooksRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	bookData, err := customerbook.GetByBillGroupCode(*pay[0].BILLGROUP, *pay[0].BOOK_NO_C)
	if err != nil {
		return nil, err
	}
	bookDesc := ""
	if len(bookData) > 0 {
		if bookData[0].DESCRIBE != nil {
			bookDesc = *bookData[0].DESCRIBE
		}
	}
	if pay[0].BOOK_NO_C != nil {
		bookDesc = *pay[0].BOOK_NO_C + " " + bookDesc
	}

	var colAmt *float64
	calType := ""
	readType := int32(0)
	activ := ""
	usernamee := ""
	var installDate *time.Time
	if pay[0].COLLECTION_AMT != nil && *pay[0].COLLECTION_AMT == 0 {
		if pay[0].CUR_PAYMNTS != nil {
			colAmt = pay[0].CUR_PAYMNTS
		}
	} else {
		colAmt = pay[0].COLLECTION_AMT
	}
	if pay[0].CALC_TYPE != nil {
		calType = strings.TrimSpace(*pay[0].CALC_TYPE)
	}
	if pay[0].READ_TYPE != nil {
		readType = *pay[0].READ_TYPE
	}
	if userName != nil {
		usernamee = *userName
	}
	if pay[0].INSTALMENT_DATE != nil && pay[0].BILNG_DATE != nil && ((*pay[0].INSTALMENT_DATE).After(*pay[0].BILNG_DATE)) {
		installDate = pay[0].BILNG_DATE
	} else {
		installDate = pay[0].INSTALMENT_DATE
	}
	if pay[0].Ctypegrp_id != nil {
		for idx := range ctg {
			ctgUse := ctg[idx]
			if ctgUse.CTYPEGRP_ID == *pay[0].Ctypegrp_id {
				activ = *ctgUse.DESCRIPTION
				break
			}
		}
	}
	timeN := time.Now()
	collectedB4 := float64(0)
	reminderValue := float64(0)
	var recep irespo.IReciptsRepository = &respo.ReciptsRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	recepData, err := recep.GetByCustKeyCycleIDCancelled(*pay[0].CUSTKEY, pay[0].CYCLE_ID, false)
	if err != nil {
		return nil, err
	}
	for idx := range recepData {
		recepDataUse := recepData[idx]
		collectedB4 += recepDataUse.AMOUNT
	}
	reminderValue = float64(0)
	reminderValue = tools.RoundTo(pay[0].Cl_blnce-collectedB4, 3)
	ctypgrid := ""
	if pay[0].Ctypegrp_id != nil {
		ctypgrid = *pay[0].Ctypegrp_id
	}
	DataJ := &serverhostmessages.CollectionDestributionItem{
		PAYMENT_NO:            pay[0].Payment_no,
		CUSTKEY:               pay[0].CUSTKEY,
		DELIVERY_ST:           pay[0].Delivery_st,
		BILLGROUP:             pay[0].BILLGROUP,
		BILNG_DATE:            create_timestamp(pay[0].BILNG_DATE),
		CL_BLNCE:              tools.ToFloatPointer(pay[0].Cl_blnce),
		ISSUED_AMOUNT:         tools.ToFloatPointer(pay[0].Cl_blnce),
		SURNAME:               pay[0].Tent_name,
		BOOK_NO:               &bookDesc,
		WALK_NO:               pay[0].WALK_NO_C,
		SEQ_NO:                tools.Int32PtrToInt64Ptr(pay[0].SEQ_NO_C),
		EMP_ID:                pay[0].EMPID_C,
		ISSUED_COUNT:          tools.Int32ToInt32Ptr(1),
		WATER_AMT:             pay[0].WATER_AMT,
		SEWER_AMT:             pay[0].SEWER_AMT,
		BASIC_AMT:             pay[0].BASIC_AMT,
		TAX_AMT:               pay[0].TAX_AMT,
		ROUND_AMT:             pay[0].ROUND_AMT,
		AGREEM_AMT:            pay[0].AGREEM_AMT,
		CONN_INSTALLS_AMT:     pay[0].CONN_INSTALLS_AMT,
		CRDT_AMT:              pay[0].CRDT_AMT,
		DBT_AMT:               pay[0].DBT_AMT,
		INSTALLS_AMT:          pay[0].INSTALLS_AMT,
		METER_INSTALLS_AMT:    pay[0].METER_INSTALLS_AMT,
		OTHER_AMT:             pay[0].OTHER_AMT,
		OTHER_AMT1:            pay[0].OTHER_AMT1,
		OTHER_AMT2:            pay[0].OTHER_AMT2,
		OTHER_AMT3:            pay[0].OTHER_AMT3,
		OTHER_AMT4:            pay[0].OTHER_AMT4,
		OTHER_AMT5:            pay[0].OTHER_AMT5,
		TAKAFUL_AMT:           pay[0].TAKAFUL_AMT,
		TANZEEM_AMT:           pay[0].TANZEEM_AMT,
		OP_BLNCE:              pay[0].OP_BLNCE,
		INSTALMENT:            pay[0].INSTALMENT,
		COMPUTER_AMT:          pay[0].COMPUTER_AMT,
		CONN_AMT:              pay[0].CONN_AMT,
		CONTRACT_AMT:          pay[0].CONTRACT_AMT,
		GOV_AMT:               pay[0].GOV_AMT,
		METER_AMT:             pay[0].METER_AMT,
		METER_MAN_AMT:         pay[0].METER_MAN_AMT,
		UNI_AMT:               pay[0].UNI_AMT,
		CLEAN_AMT:             pay[0].CLEAN_AMT,
		CUR_CHARGES:           pay[0].CUR_CHARGES,
		CUR_PAYMNTS:           colAmt,
		NO_UNITS:              pay[0].No_units,
		CALC_TYPE:             &calType,
		CR_REAING:             pay[0].S_CR_READING, //القراءة الحالية في الطباعة ليس لعا علاقة بقراءات العداد الاخيرة
		PR_READING:            pay[0].S_PR_READING,
		CONSUMP:               pay[0].S_CONSUMP,
		ADDRESS:               pay[0].Ua_adress1,
		OLD_KEY:               pay[0].OLD_KEY,
		READ_TYPE:             &readType,
		PR_READ1:              tools.FloatPtrToInt32Ptr(pay[0].Pr_read1), //القراءة السابقة في حال متعذر والمغلق والمعطل
		METER_REF:             pay[0].Meter_ref,
		ACTIVITY:              &activ,
		INSTALMENT_DATE:       create_timestamp(installDate),
		CYCLE_ID:              pay[0].CYCLE_ID,
		DUE_AMOUNT:            pay[0].DUE_AMOUNT,
		BILL_AMOUNT:           pay[0].BILL_AMOUNT,
		BILL_ADJ_AMOUNT:       pay[0].BILL_ADJ_AMOUNT,
		COLLECTION_AMT:        colAmt,
		CTG:                   &ctypgrid,
		STAMP_DATE:            create_timestamp(&timeN),
		USER:                  &usernamee,
		TotalAmountCollected:  &collectedB4,
		REMINDER_VALUE:        &reminderValue,
		CTYPES_DTL:            []*serverhostmessages.CTYPE_DTL{},
		IS_CTYPES:             tools.ToBoolPointer(false),
		TotalCountCollected:   tools.ToFloatPointer(0),
		IS_COLLECTED_BY_OWNER: tools.ToBoolPointer(false),
		IS_COLLECTED_BY_OTHER: tools.ToBoolPointer(false),
	}
	var ctydetail irespo.IBillCtypesRepository = &respo.BillCtypesRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	ctydetailData, err := ctydetail.GetSumsByCustKeyCycleID(*pay[0].CUSTKEY, *pay[0].CYCLE_ID)
	if len(ctydetailData) > 1 {
		for idx := range ctydetailData {
			ctydetailDataUse := ctydetailData[idx]
			dec := ""
			if ctydetailDataUse.DESCRIPTION == nil {
				dec = *ctydetailDataUse.DESCRIPTION
			}
			for idx := range ctg {
				ctgUse := ctg[idx]
				if ctgUse.CTYPEGRP_ID == dec {
					ctydetailDataUse.DESCRIPTION = ctgUse.DESCRIPTION
					break
				}
			}
			totalV := tools.RoundTo(*ctydetailDataUse.OTHER_AMT+*ctydetailDataUse.WATER_AMT+*ctydetailDataUse.SEWER_AMT, 2)
			ctcyJ := serverhostmessages.CTYPE_DTL{
				DESCRIPTION:  ctydetailDataUse.DESCRIPTION,
				OTHER_AMT:    ctydetailDataUse.OTHER_AMT,
				WATER_AMT:    ctydetailDataUse.WATER_AMT,
				SEWER_AMT:    ctydetailDataUse.SEWER_AMT,
				TOTAL_AMOUNT: &totalV,
			}
			DataJ.CTYPES_DTL = append(DataJ.CTYPES_DTL, &ctcyJ)
		}
		mokhtalat := "مختلط"
		DataJ.ACTIVITY = &mokhtalat
		DataJ.IS_CTYPES = tools.ToBoolPointer(true)
	}
	log.Println("End Payment ..", *paymentNo)
	return DataJ, nil
}
func cancelledBillRequestPP(ctx *context.Context, in *pbMessages.CancelledBillRequestRequest, perfectreq *bool) (rsp *pbMessages.CancelledBillRequestResponse, err error) {
	if in.FormNo == nil {
		return nil, sendError(codes.Internal, "قم بادخال رقم الطلب", err.Error(), nil)
	}
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	if user.CANCEL_BILL == nil || !*user.CANCEL_BILL {
		return nil, errors.New("المستخدم لا يمتلك الصلاحية الكافية")
	}
	var cancelledBills irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledreqData []*dbmodels.CANCELLED_REQUEST
	cancelledreqData, err = cancelledBills.GetByFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}
	if len(cancelledreqData) == 0 {
		return nil, errors.New("لا يوجد مستند بالرقم " + *(tools.Int64ToString(in.FormNo)) + " للعميل ")
	} else if len(cancelledreqData) > 1 {
		return nil, errors.New("تكرار في رقم المستند")
	}
	var cancelledBillsData []*dbmodels.CANCELLED_BILL
	cancelledBillsData, err = cancelledBills.GetByBillsFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}
	var cancelledBillsAction irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledBillsActionData []*dbmodels.CANCELLED_BILLS_ACTION
	cancelledBillsActionData, err = cancelledBillsAction.GetByFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}
	var lucancelledBillsAction irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	DataJ := &pbMessages.CancelledBillRequestResponse{}
	REQUESTBY := ""
	if cancelledreqData[0].REQUEST_BY != nil {
		REQUESTBY = *cancelledreqData[0].REQUEST_BY
	}
	STATUS := ""
	if cancelledreqData[0].STATUS != nil {
		STATUS = *cancelledreqData[0].STATUS
	}
	SURNAME := ""
	if cancelledreqData[0].SURNAME != nil {
		SURNAME = *cancelledreqData[0].SURNAME
	}
	COMMENT := ""
	if cancelledreqData[0].COMMENT != nil {
		COMMENT = *cancelledreqData[0].COMMENT
	}
	DataJ.Content = &pbdbMessages.CANCELLED_REQUEST{
		STATION_NO:   tools.StringToInt32(&cancelledreqData[0].STATION_NO),
		FORM_NO:      &cancelledreqData[0].FORM_NO,
		CUSTKEY:      &cancelledreqData[0].CUSTKEY,
		DOCUMENT_NO:  &cancelledreqData[0].DOCUMENT_NO,
		REQUEST_BY:   &REQUESTBY,
		REQUEST_DATE: create_timestamp(cancelledreqData[0].REQUEST_DATE),
		STAMP_DATE:   create_timestamp(cancelledreqData[0].STAMP_DATE),
		COUNTER:      cancelledreqData[0].COUNTER,
		STATE:        cancelledreqData[0].STATE,
		CLOSED:       cancelledreqData[0].CLOSED,
		STATUS:       &STATUS,
		SURNAME:      &SURNAME,
		COMMENT:      &COMMENT,
		Actions:      []*pbdbMessages.CANCELLED_BILL_ACTION{},
		Bills:        []*pbdbMessages.CANCELLED_BILL{},
	}
	for idx := range cancelledBillsActionData {
		cancelledBillsActionDataUse := cancelledBillsActionData[idx]
		var lucancelledBillsActionData []*dbmodels.LU_CANCELLED_BILLS_ACTION
		lucancelledBillsActionData, err = lucancelledBillsAction.GetByID(cancelledBillsActionDataUse.ACTION_ID)
		if err != nil {
			return nil, err
		}
		if len(lucancelledBillsActionData) == 0 {
			return nil, errors.New("لا يوجد اكشن ")
		}
		desc := ""
		if lucancelledBillsActionData[0].DESCRIPTION != nil {
			desc = *lucancelledBillsActionData[0].DESCRIPTION
		}
		STAMP_USER := ""
		if cancelledBillsActionDataUse.STAMP_USER != nil {
			STAMP_USER = *cancelledBillsActionDataUse.STAMP_USER
		}
		COMMENT := ""
		if cancelledBillsActionDataUse.COMMENT != nil {
			COMMENT = *cancelledBillsActionDataUse.COMMENT
		}
		act := &pbdbMessages.CANCELLED_BILL_ACTION{
			FORM_NO:     &cancelledBillsActionDataUse.FORM_NO,
			ACTION_ID:   &cancelledBillsActionDataUse.ACTION_ID,
			DOCUMENT_NO: &cancelledBillsActionDataUse.DOCUMENT_NO,
			CUSTKEY:     &cancelledBillsActionDataUse.CUSTKEY,
			STAMP_DATE:  create_timestamp(cancelledBillsActionDataUse.STAMP_DATE),
			STAMP_USER:  &STAMP_USER,
			COMMENT:     &COMMENT,
			USER_ID:     cancelledBillsActionDataUse.USER_ID,
			DESCRIPTION: &desc,
		}
		DataJ.Content.Actions = append(DataJ.Content.Actions, act)
	}
	for idx := range cancelledBillsData {
		cancelledBillsDataUse := cancelledBillsData[idx]
		DOCUMENTNO := ""
		if cancelledBillsDataUse.DOCUMENT_NO != nil {
			DOCUMENTNO = *cancelledBillsDataUse.DOCUMENT_NO
		}
		COMMENT := ""
		if cancelledBillsDataUse.COMMENT != nil {
			COMMENT = *cancelledBillsDataUse.COMMENT
		}
		CANCELLEDBY := ""
		if cancelledBillsDataUse.CANCELLED_BY != nil {
			CANCELLEDBY = *cancelledBillsDataUse.CANCELLED_BY
		}
		SURNAME := ""
		if cancelledBillsDataUse.SURNAME != nil {
			SURNAME = *cancelledBillsDataUse.SURNAME
		}
		bill := &pbdbMessages.CANCELLED_BILL{
			FORM_NO:        &cancelledBillsDataUse.FORM_NO,
			DOCUMENT_NO:    &DOCUMENTNO,
			CUSTKEY:        &cancelledBillsDataUse.CUSTKEY,
			COMMENT:        &COMMENT,
			PAYMENT_NO:     &cancelledBillsDataUse.PAYMENT_NO,
			CL_BLNCE:       cancelledBillsDataUse.CL_BLNCE,
			CANCELLED_BY:   &CANCELLEDBY,
			CANCELLED_DATE: create_timestamp(cancelledBillsDataUse.BILNG_DATE),
			STATION_NO:     cancelledBillsDataUse.STATION_NO,
			SURNAME:        &SURNAME,
			BILNG_DATE:     create_timestamp(cancelledBillsDataUse.BILNG_DATE),
		}
		DataJ.Content.Bills = append(DataJ.Content.Bills, bill)
	}
	log.Println("End CancelledBillRequest..")
	*perfectreq = true
	return DataJ, nil
}
func cancelledBillActionPP(ctx *context.Context, in *pbMessages.CancelledBillActionRequest, perfectreq *bool) (rsp *pbMessages.CancelledBillActionResponse, err error) {
	if in.FormNo == nil {
		return nil, errors.New("رقم الطلب غير صحيح")
	}
	if in.Action == nil {
		return nil, errors.New("برجاء تحديد الاجراء")
	}
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	if user.CANCEL_BILL == nil || !*user.CANCEL_BILL {
		return nil, errors.New("المستخدم لا يمتلك الصلاحية الكافية")
	}
	var cancelledBills irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledreqData []*dbmodels.CANCELLED_REQUEST
	cancelledreqData, err = cancelledBills.GetByFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}
	if len(cancelledreqData) == 0 {
		return nil, errors.New("لا يوجد مستند بالرقم " + *(tools.Int64ToString(in.FormNo)) + " للعميل ")
	} else if len(cancelledreqData) > 1 {
		return nil, errors.New("تكرار في رقم المستند")
	} else if cancelledreqData[0].CLOSED != nil && *cancelledreqData[0].CLOSED {
		return nil, errors.New("الطلب مغلق")
	}
	var lucancelledBillsAction irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsActionData []*dbmodels.LU_CANCELLED_BILLS_ACTION
	lucancelledBillsActionData, err = lucancelledBillsAction.GetByID(*in.Action)
	if err != nil {
		return nil, err
	}
	if len(lucancelledBillsActionData) == 0 {
		return nil, errors.New("الاجراء غير معروف")
	}
	if lucancelledBillsActionData[0].CURRENT_STATE != nil && cancelledreqData[0].STATE != nil && *cancelledreqData[0].STATE != *lucancelledBillsActionData[0].CURRENT_STATE {
		return nil, errors.New("غير مسموح ب " + *lucancelledBillsActionData[0].DESCRIPTION + " الحالة الحالية لا تقبل الاجراء ")
	}
	if lucancelledBillsActionData[0].DEPARTMENT != nil {
		if user.DEPARTMENT == nil || *user.DEPARTMENT != *lucancelledBillsActionData[0].DEPARTMENT {
			return nil, errors.New("غير مسموح ب " + *lucancelledBillsActionData[0].DESCRIPTION + " ادارة مختلفة")
		}
	}
	var lucancelledBillsState irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsStateData []*dbmodels.LU_CANCELLED_BILL_STATE
	lucancelledBillsStateData, err = lucancelledBillsState.GetByID(*lucancelledBillsActionData[0].NEXT_STATE)
	if err != nil {
		return nil, err
	}
	if len(lucancelledBillsStateData) == 0 {
		return nil, errors.New("حالة غير معرفة للفاتورة " + *lucancelledBillsActionData[0].DESCRIPTION)
	}
	var cancelledBillsData []*dbmodels.CANCELLED_BILL
	cancelledBillsData, err = cancelledBills.GetByBillsFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}
	if len(cancelledBillsData) == 0 {
		return nil, errors.New("الطلب لا يوجد به فواتير")
	}
	var hand irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var hhcybc irespo.IHhhcycRepository = &respo.HhhcycRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	for idx := range cancelledBillsData {
		cancelledBillsDataUse := cancelledBillsData[idx]
		var handData []*dbmodels.HAND_MH_ST
		handData, err = hand.GetByPaymentNo(cancelledBillsDataUse.PAYMENT_NO)
		if err != nil {
			return nil, err
		}
		if len(handData) > 1 {
			return nil, errors.New("رقم القاتورة مكرر بقواعد البيانات " + cancelledBillsDataUse.PAYMENT_NO)
		}
		if len(handData) == 0 {
			return nil, errors.New("رقم القاتورة غير موجود " + cancelledBillsDataUse.PAYMENT_NO)
		}
		if handData[0] == nil {
			return nil, errors.New("رقم القاتورة غير موجود " + cancelledBillsDataUse.PAYMENT_NO)
		}
		if handData[0].IS_COLLECTION_ROW != nil && *handData[0].IS_COLLECTION_ROW == 1 {
			return nil, errors.New("الفاتورة غير ملغاه " + cancelledBillsDataUse.PAYMENT_NO)
		}
		if (tools.Int32PtrToInt64Ptr(handData[0].STATION_NO) != user.STATION_NO) && (station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS != 1) {
			return nil, errors.New("الفاتورة  تخص فرع اخر " + cancelledBillsDataUse.PAYMENT_NO)
		}
		var hhcybcData []*dbmodels.HH_BCYC
		hhcybcData, err = hhcybc.GetByStationNoBillGroupBookCWalkCCycleID(*handData[0].STATION_NO, *handData[0].BILLGROUP, *handData[0].BOOK_NO_C, *handData[0].WALK_NO_C, *handData[0].CYCLE_ID)
		if err != nil {
			return nil, err
		}
		if len(hhcybcData) != 0 {
			if hhcybcData[0] != nil {
				if hhcybcData[0].ISCYCLE_COMPLETED_C != nil && *hhcybcData[0].ISCYCLE_COMPLETED_C == 1 {
					return nil, errors.New("دورة التحصيل مغلقة " + cancelledBillsDataUse.PAYMENT_NO)
				}
			}
		}
	}
	cancelledreqData[0].STATE = lucancelledBillsActionData[0].NEXT_STATE
	cancelledreqData[0].STATUS = lucancelledBillsStateData[0].DESCRIPTION
	if lucancelledBillsActionData[0].CLOSED != nil && *lucancelledBillsActionData[0].CLOSED {
		cancelledreqData[0].CLOSED = lucancelledBillsActionData[0].CLOSED
	}
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
	}
	defer dbr.Rollback()
	//defer dbr.Close()
	// var cancelBillactionr irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	// err = cancelBillactionr.Upsert(&dbmodels.CANCELLED_BILLS_ACTION{
	// 	ACTION_ID:   lucancelledBillsActionData[0].ID,
	// 	DOCUMENT_NO: cancelledreqData[0].DOCUMENT_NO,
	// 	CUSTKEY:     cancelledreqData[0].CUSTKEY,
	// 	STAMP_DATE:  tools.ToTimePrt(time.Now()),
	// 	STAMP_USER:  user.USER_NAME,
	// 	USER_ID:     &user.ID,
	// 	COMMENT:     in.Comment,
	// 	FORM_NO:     cancelledreqData[0].FORM_NO,
	// })
	err = dbr.Add(&dbmodels.CANCELLED_BILLS_ACTION{
		ACTION_ID:   lucancelledBillsActionData[0].ID,
		DOCUMENT_NO: cancelledreqData[0].DOCUMENT_NO,
		CUSTKEY:     cancelledreqData[0].CUSTKEY,
		STAMP_DATE:  tools.ToTimePrt(time.Now()),
		STAMP_USER:  user.USER_NAME,
		USER_ID:     &user.ID,
		COMMENT:     in.Comment,
		FORM_NO:     cancelledreqData[0].FORM_NO,
	})
	if err != nil {
		return nil, err
	}
	err = dbr.Save(cancelledreqData[0])
	if err != nil {
		return nil, err
	}
	err = dbr.Commit()
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.CancelledBillActionResponse{}
	DataJ.Message = tools.ToStringPointer("Done")
	*perfectreq = true
	return DataJ, nil
}
func billActionsPP(ctx *context.Context, in *pbMessages.Empty, perfectreq *bool) (rsp *pbMessages.BillActionsResponse, err error) {
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var lucancelledBillsAction irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsActionData []*dbmodels.LU_CANCELLED_BILLS_ACTION
	lucancelledBillsActionData, err = lucancelledBillsAction.GetAll()
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.BillActionsResponse{Items: []*pbdbMessages.LU_CANCELLED_BILL_ACTION{}}
	for idx := range lucancelledBillsActionData {
		lucancelledBillsActionDataUse := lucancelledBillsActionData[idx]
		DataJ.Items = append(DataJ.Items, &pbdbMessages.LU_CANCELLED_BILL_ACTION{
			ID:            &lucancelledBillsActionDataUse.ID,
			DESCRIPTION:   lucancelledBillsActionDataUse.DESCRIPTION,
			CURRENT_STATE: lucancelledBillsActionDataUse.CURRENT_STATE,
			NEXT_STATE:    lucancelledBillsActionDataUse.NEXT_STATE,
			CLOSED:        lucancelledBillsActionDataUse.CLOSED,
			START_UP:      lucancelledBillsActionDataUse.START_UP,
			DEPARTMENT:    lucancelledBillsActionDataUse.DEPARTMENT,
		})
	}
	log.Println("End BillActions..")
	*perfectreq = true
	return DataJ, nil
}
func billStatesPP(ctx *context.Context, in *pbMessages.Empty, perfectreq *bool) (rsp *pbMessages.BillStatesResponse, err error) {
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error(), nil)
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var lucancelledStates irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	lucancelledStatesData, err := lucancelledStates.GetAll()
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.BillStatesResponse{Items: []*pbdbMessages.LU_CANCELLED_BILL_STATE{}}
	for idx := range lucancelledStatesData {
		lucancelledStatesDataUse := lucancelledStatesData[idx]
		DataJ.Items = append(DataJ.Items, &pbdbMessages.LU_CANCELLED_BILL_STATE{
			ID:          &lucancelledStatesDataUse.ID,
			DESCRIPTION: lucancelledStatesDataUse.DESCRIPTION,
			RECAL_READY: lucancelledStatesDataUse.RECAL_READY,
		})
	}
	log.Println("End BillStates..")
	*perfectreq = true
	return DataJ, nil
}
func saveBillCancelRequestPP(ctx *context.Context, in *pbMessages.SaveBillCancelRequestRequest, perfectreq *bool) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Aborted, err.Error(), err.Error(), nil)
	} else {
		//conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, sendError(codes.Aborted, err.Error(), err.Error(), nil)
	}
	if user.CANCEL_BILL == nil || !*user.CANCEL_BILL {
		return nil, errors.New("المستخدم لا يمتلك الصلاحية الكافية")
	}
	if in.Request == nil {
		return nil, errors.New("طلب خاطئ")
	}
	if len(in.Request.Bills) == 0 {
		return nil, errors.New("طلب خاطئ : على الاقل فاتورة في الطلب")
	}
	if in.Request.CUSTKEY == nil || strings.TrimSpace(*in.Request.CUSTKEY) == "" {
		return nil, errors.New("رقم الحساب غير صحيح")
	}
	if in.Request.COMMENT == nil || strings.TrimSpace(*in.Request.COMMENT) == "" {
		return nil, errors.New("برجاء كتابة تعليق")
	}
	if in.Request.DOCUMENT_NO == nil || strings.TrimSpace(*in.Request.DOCUMENT_NO) == "" {
		return nil, errors.New("رقم المستند غير صحيح")
	}

	var cancelBillReq irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	cancelledBillsReqData, err := cancelBillReq.GetByCustKeyDocNoNotFormNo(*in.Request.CUSTKEY, *in.Request.DOCUMENT_NO, *in.Request.FORM_NO)
	if err != nil {
		return nil, err
	}
	if len(cancelledBillsReqData) > 0 {
		return nil, errors.New("رقم المستند مستخدم بالفعل")
	}

	cancelledBillsReqSData, err := cancelBillReq.GetByCustKeyNotFormNo(*in.Request.CUSTKEY, *in.Request.FORM_NO)
	if err != nil {
		return nil, err
	}
	if len(cancelledBillsReqSData) > 0 {
		//var isCrossed []*string
		isCrossed := ""
		for idx := range cancelledBillsReqSData {
			cancelledBillsReqSDataUse := cancelledBillsReqSData[idx]
			for idxx := range in.Request.Bills {
				reqBill := in.Request.Bills[idxx]
				reqBillPAYMENT_NO := ""
				if reqBill.PAYMENT_NO != nil {
					reqBillPAYMENT_NO = *reqBill.PAYMENT_NO
				}
				fnd, err := cancelBillReq.GetByDocNoPaymentNo(cancelledBillsReqSDataUse.DOCUMENT_NO, reqBillPAYMENT_NO)
				if err != nil {
					return nil, err
				}
				if len(fnd) > 0 {
					if isCrossed != "" {
						isCrossed = isCrossed + "," + reqBillPAYMENT_NO
					} else {
						isCrossed = *reqBill.PAYMENT_NO
					}
				}
			}
		}
		if isCrossed != "" {
			return nil, errors.New("يوجد للعميل طلبات اخرى قيد التشغيل لا يمكن بدء الطلب : تخص نفس الفواتير او بعض منها " + isCrossed)
		}
	}

	log.Println("Done .... 1")
	var handbill irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	for idxx := range in.Request.Bills {
		reqBill := in.Request.Bills[idxx]
		cleanString(reqBill.PAYMENT_NO, nil, nil, nil)
		if reqBill.PAYMENT_NO == nil || strings.TrimSpace(*reqBill.PAYMENT_NO) == "" {
			return nil, errors.New("برجاء تحديد كود الفاتورة")
		}
		handbillData, err := handbill.GetPayment(reqBill.PAYMENT_NO, nil, in.Request.CUSTKEY, nil)
		if err != nil {
			return nil, err
		}
		if len(handbillData) > 1 {
			return nil, errors.New("رقم القاتورة مكرر بقواعد البيانات " + *reqBill.PAYMENT_NO)
		}
		if len(handbillData) == 0 {
			return nil, errors.New("رقم القاتورة غير موجود  " + *reqBill.PAYMENT_NO)
		}
	}
	log.Println("Done .... 2")
	handcstData, err := handbill.GetAllByCustkey(*in.Request.CUSTKEY)
	if err != nil {
		return nil, err
	}
	if len(handcstData) == 0 {
		return nil, sendError(codes.InvalidArgument, "لا يوجد رقم حساب عميل مسجل", "لا يوجد رقم حساب عميل مسجل", perfectreq)
	}
	timeNow := time.Now()
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
	}
	defer dbr.Rollback()
	//defer dbr.Close()
	ReqFormNo := int64(0)
	if in.Request.FORM_NO != nil {
		ReqFormNo = *in.Request.FORM_NO
	}
	ReqCUSTKEY := ""
	if in.Request.CUSTKEY != nil {
		ReqCUSTKEY = *in.Request.CUSTKEY
	}
	ReqStationNo := ""
	if in.Request.STATION_NO != nil {
		ReqStationNo = *tools.Int32ToString(in.Request.STATION_NO)
	}
	ReqDocNo := ""
	if in.Request.DOCUMENT_NO != nil {
		ReqDocNo = *in.Request.DOCUMENT_NO
	}
	reqsave := &dbmodels.CANCELLED_REQUEST{
		FORM_NO:      ReqFormNo,
		CUSTKEY:      ReqCUSTKEY,
		STATION_NO:   ReqStationNo,
		DOCUMENT_NO:  ReqDocNo,
		REQUEST_DATE: create_time(in.Request.REQUEST_DATE),
		REQUEST_BY:   in.Request.REQUEST_BY,
		STATE:        in.Request.STATE,
		CLOSED:       in.Request.CLOSED,
		STATUS:       in.Request.STATUS,
		COMMENT:      in.Request.COMMENT,
		COUNTER:      in.Request.COUNTER,
		SURNAME:      in.Request.SURNAME,
		STAMP_DATE:   create_time(in.Request.STAMP_DATE),
	}
	log.Println("Done .... 3")
	if in.Request.FORM_NO == nil && *in.Request.FORM_NO == 0 {
		nextFormNo, err := cancelBillReq.GetMax("FORM_NO")
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if nextFormNo == nil {
			return nil, sendError(codes.InvalidArgument, "لم يتم احتساب رقم الطلب", "لم يتم احتساب رقم الطلب", perfectreq)
		}
		reqsave.STAMP_DATE = &timeNow
		reqsave.FORM_NO = 1 + *nextFormNo
		reqsave.CLOSED = tools.ToBoolPointer(false)
		reqsave.COUNTER = tools.Int32ToInt32Ptr(0)
		err = dbr.Add(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		log.Println("Done .... 4", nextFormNo)
	} else {
		prevStms, err := cancelBillReq.GetByBillsFormNo(reqsave.FORM_NO)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		for idxb := range prevStms {
			prevStmsUse := prevStms[idxb]
			removeb := true
			for idxs := range in.Request.Bills {
				rebill := in.Request.Bills[idxs]
				if rebill.PAYMENT_NO != nil && *rebill.PAYMENT_NO == prevStmsUse.PAYMENT_NO {
					removeb = false
					break
				}
			}
			if removeb {
				err := dbr.Delete(prevStmsUse)
				if err != nil {
					return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
				}
			}
		}
	}
	cstSTATION_NO := ""
	if handcstData[0].STATION_NO != nil {
		cstSTATION_NO = *tools.Int32ToString(handcstData[0].STATION_NO)
	}
	reqsave.REQUEST_BY = user.USER_NAME
	reqsave.STATION_NO = cstSTATION_NO
	reqsave.SURNAME = handcstData[0].Tent_name
	if reqsave.REQUEST_DATE == nil {
		reqsave.REQUEST_DATE = &timeNow
	}

	var satatesr irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var intiActr irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var actionr irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var bcycr irespo.IHhhcycRepository = &respo.HhhcycRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var tracr irespo.IStatmenrTracerRepository = &respo.StatmenrTracerRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var archand irespo.IArcHandMhStRepository = &respo.ArcHandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	actionsData, err := actionr.GetByFormNo(reqsave.FORM_NO)
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
	}
	if len(actionsData) > 2 {
		return nil, sendError(codes.AlreadyExists, "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب", "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب", perfectreq)
	}
	if reqsave.STATE == nil || *reqsave.STATE == 0 {
		stateId := int32(0)
		intiActData, err := intiActr.GetByStartUp(true)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(actionsData) == 0 {
			if len(intiActData) != 0 {
				if intiActData[0].NEXT_STATE != nil {
					stateId = *intiActData[0].NEXT_STATE
				}
				cancelbillactionSave := &dbmodels.CANCELLED_BILLS_ACTION{
					FORM_NO:     reqsave.FORM_NO,
					CUSTKEY:     ReqCUSTKEY,
					STAMP_DATE:  &timeNow,
					STAMP_USER:  user.USER_NAME,
					COMMENT:     tools.ToStringPointer("تم ايقاف الفاتورة"),
					ACTION_ID:   intiActData[0].ID,
					DOCUMENT_NO: ReqDocNo,
				}
				dbr.Add(cancelbillactionSave)
				if err != nil {
					return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
				}
			} else {
				return nil, sendError(codes.InvalidArgument, "برجاء تعريف الاجراء الابتدائي للعملية", "برجاء تعريف الاجراء الابتدائي للعملية", perfectreq)
			}
		} else {
			stateIdData, err := intiActr.GetByID(actionsData[0].ACTION_ID)
			if err != nil {
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
			}
			if len(stateIdData) > 0 {
				if stateIdData[0].CURRENT_STATE != nil {
					stateId = *stateIdData[0].CURRENT_STATE
				}
			}
		}
		stateData, err := satatesr.GetByID(stateId)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(stateData) == 0 {
			return nil, sendError(codes.InvalidArgument, "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId), "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId), perfectreq)
		}
		reqsave.STATUS = stateData[0].DESCRIPTION
		reqsave.STATE = &stateData[0].ID
		err = dbr.Save(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
	}
	for idxb := range in.Request.Bills {
		pay := in.Request.Bills[idxb]
		payPAYMENT_NO := ""
		if pay.PAYMENT_NO != nil {
			payPAYMENT_NO = *pay.PAYMENT_NO
		}
		handData, err := handbill.GetByPaymentNoCustkey(payPAYMENT_NO, ReqCUSTKEY)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(handData) == 0 {
			return nil, sendError(codes.InvalidArgument, "رقم القاتورة غير موجود  "+*pay.PAYMENT_NO, "رقم القاتورة غير موجود  "+*pay.PAYMENT_NO, perfectreq)
		}
		err = throwsIfStationNoInvalied(user, handData[0].STATION_NO, conn)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		// Check
		bcycData, err := bcycr.GetByStationNoBillGroupBookCWalkCCycleID(*handData[0].STATION_NO, *handData[0].BILLGROUP, *handData[0].BOOK_NO_C, *handData[0].WALK_NO_C, *handData[0].CYCLE_ID)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(bcycData) > 0 {
			if bcycData[0].ISCYCLE_COMPLETED_C != nil && *bcycData[0].ISCYCLE_COMPLETED_C == 1 {
				return nil, sendError(codes.Unavailable, "دورة التحصيل مغلقة   "+*pay.PAYMENT_NO, "دورة التحصيل مغلقة   "+*pay.PAYMENT_NO, perfectreq)
			}
		}
		cancelledBillsData, err := cancelBillReq.GetByDocNoCustKeyPaymentNo(*in.Request.DOCUMENT_NO, handData[0].CUSTKEY, *pay.PAYMENT_NO)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(cancelledBillsData) > 0 {
			cancRow := cancelledBillsData[0]
			cancRow.COMMENT = pay.COMMENT
			cancRow.CANCELLED_DATE = &timeNow
			cancRow.CANCELLED_BY = user.USER_NAME
			dbr.Save(cancRow)
		} else {
			billRow := &dbmodels.CANCELLED_BILL{
				DOCUMENT_NO:    in.Request.DOCUMENT_NO,
				COMMENT:        pay.COMMENT,
				CUSTKEY:        handData[0].CUSTKEY,
				PAYMENT_NO:     *handData[0].Payment_no,
				CL_BLNCE:       handData[0].Cl_blnce,
				CANCELLED_DATE: &timeNow,
				CANCELLED_BY:   user.USER_NAME,
				STATION_NO:     handData[0].STATION_NO,
				SURNAME:        handData[0].Tent_name,
				BILNG_DATE:     handData[0].BILNG_DATE,
				FORM_NO:        reqsave.FORM_NO,
			}
			dbr.Add(billRow)
			if err != nil {
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
			}
		}
		hrecData, err := handbill.GetByPaymentNoCustkey(*handData[0].Payment_no, handData[0].CUSTKEY)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
		if len(hrecData) > 0 {
			hrec := hrecData[0]
			hrec.IS_COLLECTION_ROW = tools.Int32ToInt32Ptr(0)
			hrec.NOTE_C = tools.ToStringPointer("ملغاة")
			hrec.GARD = nil
			hrec.Delivery_st = nil
			hrec.AMOUNT_COLLECTED = nil
			hrec.COLLECTION_DATE = nil
			hrec.COLLECTION_DEVICEID = nil
			hrec.GARD_PAYMENT_NO = nil
			dbr.Save(hrec)
		} else {
			archandData, err := archand.GetByPaymentNoCustkey(*handData[0].Payment_no, handData[0].CUSTKEY)
			if err != nil {
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
			}
			arc_hrec := archandData[0]
			arc_hrec.IS_COLLECTION_ROW = tools.Int32ToInt32Ptr(0)
			arc_hrec.NOTE_C = tools.ToStringPointer("ملغاة")
			arc_hrec.GARD = nil
			arc_hrec.Delivery_st = nil
			arc_hrec.AMOUNT_COLLECTED = nil
			arc_hrec.COLLECTION_DATE = nil
			arc_hrec.COLLECTION_DEVICEID = nil
			dbr.Save(arc_hrec)
		}
		err = tracr.AddStatmentAction(dbr, handData[0].Payment_no, &handData[0].CUSTKEY, tools.Int32ToInt32Ptr(int32(*handData[0].EMPID_C)), true)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
		}
	}
	err = dbr.Commit()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error(), perfectreq)
	}
	DataJ := &pbMessages.SaveBillCancelRequestResponse{Message: tools.ToStringPointer("Done")}

	log.Println("End SaveBillCancelRequest..")
	*perfectreq = true
	return DataJ, nil
}
