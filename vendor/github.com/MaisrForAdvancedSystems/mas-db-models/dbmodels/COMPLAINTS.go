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

type COMPLAINTS struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] DEVICE_COMPLAINT_ID                            BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	DEVICE_COMPLAINT_ID int64 `gorm:"column:DEVICE_COMPLAINT_ID;type:BIGINT;" json:"DEVICE_COMPLAINT_ID" db:"DEVICE_COMPLAINT_ID"`
	//[ 2] DEVICE_ID                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 3] BOOK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 4] WALK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NVARCHAR;" json:"WALK_NO" db:"WALK_NO"`
	//[ 5] SHAKWA_TYPE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SHAKWA_TYPE *int64 `gorm:"column:SHAKWA_TYPE;type:INT;" json:"SHAKWA_TYPE" db:"SHAKWA_TYPE"`
	//[ 6] SHAKWA_DEGREE                                  INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SHAKWA_DEGREE *int64 `gorm:"column:SHAKWA_DEGREE;type:INT;" json:"SHAKWA_DEGREE" db:"SHAKWA_DEGREE"`
	//[ 7] NOTES                                          NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	NOTES *string `gorm:"column:NOTES;type:NVARCHAR;" json:"NOTES" db:"NOTES"`
	//[ 8] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT *float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 9] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG *float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[10] STAMP_USER                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;" json:"STAMP_USER" db:"STAMP_USER"`
	//[11] SYNC_ST                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SYNC_ST *int64 `gorm:"column:SYNC_ST;type:INT;" json:"SYNC_ST" db:"SYNC_ST"`
	//[12] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[13] LOCATION_SOURCE                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	LOCATION_SOURCE *int64 `gorm:"column:LOCATION_SOURCE;type:INT;" json:"LOCATION_SOURCE" db:"LOCATION_SOURCE"`
	//[14] LOCATION_DATE                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	LOCATION_DATE *time.Time `gorm:"column:LOCATION_DATE;type:DATETIME;" json:"LOCATION_DATE" db:"LOCATION_DATE"`
	//[15] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY *float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
}

// TableName sets the insert table name for this struct type
func (c *COMPLAINTS) TableName() string {
	return "COMPLAINTS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *COMPLAINTS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *COMPLAINTS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *COMPLAINTS) Validate(action Action) error {
	return nil
}
