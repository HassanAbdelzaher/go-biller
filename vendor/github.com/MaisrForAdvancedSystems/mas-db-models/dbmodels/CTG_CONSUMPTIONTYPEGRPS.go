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

type CTG_CONSUMPTIONTYPEGRPS struct {
	//[ 0] CTYPEGRP_ID                                    NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	CTYPEGRP_ID string `gorm:"primary_key;column:CTYPEGRP_ID;type:NVARCHAR;size:256;" json:"CTYPEGRP_ID" db:"CTYPEGRP_ID"`
	//[ 1] DESCRIPTION                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;" json:"DESCRIPTION" db:"DESCRIPTION"`
}

// TableName sets the insert table name for this struct type
func (c *CTG_CONSUMPTIONTYPEGRPS) TableName() string {
	return "CTG_CONSUMPTIONTYPEGRPS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CTG_CONSUMPTIONTYPEGRPS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CTG_CONSUMPTIONTYPEGRPS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CTG_CONSUMPTIONTYPEGRPS) Validate(action Action) error {
	return nil
}
