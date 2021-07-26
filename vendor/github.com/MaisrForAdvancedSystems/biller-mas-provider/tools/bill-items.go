package tools

import (
	"errors"
	"fmt"
	structs "github.com/HassanAbdelzaher/lama/structs"
	billing "github.com/MaisrForAdvancedSystems/go-biller-proto/go"
	dbmodels "github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"
	"log"
	"math"
	"reflect"
)

func GetBillItemValues(bi *dbmodels.BILL_ITEMS) (valMap map[string]float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recover")
			err = errors.New(fmt.Sprintf("panic %v", r))
		}
	}()
	resp:=make(map[string]float64)
	valus:=structs.New(bi,structs.MapOptions{
		SkipZeroValue:      true,
		UseFieldName:       false,
		SkipUnTaged:        false,
		SkipComputed:       false,
		Flatten:            true,
		SelectedZeroValues: nil,
	}).Map()
	for k:=range valus{
		f,ok:=valus[k].(*float64)
		if ok{
			if f!=nil && *f!=0{
				resp[k]=*f
			}
		}else {
			f,ok:=valus[k].(float64)
			if ok{
				if f!=0{
					resp[k]=f
				}
			}
		}
	}
	return resp,nil

}

func SetBillItemValues(bi *dbmodels.BILL_ITEMS, values map[string]float64) (err error) {
	log.Println("In Reflect")
	defer func() {
		if r := recover(); r != nil {
			log.Println("recover")
			err = errors.New(fmt.Sprintf("panic %v", r))
		}
	}()
	//typ := reflect.ValueOf(bi).Elem()
	vOf := reflect.ValueOf(bi)
	typ := reflect.Indirect(vOf)
	for _k, _v := range values {
		k := _k
		v := _v //copy
		/*if v==0{
			continue
		}*/
		fv := typ.FieldByName(k)
		if !fv.IsValid() {
			log.Println("k", k, "v", v)
			return errors.New("invalied financial item:" + k)
		}
		if !fv.CanSet() {
			log.Println("k", k, "v", v)
			return errors.New("can not st financial item:" + k)
		}
		val := reflect.ValueOf(&v)
		fv.Set(val)
	}
	return nil
}

// CreateBillItems Imps
func CreateBillItems(custkey string, cycle_id int64, stationNo *int32, trans []*billing.FinantialTransaction) (*dbmodels.BILL_ITEM, *float64, error) {
	bi := dbmodels.BILL_ITEM{CUSTKEY: custkey, CYCLE_ID: ToInt32Pointer(int32(cycle_id)), STATION_NO: stationNo}
	values := make(map[string]float64)
	log.Println("Trans Count, ", len(trans))
	sumAmount := float64(0)
	for _, t := range trans {
		if t == nil {
			return nil, &sumAmount, errors.New("invalied transaction")
		}
		if t.Code == nil {
			return nil, &sumAmount, errors.New("invalied transaction code")
		}
		if t.Amount == nil {
			continue
		}
		amt:=math.Round(1000*(*t.Amount))/1000
		values[*t.Code] = values[*t.Code]+amt
		sumAmount = sumAmount + *t.Amount
	}
	sumAmount=math.Round(1000*sumAmount)/1000
	log.Println("Values, ", values)
	err := SetBillItemValues(&bi.BILL_ITEMS, values)
	return &bi, &sumAmount, err
}

func UpdateBillItems(bi *dbmodels.BILL_ITEM, trans []*billing.FinantialTransaction) (*dbmodels.BILL_ITEM, error) {
	//bi := dbmodels.BILL_ITEM{CUSTKEY: custkey, CYCLE_ID: ToInt32Pointer(int32(cycle_id))}
	values := make(map[string]float64)
	log.Println("Trans Count, ", len(trans))
	for _, t := range trans {
		if t == nil {
			return nil, errors.New("invalied transaction")
		}
		if t.Code == nil {
			return nil, errors.New("invalied transaction code")
		}
		if t.Amount == nil {
			continue
		}
		values[*t.Code] = *t.Amount
	}
	log.Println("Values, ", values)
	err := SetBillItemValues(&bi.BILL_ITEMS, values)
	return bi, err
}
