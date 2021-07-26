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

type LU_CANCELLED_BILL_STATE struct {
	//[ 0] CODE                                           INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIBE                                       NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:100;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 2] ACTION                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	RECAL_READY *bool `gorm:"column:RECAL_READY;type:INT;" json:"RECAL_READY" db:"RECAL_READY"`

}

// TableName sets the insert table name for this struct type
func (l *LU_CANCELLED_BILL_STATE) TableName() string {
	return "LU_CANCELLED_BILL_STATES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LU_CANCELLED_BILL_STATE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LU_CANCELLED_BILL_STATE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LU_CANCELLED_BILL_STATE) Validate(action Action) error {
	return nil
}
