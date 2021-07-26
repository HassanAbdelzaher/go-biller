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

/*STATION_NO,CUSTKEY,CYCLE_ID,RECALC_ID,BILNG_DATE,PAYMENT_NO,NO_UNITS,C_TYPE,CONN_AVRG_CONSUMP,MODIFIED_AVRG_CONSUMP,READ_TYPE,
cl_blnce,OP_BLNCE,CUR_CHARGES,S_CONSUMP,S_CR_READING,S_PR_READING*/
//using with database version less >= 20
type HST_HAND_MH_ST struct {
	RECALC_ID *int64 `gorm:"primary_key;column:RECALC_ID;type:INT;" json:"RECALC_ID" db:"RECALC_ID"`
	HAND_MH_ST
}

// TableName sets the insert table name for this struct type
func (h *HST_HAND_MH_ST) TableName() string {
	return "HST_HAND_MH_ST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HST_HAND_MH_ST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HST_HAND_MH_ST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HST_HAND_MH_ST) Validate(action Action) error {
	return nil
}
