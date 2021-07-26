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

type LU_COMPLAINT_TYPE struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIPE                                       NVARCHAR(500)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 500     default: []
	DESCRIPE *string `gorm:"column:DESCRIPE;type:NVARCHAR;size:500;" json:"DESCRIPE" db:"DESCRIPE"`
	//[ 2] ACTION_TYPE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTION_TYPE *int64 `gorm:"column:ACTION_TYPE;type:INT;" json:"ACTION_TYPE" db:"ACTION_TYPE"`
}

// TableName sets the insert table name for this struct type
func (l *LU_COMPLAINT_TYPE) TableName() string {
	return "LU_COMPLAINT_TYPE"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LU_COMPLAINT_TYPE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LU_COMPLAINT_TYPE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LU_COMPLAINT_TYPE) Validate(action Action) error {
	return nil
}
