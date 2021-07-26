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

type EXT_READINGS struct {
	//[ 0] NAME                                           VARCHAR(100)         null: false  primary: true   isArray: false  auto: false  col: VARCHAR         len: 100     default: []
	NAME string `gorm:"primary_key;column:NAME;type:VARCHAR;size:100;" json:"NAME" db:"NAME"`
	//[ 1] ADDRESS                                        VARCHAR(150)         null: false  primary: false  isArray: false  auto: false  col: VARCHAR         len: 150     default: []
	ADDRESS string `gorm:"column:ADDRESS;type:VARCHAR;size:150;" json:"ADDRESS" db:"ADDRESS"`
	//[ 2] CUSTKEY                                        VARCHAR(30)          null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 30      default: []
	CUSTKEY *string `gorm:"column:CUSTKEY;type:VARCHAR;size:30;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 3] CR_READING                                     BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	CR_READING int64 `gorm:"column:CR_READING;type:BIGINT;" json:"CR_READING" db:"CR_READING"`
	//[ 4] STAMP_DATE                                     DATETIME             null: false  primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
}

// TableName sets the insert table name for this struct type
func (e *EXT_READINGS) TableName() string {
	return "EXT_READINGS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EXT_READINGS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EXT_READINGS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *EXT_READINGS) Validate(action Action) error {
	return nil
}
