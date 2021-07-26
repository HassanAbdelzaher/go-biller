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

type UserPosts struct {
	//[ 0] id                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	Id int32 `gorm:"primary_key;column:id;type:INT;" json:"id" db:"id"`
	//[ 1] userId                                         INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	UserId int32 `gorm:"column:userId;type:INT;" json:"userId" db:"userId"`
	//[ 2] descreption                                    NVARCHAR(400)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 400     default: []
	Descreption *string `gorm:"column:descreption;type:NVARCHAR;size:400;" json:"descreption" db:"descreption"`
}

// TableName sets the insert table name for this struct type
func (u *UserPosts) TableName() string {
	return "UserPosts"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *UserPosts) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *UserPosts) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *UserPosts) Validate(action Action) error {
	return nil
}
