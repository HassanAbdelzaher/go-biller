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

type EMP_TYPES struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIPTION                                    NVARCHAR(120)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	DESCRIPTION string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:120;" json:"DESCRIPTION" db:"DESCRIPTION"`
}

// TableName sets the insert table name for this struct type
func (e *EMP_TYPES) TableName() string {
	return "EMP_TYPES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EMP_TYPES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EMP_TYPES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *EMP_TYPES) Validate(action Action) error {
	return nil
}
