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

type NA_COLLECTIONS struct {
	//[ 0] CUSTKEY                                        NVARCHAR(100)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:100;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 1] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 2] DEVICE_ID                                      NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:100;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 3] PAYMENT_NO                                     NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	PAYMENT_NO *string `gorm:"column:PAYMENT_NO;type:NVARCHAR;size:100;" json:"PAYMENT_NO" db:"PAYMENT_NO"`
	//[ 4] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[ 5] MESSAGE                                        NVARCHAR(300)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 300     default: []
	MESSAGE *string `gorm:"column:MESSAGE;type:NVARCHAR;size:300;" json:"MESSAGE" db:"MESSAGE"`
	//[ 6] DELIVERY_ST                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DELIVERY_ST *int64 `gorm:"column:DELIVERY_ST;type:INT;" json:"DELIVERY_ST" db:"DELIVERY_ST"`
}

// TableName sets the insert table name for this struct type
func (n *NA_COLLECTIONS) TableName() string {
	return "NA_COLLECTIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (n *NA_COLLECTIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (n *NA_COLLECTIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (n *NA_COLLECTIONS) Validate(action Action) error {
	return nil
}
