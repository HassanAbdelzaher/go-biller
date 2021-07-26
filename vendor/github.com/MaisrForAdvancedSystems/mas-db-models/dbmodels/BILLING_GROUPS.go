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

type BILLING_GROUPS struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] GROUP_ID                                       NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	GROUP_ID string `gorm:"primary_key;column:GROUP_ID;type:NVARCHAR;size:256;" json:"GROUP_ID" db:"GROUP_ID"`
	//[ 2] DESCRIPTION                                    NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:512;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 3] UNUSED                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	UNUSED *bool `gorm:"column:UNUSED;type:BIT;" json:"UNUSED" db:"UNUSED"`
}

// TableName sets the insert table name for this struct type
func (b *BILLING_GROUPS) TableName() string {
	return "BILLING_GROUPS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BILLING_GROUPS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BILLING_GROUPS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BILLING_GROUPS) Validate(action Action) error {
	return nil
}
