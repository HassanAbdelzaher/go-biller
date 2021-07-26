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

type RECEIPT_DTLS struct {
	//[ 0] RECEIPT_NO                                     NVARCHAR(100)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	RECEIPT_NO string `gorm:"primary_key;column:RECEIPT_NO;type:NVARCHAR;size:100;" json:"RECEIPT_NO" db:"RECEIPT_NO"`
	//[ 1] CUSTKEY                                        NVARCHAR(100)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CUSTKEY string `gorm:"column:CUSTKEY;type:NVARCHAR;size:100;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 2] TARGET_PAYMENT_NO                              NVARCHAR(100)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	TARGET_PAYMENT_NO string `gorm:"primary_key;column:TARGET_PAYMENT_NO;type:NVARCHAR;size:100;" json:"TARGET_PAYMENT_NO" db:"TARGET_PAYMENT_NO"`
	//[ 3] AMOUNT                                         FLOAT                null: false  primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	AMOUNT float64 `gorm:"column:AMOUNT;type:FLOAT;" json:"AMOUNT" db:"AMOUNT"`
}

// TableName sets the insert table name for this struct type
func (r *RECEIPT_DTLS) TableName() string {
	return "RECEIPT_DTLS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RECEIPT_DTLS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RECEIPT_DTLS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RECEIPT_DTLS) Validate(action Action) error {
	return nil
}
