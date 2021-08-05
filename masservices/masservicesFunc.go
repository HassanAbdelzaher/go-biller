package masservices

import (
	"MaisrForAdvancedSystems/go-biller/tools"
	"errors"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/HassanAbdelzaher/lama"
	"github.com/MaisrForAdvancedSystems/go-biller-proto/go/serverhostmessages"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	irespo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/interfaces"
	respo "github.com/MaisrForAdvancedSystems/mas-db-models/repositories/repositories"
)

func getPayment(paymentNo *string, custKey *string, skipBracodTrim *bool, forQuery *bool, cycle_id *int32, stationNo *int32, hand *irespo.IHandMhStRepository, user *dbmodels.USERS, ctg []*dbmodels.CTG_CONSUMPTIONTYPEGRPS, conn *lama.Lama, station *dbmodels.STATIONS, formNo *int64) (rsp *serverhostmessages.CollectionDestributionItem, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("recover error:%v", r))
		}
	}()
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
	pay, err := (*hand).GetPayment(paymentNo, nil, custKey, cycle_id)
	if err != nil {
		return nil, err
	}
	if len(pay) == 0 {
		return nil, errors.New("لا يوجد فاتورة بهذا الرقم")
	}
	if custKey == nil {
		custKey = pay[0].CUSTKEY
	}
	if paymentNo == nil {
		paymentNo = pay[0].Payment_no
	}
	if len(pay) > 1 {
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
	if !(station.IS_HEADQUARTERS != nil && *station.IS_HEADQUARTERS == 1) {
		if pay[0].STATION_NO != nil && user.STATION_NO != nil && int64(*pay[0].STATION_NO) != *user.STATION_NO {
			return nil, errors.New("الفاتورة تخص فرع اخر")
		}
	}

	var cancelreq irespo.ICancelledBillsRepository = &respo.CancelledBillsRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	cancelData, err := cancelreq.GetByCustKeyClosed(*custKey, false)
	if err != nil {
		return nil, err
	}
	commentbill := ""
	for idx := range cancelData {
		cancelDataUse := cancelData[idx]
		cancelBillData, err := cancelreq.GetByFormNoPaymentNo(cancelDataUse.FORM_NO, *pay[0].Payment_no)
		if err != nil {
			return nil, err
		}
		if formNo != nil && cancelDataUse.FORM_NO == *formNo {
			if len(cancelBillData) > 0 {
				if cancelBillData[0].COMMENT != nil {
					commentbill = *cancelBillData[0].COMMENT
				}
			}
			continue
		}
		if len(cancelBillData) > 0 {
			return nil, errors.New("يوجد طلب مفتوح على الفاتورة قيد المراجعة")
		}
	}
	var customerbook irespo.ICustomerBooksRepository = &respo.CustomerBooksRepository{CommonRepository: respo.CommonRepository{Lama: (*hand).GetUnderLineConnection()}}
	if pay[0].BILLGROUP == nil || pay[0].BOOK_NO_C == nil {
		return nil, errors.New("لم يتم تحديد BILLGROUP او BOOK_NO_C")
	}
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
	if user.USER_NAME != nil {
		usernamee = *user.USER_NAME
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
	clbance := float64(0)
	if pay[0].Cl_blnce != nil {
		clbance = *pay[0].Cl_blnce
	}
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
	reminderValue = tools.RoundTo(clbance-collectedB4, 3)
	var deliverst *int32 = pay[0].Delivery_st
	if reminderValue <= 0 {
		reminderValue = 0
		deliverst = tools.Int32ToInt32Ptr(1)
	} else {
		isBillColl, err := recep.GetByCustKeyCycleIDCancelledCollectionType(*pay[0].CUSTKEY, pay[0].CYCLE_ID, false)
		if err != nil {
			return nil, err
		}
		if len(isBillColl) > 0 {
			deliverst = tools.Int32ToInt32Ptr(1)
		}
	}
	ctypgrid := ""
	if pay[0].Ctypegrp_id != nil {
		ctypgrid = *pay[0].Ctypegrp_id
	}
	DataJ := &serverhostmessages.CollectionDestributionItem{
		PAYMENT_NO:            pay[0].Payment_no,
		CUSTKEY:               pay[0].CUSTKEY,
		DELIVERY_ST:           deliverst,
		BILLGROUP:             pay[0].BILLGROUP,
		BILNG_DATE:            create_timestamp(pay[0].BILNG_DATE),
		CL_BLNCE:              pay[0].Cl_blnce,
		ISSUED_AMOUNT:         pay[0].Cl_blnce,
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
		COMMENT:               &commentbill,
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
			otherAmount := float64(0)
			waterAmount := float64(0)
			sewerAmount := float64(0)
			if ctydetailDataUse.OTHER_AMT != nil {
				otherAmount = *ctydetailDataUse.OTHER_AMT
			}
			if ctydetailDataUse.WATER_AMT != nil {
				waterAmount = *ctydetailDataUse.WATER_AMT
			}
			if ctydetailDataUse.SEWER_AMT != nil {
				sewerAmount = *ctydetailDataUse.SEWER_AMT
			}
			totalV := tools.RoundTo(otherAmount+waterAmount+sewerAmount, 2)
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
