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

type PATCHES struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: true   col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] DEVICE_ID                                      NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NCHAR;size:60;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 2] BILLGROUP                                      NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NCHAR;size:20;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 3] BOOK_NO                                        NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NCHAR;size:20;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 4] WALK_NO                                        NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NCHAR;size:20;" json:"WALK_NO" db:"WALK_NO"`
	//[ 5] STAM_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAM_DATE *time.Time `gorm:"column:STAM_DATE;type:DATETIME;" json:"STAM_DATE" db:"STAM_DATE"`
	//[ 6] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[ 7] STATION_NO                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int64 `gorm:"column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 8] CLIENT_VERSION                                 NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	CLIENT_VERSION *string `gorm:"column:CLIENT_VERSION;type:NCHAR;size:20;" json:"CLIENT_VERSION" db:"CLIENT_VERSION"`
}

// TableName sets the insert table name for this struct type
func (p *PATCHES) TableName() string {
	return "PATCHES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PATCHES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PATCHES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PATCHES) Validate(action Action) error {
	return nil
}
