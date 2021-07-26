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

//using with database version less >= 20
type HST_BILL_ITEM struct {
	RECALC_ID *int64 `gorm:"primary_key;column:RECALC_ID;type:INT;" json:"RECALC_ID" db:"RECALC_ID"`
	BILL_ITEM
}

// TableName sets the insert table name for this struct type
func (h *HST_BILL_ITEM) TableName() string {
	return "HST_BILL_ITEMS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HST_BILL_ITEM) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HST_BILL_ITEM) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HST_BILL_ITEM) Validate(action Action) error {
	return nil
}
