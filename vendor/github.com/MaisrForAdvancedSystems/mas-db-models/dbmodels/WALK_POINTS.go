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

type WALK_POINTS struct {
	//[ 0] BOOK_NO                                        NCHAR(20)            null: false  primary: true   isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NCHAR;size:20;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 1] WALK_NO                                        INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	WALK_NO int32 `gorm:"primary_key;column:WALK_NO;type:INT;" json:"WALK_NO" db:"WALK_NO"`
	//[ 2] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 3] LAT                                            FLOAT                null: false  primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 4] LNG                                            FLOAT                null: false  primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[ 5] ANGLE                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ANGLE *float64 `gorm:"column:ANGLE;type:FLOAT;" json:"ANGLE" db:"ANGLE"`
	//[ 6] DATA                                           NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DATA *string `gorm:"column:DATA;type:NVARCHAR;size:100;" json:"DATA" db:"DATA"`
}

// TableName sets the insert table name for this struct type
func (w *WALK_POINTS) TableName() string {
	return "WALK_POINTS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WALK_POINTS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WALK_POINTS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WALK_POINTS) Validate(action Action) error {
	return nil
}
