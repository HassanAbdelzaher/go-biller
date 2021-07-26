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

type CUSTOMER_SEQUENCE struct {
	//[ 0] CUSTKEY                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 1] BOOK_NO_R                                      NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	BOOK_NO_R *string `gorm:"column:BOOK_NO_R;type:NCHAR;size:60;" json:"BOOK_NO_R" db:"BOOK_NO_R"`
	//[ 2] WALK_NO_R                                      NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	WALK_NO_R *string `gorm:"column:WALK_NO_R;type:NCHAR;size:20;" json:"WALK_NO_R" db:"WALK_NO_R"`
	//[ 3] SEQ_NO_R                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQ_NO_R *int64 `gorm:"column:SEQ_NO_R;type:INT;" json:"SEQ_NO_R" db:"SEQ_NO_R"`
	//[ 4] BOOK_NO_C                                      NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	BOOK_NO_C *string `gorm:"column:BOOK_NO_C;type:NCHAR;size:60;" json:"BOOK_NO_C" db:"BOOK_NO_C"`
	//[ 5] WALK_NO_C                                      NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	WALK_NO_C *string `gorm:"column:WALK_NO_C;type:NCHAR;size:20;" json:"WALK_NO_C" db:"WALK_NO_C"`
	//[ 6] SEQ_NO_C                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQ_NO_C *int64 `gorm:"column:SEQ_NO_C;type:INT;" json:"SEQ_NO_C" db:"SEQ_NO_C"`
	//[ 7] UPDATEDATE_R                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	UPDATEDATE_R *time.Time `gorm:"column:UPDATEDATE_R;type:DATETIME;" json:"UPDATEDATE_R" db:"UPDATEDATE_R"`
	//[ 8] UPDATEDATE_C                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	UPDATEDATE_C *time.Time `gorm:"column:UPDATEDATE_C;type:DATETIME;" json:"UPDATEDATE_C" db:"UPDATEDATE_C"`
	//[ 9] IS_POSTED                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_POSTED *bool `gorm:"column:IS_POSTED;type:BIT;" json:"IS_POSTED" db:"IS_POSTED"`
	//[10] CYCLE_ID                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID *int64 `gorm:"column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[11] POST_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	POST_DATE *time.Time `gorm:"column:POST_DATE;type:DATETIME;" json:"POST_DATE" db:"POST_DATE"`
	//[12] PROP_REF                                       NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	PROP_REF *string `gorm:"column:PROP_REF;type:NVARCHAR;size:100;" json:"PROP_REF" db:"PROP_REF"`
	//[13] IS_POSTED_R                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_POSTED_R *bool `gorm:"column:IS_POSTED_R;type:BIT;" json:"IS_POSTED_R" db:"IS_POSTED_R"`
	//[14] IS_POSTED_C                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_POSTED_C *bool `gorm:"column:IS_POSTED_C;type:BIT;" json:"IS_POSTED_C" db:"IS_POSTED_C"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTOMER_SEQUENCE) TableName() string {
	return "CUSTOMER_SEQUENCE"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTOMER_SEQUENCE) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTOMER_SEQUENCE) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTOMER_SEQUENCE) Validate(action Action) error {
	return nil
}
