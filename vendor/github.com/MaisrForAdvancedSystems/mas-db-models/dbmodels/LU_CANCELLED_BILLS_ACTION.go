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

type LU_CANCELLED_BILLS_ACTION struct {
	//[ 0] CODE                                           INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIBE                                       NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:100;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 2] ACTION                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CURRENT_STATE *int32 `gorm:"column:CURRENT_STATE;type:INT;" json:"CURRENT_STATE" db:"CURRENT_STATE"`

	NEXT_STATE *int32 `gorm:"column:NEXT_STATE;type:INT;" json:"NEXT_STATE" db:"NEXT_STATE"`

	CLOSED *bool `gorm:"column:CLOSED;type:BIT;" json:"CLOSED" db:"CLOSED"`

	START_UP *bool `gorm:"column:START_UP;type:BIT;" json:"START_UP" db:"START_UP"`

	DEPARTMENT *int32 `gorm:"column:DEPARTMENT;type:INT;" json:"DEPARTMENT" db:"DEPARTMENT"`

	RECALC_DONE_ACTION *bool `gorm:"column:RECALC_DONE_ACTION;type:BIT;" json:"RECALC_DONE_ACTION" db:"RECALC_DONE_ACTION"`

}

// TableName sets the insert table name for this struct type
func (l *LU_CANCELLED_BILLS_ACTION) TableName() string {
	return "LU_CANCELLED_BILLS_ACTIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LU_CANCELLED_BILLS_ACTION) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LU_CANCELLED_BILLS_ACTION) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LU_CANCELLED_BILLS_ACTION) Validate(action Action) error {
	return nil
}
