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

type IDS struct {
	//[ 0] TABLE_NAME                                     NVARCHAR(200)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TABLE_NAME string `gorm:"primary_key;column:TABLE_NAME;type:NVARCHAR;size:200;" json:"TABLE_NAME" db:"TABLE_NAME"`
	//[ 1] COLUMN_NAME                                    NVARCHAR(200)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	COLUMN_NAME string `gorm:"primary_key;column:COLUMN_NAME;type:NVARCHAR;size:200;" json:"COLUMN_NAME" db:"COLUMN_NAME"`
	//[ 2] CURRENT_VALUE                                  BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	CURRENT_VALUE int64 `gorm:"column:CURRENT_VALUE;type:BIGINT;" json:"CURRENT_VALUE" db:"CURRENT_VALUE"`
	//[ 3] MIN_VALUE                                      BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	MIN_VALUE int64 `gorm:"column:MIN_VALUE;type:BIGINT;" json:"MIN_VALUE" db:"MIN_VALUE"`
	//[ 4] MAX_VALUE                                      BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	MAX_VALUE int64 `gorm:"column:MAX_VALUE;type:BIGINT;" json:"MAX_VALUE" db:"MAX_VALUE"`
	//[ 5] STEP                                           INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STEP int32 `gorm:"column:STEP;type:INT;" json:"STEP" db:"STEP"`
	//[ 6] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
}

// TableName sets the insert table name for this struct type
func (i *IDS) TableName() string {
	return "IDS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (i *IDS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (i *IDS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (i *IDS) Validate(action Action) error {
	return nil
}
