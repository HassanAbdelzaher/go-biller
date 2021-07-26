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

type CONN_DTL_TARIFF_ALLOC struct {
	CUSTKEY          string     `gorm:"primary_key;column:CUSTKEY;type:nvarchar;" json:"CUSTKEY" db:"CUSTKEY"`
	CYCLE_ID         int64      `gorm:"primary_key;column:CYCLE_ID;type:NVARCHAR;" json:"CYCLE_ID" db:"CYCLE_ID"`
	C_TYPE           string     `gorm:"primary_key;column:C_TYPE;type:NVARCHAR;" json:"C_TYPE" db:"C_TYPE"`
	PROP_REF         *string    `gorm:"column:PROP_REF;type:NVARCHAR;" json:"PROP_REF" db:"PROP_REF"`
	BILNG_DATE       *time.Time `gorm:"column:BILNG_DATE;type:NVARCHAR;" json:"BILNG_DATE" db:"BILNG_DATE"`
	ALLOC_PERC       *float64   `gorm:"column:ALLOC_PERC;type:decimal;" json:"ALLOC_PERC" db:"ALLOC_PERC"`
	ALLOC_PERC_SEWER *float64   `gorm:"column:ALLOC_PERC_SEWER;type:decimal;" json:"ALLOC_PERC_SEWER" db:"ALLOC_PERC_SEWER"`
	ESTIM_CONS_PU    *float64   `gorm:"column:ESTIM_CONS_PU;type:decimal;" json:"ESTIM_CONS_PU" db:"ESTIM_CONS_PU"`
	NO_UNITS         *int64     `gorm:"column:NO_UNITS;type:int;" json:"NO_UNITS" db:"NO_UNITS"`
	DESCRIPTION      *string    `gorm:"column:DESCRIPTION;type:nvarchar;" json:"DESCRIPTION" db:"DESCRIPTION"`
	UPDATE_DATE      *time.Time `gorm:"column:UPDATE_DATE;type:nvarchar;" json:"UPDATE_DATE" db:"UPDATE_DATE"`
}

// TableName sets the insert table name for this struct type
func (a *CONN_DTL_TARIFF_ALLOC) TableName() string {
	return "CONN_DTL_TARIFF_ALLOC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *CONN_DTL_TARIFF_ALLOC) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *CONN_DTL_TARIFF_ALLOC) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *CONN_DTL_TARIFF_ALLOC) Validate(action Action) error {
	return nil
}
