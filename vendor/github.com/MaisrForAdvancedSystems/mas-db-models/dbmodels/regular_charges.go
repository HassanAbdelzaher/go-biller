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

// RegularCharges : Is Table In SQL For Model
type RegularCharges struct {
	RegularChargeID     int32      `gorm:"primary_key;column:REGULAR_CHARGE_ID;type:INT;" db:"REGULAR_CHARGE_ID"`
	RegularChargeCode   string     `gorm:"column:REGULAR_CHARE_CODE;type:NVARCHAR;size:50;" db:"REGULAR_CHARE_CODE"`
	EffectDate          time.Time  `gorm:"column:EFFECT_DATE;type:DATETIME;" db:"EFFECT_DATE"`
	EffectDateTo        *time.Time `gorm:"column:EFFECT_DATE_TO;type:DATETIME;" db:"EFFECT_DATE_TO"`
	ServiceType         int32      `gorm:"column:SERVICE_TYPE;type:INT;" db:"SERVICE_TYPE"`
	TransCode           string     `gorm:"column:TRANS_CODE;type:NVARCHAR;size:50;" db:"TRANS_CODE"`
	Title               string     `gorm:"column:TITLE;type:NVARCHAR;size:250;" db:"TITLE"`
	ChargeType          int32      `gorm:"column:CHARGE_TYPE;type:INT;" db:"CHARGE_TYPE"`
	IsChargable         bool       `gorm:"column:IS_CHARGABLE;type:BIT;" db:"IS_CHARGABLE"`
	ChargeCalcPeriod    int32      `gorm:"column:CHARGE_CALC_PERIOD;type:INT;" db:"CHARGE_CALC_PERIOD"`
	ChargeInterval      *int64     `gorm:"column:CHARGE_INTERVAL;type:BIGINT;" db:"CHARGE_INTERVAL"`
	ChargeMonthlyDay    *int64     `gorm:"column:CHARGE_MONTHLY_DAY;type:BIGINT;" db:"CHARGE_MONTHLY_DAY"`
	FixedCharge         *float64   `gorm:"column:FIXED_CHARGE;type:FLOAT;" db:"FIXED_CHARGE"`
	FixedChargeDiscount *float64   `gorm:"column:FIXED_CHARGE_DISCOUNT;type:FLOAT;" db:"FIXED_CHARGE_DISCOUNT"`
	MinCharge           *float64   `gorm:"column:MIN_CHARGE;type:FLOAT;" db:"MIN_CHARGE"`
	VatPercentage       *float64   `gorm:"column:VAT_PERCENTAGE;type:FLOAT;" db:"VAT_PERCENTAGE"`
	//RelationChargeCode       *string    `gorm:"column:RELATION_CHARGE_CODE;type:NVARCHAR;size:128;" db:"RELATION_CHARGE_CODE"`
	//RelationChargeEntityType *int32     `gorm:"column:RELATION_CHARGE_ENTITY_TYPE;type:INT;" db:"RELATION_CHARGE_ENTITY_TYPE"`
	//RelationEnableCode       *string    `gorm:"column:RELATION_ENABLE_CODE;type:NVARCHAR;size:128;" db:"RELATION_ENABLE_CODE"`
	//RelationEnableEntityType *int32     `gorm:"column:RELATION_ENABLE_ENTITY_TYPE;type:INT;" db:"RELATION_ENABLE_ENTITY_TYPE"`
	ByPass       *bool  `gorm:"column:BY_PASS;type:BIT;" db:"BY_PASS"`
	CalcStartegy *int32 `gorm:"column:CALC_STARTEGY;type:INT;" db:"CALC_STARTEGY"`
	PerUnit      *bool  `gorm:"column:PER_UNIT;type:BIT;" db:"PER_UNIT"`
}

// TableName sets the insert table name for this RegularCharges
func (t *RegularCharges) TableName() string {
	return "REGULAR_CHARGES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *RegularCharges) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *RegularCharges) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *RegularCharges) Validate(action Action) error {
	return nil
}
