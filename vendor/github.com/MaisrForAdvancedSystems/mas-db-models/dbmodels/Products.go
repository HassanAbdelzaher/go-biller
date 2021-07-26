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

type Products struct {
	//[ 0] name                                           NVARCHAR(510)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	Name string `gorm:"primary_key;column:name;type:NVARCHAR;size:510;" json:"name" db:"name"`
	//[ 1] description                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	Description *string `gorm:"column:description;type:NVARCHAR;" json:"description" db:"description"`
}

// TableName sets the insert table name for this struct type
func (p *Products) TableName() string {
	return "Products"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Products) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Products) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Products) Validate(action Action) error {
	return nil
}
