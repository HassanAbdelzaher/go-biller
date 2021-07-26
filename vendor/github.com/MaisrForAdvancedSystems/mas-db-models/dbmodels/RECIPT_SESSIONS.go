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

type RECIPT_SESSIONS struct {
	//[ 0] SESSION_ID                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	SESSION_ID int32 `gorm:"primary_key;column:SESSION_ID;type:INT;" json:"SESSION_ID" db:"SESSION_ID"`
	//[ 1] CASHER                                         INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CASHER int32 `gorm:"column:CASHER;type:INT;" json:"CASHER" db:"CASHER"`
	//[ 2] START_DATE                                     DATETIME             null: false  primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	START_DATE *time.Time `gorm:"column:START_DATE;type:DATETIME;" json:"START_DATE" db:"START_DATE"`
	//[ 3] END_DATE                                       DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	END_DATE *time.Time `gorm:"column:END_DATE;type:DATETIME;" json:"END_DATE" db:"END_DATE"`
	//[ 4] IS_CLOSED                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_CLOSED *bool `gorm:"column:IS_CLOSED;type:BIT;" json:"IS_CLOSED" db:"IS_CLOSED"`
	//[ 5] DEPOSIT_ID                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DEPOSIT_ID *int64 `gorm:"column:DEPOSIT_ID;type:INT;" json:"DEPOSIT_ID" db:"DEPOSIT_ID"`
	//[ 6] DEPOSIT_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	DEPOSIT_DATE *time.Time `gorm:"column:DEPOSIT_DATE;type:DATETIME;" json:"DEPOSIT_DATE" db:"DEPOSIT_DATE"`
	//[ 7] RECIPT_NO                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	RECIPT_NO *string `gorm:"column:RECIPT_NO;type:NVARCHAR;" json:"RECIPT_NO" db:"RECIPT_NO"`
	//[ 8] AMT_COLLECTED                                  DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	AMT_COLLECTED *float64 `gorm:"column:AMT_COLLECTED;type:DECIMAL;" json:"AMT_COLLECTED" db:"AMT_COLLECTED"`
	//[ 9] COUNT_COLLECTED                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COUNT_COLLECTED *int64 `gorm:"column:COUNT_COLLECTED;type:INT;" json:"COUNT_COLLECTED" db:"COUNT_COLLECTED"`
}

// TableName sets the insert table name for this struct type
func (r *RECIPT_SESSIONS) TableName() string {
	return "RECIPT_SESSIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RECIPT_SESSIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RECIPT_SESSIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RECIPT_SESSIONS) Validate(action Action) error {
	return nil
}
