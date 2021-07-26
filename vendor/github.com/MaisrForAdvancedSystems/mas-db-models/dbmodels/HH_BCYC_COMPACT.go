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

type HH_BCYC_COMPACT struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] BILLGROUP                                      NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 2] BOOK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NVARCHAR;size:60;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 3] WALK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	WALK_NO string `gorm:"primary_key;column:WALK_NO;type:NVARCHAR;size:60;" json:"WALK_NO" db:"WALK_NO"`
	//[ 4] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 5] IS_COLLECTION                                  BIT                  null: false  primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_COLLECTION bool `gorm:"column:IS_COLLECTION;type:BIT;" json:"IS_COLLECTION" db:"IS_COLLECTION"`
	//[ 6] IS_READING                                     BIT                  null: false  primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_READING bool `gorm:"column:IS_READING;type:BIT;" json:"IS_READING" db:"IS_READING"`
	//[ 7] BRANCH                                         NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	BRANCH *string `gorm:"column:BRANCH;type:NVARCHAR;size:200;" json:"BRANCH" db:"BRANCH"`
	//[ 8] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE  *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	ALLOW_FAWRY *int64     `gorm:"column:ALLOW_FAWRY;type:INT;" json:"ALLOW_FAWRY" db:"ALLOW_FAWRY"`
	//[96] FAWRY_OPEN_BY                                  VARCHAR(128)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
	FAWRY_OPEN_BY *string `gorm:"column:FAWRY_OPEN_BY;type:VARCHAR;size:128;" json:"FAWRY_OPEN_BY" db:"FAWRY_OPEN_BY"`
	//[97] FAWRY_CLOSE_BY                                 VARCHAR(128)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
	FAWRY_CLOSE_BY *string `gorm:"column:FAWRY_CLOSE_BY;type:VARCHAR;size:128;" json:"FAWRY_CLOSE_BY" db:"FAWRY_CLOSE_BY"`
	//[98] FAWRY_OPEN_DATE                                DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	FAWRY_OPEN_DATE *time.Time `gorm:"column:FAWRY_OPEN_DATE;type:DATETIME;" json:"FAWRY_OPEN_DATE" db:"FAWRY_OPEN_DATE"`
	//[99] FAWRY_CLOSE_DATE                               DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	FAWRY_CLOSE_DATE    *time.Time `gorm:"column:FAWRY_CLOSE_DATE;type:DATETIME;" json:"FAWRY_CLOSE_DATE" db:"FAWRY_CLOSE_DATE"`
	ISCYCLE_COMPLETED_C *int64     `gorm:"column:ISCYCLE_COMPLETED_C;type:INT;" json:"ISCYCLE_COMPLETED_C" db:"ISCYCLE_COMPLETED_C"`
	//[34] ISCYCLE_COMPLETED_R                            INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ISCYCLE_COMPLETED_R *int64 `gorm:"column:ISCYCLE_COMPLETED_R;type:INT;" json:"ISCYCLE_COMPLETED_R" db:"ISCYCLE_COMPLETED_R"`
}

// TableName sets the insert table name for this struct type
func (h *HH_BCYC_COMPACT) TableName() string {
	return "HH_BCYC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HH_BCYC_COMPACT) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HH_BCYC_COMPACT) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HH_BCYC_COMPACT) Validate(action Action) error {
	return nil
}

func (h *HH_BCYC_COMPACT) IsCycleComplatedC() bool {
	if h.ISCYCLE_COMPLETED_C == nil || *h.ISCYCLE_COMPLETED_C == 0 {
		return false
	}
	return true
}

func (h *HH_BCYC_COMPACT) IsCycleComplated() bool {
	if h.ISCYCLE_COMPLETED_R == nil || *h.ISCYCLE_COMPLETED_R == 0 {
		return false
	}
	return true
}
