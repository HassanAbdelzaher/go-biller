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

type DEVICE_LOCATIONS struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: true   col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] DEVICE_ID                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:200;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 2] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT *float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 3] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG *float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[ 4] START_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	START_DATE *time.Time `gorm:"column:START_DATE;type:DATETIME;" json:"START_DATE" db:"START_DATE"`
	//[ 5] GPS_SOURCE                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	GPS_SOURCE *int64 `gorm:"column:GPS_SOURCE;type:INT;" json:"GPS_SOURCE" db:"GPS_SOURCE"`
	//[ 6] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY *float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
	//[ 7] REPEAT                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	REPEAT *int64 `gorm:"column:REPEAT;type:INT;" json:"REPEAT" db:"REPEAT"`
	//[ 8] END_DATE                                       DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	END_DATE *time.Time `gorm:"column:END_DATE;type:DATETIME;" json:"END_DATE" db:"END_DATE"`
	//[ 9] CHANGE                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CHANGE *int64 `gorm:"column:CHANGE;type:INT;" json:"CHANGE" db:"CHANGE"`
}

// TableName sets the insert table name for this struct type
func (d *DEVICE_LOCATIONS) TableName() string {
	return "DEVICE_LOCATIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEVICE_LOCATIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEVICE_LOCATIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEVICE_LOCATIONS) Validate(action Action) error {
	return nil
}
