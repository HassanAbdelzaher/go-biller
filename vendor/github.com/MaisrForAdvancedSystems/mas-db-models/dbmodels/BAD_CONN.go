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

type BAD_CONN struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] BOOK_NO                                        NVARCHAR(20)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 20      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;size:20;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 2] WALK_NO                                        NVARCHAR(20)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 20      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NVARCHAR;size:20;" json:"WALK_NO" db:"WALK_NO"`
	//[ 3] CONN_TYPE                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CONN_TYPE *int64 `gorm:"column:CONN_TYPE;type:INT;" json:"CONN_TYPE" db:"CONN_TYPE"`
	//[ 4] NOTES                                          NVARCHAR(400)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 400     default: []
	NOTES *string `gorm:"column:NOTES;type:NVARCHAR;size:400;" json:"NOTES" db:"NOTES"`
	//[ 5] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT *float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[ 6] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG *float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[ 7] SYNC_ST                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SYNC_ST *int64 `gorm:"column:SYNC_ST;type:INT;" json:"SYNC_ST" db:"SYNC_ST"`
	//[ 8] SURNAME                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	SURNAME *string `gorm:"column:SURNAME;type:NVARCHAR;" json:"SURNAME" db:"SURNAME"`
	//[ 9] NO_FLOORS                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NO_FLOORS *int64 `gorm:"column:NO_FLOORS;type:INT;" json:"NO_FLOORS" db:"NO_FLOORS"`
	//[10] DEVICE_ID                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[11] DEVICE_CONNID                                  BIGINT               null: false  primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	DEVICE_CONNID int64 `gorm:"column:DEVICE_CONNID;type:BIGINT;" json:"DEVICE_CONNID" db:"DEVICE_CONNID"`
	//[12] STAMP_USER                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;" json:"STAMP_USER" db:"STAMP_USER"`
	//[13] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[14] STATUS                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATUS *int64 `gorm:"column:STATUS;type:INT;" json:"STATUS" db:"STATUS"`
	//[15] LOCATION_SOURCE                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	LOCATION_SOURCE *int64 `gorm:"column:LOCATION_SOURCE;type:INT;" json:"LOCATION_SOURCE" db:"LOCATION_SOURCE"`
	//[16] LOCATION_DATE                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	LOCATION_DATE *time.Time `gorm:"column:LOCATION_DATE;type:DATETIME;" json:"LOCATION_DATE" db:"LOCATION_DATE"`
	//[17] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY *float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
}

// TableName sets the insert table name for this struct type
func (b *BAD_CONN) TableName() string {
	return "BAD_CONN"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BAD_CONN) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BAD_CONN) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BAD_CONN) Validate(action Action) error {
	return nil
}
