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

type CANCELLED_REQUEST struct {
	FORM_NO      int64      `gorm:"primary_key;column:FORM_NO;type:bigint;" json:"FORM_NO" db:"FORM_NO"`
	CUSTKEY      string     `gorm:"column:CUSTKEY;type:NVARCHAR;size:256;" json:"CUSTKEY" db:"CUSTKEY"`
	STATION_NO   string     `gorm:"column:STATION_NO;type:nvarchar;" json:"STATION_NO" db:"STATION_NO"`
	DOCUMENT_NO  string     `gorm:"column:DOCUMENT_NO;type:nvarchar;" json:"DOCUMENT_NO" db:"DOCUMENT_NO"`
	REQUEST_DATE *time.Time `gorm:"column:REQUEST_DATE;type:datetime;" json:"REQUEST_DATE" db:"REQUEST_DATE"`
	REQUEST_BY   *string    `gorm:"column:REQUEST_BY;type:nvarchar;" json:"REQUEST_BY" db:"REQUEST_BY"`
	STATE        *int32     `gorm:"column:STATE;type:int;" json:"STATE" db:"STATE"`
	CLOSED       *bool      `gorm:"column:CLOSED;type:bit;" json:"CLOSED" db:"CLOSED"`
	STATUS       *string    `gorm:"column:STATUS;type:nvarchar;" json:"STATUS" db:"STATUS"`
	COMMENT      *string    `gorm:"column:COMMENT;type:nvarchar;" json:"COMMENT" db:"COMMENT"`
	COUNTER      *int32     `gorm:"column:COUNTER;type:int;" json:"COUNTER" db:"COUNTER"`
	SURNAME      *string    `gorm:"column:SURNAME;type:nvarchar;" json:"SURNAME" db:"SURNAME"`
	STAMP_DATE   *time.Time `gorm:"column:STAMP_DATE;type:datetime;" json:"STAMP_DATE" db:"STAMP_DATE"`
}

// TableName sets the insert table name for this struct type
func (b *CANCELLED_REQUEST) TableName() string {
	return "CANCELLED_REQUESTS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *CANCELLED_REQUEST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *CANCELLED_REQUEST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *CANCELLED_REQUEST) Validate(action Action) error {
	return nil
}
