package masservices

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"context"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MaisrForAdvancedSystems/go-biller-proto/go/dbmessages"
	pbdbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/dbmessages"
	pbMessages "github.com/MaisrForAdvancedSystems/go-biller-proto/go/messages"
	"github.com/MaisrForAdvancedSystems/go-biller-proto/go/serverhostmessages"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbpool"
	irespo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/interfaces"
	respo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/repositories"
	"golang.org/x/sync/syncmap"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func cancelledBillListP(ctx *context.Context, in *pbMessages.CancelledBillListRequest) (rsp *pbMessages.CancelledBillListResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
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
	var cancelledBillsAction irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsAction irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsState irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledBillsData []*dbmodels.CANCELLED_REQUEST
	inclose := false
	if in.State == nil {
		inclose = true
	} else {
		checkOpened, err := lucancelledBillsAction.GetByNextState(*in.State)
		if err != nil {
			return nil, err
		}
		for idxcheckes := range checkOpened {
			if checkOpened[idxcheckes].CLOSED != nil && *checkOpened[idxcheckes].CLOSED {
				inclose = true
				break
			}
		}
	}
	if in.State != nil && (station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0) {
		cancelledBillsData, err = cancelledBills.GetByClosedStatusStation(false, *in.State, station.STATION_NO, inclose, 1)
	} else if in.State != nil {
		cancelledBillsData, err = cancelledBills.GetByClosedStatus(false, *in.State, inclose, 1)
	} else if station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0 {
		cancelledBillsData, err = cancelledBills.GetByClosedStation(false, station.STATION_NO, inclose, 1)
	} else {
		cancelledBillsData, err = cancelledBills.GetByClosed(false, inclose, 1)
	}

	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.CancelledBillListResponse{}
	stringEmpty := ""
	// for idx := range cancelledBillsData {
	// 	usj := &pbdbMessages.CANCELLED_REQUEST{}
	// 	cancelledBillsUse := cancelledBillsData[idx]
	// 	lastaction, err := cancelledBillsAction.GetByFormNoWithState(cancelledBillsUse.FORM_NO, in.State)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if len(lastaction) > 0 {
	// 		usj.STAMP_DATE = create_timestamp(lastaction[0].STAMP_DATE)
	// 	} else {
	// 		usj.STAMP_DATE = create_timestamp(cancelledBillsUse.STAMP_DATE)
	// 	}
	// 	usj.CLOSED = cancelledBillsUse.CLOSED
	// 	if cancelledBillsUse.COMMENT == nil {
	// 		usj.COMMENT = &stringEmpty
	// 	} else {
	// 		usj.COMMENT = cancelledBillsUse.COMMENT
	// 	}
	// 	usj.CUSTKEY = &cancelledBillsUse.CUSTKEY
	// 	usj.DOCUMENT_NO = &cancelledBillsUse.DOCUMENT_NO
	// 	usj.FORM_NO = &cancelledBillsUse.FORM_NO
	// 	usj.REQUEST_BY = cancelledBillsUse.REQUEST_BY
	// 	usj.REQUEST_DATE = create_timestamp(cancelledBillsUse.REQUEST_DATE)
	// 	usj.STATE = cancelledBillsUse.STATE
	// 	usj.STATION_NO = tools.StringToInt32(&cancelledBillsUse.STATION_NO)
	// 	usj.STATUS = cancelledBillsUse.STATUS
	// 	usj.SURNAME = cancelledBillsUse.SURNAME

	// 	DataJ.CancelledBillList = append(DataJ.CancelledBillList, usj)
	// }
	//lastactions := make(map[int64]*timestamppb.Timestamp)
	lastactions := syncmap.Map{}
	lastactionsStates := syncmap.Map{}
	var wg sync.WaitGroup
	var erro error
	type formState struct {
		cancel *bool
		edit   *bool
	}
	for idx := range cancelledBillsData {
		cancelledBillsUsee := cancelledBillsData[idx]
		wg.Add(1)
		go func(wgg *sync.WaitGroup, cancelledBillsUse *dbmodels.CANCELLED_REQUEST) {
			defer wgg.Done()
			if erro != nil {
				return
			}
			lastaction, err := cancelledBillsAction.GetByFormNoWithState(cancelledBillsUse.FORM_NO, in.State)
			if err != nil {
				erro = err
				return
			}
			formStateReq := formState{cancel: nil, edit: nil}
			if cancelledBillsUse.STATE != nil {
				laststate, err := lucancelledBillsState.GetByID(*cancelledBillsUse.STATE)
				if err != nil {
					erro = err
					return
				}
				formStateReq = formState{cancel: laststate[0].CANCELLED, edit: laststate[0].EDITED}
			}
			var stamp *timestamppb.Timestamp
			if len(lastaction) > 0 {
				stamp = create_timestamp(lastaction[0].STAMP_DATE)
			} else {
				stamp = create_timestamp(cancelledBillsUse.STAMP_DATE)
			}
			//lastactions[cancelledBillsUse.FORM_NO] = stamp
			lastactions.Store(cancelledBillsUse.FORM_NO, stamp)
			lastactionsStates.Store(cancelledBillsUse.FORM_NO, formStateReq)
		}(&wg, cancelledBillsUsee)
	}
	wg.Wait()
	if erro != nil {
		return nil, erro
	}
	for idx := range cancelledBillsData {
		usj := &pbdbMessages.CANCELLED_REQUEST{}
		cancelledBillsUse := cancelledBillsData[idx]
		//usj.STAMP_DATE = lastactions[cancelledBillsUse.FORM_NO]
		v, ok := lastactions.Load(cancelledBillsUse.FORM_NO)
		if ok {
			usj.STAMP_DATE = v.(*timestamppb.Timestamp)
		}
		vstate, ok := lastactionsStates.Load(cancelledBillsUse.FORM_NO)
		if ok {
			usj.CANCELLED = vstate.(formState).cancel
			usj.EDITED = vstate.(formState).edit
		}
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
		usj.STATE = cancelledBillsUse.STATE
		usj.STATION_NO = tools.StringToInt32(&cancelledBillsUse.STATION_NO)
		usj.STATUS = cancelledBillsUse.STATUS
		usj.SURNAME = cancelledBillsUse.SURNAME

		DataJ.CancelledBillList = append(DataJ.CancelledBillList, usj)
	}
	sort.SliceStable(DataJ.CancelledBillList, func(i, j int) bool {
		if DataJ.CancelledBillList[i].STAMP_DATE == nil && DataJ.CancelledBillList[j].STAMP_DATE == nil {
			return false
		} else if DataJ.CancelledBillList[i].STAMP_DATE == nil {
			return false
		} else if DataJ.CancelledBillList[j].STAMP_DATE == nil {
			return true
		}
		return (*DataJ.CancelledBillList[i].STAMP_DATE).AsTime().After((*DataJ.CancelledBillList[j].STAMP_DATE).AsTime())
	})
	log.Println("End cancelledBillList..")
	return DataJ, nil
}
func getPaymentP(ctx *context.Context, in *pbMessages.GetPaymentRequest) (rsp *pbMessages.GetPaymentResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	var hand irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var ctgConTypeGr irespo.ICtgConsumptionTypeGroupsRepository = &respo.CtgConsumptionTypeGroupsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	ctgData, err := ctgConTypeGr.GetAll()
	if err != nil {
		return nil, err
	}
	usj, err := getPayment(in.PaymentNo, in.Custkey, in.SkipBracodTrim, in.ForQuery, in.CycleId, &hand, user, ctgData, conn, station, nil, tools.ToBoolPointer(false))
	if err != nil {
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	}
	DataJ := &pbMessages.GetPaymentResponse{}
	DataJ.Item = usj
	return DataJ, nil
}
func getCustomerPaymentsP(ctx *context.Context, in *pbMessages.GetCustomerPaymentsRequest) (rsp *pbMessages.GetCustomerPaymentsResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
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
	if len(handData) == 0 {
		return nil, errors.New("رقم الحساب غير صحيح او لا توجد اي فواتير للعميل " + *in.Custkey)
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	/*var stationNo *int32 = nil
	if !(station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 1) {
		stationNo = &station.STATION_NO
	}*/
	var cancelbill irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	openRequests, err := cancelbill.GetByCustKeyClosed(*in.Custkey, false, 1)
	if err != nil {
		return nil, err
	}
	//openRequestsList := []int64{}
	finalpayment := []*dbmodels.HAND_MH_ST{}
	if len(openRequests) > 0 {
		formNoString := ""
		for idxform := range openRequests {
			openreqUse := openRequests[idxform]
			if in.FormNo != nil && openreqUse.FORM_NO == *in.FormNo {
				continue
			}
			if formNoString != "" {
				formNoString += "," + *tools.Int64ToString(&openreqUse.FORM_NO)
			} else {
				formNoString = *tools.Int64ToString(&openreqUse.FORM_NO)
			}
		}
		if formNoString != "" {
			for idxpay := range handData {
				handpay := handData[idxpay]
				isCanp, err := cancelbill.ExistCancelBill(handpay.CUSTKEY, *handpay.Payment_no, formNoString)
				if err != nil {
					return nil, err
				}
				if !isCanp {
					finalpayment = append(finalpayment, handpay)
				} else {
					log.Println(handpay.Payment_no)
				}
			}
		} else {
			finalpayment = handData
		}
	} else {
		finalpayment = handData
	}
	DataJ := &pbMessages.GetCustomerPaymentsResponse{}
	//var wg sync.WaitGroup
	var recep irespo.IReciptsRepository = &respo.ReciptsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	for idx := range finalpayment {
		//handDataUsee := finalpayment[idx]
		// wg.Add(1)
		// go func(wge *sync.WaitGroup, handDataUse *dbmodels.HAND_MH_ST) error {
		// 	defer wge.Done()
		// 	usj, err := getPayment(handDataUse.Payment_no, &handDataUse.CUSTKEY, tools.ToBoolPointer(true), nil, nil, stationNo, &hand, user, ctgData, conn, station, in.FormNo)
		// 	if err != nil {
		// 		return sendError(codes.Internal, err.Error(), err.Error())
		// 	}
		// 	DataJ.Items = append(DataJ.Items, usj)
		// 	return nil
		// }(&wg, handDataUsee)
		handDataUse := finalpayment[idx]
		if handDataUse.Payment_no == nil {
			continue
		}
		countReceipts, err := recep.GetCountByPaymentNoCancelled(*handDataUse.Payment_no, false)
		if err != nil {
			return nil, err
		}
		if countReceipts > 0 {
			continue
		}
		usj, err := getPayment(handDataUse.Payment_no, &handDataUse.CUSTKEY, tools.ToBoolPointer(true), nil, nil, &hand, user, ctgData, conn, station, in.FormNo, tools.ToBoolPointer(false))
		if err != nil {
			return nil, sendError(codes.Internal, err.Error(), err.Error())
		}
		DataJ.Items = append(DataJ.Items, usj)
	}
	//wg.Wait()
	log.Println("end ..")
	return DataJ, nil
}
func cancelledBillRequestP(ctx *context.Context, in *pbMessages.CancelledBillRequestRequest) (rsp *pbMessages.CancelledBillRequestResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	if in.FormNo == nil {
		return nil, sendError(codes.Internal, "قم بادخال رقم الطلب", err.Error())
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
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
	cancelledreqData, err = cancelledBills.GetByFormNo(*in.FormNo, 1)
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
	return DataJ, nil
}
func cancelledBillActionP(ctx *context.Context, in *pbMessages.CancelledBillActionRequest) (rsp *pbMessages.CancelledBillActionResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
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
	cancelledreqData, err = cancelledBills.GetByFormNo(*in.FormNo, 1)
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
	var customerwalks irespo.ICustomerWalksRepository = &respo.CustomerWalksRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	defer dbr.Rollback()
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
		//open payment for collection with last action
		actClose := false
		stmCl_blnce := float64(-9999)
		if lucancelledBillsActionData[0].CLOSED != nil {
			actClose = *lucancelledBillsActionData[0].CLOSED
		}
		if handData[0].Cl_blnce != nil {
			stmCl_blnce = *handData[0].Cl_blnce
		}
		if actClose && stmCl_blnce >= 0 {
			handDataCheck, err := hand.GetByCustKeyStationNoCycleID(handData[0].STATION_NO, handData[0].CUSTKEY, handData[0].CYCLE_ID)
			if err != nil {
				return nil, err
			}
			if handDataCheck != nil {
				handDataCheck.IS_COLLECTION_ROW = tools.Int32ToInt32Ptr(1)
				handDataCheck.Delivery_st = tools.ToIntPointer(0)
				//always update empid
				hhcybcChechData, err := hhcybc.GetByBillGroupBookCWalkCCycleID(*handDataCheck.BILLGROUP, *handDataCheck.BOOK_NO_C, *handDataCheck.WALK_NO_C, *handDataCheck.CYCLE_ID)
				if err != nil {
					return nil, err
				}
				var empic *int64
				if len(hhcybcChechData) > 0 {
					empic = hhcybcChechData[0].EMPID_C
				}
				if empic != nil {
					handDataCheck.EMPID_C = empic
				} else {
					customerwalksData, err := customerwalks.GetByBillGroupBookNoWalkNo(*handDataCheck.BILLGROUP, *handDataCheck.BOOK_NO_C, *handDataCheck.WALK_NO_C)
					if err != nil {
						return nil, err
					}
					var assignto *int64
					if len(customerwalksData) > 0 {
						assignto = customerwalksData[0].ASSIGNED_TO_HH
					}
					if assignto != nil {
						handDataCheck.EMPID_C = assignto
					} else {
						return nil, errors.New("برجاء مراجعة العهدة للمسار حيث لم يتمكن النظام من تخصيص الفاتورة لمحصل")
					}
				}
				dbr.Save(handDataCheck)
			}
		}
	}
	dateNow := time.Now()
	cancelledreqData[0].STATE = lucancelledBillsActionData[0].NEXT_STATE
	cancelledreqData[0].STATUS = lucancelledBillsStateData[0].DESCRIPTION
	cancelledreqData[0].STAMP_DATE = tools.ToTimePrt(dateNow)
	if lucancelledBillsActionData[0].CLOSED != nil && *lucancelledBillsActionData[0].CLOSED {
		cancelledreqData[0].CLOSED = lucancelledBillsActionData[0].CLOSED
	}

	err = dbr.Add(&dbmodels.CANCELLED_BILLS_ACTION{
		ACTION_ID:   lucancelledBillsActionData[0].ID,
		DOCUMENT_NO: cancelledreqData[0].DOCUMENT_NO,
		CUSTKEY:     cancelledreqData[0].CUSTKEY,
		STAMP_DATE:  tools.ToTimePrt(dateNow),
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
	return DataJ, nil
}
func billActionsP(ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillActionsResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var lucancelledBillsAction irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var lucancelledBillsActionData []*dbmodels.LU_CANCELLED_BILLS_ACTION
	lucancelledBillsActionData, err = lucancelledBillsAction.GetByApplicationTypeID(1)
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.BillActionsResponse{Items: []*pbdbMessages.LU_CANCELLED_BILL_ACTION{}}
	for idx := range lucancelledBillsActionData {
		lucancelledBillsActionDataUse := lucancelledBillsActionData[idx]
		DataJ.Items = append(DataJ.Items, &pbdbMessages.LU_CANCELLED_BILL_ACTION{
			ID:              &lucancelledBillsActionDataUse.ID,
			DESCRIPTION:     lucancelledBillsActionDataUse.DESCRIPTION,
			CURRENT_STATE:   lucancelledBillsActionDataUse.CURRENT_STATE,
			NEXT_STATE:      lucancelledBillsActionDataUse.NEXT_STATE,
			CLOSED:          lucancelledBillsActionDataUse.CLOSED,
			START_UP:        lucancelledBillsActionDataUse.START_UP,
			DEPARTMENT:      lucancelledBillsActionDataUse.DEPARTMENT,
			ApplicationType: lucancelledBillsActionDataUse.APPLICATION_TYPE_ID,
		})
	}
	log.Println("End BillActions..")
	return DataJ, nil
}
func billStatesP(ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.BillStatesResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var lucancelledStates irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	lucancelledStatesData, err := lucancelledStates.GetByApplicationTypeID(1)
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.BillStatesResponse{Items: []*pbdbMessages.LU_CANCELLED_BILL_STATE{}}
	for idx := range lucancelledStatesData {
		lucancelledStatesDataUse := lucancelledStatesData[idx]
		DataJ.Items = append(DataJ.Items, &pbdbMessages.LU_CANCELLED_BILL_STATE{
			ID:              &lucancelledStatesDataUse.ID,
			DESCRIPTION:     lucancelledStatesDataUse.DESCRIPTION,
			RECAL_READY:     lucancelledStatesDataUse.RECAL_READY,
			CANCELLED:       lucancelledStatesDataUse.CANCELLED,
			EDITED:          lucancelledStatesDataUse.EDITED,
			ApplicationType: lucancelledStatesDataUse.APPLICATION_TYPE_ID,
		})
	}
	log.Println("End BillStates..")
	return DataJ, nil
}
func saveBillCancelRequestP(ctx *context.Context, in *pbMessages.SaveBillCancelRequestRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	if in.Request == nil {
		return nil, errors.New("طلب خاطئ")
	}
	if in.Request.ApplicationType == nil {
		in.Request.ApplicationType = tools.Int32ToInt32Ptr(1)
	}
	// if in.Request.ApplicationType == nil {
	// 	return nil, errors.New("نوع الطلب غير محدد")
	// }
	if *in.Request.ApplicationType == 1 {
		return saveBillCancelRequestType(ctx, in)
	}
	return saveAppication(ctx, in)
}
func saveBillCancelRequestType(ctx *context.Context, in *pbMessages.SaveBillCancelRequestRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	username, _ := (*ctx).Value("username").(string)
	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Aborted, err.Error(), err.Error())
	} else {
		//conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, sendError(codes.Aborted, err.Error(), err.Error())
	}
	if user.CANCEL_BILL == nil || !*user.CANCEL_BILL {
		return nil, errors.New("المستخدم لا يمتلك الصلاحية الكافية")
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
	var apptype irespo.IApplicationTypesRepository = &respo.ApplicationTypesRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	formN := int64(0)
	if in.Request.FORM_NO != nil {
		formN = *in.Request.FORM_NO
	}
	apptypeData, err := apptype.GetTypeByID(*in.Request.ApplicationType)
	if err != nil {
		return nil, err
	}
	if apptypeData == nil {
		return nil, errors.New("نوع الطلب غير معرف")
	}
	cancelledBillsReqData, err := cancelBillReq.GetByCustKeyDocNoNotFormNo(*in.Request.CUSTKEY, *in.Request.DOCUMENT_NO, formN, 1)
	if err != nil {
		return nil, err
	}
	if len(cancelledBillsReqData) > 0 {
		return nil, errors.New("رقم المستند مستخدم بالفعل")
	}

	cancelledBillsReqSData, err := cancelBillReq.GetByCustKeyNotFormNo(*in.Request.CUSTKEY, formN, 1)
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

	var handbill irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var recep irespo.IReciptsRepository = &respo.ReciptsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
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
		if handbillData[0] == nil {
			return nil, errors.New("رقم القاتورة غير موجود  " + *reqBill.PAYMENT_NO)
		}
		countReceipts, err := recep.GetCountByPaymentNoCancelled(*reqBill.PAYMENT_NO, false)
		if err != nil {
			return nil, err
		}
		if countReceipts > 0 {
			return nil, errors.New("الفاتورة تمت عليها عملية تحصيل  " + *reqBill.PAYMENT_NO)
		}
		// collectedAmount := float64(0)
		// collectedAmountData, err := recep.GetByCustKeyCycleIDCancelled(*handbillData[0].CUSTKEY, handbillData[0].CYCLE_ID, false)
		// if err != nil {
		// 	return nil, err
		// }
		// for idx := range collectedAmountData {
		// 	recepDataUse := collectedAmountData[idx]
		// 	collectedAmount += recepDataUse.AMOUNT
		// }
		// if collectedAmount > 0.1 {
		// 	return nil, errors.New("الفاتورة تمت عليها عملية تحصيل  " + *reqBill.PAYMENT_NO)
		// }
	}

	handcstData, err := handbill.GetAllByCustkey(*in.Request.CUSTKEY)
	if err != nil {
		return nil, err
	}
	if len(handcstData) == 0 {
		return nil, sendError(codes.InvalidArgument, "لا يوجد رقم حساب عميل مسجل", "لا يوجد رقم حساب عميل مسجل")
	}
	timeNow := time.Now()
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
	if in.Request.ApplicationType == nil {
		in.Request.ApplicationType = tools.Int32ToInt32Ptr(1)
	}
	reqsave := &dbmodels.CANCELLED_REQUEST{
		FORM_NO:             ReqFormNo,
		APPLICATION_TYPE_ID: *in.Request.ApplicationType,
		CANCELLED:           tools.ToBoolPointer(false),
		CUSTKEY:             ReqCUSTKEY,
		STATION_NO:          ReqStationNo,
		DOCUMENT_NO:         ReqDocNo,
		REQUEST_DATE:        create_time(in.Request.REQUEST_DATE),
		REQUEST_BY:          in.Request.REQUEST_BY,
		STATE:               in.Request.STATE,
		CLOSED:              in.Request.CLOSED,
		STATUS:              in.Request.STATUS,
		COMMENT:             in.Request.COMMENT,
		COUNTER:             in.Request.COUNTER,
		SURNAME:             in.Request.SURNAME,
		STAMP_DATE:          create_time(in.Request.STAMP_DATE),
	}

	if in.Request.FORM_NO == nil || (in.Request.FORM_NO != nil && *in.Request.FORM_NO == 0) {
		nextFormNo, err := cancelBillReq.GetMax("FORM_NO", 1)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if nextFormNo == nil {
			return nil, sendError(codes.InvalidArgument, "لم يتم احتساب رقم الطلب", "لم يتم احتساب رقم الطلب")
		}
		reqsave.STAMP_DATE = &timeNow
		reqsave.FORM_NO = 1 + *nextFormNo
		reqsave.CLOSED = tools.ToBoolPointer(false)
		reqsave.COUNTER = tools.Int32ToInt32Ptr(0)

	} else {
		prevStms, err := cancelBillReq.GetByBillsFormNo(reqsave.FORM_NO)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
					return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	if len(actionsData) > 2 {
		return nil, sendError(codes.AlreadyExists, "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب", "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب")
	}
	if reqsave.STATE == nil || *reqsave.STATE == 0 {
		stateId := int32(0)
		intiActData, err := intiActr.GetByStartUp(true, *in.Request.ApplicationType)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
					return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
				}
			} else {
				return nil, sendError(codes.InvalidArgument, "برجاء تعريف الاجراء الابتدائي للعملية", "برجاء تعريف الاجراء الابتدائي للعملية")
			}
		} else {
			stateIdData, err := intiActr.GetByID(actionsData[0].ACTION_ID)
			if err != nil {
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
			}
			if len(stateIdData) > 0 {
				if stateIdData[0].CURRENT_STATE != nil {
					stateId = *stateIdData[0].CURRENT_STATE
				}
			}
		}
		stateData, err := satatesr.GetByID(stateId)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if len(stateData) == 0 {
			return nil, sendError(codes.InvalidArgument, "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId), "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId))
		}
		reqsave.STATUS = stateData[0].DESCRIPTION
		reqsave.STATE = &stateData[0].ID
	}
	if in.Request.FORM_NO == nil || (in.Request.FORM_NO != nil && *in.Request.FORM_NO == 0) {
		err = dbr.Add(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
	} else {
		err = dbr.Save(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if len(handData) == 0 {
			return nil, sendError(codes.InvalidArgument, "رقم القاتورة غير موجود  "+*pay.PAYMENT_NO, "رقم القاتورة غير موجود  "+*pay.PAYMENT_NO)
		}
		err = throwsIfStationNoInvalied(user, handData[0].STATION_NO, conn)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		// Check
		bcycData, err := bcycr.GetByStationNoBillGroupBookCWalkCCycleID(*handData[0].STATION_NO, *handData[0].BILLGROUP, *handData[0].BOOK_NO_C, *handData[0].WALK_NO_C, *handData[0].CYCLE_ID)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if len(bcycData) > 0 {
			if bcycData[0].ISCYCLE_COMPLETED_C != nil && *bcycData[0].ISCYCLE_COMPLETED_C == 1 {
				return nil, sendError(codes.Unavailable, "دورة التحصيل مغلقة   "+*pay.PAYMENT_NO, "دورة التحصيل مغلقة   "+*pay.PAYMENT_NO)
			}
		}
		cancelledBillsData, err := cancelBillReq.GetByDocNoCustKeyPaymentNo(*in.Request.DOCUMENT_NO, handData[0].CUSTKEY, *pay.PAYMENT_NO)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
			}
		}
		hrecData, err := handbill.GetByPaymentNoCustkey(*handData[0].Payment_no, handData[0].CUSTKEY)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
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
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
	}

	// Add Entries
	// for k, v := range in.Request.Entries {

	// }
	err = dbr.Commit()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	DataJ := &pbMessages.SaveBillCancelRequestResponse{Message: tools.ToStringPointer("Done")}

	log.Println("End SaveBillCancelRequest..")
	return DataJ, nil
}
func saveAppication(ctx *context.Context, in *pbMessages.SaveBillCancelRequestRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	username, _ := (*ctx).Value("username").(string)
	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Aborted, err.Error(), err.Error())
	} else {
		//conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, sendError(codes.Aborted, err.Error(), err.Error())
	}

	var apptype irespo.IApplicationTypesRepository = &respo.ApplicationTypesRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var req irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	formN := int64(0)
	if in.Request.FORM_NO != nil {
		formN = *in.Request.FORM_NO
	}
	apptypeData, err := apptype.GetTypeByID(*in.Request.ApplicationType)
	if err != nil {
		return nil, err
	}
	if apptypeData == nil {
		return nil, errors.New("نوع الطلب غير معرف")
	}
	timeNow := time.Now()
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	defer dbr.Rollback()
	//defer dbr.Close()
	reqsave := &dbmodels.CANCELLED_REQUEST{}
	if formN > 0 {
		reqData, err := req.GetByFormNo(formN, *in.Request.ApplicationType)
		if err != nil {
			return nil, err
		}
		if len(reqData) > 0 {
			reqsave = reqData[0]
		} else {
			return nil, errors.New("لا يوجد طلب بهذا الرقم")
		}
		reqsave.STATE = in.Request.STATE
		reqsave.CLOSED = in.Request.CLOSED
		reqsave.STATUS = in.Request.STATUS
		reqsave.STAMP_DATE = create_time(in.Request.STAMP_DATE)
	} else {
		reqno, err := req.GetMax("FORM_NO", *in.Request.ApplicationType)
		if err != nil {
			return nil, err
		}
		if reqno == nil {
			return nil, errors.New("لم يتم تحديد رقم الطلب")
		}
		reqsave = &dbmodels.CANCELLED_REQUEST{
			FORM_NO:             *reqno,
			APPLICATION_TYPE_ID: *in.Request.ApplicationType,
			CANCELLED:           tools.ToBoolPointer(false),
			CUSTKEY:             "",
			STATION_NO:          "",
			DOCUMENT_NO:         "",
			REQUEST_DATE:        &timeNow,
			REQUEST_BY:          &username,
			//STATE:               in.Request.STATE,
			CLOSED:     tools.ToBoolPointer(false),
			STATUS:     in.Request.STATUS,
			COMMENT:    tools.ToStringPointer(""),
			COUNTER:    tools.Int32ToInt32Ptr(0),
			SURNAME:    tools.ToStringPointer(""),
			STAMP_DATE: &timeNow,
		}
	}

	var satatesr irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var intiActr irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var actionr irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	actionsData, err := actionr.GetByFormNo(reqsave.FORM_NO)
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	if len(actionsData) > 2 {
		return nil, sendError(codes.AlreadyExists, "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب", "لا يمكن حفظ الطلب لوجود اجراءات تمت على الطلب")
	}
	if reqsave.STATE == nil || *reqsave.STATE == 0 {
		stateId := int32(0)
		intiActData, err := intiActr.GetByStartUp(true, *in.Request.ApplicationType)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if len(actionsData) == 0 {
			if len(intiActData) != 0 {
				if intiActData[0].NEXT_STATE != nil {
					stateId = *intiActData[0].NEXT_STATE
				}
				cancelbillactionSave := &dbmodels.CANCELLED_BILLS_ACTION{
					FORM_NO:     reqsave.FORM_NO,
					CUSTKEY:     "",
					STAMP_DATE:  &timeNow,
					STAMP_USER:  user.USER_NAME,
					COMMENT:     tools.ToStringPointer("تم ايقاف الفاتورة"),
					ACTION_ID:   intiActData[0].ID,
					DOCUMENT_NO: "",
				}
				dbr.Add(cancelbillactionSave)
				if err != nil {
					return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
				}
			} else {
				return nil, sendError(codes.InvalidArgument, "برجاء تعريف الاجراء الابتدائي للعملية", "برجاء تعريف الاجراء الابتدائي للعملية")
			}
		} else {
			stateIdData, err := intiActr.GetByID(actionsData[0].ACTION_ID)
			if err != nil {
				return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
			}
			if len(stateIdData) > 0 {
				if stateIdData[0].CURRENT_STATE != nil {
					stateId = *stateIdData[0].CURRENT_STATE
				}
			}
		}
		stateData, err := satatesr.GetByID(stateId)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
		if len(stateData) == 0 {
			return nil, sendError(codes.InvalidArgument, "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId), "حالة غير معرفة للفاتورة "+*tools.Int32ToString(&stateId))
		}
		reqsave.STATUS = stateData[0].DESCRIPTION
		reqsave.STATE = &stateData[0].ID
	}
	if in.Request.FORM_NO == nil || (in.Request.FORM_NO != nil && *in.Request.FORM_NO == 0) {
		err = dbr.Add(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
	} else {
		err = dbr.Save(reqsave)
		if err != nil {
			return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
		}
	}

	// Add Entries
	// for k, v := range in.Request.Entries {

	// }
	err = dbr.Commit()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	DataJ := &pbMessages.SaveBillCancelRequestResponse{Message: tools.ToStringPointer("Done")}

	log.Println("End SaveBillCancelRequest..")
	return DataJ, nil
}
func cancelBillsReportP(ctx *context.Context, in *pbMessages.CancelBillsReportRequest) (rsp *pbMessages.CancelBillsReportResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	var cancelledBills irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var cancelledBillsAction irespo.ICancelledBillActionsRepository = &respo.CancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	cleanString(in.Custkey, nil, nil, nil)
	var stationNo *int32
	if station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0 {
		stationNo = &station.STATION_NO
	}
	if station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 1 {
		if in.StationNo != nil {
			stationNo = in.StationNo
		}
	}
	countReport, err := cancelledBills.GetCountRequestsBills(stationNo, in.Custkey, in.State, in.FormNo, create_time(in.RequestFrom), create_time(in.RequestTo), create_time(in.StampFrom), create_time(in.StampTo))
	if err != nil {
		return nil, err
	}
	if countReport > 10000 {
		return nil, errors.New("الحد الاقصى للبيانات هو 10000 برجاء تحديد عامل تصفية")
	}
	cancelledBillsData, err := cancelledBills.GetRequestsBills(stationNo, in.Custkey, in.State, in.FormNo, create_time(in.RequestFrom), create_time(in.RequestTo), create_time(in.StampFrom), create_time(in.StampTo))
	DataJ := &pbMessages.CancelBillsReportResponse{}
	for idx := range cancelledBillsData {
		usj := &serverhostmessages.CollectionDestributionItem{}
		cancelledBillsUse := cancelledBillsData[idx]
		if cancelledBillsUse.STATE == nil || cancelledBillsUse.FORM_NO == nil {
			continue
		}
		lastaction, err := cancelledBillsAction.GetByFormNoWithState(*cancelledBillsUse.FORM_NO, cancelledBillsUse.STATE)
		if err != nil {
			return nil, err
		}
		if cancelledBillsUse.CUSTKEY != nil {
			usj.CUSTKEY = cancelledBillsUse.CUSTKEY
		} else {
			usj.CUSTKEY = tools.ToStringPointer("")
		}
		if cancelledBillsUse.PAYMENT_NO != nil {
			usj.PAYMENT_NO = cancelledBillsUse.PAYMENT_NO
		} else {
			usj.PAYMENT_NO = tools.ToStringPointer("")
		}
		if cancelledBillsUse.SURNAME != nil {
			usj.SURNAME = cancelledBillsUse.SURNAME
		} else {
			usj.SURNAME = tools.ToStringPointer("")
		}
		usj.BILNG_DATE = create_timestamp(cancelledBillsUse.BILNG_DATE)
		usj.CL_BLNCE = cancelledBillsUse.CL_BLNCE
		usj.FORM_NO = cancelledBillsUse.FORM_NO
		usj.STAMP_DATE = create_timestamp(cancelledBillsUse.STAMP_DATE)
		usj.COMMENT = cancelledBillsUse.BILLCOMMENT
		usj.REQUEST_COMMENT = cancelledBillsUse.COMMENT
		if len(lastaction) > 0 {
			usj.ACTION_COMMENT = lastaction[0].COMMENT
		}
		usj.ACTIVITY = tools.ToStringPointer("")
		usj.CALC_TYPE = tools.ToStringPointer("")
		usj.ADDRESS = tools.ToStringPointer("")
		usj.TotalAmountCollected = tools.ToFloatPointer(0)
		usj.TotalCountCollected = tools.ToFloatPointer(0)
		usj.IS_COLLECTED_BY_OTHER = tools.ToBoolPointer(false)
		usj.IS_COLLECTED_BY_OWNER = tools.ToBoolPointer(false)
		usj.BILLGROUP = tools.ToStringPointer("")
		usj.BOOK_NO = tools.ToStringPointer("")
		usj.WALK_NO = tools.ToStringPointer("")
		usj.CTG = tools.ToStringPointer("")
		usj.OLD_KEY = tools.ToStringPointer("")
		usj.USER = tools.ToStringPointer("")
		DataJ.Items = append(DataJ.Items, usj)
	}
	log.Println("End cancelBillsReportP..")
	return DataJ, nil
}
func getStationsP(ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.GetStationsResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	if station == nil {
		return nil, errors.New("Station Not Found")
	}

	DataJ := &pbMessages.GetStationsResponse{}
	if station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 0 {
		usj := &pbMessages.Station{}
		usj.Description = station.DESCRIPTION
		usj.StationNo = &station.STATION_NO
		usj.IsHead = tools.ToBoolPointer(false)
		DataJ.Stations = append(DataJ.Stations, usj)
	} else {
		var stationr irespo.ICommonRepository = &respo.CommonRepository{Lama: conn}
		stations, err := stationr.GetStations()
		if err != nil {
			return nil, err
		}
		for idx := range stations {
			stationData := stations[idx]
			usj := &pbMessages.Station{}
			usj.Description = stationData.DESCRIPTION
			usj.StationNo = &stationData.STATION_NO
			if stationData.IS_HEADQUARTERS != nil && *stationData.IS_HEADQUARTERS == 1 {
				usj.IsHead = tools.ToBoolPointer(true)
			} else {
				usj.IsHead = tools.ToBoolPointer(false)
			}
			DataJ.Stations = append(DataJ.Stations, usj)
		}
	}
	log.Println("End GetStationsP..")
	return DataJ, nil
}
func getFormNoPaymentsP(ctx *context.Context, in *pbMessages.GetFormNoPaymentsRequest) (rsp *pbMessages.GetFormNoPaymentsResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	if in.FormNo == nil {
		return nil, errors.New("رقم الطلب غير صحيح")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var cancelbill irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	req, err := cancelbill.GetByFormNo(*in.FormNo, 1)
	if err != nil {
		return nil, err
	}
	if len(req) == 0 {
		return nil, errors.New("لا يوجد طلب بهذا الرقم")
	}
	if req[0].CLOSED != nil && *req[0].CLOSED {
		return nil, errors.New("الطلب تم اغلاقه")
	}
	cleanString(&req[0].CUSTKEY, tools.ToStringPointer(""), nil, nil)
	if strings.TrimSpace(req[0].CUSTKEY) == "" {
		return nil, errors.New("رقم حساب العميل غير صحيح")
	}
	var ctgConTypeGr irespo.ICtgConsumptionTypeGroupsRepository = &respo.CtgConsumptionTypeGroupsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	ctgData, err := ctgConTypeGr.GetAll()
	if err != nil {
		return nil, err
	}
	var hand irespo.IHandMhStRepository = &respo.HandMhStRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	user, err := getUser(&username, conn)
	if err != nil {
		return nil, err
	}
	station, err := getStation(user.STATION_NO, conn)
	if err != nil {
		return nil, err
	}
	bills, err := cancelbill.GetByBillsFormNo(*in.FormNo)
	if err != nil {
		return nil, err
	}

	DataJ := &pbMessages.GetFormNoPaymentsResponse{}
	//var recep irespo.IReciptsRepository = &respo.ReciptsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	for idx := range bills {
		billUse := bills[idx]
		/*countReceipts, err := recep.GetCountByPaymentNoCancelled(billUse.PAYMENT_NO, false)
		if err != nil {
			return nil, err
		}
		if countReceipts > 0 {
			continue
		}*/
		usjOld, err := getPayment(&billUse.PAYMENT_NO, &billUse.CUSTKEY, tools.ToBoolPointer(true), nil, nil, &hand, user, ctgData, conn, station, in.FormNo, tools.ToBoolPointer(true))
		if err != nil {
			return nil, sendError(codes.Internal, err.Error(), err.Error())
		}
		usjNew, err := getPayment(&billUse.PAYMENT_NO, &billUse.CUSTKEY, tools.ToBoolPointer(true), nil, nil, &hand, user, ctgData, conn, station, in.FormNo, tools.ToBoolPointer(false))
		if err != nil {
			return nil, sendError(codes.Internal, err.Error(), err.Error())
		}
		usj := &serverhostmessages.OldNewItem{OldItem: usjOld, NewItem: usjNew}
		DataJ.Items = append(DataJ.Items, usj)
	}
	log.Println("end ..")
	return DataJ, nil
}
func getApplicationTypesP(ctx *context.Context, in *pbMessages.Empty) (rsp *pbMessages.ApplicationTypesRs, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	var apptype irespo.IApplicationTypesRepository = &respo.ApplicationTypesRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	apptypesData, err := apptype.GetAllTypes()
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.ApplicationTypesRs{}
	var appStates irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var appActions irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	for idx := range apptypesData {
		applicationtypeUse := apptypesData[idx]
		Usj := &dbmessages.ApplicationType{}
		Usj.Id = &applicationtypeUse.ID
		Usj.Description = applicationtypeUse.DESCRIPTION
		Usj.Seq = applicationtypeUse.SEQ
		appStatesData, err := appStates.GetByApplicationTypeID(applicationtypeUse.ID)
		if err != nil {
			return nil, err
		}
		for idxState := range appStatesData {
			appStateUse := appStatesData[idxState]
			Usj.States = append(Usj.States, &pbdbMessages.LU_CANCELLED_BILL_STATE{
				ID:              &appStateUse.ID,
				DESCRIPTION:     appStateUse.DESCRIPTION,
				RECAL_READY:     appStateUse.RECAL_READY,
				CANCELLED:       appStateUse.CANCELLED,
				EDITED:          appStateUse.EDITED,
				ApplicationType: appStateUse.APPLICATION_TYPE_ID,
			})
		}
		appActionsData, err := appActions.GetByApplicationTypeID(applicationtypeUse.ID)
		if err != nil {
			return nil, err
		}
		for idxAction := range appActionsData {
			appActionUse := appActionsData[idxAction]
			UsjAction := &pbdbMessages.LU_CANCELLED_BILL_ACTION{
				ID:              &appActionUse.ID,
				DESCRIPTION:     appActionUse.DESCRIPTION,
				CURRENT_STATE:   appActionUse.CURRENT_STATE,
				NEXT_STATE:      appActionUse.NEXT_STATE,
				CLOSED:          appActionUse.CLOSED,
				START_UP:        appActionUse.START_UP,
				ApplicationType: appActionUse.APPLICATION_TYPE_ID,
			}
			// Action Fields
			fieldGroupActionData, err := apptype.GetAllGroupFieldsByActionID(appActionUse.ID)
			if err != nil {
				return nil, err
			}
			for idxFieldGroup := range fieldGroupActionData {
				fieldGroupUse := fieldGroupActionData[idxFieldGroup]
				UsjFieldGroup := &pbdbMessages.FieldGroup{
					Title: fieldGroupUse.TITLE,
					Seq:   fieldGroupUse.SEQ,
				}
				fieldActionData, err := apptype.GetAllFieldsByGroupID(fieldGroupUse.ID)
				if err != nil {
					return nil, err
				}
				for idxField := range fieldActionData {
					fieldUse := fieldActionData[idxField]
					var fieldType pbdbMessages.DataType
					fieldType = pbdbMessages.DataType(fieldUse.DATA_TYPE)
					var fieldKind pbdbMessages.FieldKind
					fieldKind = pbdbMessages.FieldKind(*fieldUse.KIND)
					UsjField := &pbdbMessages.Field{
						Title:      fieldUse.TITLE,
						Seq:        fieldUse.SEQ,
						Name:       &fieldUse.NAME,
						IsRequired: fieldUse.IS_REQUIRED,
						Format:     fieldUse.FORMAT,
						DataType:   &fieldType,
						Kind:       &fieldKind,
					}
					fieldValueData, err := apptype.GetAllListValuesByFieldID(fieldUse.ID)
					if err != nil {
						return nil, err
					}
					if len(fieldValueData) > 0 {
						listVal := &pbdbMessages.ListValues{}
						for idxVal := range fieldValueData {
							valUse := fieldValueData[idxVal]
							listVal.ListValues = append(listVal.ListValues, &pbdbMessages.ListValue{
								Key:   &valUse.KEY_SOURCE,
								Value: &valUse.VALUE_SOURCE,
							})
						}
						UsjField.ListValues = listVal
					}
					UsjFieldGroup.Fields = append(UsjFieldGroup.Fields, UsjField)
				}
				UsjAction.Fields = append(UsjAction.Fields, UsjFieldGroup)
			}
			Usj.Actions = append(Usj.Actions, UsjAction)
		}
		// ApplicationType Fields
		fieldGroupAppTypeData, err := apptype.GetAllGroupFieldsByApplicationTypeID(applicationtypeUse.ID)
		if err != nil {
			return nil, err
		}
		for idxFieldGroup := range fieldGroupAppTypeData {
			fieldGroupUse := fieldGroupAppTypeData[idxFieldGroup]
			if fieldGroupUse.LU_ACTIONS_ID != nil {
				continue
			}
			UsjFieldGroup := &pbdbMessages.FieldGroup{
				Title: fieldGroupUse.TITLE,
				Seq:   fieldGroupUse.SEQ,
			}
			fieldActionData, err := apptype.GetAllFieldsByGroupID(fieldGroupUse.ID)
			if err != nil {
				return nil, err
			}
			for idxField := range fieldActionData {
				fieldUse := fieldActionData[idxField]
				var fieldType pbdbMessages.DataType
				fieldType = pbdbMessages.DataType(fieldUse.DATA_TYPE)
				var fieldKind pbdbMessages.FieldKind
				fieldKind = pbdbMessages.FieldKind(*fieldUse.KIND)
				UsjField := &pbdbMessages.Field{
					Title:      fieldUse.TITLE,
					Seq:        fieldUse.SEQ,
					Name:       &fieldUse.NAME,
					IsRequired: fieldUse.IS_REQUIRED,
					Format:     fieldUse.FORMAT,
					DataType:   &fieldType,
					Kind:       &fieldKind,
				}
				fieldValueData, err := apptype.GetAllListValuesByFieldID(fieldUse.ID)
				if err != nil {
					return nil, err
				}
				if len(fieldValueData) > 0 {
					listVal := &pbdbMessages.ListValues{}
					for idxVal := range fieldValueData {
						valUse := fieldValueData[idxVal]
						listVal.ListValues = append(listVal.ListValues, &pbdbMessages.ListValue{
							Key:   &valUse.KEY_SOURCE,
							Value: &valUse.VALUE_SOURCE,
						})
					}
					UsjField.ListValues = listVal
				}
				UsjFieldGroup.Fields = append(UsjFieldGroup.Fields, UsjField)
			}
			Usj.Fields = append(Usj.Fields, UsjFieldGroup)
		}
		DataJ.ApplicationTypes = append(DataJ.ApplicationTypes, Usj)
	}
	log.Println("end ..")
	return DataJ, nil
}
func saveApplicationTypeP(ctx *context.Context, in *pbMessages.SaveApplicationTypeRequest) (rsp *pbMessages.SaveBillCancelRequestResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
	username, ok := (*ctx).Value("username").(string)
	if !ok {
		return nil, errors.New("can not parse username")
	}
	if username == "" {
		return nil, errors.New("missing username")
	}
	if in.ApplicationTypes == nil || in.ApplicationTypes.Id == nil {
		return nil, errors.New("لابد من ارسال بيانات لنوع الطلب")
	}
	if in.ApplicationTypes.Id == nil {
		return nil, errors.New("لابد من ارسال بيانات لنوع الطلب")
	}
	if len(in.ApplicationTypes.Actions) <= 0 {
		return nil, errors.New("لابد من وجود اجراء واحد علي الاقل")
	}
	if len(in.ApplicationTypes.States) <= 0 {
		return nil, errors.New("لابد من وجود مرحله واحده علي الاقل")
	}

	conn, err := dbpool.GetConnection()
	if err != nil {
		log.Println(err)
		return nil, sendError(codes.Internal, err.Error(), err.Error())
	} else {
		conn.Debug = true
		log.Println("connected")
	}
	dbr, err := conn.Begin()
	if err != nil {
		return nil, sendError(codes.InvalidArgument, err.Error(), err.Error())
	}
	defer dbr.Rollback()
	var appreq irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var appActions irespo.ILuCancelledBillActionsRepository = &respo.LuCancelledBillActionsRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var appStates irespo.ILuCancelledBillStatessRepository = &respo.LuCancelledBillStatessRepository{CommonRepository: respo.CommonRepository{Lama: conn}}
	var apptype irespo.IApplicationTypesRepository = &respo.ApplicationTypesRepository{CommonRepository: respo.CommonRepository{Lama: conn}}

	// Add Or Save Application Type
	apptypeData, err := apptype.GetTypeByID(*in.ApplicationTypes.Id)
	if err != nil {
		return nil, err
	}
	adding := true
	if apptypeData == nil {
		apptypeData = &dbmodels.APPLICATION_TYPES{}
	} else {
		adding = false
	}
	apptypeData.DESCRIPTION = in.ApplicationTypes.Description
	apptypeData.SEQ = in.ApplicationTypes.Seq
	if in.ApplicationTypes.Icon != nil {
		imag := []byte(*in.ApplicationTypes.Icon)
		apptypeData.ICON = &imag
	} else {
		apptypeData.ICON = &[]byte{}
	}
	if !adding {
		// Check old Req
		oldReq, err := appreq.GetCountByApplicationTypeID(apptypeData.ID)
		if err != nil {
			return nil, err
		}
		if oldReq > 0 {
			return nil, errors.New("يوجد طلبات لهذا النوع لا يمكن التعديل")
		}
		err = dbr.Save(apptypeData)
		if err != nil {
			return nil, err
		}
		// delete All Group Fields & Fields & States & Actoins
		// Delete All Old States
		oldAppState, err := appStates.GetByApplicationTypeID(apptypeData.ID)
		if err != nil {
			return nil, err
		}
		for idx := range oldAppState {
			oldappState := oldAppState[idx]
			err = dbr.Delete(oldappState)
			if err != nil {
				return nil, err
			}
		}
		// Delete All Old Actions
		oldAppAction, err := appActions.GetByApplicationTypeID(apptypeData.ID)
		if err != nil {
			return nil, err
		}
		for idx := range oldAppAction {
			oldappAction := oldAppAction[idx]
			// Delete All Old fields Group
			oldAppGroups, err := apptype.GetAllGroupFieldsByActionID(oldappAction.ID)
			if err != nil {
				return nil, err
			}
			for idxv := range oldAppGroups {
				oldappGroup := oldAppGroups[idxv]
				// Delete All Old fields
				oldAppFields, err := apptype.GetAllFieldsByGroupID(oldappGroup.ID)
				if err != nil {
					return nil, err
				}
				for idxf := range oldAppFields {
					oldappField := oldAppFields[idxf]
					// Delete All Old list values
					oldAppListVal, err := apptype.GetAllListValuesByFieldID(oldappField.ID)
					if err != nil {
						return nil, err
					}
					for idxval := range oldAppListVal {
						oldappList := oldAppListVal[idxval]
						err = dbr.Delete(oldappList)
						if err != nil {
							return nil, err
						}
					}
					err = dbr.Delete(oldappField)
					if err != nil {
						return nil, err
					}
				}
				err = dbr.Delete(oldappGroup)
				if err != nil {
					return nil, err
				}
			}
			err = dbr.Delete(oldappAction)
			if err != nil {
				return nil, err
			}
		}
		// Fields Application Type
		// Delete All Old fields Group
		oldTypeGroups, err := apptype.GetAllGroupFieldsByApplicationTypeID(apptypeData.ID)
		if err != nil {
			return nil, err
		}
		for idxv := range oldTypeGroups {
			oldTypeGroup := oldTypeGroups[idxv]
			if oldTypeGroup.LU_ACTIONS_ID != nil {
				continue
			}
			// Delete All Old fields
			oldTypeFields, err := apptype.GetAllFieldsByGroupID(oldTypeGroup.ID)
			if err != nil {
				return nil, err
			}
			for idxf := range oldTypeFields {
				oldTypeField := oldTypeFields[idxf]
				// Delete All Old list values
				oldTypeListVal, err := apptype.GetAllListValuesByFieldID(oldTypeField.ID)
				if err != nil {
					return nil, err
				}
				for idxval := range oldTypeListVal {
					oldTypeList := oldTypeListVal[idxval]
					err = dbr.Delete(oldTypeList)
					if err != nil {
						return nil, err
					}
				}
				err = dbr.Delete(oldTypeField)
				if err != nil {
					return nil, err
				}
			}
			err = dbr.Delete(oldTypeGroup)
			if err != nil {
				return nil, err
			}
		}
	} else {
		typeID, err := apptype.GetMaxTypeID()
		if err != nil {
			return nil, err
		}
		apptypeData.ID = typeID + 1
		err = dbr.Add(apptypeData)
		if err != nil {
			return nil, err
		}
	}
	// Check Validation
	mapStates := make(map[int32]int32)
	closedOne := false
	startOne := false
	firstState := int32(0)
	lastState := int32(0)
	allActionState := []*ActStates{}
	mapRepeatedState := make(map[int32]int32)
	for idx := range in.ApplicationTypes.States {
		sta := in.ApplicationTypes.States[idx]
		if sta.ID == nil || *sta.ID <= 0 {
			return nil, errors.New("لابد من تحديد رقم المرحلة")
		}
		stateID := 0
		if *sta.ID >= 1 && *sta.ID <= 9 {
			stateID, err = strconv.Atoi(*tools.Int32ToString(&apptypeData.ID) + "00" + *tools.Int32ToString(sta.ID))
			if err != nil {
				return nil, err
			}
		} else if *sta.ID >= 10 && *sta.ID <= 99 {
			stateID, err = strconv.Atoi(*tools.Int32ToString(&apptypeData.ID) + "0" + *tools.Int32ToString(sta.ID))
			if err != nil {
				return nil, err
			}
		} else if *sta.ID >= 100 && *sta.ID <= 999 {
			stateID, err = strconv.Atoi(*tools.Int32ToString(&apptypeData.ID) + *tools.Int32ToString(sta.ID))
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("لابد من تحديد رقم المرحلة اقل  من 1000")
		}

		_, ok := mapRepeatedState[*sta.ID]
		if ok {
			return nil, errors.New("رقم المرحله مكرر")
		} else {
			mapRepeatedState[*sta.ID] = int32(stateID)
		}
		sta.ID = tools.Int32ToInt32Ptr(int32(stateID))
	}
	for idx := range in.ApplicationTypes.Actions {
		act := in.ApplicationTypes.Actions[idx]
		if act.CURRENT_STATE == nil {
			return nil, errors.New("لابد من تحديد المرحله الحاليه للاجراء")
		}
		if act.NEXT_STATE == nil {
			return nil, errors.New("لابد من تحديد المرحله التاليه للاجراء")
		}
		if act.START_UP == nil {
			return nil, errors.New("لابد من تحديد هل الاجراء بدايه ام لا")
		}
		if *act.CURRENT_STATE == *act.NEXT_STATE {
			return nil, errors.New("لا يمكن ان يكون المرحله الحاليه والتاليه متساوي للاجراء")
		}
		currentState, ok := mapRepeatedState[*act.CURRENT_STATE]
		if !ok {
			return nil, errors.New("رقم المرحله غير معرف")
		}
		nextState, ok := mapRepeatedState[*act.NEXT_STATE]
		if !ok {
			return nil, errors.New("رقم المرحله غير معرف")
		}
		mapStates[*act.CURRENT_STATE] = currentState
		mapStates[*act.NEXT_STATE] = nextState
		act.CURRENT_STATE = &currentState
		act.NEXT_STATE = &nextState
		if act.CLOSED != nil && *act.CLOSED {
			if closedOne {
				return nil, errors.New("لابد من وجود اجراء واحد لاغلاق الطلب")
			}
			lastState = *act.NEXT_STATE
			closedOne = true
		}
		if act.START_UP != nil && *act.START_UP {
			if startOne {
				return nil, errors.New("لابد من وجود اجراء واحد لبداية الطلب")
			}
			firstState = *act.CURRENT_STATE
			startOne = true
		}
		actState := &ActStates{From_State: *act.CURRENT_STATE, To_State: *act.NEXT_STATE}
		extact := Exists(allActionState, func(val interface{}) bool {
			return (val.(*ActStates)).From_State == actState.From_State && (val.(*ActStates)).To_State == actState.To_State
		})
		if extact {
			return nil, errors.New("يوجد تكرار في الاجراءات")
		}
		allActionState = append(allActionState, actState)
	}
	if len(mapStates) != len(in.ApplicationTypes.States) {
		return nil, errors.New("يوجد مراحل غير مستخدمه في الاجراءات")
	}
	if !closedOne {
		return nil, errors.New("لابد من وجود اجراء لاغلاق الطلب")
	}
	if !startOne {
		return nil, errors.New("لابد من وجود اجراء لبداية الطلب")
	}
	firstnodes := []*ActStates{} // will start from finshed
	fnode := ActStates{From_State: lastState}
	firstnodes = append(firstnodes, &fnode)
	outf := []*int32{}
	recursiveFlow(&outf, allActionState, firstnodes)
	allfrom := []int32{}
	allfromdone := []int32{firstState, lastState}
	alltodone := []int32{firstState, lastState}
	for _, vbc := range allActionState {
		extactvbc := Exists(allfrom, func(val interface{}) bool {
			return (val.(int32)) == vbc.From_State
		})
		if !extactvbc {
			allfrom = append(allfrom, vbc.From_State)
		}

		extactvbc = Exists(allfromdone, func(val interface{}) bool {
			return (val.(int32)) == vbc.From_State
		})
		if !extactvbc {
			allfromdone = append(allfromdone, vbc.From_State)
		}

		extactvbc = Exists(allfrom, func(val interface{}) bool {
			return (val.(int32)) == vbc.To_State
		})
		if !extactvbc {
			allfrom = append(allfrom, vbc.To_State)
		}

		extactvbc = Exists(alltodone, func(val interface{}) bool {
			return (val.(int32)) == vbc.To_State
		})
		if !extactvbc {
			alltodone = append(alltodone, vbc.To_State)
		}
	}
	if len(allfrom) <= 2 {
		return nil, errors.New("اجراءات العمل غير صحيحه")
	}
	if len(outf) < len(allfrom) {
		return nil, errors.New("اجراءات العمل غير صحيحه")
	}
	for _, cv := range outf {
		extactvbc := Exists(alltodone, func(val interface{}) bool {
			return (val.(int32)) == *cv
		})
		if !extactvbc {
			return nil, errors.New("اجراءات العمل غير صحيحه")
		}
		extactvbc = Exists(allfromdone, func(val interface{}) bool {
			return (val.(int32)) == *cv
		})
		if !extactvbc {
			return nil, errors.New("اجراءات العمل غير صحيحه")
		}
	}
	i, err := strconv.Atoi(*tools.Int32ToString(&apptypeData.ID) + "0000")
	if err != nil {
		return nil, err
	}
	actionID := int32(i)
	//stateID := int32(i)
	groupID := int32(i)
	fieldID := int32(i)
	valID := int32(i)
	// states
	for idx := range in.ApplicationTypes.States {
		appState := in.ApplicationTypes.States[idx]
		appstateData := &dbmodels.LU_CANCELLED_BILL_STATE{}
		appstateData.APPLICATION_TYPE_ID = &apptypeData.ID
		appstateData.CANCELLED = appState.CANCELLED
		appstateData.DESCRIPTION = appState.DESCRIPTION
		appstateData.EDITED = appState.EDITED
		appstateData.RECAL_READY = appState.RECAL_READY
		appstateData.ID = *appState.ID
		err = dbr.Add(appstateData)
		if err != nil {
			return nil, err
		}
	}
	// mapping For Fileds name
	chechFields := make(map[string]*string)
	// actions
	for idx := range in.ApplicationTypes.Actions {
		actionID = actionID + 1
		appAction := in.ApplicationTypes.Actions[idx]
		appactionData := &dbmodels.LU_CANCELLED_BILLS_ACTION{}
		appactionData.APPLICATION_TYPE_ID = &apptypeData.ID
		appactionData.DESCRIPTION = appAction.DESCRIPTION
		appactionData.CLOSED = appAction.CLOSED
		appactionData.CURRENT_STATE = appAction.CURRENT_STATE
		appactionData.NEXT_STATE = appAction.NEXT_STATE
		appactionData.DEPARTMENT = appAction.DEPARTMENT
		appactionData.START_UP = appAction.START_UP
		appactionData.ID = actionID
		err = dbr.Add(appactionData)
		if err != nil {
			return nil, err
		}
		// fields Group
		for idxgroup := range appAction.Fields {
			groupID = groupID + 1
			appActionGroup := appAction.Fields[idxgroup]
			appactionGroupData := &dbmodels.FIELD_GROUPS{}
			appactionGroupData.APPLICATION_TYPE_ID = &apptypeData.ID
			appactionGroupData.LU_ACTIONS_ID = &appactionData.ID
			appactionGroupData.SEQ = appActionGroup.Seq
			appactionGroupData.TITLE = appActionGroup.Title
			appactionGroupData.ID = groupID
			err = dbr.Add(appactionGroupData)
			if err != nil {
				return nil, err
			}
			// fields
			for idxfield := range appActionGroup.Fields {
				fieldID = fieldID + 1
				appActionField := appActionGroup.Fields[idxfield]
				appactionFieldData := &dbmodels.FIELDS{}
				appactionFieldData.FIELD_GROUP_ID = appactionGroupData.ID
				appactionFieldData.FORMAT = appActionField.Format
				appactionFieldData.IS_REQUIRED = appActionField.IsRequired
				if appActionField.Name == nil {
					return nil, errors.New("لابد من ادخال الاسم")
				}
				enn, namef := IsEnglish(*appActionField.Name)
				if !enn {
					return nil, errors.New("لابد من ان يكون الاسم بالانجليزيه")
				}
				_, ok := chechFields[namef]
				if ok {
					return nil, errors.New("لابد من ان يكون اسم الحقل غير مكرر")
				}
				chechFields[namef] = &namef
				appactionFieldData.NAME = namef
				appactionFieldData.SEQ = appActionField.Seq
				appactionFieldData.TITLE = appActionField.Title
				appactionFieldData.KIND = (*int32)(appActionField.Kind)
				if appActionField.DataType == nil {
					return nil, errors.New("لابد من ادخال النوع")
				}
				appactionFieldData.DATA_TYPE = int32(*appActionField.DataType)
				appactionFieldData.ID = fieldID
				err = dbr.Add(appactionFieldData)
				if err != nil {
					return nil, err
				}
				// List Values
				if appActionField.ListValues != nil {
					for idxval := range appActionField.ListValues.ListValues {
						valID = valID + 1
						listval := appActionField.ListValues.ListValues[idxval]
						appVallistData := &dbmodels.FIELD_VALUES{}
						appVallistData.FIELD_ID = appactionFieldData.ID
						if listval.Key == nil {
							return nil, errors.New("لابد من ادخال مميز القائمه")
						}
						appVallistData.KEY_SOURCE = *listval.Key
						if listval.Value == nil {
							return nil, errors.New("لابد من ادخال قيمه القائمه")
						}
						appVallistData.VALUE_SOURCE = *listval.Value
						appVallistData.ID = valID
						err = dbr.Add(appVallistData)
						if err != nil {
							return nil, err
						}
					}
				}
			}
		}
	}
	// Fields
	for idxgroup := range in.ApplicationTypes.Fields {
		groupID = groupID + 1
		appActionGroup := in.ApplicationTypes.Fields[idxgroup]
		appactionGroupData := &dbmodels.FIELD_GROUPS{}
		appactionGroupData.APPLICATION_TYPE_ID = &apptypeData.ID
		appactionGroupData.SEQ = appActionGroup.Seq
		appactionGroupData.TITLE = appActionGroup.Title
		appactionGroupData.ID = groupID
		err = dbr.Add(appactionGroupData)
		if err != nil {
			return nil, err
		}
		// fields
		for idxfield := range appActionGroup.Fields {
			fieldID = fieldID + 1
			appActionField := appActionGroup.Fields[idxfield]
			appactionFieldData := &dbmodels.FIELDS{}
			appactionFieldData.FIELD_GROUP_ID = appactionGroupData.ID
			appactionFieldData.FORMAT = appActionField.Format
			appactionFieldData.IS_REQUIRED = appActionField.IsRequired
			if appActionField.Name == nil {
				return nil, errors.New("لابد من ادخال الاسم")
			}
			enn, namef := IsEnglish(*appActionField.Name)
			if !enn {
				return nil, errors.New("لابد من ان يكون الاسم بالانجليزيه")
			}
			_, ok := chechFields[namef]
			if ok {
				return nil, errors.New("لابد من ان يكون اسم الحقل غير مكرر")
			}
			chechFields[namef] = &namef
			appactionFieldData.NAME = *appActionField.Name
			appactionFieldData.SEQ = appActionField.Seq
			appactionFieldData.TITLE = appActionField.Title
			appactionFieldData.KIND = (*int32)(appActionField.Kind)
			if appActionField.DataType == nil {
				return nil, errors.New("لابد من ادخال النوع")
			}
			appactionFieldData.DATA_TYPE = int32(*appActionField.DataType)
			appactionFieldData.ID = fieldID
			err = dbr.Add(appactionFieldData)
			if err != nil {
				return nil, err
			}
			// List Values
			if appActionField.ListValues != nil {
				for idxval := range appActionField.ListValues.ListValues {
					valID = valID + 1
					listval := appActionField.ListValues.ListValues[idxval]
					appVallistData := &dbmodels.FIELD_VALUES{}
					appVallistData.FIELD_ID = appactionFieldData.ID
					if listval.Key == nil {
						return nil, errors.New("لابد من ادخال مميز القائمه")
					}
					appVallistData.KEY_SOURCE = *listval.Key
					if listval.Value == nil {
						return nil, errors.New("لابد من ادخال قيمه القائمه")
					}
					appVallistData.VALUE_SOURCE = *listval.Value
					appVallistData.ID = valID
					err = dbr.Add(appVallistData)
					if err != nil {
						return nil, err
					}
				}
			}
		}
	}
	err = dbr.Commit()
	if err != nil {
		return nil, err
	}
	DataJ := &pbMessages.SaveBillCancelRequestResponse{}
	DataJ.Message = tools.ToStringPointer("تم حفظ نوع الطلب")
	DataJ.SaveId = tools.Int32PtrToInt64Ptr(&apptypeData.ID)
	log.Println("end ..")
	return DataJ, nil
}
