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

type CUSTROUTES_HISTORY struct {
	//[ 0] BOOK_NO                                        NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NVARCHAR;size:256;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 1] WALK_NO                                        SMALLINT             null: false  primary: true   isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	WALK_NO int32 `gorm:"primary_key;column:WALK_NO;type:SMALLINT;" json:"WALK_NO" db:"WALK_NO"`
	//[ 2] HHUNIT_ID                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	HHUNIT_ID *string `gorm:"column:HHUNIT_ID;type:NVARCHAR;" json:"HHUNIT_ID" db:"HHUNIT_ID"`
	//[ 3] CYCLE_NO                                       SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	CYCLE_NO *int64 `gorm:"column:CYCLE_NO;type:SMALLINT;" json:"CYCLE_NO" db:"CYCLE_NO"`
	//[ 4] PROCESS_ST                                     SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	PROCESS_ST *int64 `gorm:"column:PROCESS_ST;type:SMALLINT;" json:"PROCESS_ST" db:"PROCESS_ST"`
	//[ 5] HH_DATE                                        DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HH_DATE *time.Time `gorm:"column:HH_DATE;type:DATETIME;" json:"HH_DATE" db:"HH_DATE"`
	//[ 6] POST_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	POST_DATE *time.Time `gorm:"column:POST_DATE;type:DATETIME;" json:"POST_DATE" db:"POST_DATE"`
	//[ 7] DELIVERY_MAN                                   NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DELIVERY_MAN *string `gorm:"column:DELIVERY_MAN;type:NVARCHAR;" json:"DELIVERY_MAN" db:"DELIVERY_MAN"`
	//[ 8] DELIVERY_DATE                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	DELIVERY_DATE *time.Time `gorm:"column:DELIVERY_DATE;type:DATETIME;" json:"DELIVERY_DATE" db:"DELIVERY_DATE"`
	//[ 9] STAMP_USER                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;" json:"STAMP_USER" db:"STAMP_USER"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTROUTES_HISTORY) TableName() string {
	return "CUSTROUTES_HISTORY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTROUTES_HISTORY) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTROUTES_HISTORY) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTROUTES_HISTORY) Validate(action Action) error {
	return nil
}
