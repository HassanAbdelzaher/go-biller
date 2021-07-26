package dbmodels

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	uuid "github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

type BILL_CTYPES struct {
	//[ 0] CUSTKEY                                        VARCHAR(30)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 30      default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:VARCHAR;size:30;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 1] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 2] C_TYPE                                         VARCHAR(30)          null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 30      default: []
	C_TYPE string `gorm:"primary_key;column:C_TYPE;type:VARCHAR;size:30;" json:"C_TYPE" db:"C_TYPE"`

	RECALC_ID *int64 `gorm:"column:RECALC_ID;type:VARCHAR;size:30;" json:"RECALC_ID" db:"RECALC_ID"`

	CONSUMP *float64 `gorm:"column:CONSUMP;type:VARCHAR;size:30;" json:"CONSUMP" db:"CONSUMP"`

	CONSUMP_PERC *float64 `gorm:"column:CONSUMP_PERC;type:VARCHAR;size:30;" json:"CONSUMP_PERC" db:"CONSUMP_PERC"`

	CONSUMP_TYPE *int64 `gorm:"column:CONSUMP_TYPE;type:VARCHAR;size:30;" json:"CONSUMP_TYPE" db:"CONSUMP_TYPE"`

	NO_UNITS *int64 `gorm:"column:NO_UNITS;type:VARCHAR;size:30;" json:"NO_UNITS" db:"NO_UNITS"`
	//[ 3] PAYMENT_NO                                     VARCHAR(50)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 50      default: []
	PAYMENT_NO *string `gorm:"column:PAYMENT_NO;type:VARCHAR;size:50;" json:"PAYMENT_NO" db:"PAYMENT_NO"`

	BILL_ITEMS
}



type HST_BILL_CTYPES struct {
	RECALC_ID *int64 `gorm:"primary_key;column:RECALC_ID;type:INT;" json:"RECALC_ID" db:"RECALC_ID"`
	BILL_CTYPES
}

// TableName sets the insert table name for this struct type
func (b *BILL_CTYPES) TableName() string {
	return "BILL_CTYPES"
}

// TableName sets the insert table name for this struct type
func (b *HST_BILL_CTYPES) TableName() string {
	return "HST_BILL_CTYPES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BILL_CTYPES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BILL_CTYPES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BILL_CTYPES) Validate(action Action) error {
	return nil
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *HST_BILL_CTYPES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *HST_BILL_CTYPES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *HST_BILL_CTYPES) Validate(action Action) error {
	return nil
}
