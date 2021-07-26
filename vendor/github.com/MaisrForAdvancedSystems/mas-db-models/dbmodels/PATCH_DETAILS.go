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

type PATCH_DETAILS struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] CUSTKEY                                        NCHAR(60)            null: false  primary: true   isArray: false  auto: false  col: NCHAR           len: 60      default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 2] DELIVERY_ST                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DELIVERY_ST *int64 `gorm:"column:DELIVERY_ST;type:INT;" json:"DELIVERY_ST" db:"DELIVERY_ST"`
	//[ 3] CR_READING                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CR_READING *int64 `gorm:"column:CR_READING;type:INT;" json:"CR_READING" db:"CR_READING"`
	//[ 4] MHH_NOTE                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	MHH_NOTE *int64 `gorm:"column:MHH_NOTE;type:INT;" json:"MHH_NOTE" db:"MHH_NOTE"`
	//[ 5] STM_NOTE                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STM_NOTE *int64 `gorm:"column:STM_NOTE;type:INT;" json:"STM_NOTE" db:"STM_NOTE"`
	//[ 6] COLLECTION_DATE                                DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	COLLECTION_DATE *time.Time `gorm:"column:COLLECTION_DATE;type:DATETIME;" json:"COLLECTION_DATE" db:"COLLECTION_DATE"`
	//[ 7] READING_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	READING_DATE *time.Time `gorm:"column:READING_DATE;type:DATETIME;" json:"READING_DATE" db:"READING_DATE"`
	//[ 8] NO_UNITS                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NO_UNITS *int64 `gorm:"column:NO_UNITS;type:INT;" json:"NO_UNITS" db:"NO_UNITS"`
	//[ 9] C_TYPE                                         NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	C_TYPE *string `gorm:"column:C_TYPE;type:NCHAR;size:20;" json:"C_TYPE" db:"C_TYPE"`
	//[10] SEWER                                          BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	SEWER *bool `gorm:"column:SEWER;type:BIT;" json:"SEWER" db:"SEWER"`
	//[11] IS_REF                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_REF *bool `gorm:"column:IS_REF;type:BIT;" json:"IS_REF" db:"IS_REF"`
	//[12] LAT_C                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_C *float64 `gorm:"column:LAT_C;type:FLOAT;" json:"LAT_C" db:"LAT_C"`
	//[13] LAT_R                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_R *float64 `gorm:"column:LAT_R;type:FLOAT;" json:"LAT_R" db:"LAT_R"`
	//[14] LNG_C                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_C *float64 `gorm:"column:LNG_C;type:FLOAT;" json:"LNG_C" db:"LNG_C"`
	//[15] LNG_R                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_R *float64 `gorm:"column:LNG_R;type:FLOAT;" json:"LNG_R" db:"LNG_R"`
	//[16] GPS_SOURCE_C                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	GPS_SOURCE_C *int64 `gorm:"column:GPS_SOURCE_C;type:INT;" json:"GPS_SOURCE_C" db:"GPS_SOURCE_C"`
	//[17] GPS_SOURCE_R                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	GPS_SOURCE_R *int64 `gorm:"column:GPS_SOURCE_R;type:INT;" json:"GPS_SOURCE_R" db:"GPS_SOURCE_R"`
	//[18] DISTANCE_R                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DISTANCE_R *float64 `gorm:"column:DISTANCE_R;type:FLOAT;" json:"DISTANCE_R" db:"DISTANCE_R"`
	//[19] SYNC_ST                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SYNC_ST *int64 `gorm:"column:SYNC_ST;type:INT;" json:"SYNC_ST" db:"SYNC_ST"`
	//[20] ACCURECY_R                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY_R *float64 `gorm:"column:ACCURECY_R;type:FLOAT;" json:"ACCURECY_R" db:"ACCURECY_R"`
	//[21] ACCURECY_C                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY_C *float64 `gorm:"column:ACCURECY_C;type:FLOAT;" json:"ACCURECY_C" db:"ACCURECY_C"`
}

// TableName sets the insert table name for this struct type
func (p *PATCH_DETAILS) TableName() string {
	return "PATCH_DETAILS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PATCH_DETAILS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PATCH_DETAILS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PATCH_DETAILS) Validate(action Action) error {
	return nil
}
