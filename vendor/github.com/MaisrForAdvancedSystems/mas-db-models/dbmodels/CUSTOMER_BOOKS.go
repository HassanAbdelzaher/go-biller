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

type CUSTOMER_BOOKS struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] BILLGROUP                                      NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 2] CODE                                           NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	CODE string `gorm:"primary_key;column:CODE;type:NVARCHAR;size:60;" json:"CODE" db:"CODE"`
	//[ 3] DESCRIBE                                       NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	DESCRIBE *string `gorm:"column:DESCRIBE;type:NVARCHAR;size:510;" json:"DESCRIBE" db:"DESCRIBE"`
	//[ 4] NO_WALKS                                       SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	NO_WALKS *int64 `gorm:"column:NO_WALKS;type:SMALLINT;" json:"NO_WALKS" db:"NO_WALKS"`
	//[ 5] ASSIGNED_TO                                    NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	ASSIGNED_TO *string `gorm:"column:ASSIGNED_TO;type:NVARCHAR;size:510;" json:"ASSIGNED_TO" db:"ASSIGNED_TO"`
	//[ 6] ASSIGNED_TO_HH                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ASSIGNED_TO_HH *int64 `gorm:"column:ASSIGNED_TO_HH;type:INT;" json:"ASSIGNED_TO_HH" db:"ASSIGNED_TO_HH"`
	//[ 7] HANDHELD_ID                                    NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	HANDHELD_ID *string `gorm:"column:HANDHELD_ID;type:NVARCHAR;size:510;" json:"HANDHELD_ID" db:"HANDHELD_ID"`
	//[ 8] UNUSED                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	UNUSED *bool `gorm:"column:UNUSED;type:BIT;" json:"UNUSED" db:"UNUSED"`
}

// TableName sets the insert table name for this struct type
func (c *CUSTOMER_BOOKS) TableName() string {
	return "CUSTOMER_BOOKS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CUSTOMER_BOOKS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CUSTOMER_BOOKS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CUSTOMER_BOOKS) Validate(action Action) error {
	return nil
}
