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

type MESSAGES struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] MESSAGE_TYPE                                   INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	MESSAGE_TYPE int32 `gorm:"column:MESSAGE_TYPE;type:INT;" json:"MESSAGE_TYPE" db:"MESSAGE_TYPE"`
	//[ 2] MESSAGE_CONTENET                               NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	MESSAGE_CONTENET *string `gorm:"column:MESSAGE_CONTENET;type:NVARCHAR;" json:"MESSAGE_CONTENET" db:"MESSAGE_CONTENET"`
	//[ 3] FROM                                           NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	FROM *string `gorm:"column:FROM;type:NVARCHAR;" json:"FROM" db:"FROM"`
	//[ 4] TO                                             NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	TO *string `gorm:"column:TO;type:NVARCHAR;" json:"TO" db:"TO"`
	//[ 5] SUBJECT                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	SUBJECT *string `gorm:"column:SUBJECT;type:NVARCHAR;" json:"SUBJECT" db:"SUBJECT"`
	//[ 6] SEND_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	SEND_DATE *time.Time `gorm:"column:SEND_DATE;type:DATETIME;" json:"SEND_DATE" db:"SEND_DATE"`
	//[ 7] RECIVE_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	RECIVE_DATE *time.Time `gorm:"column:RECIVE_DATE;type:DATETIME;" json:"RECIVE_DATE" db:"RECIVE_DATE"`
	//[ 8] EXPIRE_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	EXPIRE_DATE *time.Time `gorm:"column:EXPIRE_DATE;type:DATETIME;" json:"EXPIRE_DATE" db:"EXPIRE_DATE"`
	//[ 9] IS_RECIVED                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_RECIVED *bool `gorm:"column:IS_RECIVED;type:BIT;" json:"IS_RECIVED" db:"IS_RECIVED"`
	//[10] IS_EXPIRED                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_EXPIRED *bool `gorm:"column:IS_EXPIRED;type:BIT;" json:"IS_EXPIRED" db:"IS_EXPIRED"`
}

// TableName sets the insert table name for this struct type
func (m *MESSAGES) TableName() string {
	return "MESSAGES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MESSAGES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MESSAGES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MESSAGES) Validate(action Action) error {
	return nil
}
