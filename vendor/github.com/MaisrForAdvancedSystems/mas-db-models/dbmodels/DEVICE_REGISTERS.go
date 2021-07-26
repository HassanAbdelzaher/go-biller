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

type DEVICE_REGISTERS struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: true   col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DEVICE_ID                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 2] ACTION_ID                                      INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ACTION_ID int32 `gorm:"column:ACTION_ID;type:INT;" json:"ACTION_ID" db:"ACTION_ID"`
	//[ 3] ACTION_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	ACTION_DATE *time.Time `gorm:"column:ACTION_DATE;type:DATETIME;" json:"ACTION_DATE" db:"ACTION_DATE"`
	//[ 4] ACTION_PERIOD                                  NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTION_PERIOD *string `gorm:"column:ACTION_PERIOD;type:NVARCHAR;" json:"ACTION_PERIOD" db:"ACTION_PERIOD"`
	//[ 5] ACTION_BY                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTION_BY *string `gorm:"column:ACTION_BY;type:NVARCHAR;" json:"ACTION_BY" db:"ACTION_BY"`
	//[ 6] ACTION_RECOMMENDATION                          NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTION_RECOMMENDATION *string `gorm:"column:ACTION_RECOMMENDATION;type:NVARCHAR;" json:"ACTION_RECOMMENDATION" db:"ACTION_RECOMMENDATION"`
	//[ 7] ACTIION_FOLLOWUP                               NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTIION_FOLLOWUP *string `gorm:"column:ACTIION_FOLLOWUP;type:NVARCHAR;" json:"ACTIION_FOLLOWUP" db:"ACTIION_FOLLOWUP"`
	//[ 8] ACTION_COST                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ACTION_COST *string `gorm:"column:ACTION_COST;type:NVARCHAR;" json:"ACTION_COST" db:"ACTION_COST"`
	//[ 9] DESCRIPTION                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;" json:"DESCRIPTION" db:"DESCRIPTION"`
}

// TableName sets the insert table name for this struct type
func (d *DEVICE_REGISTERS) TableName() string {
	return "DEVICE_REGISTERS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEVICE_REGISTERS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEVICE_REGISTERS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEVICE_REGISTERS) Validate(action Action) error {
	return nil
}
