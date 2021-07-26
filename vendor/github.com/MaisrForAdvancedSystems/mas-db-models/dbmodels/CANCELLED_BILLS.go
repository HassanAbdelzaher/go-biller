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

type CANCELLED_BILL struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	PAYMENT_NO string `gorm:"primary_key;column:PAYMENT_NO;type:nvarchar;" json:"PAYMENT_NO" db:"PAYMENT_NO"`
	//[ 1] GROUP_ID                                       NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:256;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 2] DESCRIPTION                                    NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	FORM_NO int64 `gorm:"primary_key;column:FORM_NO;type:bigint;" json:"FORM_NO" db:"FORM_NO"`
	//[ 3] UNUSED                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CL_BLNCE       *float64   `gorm:"column:CL_BLNCE;type:FLOAT;" json:"CL_BLNCE" db:"CL_BLNCE"`
	CANCELLED_DATE *time.Time `gorm:"column:CANCELLED_DATE;type:date;" json:"CANCELLED_DATE" db:"CANCELLED_DATE"`
	CANCELLED_BY   *string    `gorm:"column:CANCELLED_BY;type:FLOAT;" json:"CANCELLED_BY" db:"CANCELLED_BY"`
	COUNTER        *int32     `gorm:"column:COUNTER;type:int;" json:"COUNTER" db:"COUNTER"`
	CLOSED         *bool      `gorm:"column:CLOSED;type:bit;" json:"CLOSED" db:"CLOSED"`
	STATION_NO     *int32     `gorm:"column:STATION_NO;type:int;" json:"STATION_NO" db:"STATION_NO"`
	STATE          *int32     `gorm:"column:STATE;type:int;" json:"STATE" db:"STATE"`
	STATUS         *string    `gorm:"column:STATUS;type:nvarchar;" json:"STATUS" db:"STATUS"`
	SURNAME        *string    `gorm:"column:SURNAME;type:NVARCHAR;" json:"SURNAME" db:"SURNAME"`
	BILNG_DATE     *time.Time `gorm:"column:BILNG_DATE;type:date;" json:"BILNG_DATE" db:"BILNG_DATE"`
	DOCUMENT_NO    *string    `gorm:"column:DOCUMENT_NO;type:nvarchar;" json:"DOCUMENT_NO" db:"DOCUMENT_NO"`
	COMMENT        *string    `gorm:"column:COMMENT;type:nvarchar;" json:"COMMENT" db:"COMMENT"`
}

// TableName sets the insert table name for this struct type
func (b *CANCELLED_BILL) TableName() string {
	return "CANCELLED_BILLS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *CANCELLED_BILL) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *CANCELLED_BILL) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *CANCELLED_BILL) Validate(action Action) error {
	return nil
}
