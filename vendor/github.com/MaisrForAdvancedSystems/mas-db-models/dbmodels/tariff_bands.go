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

// TariffBands : Is Table In SQL For Model
type TariffBands struct {
	TariffID   int32     `gorm:"primary_key;column:TARIFF_ID;type:INT;" db:"TARIFF_ID"`
	EffectDate time.Time `gorm:"primary_key;column:EFFECT_DATE;type:DATETIME;" db:"EFFECT_DATE"`
	From       float64   `gorm:"primary_key;column:RANGE_FROM;type:FLOAT;" db:"RANGE_FROM"`
	To         float64   `gorm:"primary_key;column:RANGE_TO;type:FLOAT;" db:"RANGE_TO"`
	Constant   float64   `gorm:"column:CONSTANT;type:FLOAT;" db:"CONSTANT"`
	Charge     float64   `gorm:"column:CHARGE;type:FLOAT;" db:"CHARGE"`
}

// TableName sets the insert table name for this TariffBands
func (t *TariffBands) TableName() string {
	return "TARIFF_BANDS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TariffBands) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TariffBands) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TariffBands) Validate(action Action) error {
	return nil
}
