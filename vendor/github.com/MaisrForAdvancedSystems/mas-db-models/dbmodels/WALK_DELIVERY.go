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

type WALK_DELIVERY struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] STATION_NO                                     INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 2] BILLGROUP                                      NCHAR(20)            null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BILLGROUP string `gorm:"column:BILLGROUP;type:NCHAR;size:20;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 3] BOOK_NO                                        NCHAR(20)            null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BOOK_NO string `gorm:"column:BOOK_NO;type:NCHAR;size:20;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 4] WALK_NO                                        NCHAR(20)            null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	WALK_NO string `gorm:"column:WALK_NO;type:NCHAR;size:20;" json:"WALK_NO" db:"WALK_NO"`
	//[ 5] CYCLE_ID                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID *int64 `gorm:"column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 6] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 7] STAMP_DATE                                     DATE                 null: false  primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATE;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 8] STAMP_USER                                     NCHAR(60)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 60      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NCHAR;size:60;" json:"STAMP_USER" db:"STAMP_USER"`
	//[ 9] FROM_EMP                                       INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	FROM_EMP int32 `gorm:"column:FROM_EMP;type:INT;" json:"FROM_EMP" db:"FROM_EMP"`
	//[10] TO_EMP                                         INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	TO_EMP int32 `gorm:"column:TO_EMP;type:INT;" json:"TO_EMP" db:"TO_EMP"`
	//[11] FROM_NAME                                      NCHAR(200)           null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	FROM_NAME string `gorm:"column:FROM_NAME;type:NCHAR;size:200;" json:"FROM_NAME" db:"FROM_NAME"`
	//[12] TO_NAME                                        NCHAR(200)           null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	TO_NAME string `gorm:"column:TO_NAME;type:NCHAR;size:200;" json:"TO_NAME" db:"TO_NAME"`
	//[13] TYPE                                           INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	TYPE int32 `gorm:"column:TYPE;type:INT;" json:"TYPE" db:"TYPE"`
	//[14] NOTE                                           NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	NOTE *string `gorm:"column:NOTE;type:NCHAR;size:200;" json:"NOTE" db:"NOTE"`
}

// TableName sets the insert table name for this struct type
func (w *WALK_DELIVERY) TableName() string {
	return "WALK_DELIVERY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WALK_DELIVERY) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WALK_DELIVERY) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WALK_DELIVERY) Validate(action Action) error {
	return nil
}
