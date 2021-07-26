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

type LU_ACCESS_CODE struct {
	//[ 0] CODE                                           INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CODE int32 `gorm:"primary_key;column:CODE;type:INT;" json:"CODE" db:"CODE"`
	//[ 1] DESCRIBE                                       NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DESCRIBE *string `gorm:"column:DESCRIBE;type:NVARCHAR;size:100;" json:"DESCRIBE" db:"DESCRIBE"`
	//[ 2] ACTION                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTION *int64 `gorm:"column:ACTION;type:INT;" json:"ACTION" db:"ACTION"`
	//[ 3] ACTION_MSG                                     NVARCHAR(30)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 30      default: []
	ACTION_MSG *string `gorm:"column:ACTION_MSG;type:NVARCHAR;size:30;" json:"ACTION_MSG" db:"ACTION_MSG"`
	//[ 4] ACTIVITY                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTIVITY *int64 `gorm:"column:ACTIVITY;type:INT;" json:"ACTIVITY" db:"ACTIVITY"`
	//[ 5] METER_OP_STATUS                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	METER_OP_STATUS *int64 `gorm:"column:METER_OP_STATUS;type:INT;" json:"METER_OP_STATUS" db:"METER_OP_STATUS"`
	//[ 6] F_METER_CONDITION                              NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	F_METER_CONDITION *string `gorm:"column:F_METER_CONDITION;type:NVARCHAR;size:60;" json:"F_METER_CONDITION" db:"F_METER_CONDITION"`
	//[ 7] LOCK_NOTE                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	LOCK_NOTE *bool `gorm:"column:LOCK_NOTE;type:BIT;" json:"LOCK_NOTE" db:"LOCK_NOTE"`
}

// TableName sets the insert table name for this struct type
func (l *LU_ACCESS_CODE) TableName() string {
	return "LU_ACCESS_CODE"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LU_ACCESS_CODE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LU_ACCESS_CODE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LU_ACCESS_CODE) Validate(action Action) error {
	return nil
}
