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

type CUSTOMER_WALKS struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] BILLGROUP                                      NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 2] BOOK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NVARCHAR;size:60;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 3] WALK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	WALK_NO string `gorm:"primary_key;column:WALK_NO;type:NVARCHAR;size:60;" json:"WALK_NO" db:"WALK_NO"`
	//[ 4] DESCRIBE                                       NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	DESCRIBE *string `gorm:"column:DESCRIBE;type:NVARCHAR;size:510;" json:"DESCRIBE" db:"DESCRIBE"`
	//[ 5] ASSIGNED_TO_HH                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ASSIGNED_TO_HH *int64 `gorm:"column:ASSIGNED_TO_HH;type:INT;" json:"ASSIGNED_TO_HH" db:"ASSIGNED_TO_HH"`
	//[ 6] UNUSED                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	UNUSED *bool `gorm:"column:UNUSED;type:BIT;" json:"UNUSED" db:"UNUSED"`
	//[ 7] MARKETING                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	MARKETING *int64 `gorm:"column:MARKETING;type:INT;" json:"MARKETING" db:"MARKETING"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTOMER_WALKS) TableName() string {
	return "CUSTOMER_WALKS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTOMER_WALKS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTOMER_WALKS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTOMER_WALKS) Validate(action Action) error {
	return nil
}
