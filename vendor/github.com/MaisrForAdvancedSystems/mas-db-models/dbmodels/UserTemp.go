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

type UserTemp struct {
	//[ 0] id                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	Id int32 `gorm:"primary_key;column:id;type:INT;" json:"id" db:"id"`
	//[ 1] nameAr                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	NameAr *string `gorm:"column:nameAr;type:NVARCHAR;size:100;" json:"nameAr" db:"nameAr"`
	//[ 2] addressAr                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	AddressAr *string `gorm:"column:addressAr;type:NVARCHAR;" json:"addressAr" db:"addressAr"`
}

// TableName sets the insert table name for this struct type
func (u *UserTemp) TableName() string {
	return "UserTemp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserTemp) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserTemp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserTemp) Validate(action Action) error {
	return nil
}
