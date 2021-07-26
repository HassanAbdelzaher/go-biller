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

// TariffExtraFess : Is Table In SQL For Model
type TariffExtraFess struct {
	TariffID       int32    `gorm:"primary_key;column:TARIFF_ID;type:NVARCHAR;size:128;" db:"TARIFF_ID"`
	TransCode      string   `gorm:"primary_key;column:TRANS_CODE;type:NVARCHAR;size:128;" db:"TRANS_CODE"`
	Description    *string  `gorm:"column:DESCRIPTION;type:NVARCHAR;size:MAX;" db:"DESCRIPTION"`
	AmountPerMeter *float64 `gorm:"column:AMOUNT_PER_METER;type:FLOAT;" db:"AMOUNT_PER_METER"`
	TaxPercentage  *float64 `gorm:"column:TAX_PERCENTAGE;type:FLOAT;" db:"TAX_PERCENTAGE"`
	FixedAmount    *float64 `gorm:"column:FIXED_AMOUNT;type:FLOAT;" db:"FIXED_AMOUNT"`
}

// TableName sets the insert table name for this TariffExtraFess
func (t *TariffExtraFess) TableName() string {
	return "TARIFF_EXTRA_FESS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TariffExtraFess) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TariffExtraFess) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TariffExtraFess) Validate(action Action) error {
	return nil
}
