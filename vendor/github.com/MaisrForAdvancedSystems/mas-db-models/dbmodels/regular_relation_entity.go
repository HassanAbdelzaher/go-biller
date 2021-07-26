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

// RegularRelationEntity : Is Table In SQL For Model
type RegularRelationEntity struct {
	RegularChargeID int32 `gorm:"primary_key;column:REGULAR_CHARGE_ID;type:INT;" db:"REGULAR_CHARGE_ID"`
	RelationType    int32 `gorm:"primary_key;column:RELEATION_TYPE;type:INT;" db:"RELEATION_TYPE"`
	EntityType      int32 `gorm:"primary_key;column:ENTITY_TYPE;type:INT;" db:"ENTITY_TYPE"`
}

// TableName sets the insert table name for this RegularRelationEntity
func (t *RegularRelationEntity) TableName() string {
	return "REGULAR_RELATION_ENTITY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *RegularRelationEntity) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *RegularRelationEntity) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *RegularRelationEntity) Validate(action Action) error {
	return nil
}
