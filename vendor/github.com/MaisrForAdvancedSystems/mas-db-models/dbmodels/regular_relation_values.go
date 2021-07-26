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

// RegularRelationValues : Is Table In SQL For Model
type RegularRelationValues struct {
	RegularChargeID int32    `gorm:"primary_key;column:REGULAR_CHARGE_ID;type:INT;" db:"REGULAR_CHARGE_ID"`
	EntityType      int32    `gorm:"primary_key;column:ENTITY_TYPE;type:INT;" db:"ENTITY_TYPE"`
	LuKey           string   `gorm:"primary_key;column:LU_KEY;type:NVARCHAR;size:128;" db:"LU_KEY"`
	From            *float64 `gorm:"column:RANGE_FROM;type:FLOAT;" db:"RANGE_FROM"`
	To              *float64 `gorm:"column:RANGE_TO;type:FLOAT;" db:"RANGE_TO"`
	Value           *float64 `gorm:"column:VALUE;type:FLOAT;" db:"VALUE"`
	EnableValue     *bool    `gorm:"column:ENABLE_VALUE;type:BIT;" db:"ENABLE_VALUE"`
}

// TableName sets the insert table name for this RegularRelationValues
func (t *RegularRelationValues) TableName() string {
	return "REGULAR_RELATION_VALUES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *RegularRelationValues) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *RegularRelationValues) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *RegularRelationValues) Validate(action Action) error {
	return nil
}
