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

type DASHBOARD_COLLECTIONS struct {
	//[ 0] branch                                         NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	Branch *string `gorm:"column:branch;type:NCHAR;size:200;" json:"branch" db:"branch"`
	//[ 1] amount                                         FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	Amount *float64 `gorm:"column:amount;type:FLOAT;" json:"amount" db:"amount"`
	//[ 2] date                                           DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	Date *time.Time `gorm:"column:date;type:DATE;" json:"date" db:"date"`
	//[ 3] stamp_date                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	Stamp_date *time.Time `gorm:"column:stamp_date;type:DATETIME;" json:"stamp_date" db:"stamp_date"`
	//[ 4] id                                             BIGINT               null: false  primary: true   isArray: false  auto: true   col: BIGINT          len: -1      default: []
	Id int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:BIGINT;" json:"id" db:"id"`
}

// TableName sets the insert table name for this struct type
func (d *DASHBOARD_COLLECTIONS) TableName() string {
	return "DASHBOARD_COLLECTIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DASHBOARD_COLLECTIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DASHBOARD_COLLECTIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DASHBOARD_COLLECTIONS) Validate(action Action) error {
	return nil
}
