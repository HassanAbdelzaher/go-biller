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

type PAYMENT_METHODS struct {
	//[ 0] TYPE_ID                                        INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	TYPE_ID int32 `gorm:"primary_key;column:TYPE_ID;type:INT;" json:"TYPE_ID" db:"TYPE_ID"`
	//[ 1] DESCRIPTION                                    NVARCHAR(500)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 500     default: []
	DESCRIPTION string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:500;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 2] BILING_CODE                                    NVARCHAR(500)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 500     default: []
	BILING_CODE *string `gorm:"column:BILING_CODE;type:NVARCHAR;size:500;" json:"BILING_CODE" db:"BILING_CODE"`
	//[ 3] SELECTABLE                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	SELECTABLE *bool `gorm:"column:SELECTABLE;type:BIT;" json:"SELECTABLE" db:"SELECTABLE"`
	//SELECTABLE_HH *bool `gorm:"column:SELECTABLE_HH;type:BIT;" json:"SELECTABLE_HH" db:"SELECTABLE_HH"`
	//[ 4] IS_SYSTEM                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_SYSTEM *bool `gorm:"column:IS_SYSTEM;type:BIT;" json:"IS_SYSTEM" db:"IS_SYSTEM"`
	//[ 5] RECEIPT_CHARGE_AMOUNT1                         FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	/*RECEIPT_CHARGE_AMOUNT1 *float64 `gorm:"column:RECEIPT_CHARGE_AMOUNT1;type:FLOAT;" json:"RECEIPT_CHARGE_AMOUNT1" db:"RECEIPT_CHARGE_AMOUNT1"`
	//[ 6] RECEIPT_CHARGE_PERCENTAGE1                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE_PERCENTAGE1 *float64 `gorm:"column:RECEIPT_CHARGE_PERCENTAGE1;type:FLOAT;" json:"RECEIPT_CHARGE_PERCENTAGE1" db:"RECEIPT_CHARGE_PERCENTAGE1"`
	//[ 7] RECEIPT_CHARGE_AMOUNT2                         FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE_AMOUNT2 *float64 `gorm:"column:RECEIPT_CHARGE_AMOUNT2;type:FLOAT;" json:"RECEIPT_CHARGE_AMOUNT2" db:"RECEIPT_CHARGE_AMOUNT2"`
	//[ 8] RECEIPT_CHARGE_PERCENTAGE2                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE_PERCENTAGE2 *float64 `gorm:"column:RECEIPT_CHARGE_PERCENTAGE2;type:FLOAT;" json:"RECEIPT_CHARGE_PERCENTAGE2" db:"RECEIPT_CHARGE_PERCENTAGE2"`
	//[ 9] RECEIPT_CHARGE_AMOUNT3                         FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE_AMOUNT3 *float64 `gorm:"column:RECEIPT_CHARGE_AMOUNT3;type:FLOAT;" json:"RECEIPT_CHARGE_AMOUNT3" db:"RECEIPT_CHARGE_AMOUNT3"`
	//[10] RECEIPT_CHARGE_PERCENTAGE3                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE_PERCENTAGE3 *float64 `gorm:"column:RECEIPT_CHARGE_PERCENTAGE3;type:FLOAT;" json:"RECEIPT_CHARGE_PERCENTAGE3" db:"RECEIPT_CHARGE_PERCENTAGE3"`
	//[11] RECEIPT_CHARGE_TITLE3                          NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	RECEIPT_CHARGE_TITLE3 *string `gorm:"column:RECEIPT_CHARGE_TITLE3;type:NVARCHAR;size:200;" json:"RECEIPT_CHARGE_TITLE3" db:"RECEIPT_CHARGE_TITLE3"`
	//[12] RECEIPT_CHARGE_TITLE2                          NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	RECEIPT_CHARGE_TITLE2 *string `gorm:"column:RECEIPT_CHARGE_TITLE2;type:NVARCHAR;size:200;" json:"RECEIPT_CHARGE_TITLE2" db:"RECEIPT_CHARGE_TITLE2"`
	//[13] RECEIPT_CHARGE_TITLE1                          NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	RECEIPT_CHARGE_TITLE1 *string `gorm:"column:RECEIPT_CHARGE_TITLE1;type:NVARCHAR;size:200;" json:"RECEIPT_CHARGE_TITLE1" db:"RECEIPT_CHARGE_TITLE1"`
	//[14] APPLY_MIN1                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MIN1 *float64 `gorm:"column:APPLY_MIN1;type:FLOAT;" json:"APPLY_MIN1" db:"APPLY_MIN1"`
	//[15] APPLY_MAX1                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MAX1 *float64 `gorm:"column:APPLY_MAX1;type:FLOAT;" json:"APPLY_MAX1" db:"APPLY_MAX1"`
	//[16] VALUE_MIN1                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MIN1 *float64 `gorm:"column:VALUE_MIN1;type:FLOAT;" json:"VALUE_MIN1" db:"VALUE_MIN1"`
	//[17] VALUE_MAX1                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MAX1 *float64 `gorm:"column:VALUE_MAX1;type:FLOAT;" json:"VALUE_MAX1" db:"VALUE_MAX1"`
	//[18] APPLY_MIN2                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MIN2 *float64 `gorm:"column:APPLY_MIN2;type:FLOAT;" json:"APPLY_MIN2" db:"APPLY_MIN2"`
	//[19] APPLY_MAX2                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MAX2 *float64 `gorm:"column:APPLY_MAX2;type:FLOAT;" json:"APPLY_MAX2" db:"APPLY_MAX2"`
	//[20] VALUE_MIN2                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MIN2 *float64 `gorm:"column:VALUE_MIN2;type:FLOAT;" json:"VALUE_MIN2" db:"VALUE_MIN2"`
	//[21] VALUE_MAX2                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MAX2 *float64 `gorm:"column:VALUE_MAX2;type:FLOAT;" json:"VALUE_MAX2" db:"VALUE_MAX2"`
	//[22] APPLY_MIN3                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MIN3 *float64 `gorm:"column:APPLY_MIN3;type:FLOAT;" json:"APPLY_MIN3" db:"APPLY_MIN3"`
	//[23] APPLY_MAX3                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	APPLY_MAX3 *float64 `gorm:"column:APPLY_MAX3;type:FLOAT;" json:"APPLY_MAX3" db:"APPLY_MAX3"`
	//[24] VALUE_MIN3                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MIN3 *float64 `gorm:"column:VALUE_MIN3;type:FLOAT;" json:"VALUE_MIN3" db:"VALUE_MIN3"`
	//[25] VALUE_MAX3                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	VALUE_MAX3 *float64 `gorm:"column:VALUE_MAX3;type:FLOAT;" json:"VALUE_MAX3" db:"VALUE_MAX3"`
*/
}

// TableName sets the insert table name for this struct type
func (p *PAYMENT_METHODS) TableName() string {
	return "PAYMENT_METHODS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PAYMENT_METHODS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PAYMENT_METHODS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PAYMENT_METHODS) Validate(action Action) error {
	return nil
}
