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

type DEVICES struct {
	//[ 0] DEVICE_ID                                      NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	DEVICE_ID string `gorm:"primary_key;column:DEVICE_ID;type:NVARCHAR;size:256;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 1] ASSIGN_TO                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ASSIGN_TO *int64 `gorm:"column:ASSIGN_TO;type:INT;" json:"ASSIGN_TO" db:"ASSIGN_TO"`
	//[ 2] NOTE                                           NVARCHAR(1024)       null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 1024    default: []
	NOTE *string `gorm:"column:NOTE;type:NVARCHAR;size:1024;" json:"NOTE" db:"NOTE"`
	//[ 3] ID                                             INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ID *int32 `gorm:"column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 4] OP_STATUS                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OP_STATUS *int64 `gorm:"column:OP_STATUS;type:INT;" json:"OP_STATUS" db:"OP_STATUS"`
	//[ 5] STATION_ID                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_ID *int64 `gorm:"column:STATION_ID;type:INT;" json:"STATION_ID" db:"STATION_ID"`
	//[ 6] SIM_NO                                         NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	SIM_NO *string `gorm:"column:SIM_NO;type:NVARCHAR;size:60;" json:"SIM_NO" db:"SIM_NO"`
	//[ 7] STATUS                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATUS *int64 `gorm:"column:STATUS;type:INT;" json:"STATUS" db:"STATUS"`
	//[ 8] COMPANY                                        NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	COMPANY *string `gorm:"column:COMPANY;type:NVARCHAR;size:200;" json:"COMPANY" db:"COMPANY"`
	//[ 9] MODEL                                          NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	MODEL *string `gorm:"column:MODEL;type:NVARCHAR;size:200;" json:"MODEL" db:"MODEL"`
	//[10] OS                                             INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OS *int64 `gorm:"column:OS;type:INT;" json:"OS" db:"OS"`
	//[11] OS_VERSION                                     NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	OS_VERSION *string `gorm:"column:OS_VERSION;type:NCHAR;size:20;" json:"OS_VERSION" db:"OS_VERSION"`
	//[12] BATTERY_NO                                     NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	BATTERY_NO *string `gorm:"column:BATTERY_NO;type:NCHAR;size:100;" json:"BATTERY_NO" db:"BATTERY_NO"`
	//[13] PURCHASE_DATE                                  DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	PURCHASE_DATE *time.Time `gorm:"column:PURCHASE_DATE;type:DATE;" json:"PURCHASE_DATE" db:"PURCHASE_DATE"`
	//[14] WARRENTY                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WARRENTY *int64 `gorm:"column:WARRENTY;type:INT;" json:"WARRENTY" db:"WARRENTY"`
	//[15] INTERNET_START                                 DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	INTERNET_START *time.Time `gorm:"column:INTERNET_START;type:DATE;" json:"INTERNET_START" db:"INTERNET_START"`
	//[16] INTERNET_RENEW                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	INTERNET_RENEW *int64 `gorm:"column:INTERNET_RENEW;type:INT;" json:"INTERNET_RENEW" db:"INTERNET_RENEW"`
	//[17] CRADLE                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CRADLE *string `gorm:"column:CRADLE;type:NVARCHAR;size:100;" json:"CRADLE" db:"CRADLE"`
}

// TableName sets the insert table name for this struct type
func (d *DEVICES) TableName() string {
	return "DEVICES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEVICES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEVICES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEVICES) Validate(action Action) error {
	return nil
}
