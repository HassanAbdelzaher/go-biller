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

type SYSTEM struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DATABASE_VERSION                               FLOAT                null: false  primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DATABASE_VERSION float64 `gorm:"column:DATABASE_VERSION;type:FLOAT;" json:"DATABASE_VERSION" db:"DATABASE_VERSION"`
	//[ 2] PROJECT_TAG                                    NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	PROJECT_TAG *string `gorm:"column:PROJECT_TAG;type:NVARCHAR;size:100;" json:"PROJECT_TAG" db:"PROJECT_TAG"`
	//[ 3] PROJECT_KEY                                    NVARCHAR(256)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	PROJECT_KEY *string `gorm:"column:PROJECT_KEY;type:NVARCHAR;size:256;" json:"PROJECT_KEY" db:"PROJECT_KEY"`
	//[ 4] IS_HQ                                          BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_HQ *bool `gorm:"column:IS_HQ;type:BIT;" json:"IS_HQ" db:"IS_HQ"`
}

// TableName sets the insert table name for this struct type
func (s *SYSTEM) TableName() string {
	return "SYSTEM"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SYSTEM) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SYSTEM) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SYSTEM) Validate(action Action) error {
	return nil
}
