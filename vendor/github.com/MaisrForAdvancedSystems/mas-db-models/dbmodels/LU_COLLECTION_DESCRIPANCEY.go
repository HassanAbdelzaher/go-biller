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

type LU_COLLECTION_DESCRIPANCEY struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIPE                                       NVARCHAR(500)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 500     default: []
	DESCRIPE *string `gorm:"column:DESCRIPE;type:NVARCHAR;size:500;" json:"DESCRIPE" db:"DESCRIPE"`
}

// TableName sets the insert table name for this struct type
func (l *LU_COLLECTION_DESCRIPANCEY) TableName() string {
	return "LU_COLLECTION_DESCRIPANCEY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LU_COLLECTION_DESCRIPANCEY) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LU_COLLECTION_DESCRIPANCEY) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LU_COLLECTION_DESCRIPANCEY) Validate(action Action) error {
	return nil
}
