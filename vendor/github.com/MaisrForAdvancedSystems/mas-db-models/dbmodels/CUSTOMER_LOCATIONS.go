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

type CUSTOMER_LOCATIONS struct {
	//[ 0] CUSTKEY                                        NVARCHAR(100)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:100;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 1] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT * float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 2] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG * float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[ 3] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 4] GPS_SOURCE                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	GPS_SOURCE *int64 `gorm:"column:GPS_SOURCE;type:INT;" json:"GPS_SOURCE" db:"GPS_SOURCE"`
	//[ 5] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY * float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
	//[ 6] BOOK_NO_R                                      NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	BOOK_NO_R *string `gorm:"column:BOOK_NO_R;type:NVARCHAR;size:40;" json:"BOOK_NO_R" db:"BOOK_NO_R"`
	//[ 7] WALK_NO_R                                      NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	WALK_NO_R *string `gorm:"column:WALK_NO_R;type:NVARCHAR;size:40;" json:"WALK_NO_R" db:"WALK_NO_R"`
	//[ 8] SEQ_NO_R                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQ_NO_R *int64 `gorm:"column:SEQ_NO_R;type:INT;" json:"SEQ_NO_R" db:"SEQ_NO_R"`
	//[ 9] BOOK_NO_C                                      NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	BOOK_NO_C *string `gorm:"column:BOOK_NO_C;type:NVARCHAR;size:40;" json:"BOOK_NO_C" db:"BOOK_NO_C"`
	//[10] WALK_NO_C                                      NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	WALK_NO_C *string `gorm:"column:WALK_NO_C;type:NVARCHAR;size:40;" json:"WALK_NO_C" db:"WALK_NO_C"`
	//[11] SEQ_NO_C                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQ_NO_C *int64 `gorm:"column:SEQ_NO_C;type:INT;" json:"SEQ_NO_C" db:"SEQ_NO_C"`
	//[12] DEVICE_ID                                      NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:100;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[13] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[14] UPDATE_DATE                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	UPDATE_DATE *time.Time `gorm:"column:UPDATE_DATE;type:DATE;" json:"UPDATE_DATE" db:"UPDATE_DATE"`
	//[15] STAMP_USER                                     NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;size:100;" json:"STAMP_USER" db:"STAMP_USER"`
	//[16] OPERATION_FLAGE                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OPERATION_FLAGE *int64 `gorm:"column:OPERATION_FLAGE;type:INT;" json:"OPERATION_FLAGE" db:"OPERATION_FLAGE"`
	//[17] DISTANCE_REF                                   FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DISTANCE_REF * float64 `gorm:"column:DISTANCE_REF;type:FLOAT;" json:"DISTANCE_REF" db:"DISTANCE_REF"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTOMER_LOCATIONS) TableName() string {
	return "CUSTOMER_LOCATIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTOMER_LOCATIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTOMER_LOCATIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTOMER_LOCATIONS) Validate(action Action) error {
	return nil
}
