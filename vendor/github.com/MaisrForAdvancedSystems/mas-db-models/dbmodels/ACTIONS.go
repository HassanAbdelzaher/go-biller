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

type ACTIONS struct {
	//[ 0] ACTION_ID                                      INT                  null: false  primary: true   isArray: false  auto: true   col: INT             len: -1      default: []
	ACTION_ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ACTION_ID;type:INT;" json:"ACTION_ID" db:"ACTION_ID"`
	//[ 1] DESCRIPTION                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;" json:"DESCRIPTION" db:"DESCRIPTION"`
}

// TableName sets the insert table name for this struct type
func (a *ACTIONS) TableName() string {
	return "ACTIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ACTIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ACTIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ACTIONS) Validate(action Action) error {
	return nil
}
