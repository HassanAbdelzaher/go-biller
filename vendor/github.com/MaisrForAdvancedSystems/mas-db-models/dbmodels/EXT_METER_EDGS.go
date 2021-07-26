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

type EXT_METER_EDGS struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] NAME                                           NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	NAME *string `gorm:"column:NAME;type:NCHAR;size:200;" json:"NAME" db:"NAME"`
	//[ 2] ADDRESS                                        NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	ADDRESS *string `gorm:"column:ADDRESS;type:NCHAR;size:200;" json:"ADDRESS" db:"ADDRESS"`
	//[ 3] CUSTKEY                                        NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	CUSTKEY *string `gorm:"column:CUSTKEY;type:NCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 4] STAMP_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATE;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 5] REMOTE_ADDRESS                                 NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	REMOTE_ADDRESS *string `gorm:"column:REMOTE_ADDRESS;type:NCHAR;size:60;" json:"REMOTE_ADDRESS" db:"REMOTE_ADDRESS"`
	//[ 6] TEL                                            NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	TEL *string `gorm:"column:TEL;type:NCHAR;size:60;" json:"TEL" db:"TEL"`
	//[ 7] CR_READING                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CR_READING *int64 `gorm:"column:CR_READING;type:INT;" json:"CR_READING" db:"CR_READING"`
	//[ 8] CR_DATE                                        DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CR_DATE *time.Time `gorm:"column:CR_DATE;type:DATE;" json:"CR_DATE" db:"CR_DATE"`
	//[ 9] EMAIL                                          NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	EMAIL *string `gorm:"column:EMAIL;type:NCHAR;size:200;" json:"EMAIL" db:"EMAIL"`
	//[10] METER_STATUS                                   NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	METER_STATUS *string `gorm:"column:METER_STATUS;type:NCHAR;size:60;" json:"METER_STATUS" db:"METER_STATUS"`
}

// TableName sets the insert table name for this struct type
func (e *EXT_METER_EDGS) TableName() string {
	return "EXT_METER_EDGS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EXT_METER_EDGS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EXT_METER_EDGS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *EXT_METER_EDGS) Validate(action Action) error {
	return nil
}
