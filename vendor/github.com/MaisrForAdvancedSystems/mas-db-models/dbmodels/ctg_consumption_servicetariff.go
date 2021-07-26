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

// CtgConsumptionServicetariff : Is Table In SQL For Model
type CtgConsumptionServicetariff struct {
	CtypeID            string   `gorm:"primary_key;column:CTYPE_ID;type:NVARCHAR;size:50;" db:"CTYPE_ID"`
	ServiceType        int32    `gorm:"primary_key;column:SERVICE_TYPE;type:INT;" db:"SERVICE_TYPE"`
	TransCode          string   `gorm:"column:TRANS_CODE;type:NVARCHAR;size:50;" db:"TRANS_CODE"`
	TariffID           string   `gorm:"column:TARIFF_ID;type:NVARCHAR;size:50;" db:"TARIFF_ID"`
	CtypegrpID         *string  `gorm:"column:CTYPEGRP_ID;type:NVARCHAR;size:50;" db:"CTYPEGRP_ID"`
	IsZeroTarif        *bool    `gorm:"column:IS_ZERO_TARIF;type:BIT;" db:"IS_ZERO_TARIF"`
	TaxPercentage      *float64 `gorm:"column:TAX_PERCENTAGE;type:FLOAT;" db:"ENTITY_TYPE"`
	DiscountPercentage *float64 `gorm:"column:DISCOUNT_PERCENTAGE;type:FLOAT;" db:"DISCOUNT_PERCENTAGE"`
}

// TableName sets the insert table name for this CtgConsumptionServicetariff
func (t *CtgConsumptionServicetariff) TableName() string {
	return "CTG_CONSUMPTION_SERVICETARIFF"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *CtgConsumptionServicetariff) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *CtgConsumptionServicetariff) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *CtgConsumptionServicetariff) Validate(action Action) error {
	return nil
}
