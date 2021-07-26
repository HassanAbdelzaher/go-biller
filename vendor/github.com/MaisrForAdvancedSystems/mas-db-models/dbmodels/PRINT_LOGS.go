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

type PRINT_LOGS struct {
	//[ 0] PAYMENT_NO                                     NCHAR(60)            null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	PAYMENT_NO string `gorm:"column:PAYMENT_NO;type:NCHAR;size:60;" json:"PAYMENT_NO" db:"PAYMENT_NO"`
	//[ 1] CUSTKEY                                        NCHAR(60)            null: false  primary: true   isArray: false  auto: false  col: NCHAR           len: 60      default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 2] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 3] PRINT_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	PRINT_ID int32 `gorm:"primary_key;column:PRINT_ID;type:INT;" json:"PRINT_ID" db:"PRINT_ID"`
	//[ 4] PRINT_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	PRINT_DATE *time.Time `gorm:"column:PRINT_DATE;type:DATETIME;" json:"PRINT_DATE" db:"PRINT_DATE"`
	//[ 5] PRINT_BY                                       NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	PRINT_BY *string `gorm:"column:PRINT_BY;type:NCHAR;size:60;" json:"PRINT_BY" db:"PRINT_BY"`
	//[ 6] PRINT_TYPE                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PRINT_TYPE *int64 `gorm:"column:PRINT_TYPE;type:INT;" json:"PRINT_TYPE" db:"PRINT_TYPE"`
	//[ 7] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT *float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 8] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG *float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[ 9] BILNG_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATETIME;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[10] IS_CANCELED                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_CANCELED *bool `gorm:"column:IS_CANCELED;type:BIT;" json:"IS_CANCELED" db:"IS_CANCELED"`
	//[11] CANCEL_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CANCEL_DATE *time.Time `gorm:"column:CANCEL_DATE;type:DATETIME;" json:"CANCEL_DATE" db:"CANCEL_DATE"`
	//[12] CANCEL_BY                                      NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	CANCEL_BY *string `gorm:"column:CANCEL_BY;type:NCHAR;size:200;" json:"CANCEL_BY" db:"CANCEL_BY"`
	//[13] CANCEL_AMOUNT                                  FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CANCEL_AMOUNT *float64 `gorm:"column:CANCEL_AMOUNT;type:FLOAT;" json:"CANCEL_AMOUNT" db:"CANCEL_AMOUNT"`
	//[14] USER_NAME                                      NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	USER_NAME *string `gorm:"column:USER_NAME;type:NVARCHAR;size:512;" json:"USER_NAME" db:"USER_NAME"`
	//[15] MACHINE_IP                                     NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	MACHINE_IP *string `gorm:"column:MACHINE_IP;type:NVARCHAR;size:512;" json:"MACHINE_IP" db:"MACHINE_IP"`
	//[16] FROM_SERVER                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FROM_SERVER *bool `gorm:"column:FROM_SERVER;type:BIT;" json:"FROM_SERVER" db:"FROM_SERVER"`
	//[17] RECEIPT_NO                                     NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	RECEIPT_NO *string `gorm:"column:RECEIPT_NO;type:NVARCHAR;size:512;" json:"RECEIPT_NO" db:"RECEIPT_NO"`
	//[18] DEPOSIT_ID                                     BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	DEPOSIT_ID *int64 `gorm:"column:DEPOSIT_ID;type:BIGINT;" json:"DEPOSIT_ID" db:"DEPOSIT_ID"`
}

// TableName sets the insert table name for this struct type
func (p *PRINT_LOGS) TableName() string {
	return "PRINT_LOGS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PRINT_LOGS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PRINT_LOGS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PRINT_LOGS) Validate(action Action) error {
	return nil
}
