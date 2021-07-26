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

type CUSTOMER_NOTES struct {
	//[ 0] APP_ID                                         INT                  null: false  primary: true   isArray: false  auto: true   col: INT             len: -1      default: []
	APP_ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:APP_ID;type:INT;" json:"APP_ID" db:"APP_ID"`
	//[ 1] APP_TYPE                                       INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	APP_TYPE int32 `gorm:"column:APP_TYPE;type:INT;" json:"APP_TYPE" db:"APP_TYPE"`
	//[ 2] APP_DATE                                       DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	APP_DATE *time.Time `gorm:"column:APP_DATE;type:DATETIME;" json:"APP_DATE" db:"APP_DATE"`
	//[ 3] BOOK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 4] WALK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NVARCHAR;" json:"WALK_NO" db:"WALK_NO"`
	//[ 5] SEQ_NO                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQ_NO *int64 `gorm:"column:SEQ_NO;type:INT;" json:"SEQ_NO" db:"SEQ_NO"`
	//[ 6] CUSTKEY                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CUSTKEY *string `gorm:"column:CUSTKEY;type:NVARCHAR;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 7] SURNAME                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	SURNAME *string `gorm:"column:SURNAME;type:NVARCHAR;" json:"SURNAME" db:"SURNAME"`
	//[ 8] ADDRESS                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ADDRESS *string `gorm:"column:ADDRESS;type:NVARCHAR;" json:"ADDRESS" db:"ADDRESS"`
	//[ 9] LAT                                            NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	LAT *string `gorm:"column:LAT;type:NVARCHAR;" json:"LAT" db:"LAT"`
	//[10] LNG                                            NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	LNG *string `gorm:"column:LNG;type:NVARCHAR;" json:"LNG" db:"LNG"`
	//[11] OLD_CTYPE                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	OLD_CTYPE *string `gorm:"column:OLD_CTYPE;type:NVARCHAR;" json:"OLD_CTYPE" db:"OLD_CTYPE"`
	//[12] NEW_CTYPE                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	NEW_CTYPE *string `gorm:"column:NEW_CTYPE;type:NVARCHAR;" json:"NEW_CTYPE" db:"NEW_CTYPE"`
	//[13] ACTUAL_CTYPE                                   NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTUAL_CTYPE *string `gorm:"column:ACTUAL_CTYPE;type:NVARCHAR;" json:"ACTUAL_CTYPE" db:"ACTUAL_CTYPE"`
	//[14] OLD_UNITS                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OLD_UNITS *int64 `gorm:"column:OLD_UNITS;type:INT;" json:"OLD_UNITS" db:"OLD_UNITS"`
	//[15] NEW_UNITS                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NEW_UNITS *int64 `gorm:"column:NEW_UNITS;type:INT;" json:"NEW_UNITS" db:"NEW_UNITS"`
	//[16] ACTUAL_UNITS                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTUAL_UNITS *int64 `gorm:"column:ACTUAL_UNITS;type:INT;" json:"ACTUAL_UNITS" db:"ACTUAL_UNITS"`
	//[17] OLD_SEWER                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OLD_SEWER *int64 `gorm:"column:OLD_SEWER;type:INT;" json:"OLD_SEWER" db:"OLD_SEWER"`
	//[18] NEW_SEWER                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NEW_SEWER *int64 `gorm:"column:NEW_SEWER;type:INT;" json:"NEW_SEWER" db:"NEW_SEWER"`
	//[19] ACTUAL_SEWER                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTUAL_SEWER *int64 `gorm:"column:ACTUAL_SEWER;type:INT;" json:"ACTUAL_SEWER" db:"ACTUAL_SEWER"`
	//[20] STATUS                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATUS *int64 `gorm:"column:STATUS;type:INT;" json:"STATUS" db:"STATUS"`
	//[21] PREVIEW_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	PREVIEW_DATE *time.Time `gorm:"column:PREVIEW_DATE;type:DATETIME;" json:"PREVIEW_DATE" db:"PREVIEW_DATE"`
	//[22] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY *float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTOMER_NOTES) TableName() string {
	return "CUSTOMER_NOTES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTOMER_NOTES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTOMER_NOTES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTOMER_NOTES) Validate(action Action) error {
	return nil
}
