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

type OS struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] NAME                                           NVARCHAR(300)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 300     default: []
	NAME string `gorm:"column:NAME;type:NVARCHAR;size:300;" json:"NAME" db:"NAME"`
	//[ 2] IS_SUPPORTED                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_SUPPORTED *bool `gorm:"column:IS_SUPPORTED;type:BIT;" json:"IS_SUPPORTED" db:"IS_SUPPORTED"`
}

// TableName sets the insert table name for this struct type
func (o *OS) TableName() string {
	return "OS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OS) Validate(action Action) error {
	return nil
}
