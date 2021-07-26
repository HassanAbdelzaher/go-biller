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

type CANCELLED_BILLS_ACTION struct {
	FORM_NO int64 `gorm:"primary_key;column:FORM_NO;type:bigint;" json:"FORM_NO" db:"FORM_NO"`
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:256;" json:"CUSTKEY" db:"CUSTKEY"`
	DOCUMENT_NO    string    `gorm:"primary_key;column:DOCUMENT_NO;type:nvarchar;" json:"DOCUMENT_NO" db:"DOCUMENT_NO"`
	ACTION_ID       int32   `gorm:"primary_key;column:ACTION_ID;type:int;" json:"ACTION_ID" db:"ACTION_ID"`
	STAMP_DATE *time.Time `gorm:"primary_key;column:STAMP_DATE;type:FLOAT;" json:"STAMP_DATE" db:"STAMP_DATE"`
	COMMENT        *string    `gorm:"column:COMMENT;type:nvarchar;" json:"COMMENT" db:"COMMENT"`
	STAMP_USER   *string    `gorm:"column:STAMP_USER;type:nvarchar;" json:"STAMP_USER" db:"STAMP_USER"`
	USER_ID        *int32     `gorm:"column:USER_ID;type:int;" json:"USER_ID" db:"USER_ID"`
}

// TableName sets the insert table name for this struct type
func (b *CANCELLED_BILLS_ACTION) TableName() string {
return "CANCELLED_BILLS_ACTIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *CANCELLED_BILLS_ACTION) BeforeSave() error {
return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *CANCELLED_BILLS_ACTION) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *CANCELLED_BILLS_ACTION) Validate(action Action) error {
return nil
}
