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

type HAND_MH_ST_COMPACT struct {
	STATION_NO   *int32     `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	CUSTKEY      string     `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	CYCLE_ID     *int32     `gorm:"primary_key;column:CYCLE_ID;type:INT;default:0;" json:"CYCLE_ID" db:"CYCLE_ID"`
	BILNG_DATE   *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	Cl_blnce     *float64   `gorm:"column:cl_blnce;type:FLOAT;" json:"cl_blnce" db:"cl_blnce"`
	Delivery_st  *int64     `gorm:"column:delivery_st;type:INT;default:0;" json:"delivery_st" db:"delivery_st"`
	Payment_no   *string    `gorm:"column:payment_no;type:NVARCHAR;size:100;" json:"payment_no" db:"payment_no"`
	STATM_NO     *int64     `gorm:"column:STATM_NO;type:INT;" json:"STATM_NO" db:"STATM_NO"`
	Tent_name    *string    `gorm:"column:tent_name;type:NVARCHAR;size:300;" json:"tent_name" db:"tent_name"`
	STOP_ISSUE   *bool      `gorm:"column:STOP_ISSUE;type:BIT;" json:"STOP_ISSUE" db:"STOP_ISSUE"`
	S_CR_READING *float64   `gorm:"column:S_CR_READING;type:FLOAT;" json:"S_CR_READING" db:"S_CR_READING"`
	S_PR_READING *float64   `gorm:"column:S_PR_READING;type:FLOAT;" json:"S_PR_READING" db:"S_PR_READING"`
	S_CONSUMP    *float64   `gorm:"column:S_CONSUMP;type:FLOAT;" json:"S_CONSUMP" db:"S_CONSUMP"`
	BILLGROUP    *string    `gorm:"column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	BOOK_NO_C    *string    `gorm:"column:BOOK_NO_C;type:NVARCHAR;size:60;" json:"BOOK_NO_C" db:"BOOK_NO_C"`
	WALK_NO_C    *string    `gorm:"column:WALK_NO_C;type:NVARCHAR;size:60;" json:"WALK_NO_C" db:"WALK_NO_C"`
}

// TableName sets the insert table name for this struct type
func (h *HAND_MH_ST_COMPACT) TableName() string {
	return "HAND_MH_ST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HAND_MH_ST_COMPACT) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HAND_MH_ST_COMPACT) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HAND_MH_ST_COMPACT) Validate(action Action) error {
	return nil
}
