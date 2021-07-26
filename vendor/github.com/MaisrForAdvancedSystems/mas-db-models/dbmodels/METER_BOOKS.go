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

type METER_BOOKS struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] BILLGROUP                                      NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 2] CODE                                           NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	CODE string `gorm:"primary_key;column:CODE;type:NVARCHAR;size:256;" json:"CODE" db:"CODE"`
	//[ 3] DESCRIBE                                       NVARCHAR(510)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	DESCRIBE string `gorm:"column:DESCRIBE;type:NVARCHAR;size:510;" json:"DESCRIBE" db:"DESCRIBE"`
	//[ 4] NO_WALKS                                       SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	NO_WALKS *int64 `gorm:"column:NO_WALKS;type:SMALLINT;" json:"NO_WALKS" db:"NO_WALKS"`
	//[ 5] HANDHELD_ID                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	HANDHELD_ID *string `gorm:"column:HANDHELD_ID;type:NVARCHAR;" json:"HANDHELD_ID" db:"HANDHELD_ID"`
	//[ 6] MREADER_ID                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	MREADER_ID *string `gorm:"column:MREADER_ID;type:NVARCHAR;" json:"MREADER_ID" db:"MREADER_ID"`
	//[ 7] ASSIGNED_TO_HH                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ASSIGNED_TO_HH *int64 `gorm:"column:ASSIGNED_TO_HH;type:INT;" json:"ASSIGNED_TO_HH" db:"ASSIGNED_TO_HH"`
	//[ 8] UNUSED                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	UNUSED *bool `gorm:"column:UNUSED;type:BIT;" json:"UNUSED" db:"UNUSED"`
	//[ 9] DISTANCE_REF                                   FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DISTANCE_REF *float64 `gorm:"column:DISTANCE_REF;type:FLOAT;" json:"DISTANCE_REF" db:"DISTANCE_REF"`
	//[10] LAT_MIN                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_MIN *float64 `gorm:"column:LAT_MIN;type:FLOAT;" json:"LAT_MIN" db:"LAT_MIN"`
	//[11] LAT_MAX                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_MAX *float64 `gorm:"column:LAT_MAX;type:FLOAT;" json:"LAT_MAX" db:"LAT_MAX"`
	//[12] LNG_MIN                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_MIN *float64 `gorm:"column:LNG_MIN;type:FLOAT;" json:"LNG_MIN" db:"LNG_MIN"`
	//[13] LNG_MAX                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_MAX *float64 `gorm:"column:LNG_MAX;type:FLOAT;" json:"LNG_MAX" db:"LNG_MAX"`
	//[14] APPLY_REF                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	APPLY_REF *int64 `gorm:"column:APPLY_REF;type:INT;" json:"APPLY_REF" db:"APPLY_REF"`
	//[15] REF_DATE                                       DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	REF_DATE *time.Time `gorm:"column:REF_DATE;type:DATE;" json:"REF_DATE" db:"REF_DATE"`
	//[16] SHAPE                                          NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	SHAPE *string `gorm:"column:SHAPE;type:NCHAR;size:60;" json:"SHAPE" db:"SHAPE"`
	//[17] PATH                                           NCHAR(2048)          null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 2048    default: []
	PATH *string `gorm:"column:PATH;type:NCHAR;size:2048;" json:"PATH" db:"PATH"`
}

// TableName sets the insert table name for this struct type
func (m *METER_BOOKS) TableName() string {
	return "METER_BOOKS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *METER_BOOKS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *METER_BOOKS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *METER_BOOKS) Validate(action Action) error {
	return nil
}
