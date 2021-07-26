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

type DEPOSIT_DETAILS struct {
	//[ 0] DEPOSIT_ID                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	DEPOSIT_ID int32 `gorm:"primary_key;column:DEPOSIT_ID;type:INT;" json:"DEPOSIT_ID" db:"DEPOSIT_ID"`
	//[ 1] BILNG_DEPOSIT_ID                               INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	BILNG_DEPOSIT_ID int32 `gorm:"primary_key;column:BILNG_DEPOSIT_ID;type:INT;" json:"BILNG_DEPOSIT_ID" db:"BILNG_DEPOSIT_ID"`
	//[ 2] BILLGROUP                                      NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NCHAR;size:20;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 3] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 4] AMOUNT                                         FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	AMOUNT *float64 `gorm:"column:AMOUNT;type:FLOAT;" json:"AMOUNT" db:"AMOUNT"`
	//[ 5] COUNT                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	COUNT *float64 `gorm:"column:COUNT;type:FLOAT;" json:"COUNT" db:"COUNT"`
	//[ 6] OP_BLNCE                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OP_BLNCE *float64 `gorm:"column:OP_BLNCE;type:FLOAT;" json:"OP_BLNCE" db:"OP_BLNCE"`
	//[ 7] WATER_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	WATER_AMT *float64 `gorm:"column:WATER_AMT;type:FLOAT;" json:"WATER_AMT" db:"WATER_AMT"`
	//[ 8] SEWER_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	SEWER_AMT *float64 `gorm:"column:SEWER_AMT;type:FLOAT;" json:"SEWER_AMT" db:"SEWER_AMT"`
	//[ 9] BASIC_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	BASIC_AMT *float64 `gorm:"column:BASIC_AMT;type:FLOAT;" json:"BASIC_AMT" db:"BASIC_AMT"`
	//[10] TAX_AMT                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	TAX_AMT *float64 `gorm:"column:TAX_AMT;type:FLOAT;" json:"TAX_AMT" db:"TAX_AMT"`
	//[11] INSTALLS_AMT                                   FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	INSTALLS_AMT *float64 `gorm:"column:INSTALLS_AMT;type:FLOAT;" json:"INSTALLS_AMT" db:"INSTALLS_AMT"`
	//[12] DBT_AMT                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DBT_AMT *float64 `gorm:"column:DBT_AMT;type:FLOAT;" json:"DBT_AMT" db:"DBT_AMT"`
	//[13] CRDT_AMT                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CRDT_AMT *float64 `gorm:"column:CRDT_AMT;type:FLOAT;" json:"CRDT_AMT" db:"CRDT_AMT"`
	//[14] AGREEM_AMT                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	AGREEM_AMT *float64 `gorm:"column:AGREEM_AMT;type:FLOAT;" json:"AGREEM_AMT" db:"AGREEM_AMT"`
	//[15] OTHER_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT *float64 `gorm:"column:OTHER_AMT;type:FLOAT;" json:"OTHER_AMT" db:"OTHER_AMT"`
	//[16] OTHER_AMT1                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT1 *float64 `gorm:"column:OTHER_AMT1;type:FLOAT;" json:"OTHER_AMT1" db:"OTHER_AMT1"`
	//[17] OTHER_AMT2                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT2 *float64 `gorm:"column:OTHER_AMT2;type:FLOAT;" json:"OTHER_AMT2" db:"OTHER_AMT2"`
	//[18] OTHER_AMT3                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT3 *float64 `gorm:"column:OTHER_AMT3;type:FLOAT;" json:"OTHER_AMT3" db:"OTHER_AMT3"`
	//[19] OTHER_AMT4                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT4 *float64 `gorm:"column:OTHER_AMT4;type:FLOAT;" json:"OTHER_AMT4" db:"OTHER_AMT4"`
	//[20] OTHER_AMT5                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	OTHER_AMT5 *float64 `gorm:"column:OTHER_AMT5;type:FLOAT;" json:"OTHER_AMT5" db:"OTHER_AMT5"`
	//[21] TAKAFUL_AMT                                    FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	TAKAFUL_AMT *float64 `gorm:"column:TAKAFUL_AMT;type:FLOAT;" json:"TAKAFUL_AMT" db:"TAKAFUL_AMT"`
	//[22] TANZEEM_AMT                                    FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	TANZEEM_AMT *float64 `gorm:"column:TANZEEM_AMT;type:FLOAT;" json:"TANZEEM_AMT" db:"TANZEEM_AMT"`
	//[23] METER_INSTALLS_AMT                             FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	METER_INSTALLS_AMT *float64 `gorm:"column:METER_INSTALLS_AMT;type:FLOAT;" json:"METER_INSTALLS_AMT" db:"METER_INSTALLS_AMT"`
	//[24] CONN_INSTALLS_AMT                              FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CONN_INSTALLS_AMT *float64 `gorm:"column:CONN_INSTALLS_AMT;type:FLOAT;" json:"CONN_INSTALLS_AMT" db:"CONN_INSTALLS_AMT"`
	//[25] INSTALMENT                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	INSTALMENT *float64 `gorm:"column:INSTALMENT;type:FLOAT;" json:"INSTALMENT" db:"INSTALMENT"`
	//[26] ROUND_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ROUND_AMT *float64 `gorm:"column:ROUND_AMT;type:FLOAT;" json:"ROUND_AMT" db:"ROUND_AMT"`
	//[27] METER_AMT                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	METER_AMT *float64 `gorm:"column:METER_AMT;type:FLOAT;" json:"METER_AMT" db:"METER_AMT"`
	//[28] CONN_AMT                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CONN_AMT *float64 `gorm:"column:CONN_AMT;type:FLOAT;" json:"CONN_AMT" db:"CONN_AMT"`
	//[29] METER_MAN_AMT                                  FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	METER_MAN_AMT *float64 `gorm:"column:METER_MAN_AMT;type:FLOAT;" json:"METER_MAN_AMT" db:"METER_MAN_AMT"`
	//[30] COMPUTER_AMT                                   FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	COMPUTER_AMT *float64 `gorm:"column:COMPUTER_AMT;type:FLOAT;" json:"COMPUTER_AMT" db:"COMPUTER_AMT"`
	//[31] CONTRACT_AMT                                   FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CONTRACT_AMT *float64 `gorm:"column:CONTRACT_AMT;type:FLOAT;" json:"CONTRACT_AMT" db:"CONTRACT_AMT"`
	//[32] GOV_AMT                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	GOV_AMT *float64 `gorm:"column:GOV_AMT;type:FLOAT;" json:"GOV_AMT" db:"GOV_AMT"`
	//[33] UNI_AMT                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	UNI_AMT *float64 `gorm:"column:UNI_AMT;type:FLOAT;" json:"UNI_AMT" db:"UNI_AMT"`
	//[34] CUR_CHARGES                                    FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CUR_CHARGES *float64 `gorm:"column:CUR_CHARGES;type:FLOAT;" json:"CUR_CHARGES" db:"CUR_CHARGES"`
	//[35] CUR_PAYMNTS                                    FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CUR_PAYMNTS *float64 `gorm:"column:CUR_PAYMNTS;type:FLOAT;" json:"CUR_PAYMNTS" db:"CUR_PAYMNTS"`
	//[36] RECEIPT_NO                                     NVARCHAR(120)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	RECEIPT_NO *string `gorm:"column:RECEIPT_NO;type:NVARCHAR;size:120;" json:"RECEIPT_NO" db:"RECEIPT_NO"`
}

// TableName sets the insert table name for this struct type
func (d *DEPOSIT_DETAILS) TableName() string {
	return "DEPOSIT_DETAILS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEPOSIT_DETAILS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEPOSIT_DETAILS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEPOSIT_DETAILS) Validate(action Action) error {
	return nil
}
