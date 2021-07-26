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

type FAWRY_TRANS struct {
	//[ 0] ID                                             BIGINT               null: false  primary: true   isArray: false  auto: true   col: BIGINT          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:ID;type:BIGINT;" json:"ID" db:"ID"`
	//[ 1] REQ_TYPE                                       NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	REQ_TYPE *string `gorm:"column:REQ_TYPE;type:NVARCHAR;size:200;" json:"REQ_TYPE" db:"REQ_TYPE"`
	//[ 2] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 3] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 4] MSG_CODE                                       NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	MSG_CODE *string `gorm:"column:MSG_CODE;type:NVARCHAR;size:200;" json:"MSG_CODE" db:"MSG_CODE"`
	//[ 5] CUSTKEY                                        NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CUSTKEY *string `gorm:"column:CUSTKEY;type:NVARCHAR;size:200;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 6] STATION_NO                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	STATION_NO *string `gorm:"column:STATION_NO;type:NVARCHAR;size:200;" json:"STATION_NO" db:"STATION_NO"`
	//[ 7] PROJECT_TAG                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	PROJECT_TAG *string `gorm:"column:PROJECT_TAG;type:NVARCHAR;size:200;" json:"PROJECT_TAG" db:"PROJECT_TAG"`
	//[ 8] PAYMENT_NO                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	PAYMENT_NO *string `gorm:"column:PAYMENT_NO;type:NVARCHAR;size:200;" json:"PAYMENT_NO" db:"PAYMENT_NO"`
	//[ 9] CL_BLNCE                                       DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	CL_BLNCE *float64 `gorm:"column:CL_BLNCE;type:DECIMAL;" json:"CL_BLNCE" db:"CL_BLNCE"`
	//[10] AMOUNT_COLLECTED                               DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	AMOUNT_COLLECTED *float64 `gorm:"column:AMOUNT_COLLECTED;type:DECIMAL;" json:"AMOUNT_COLLECTED" db:"AMOUNT_COLLECTED"`
	//[11] REMOTE_IP                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	REMOTE_IP *string `gorm:"column:REMOTE_IP;type:NVARCHAR;size:200;" json:"REMOTE_IP" db:"REMOTE_IP"`
	//[12] REMOTE_PORT                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	REMOTE_PORT *int64 `gorm:"column:REMOTE_PORT;type:INT;" json:"REMOTE_PORT" db:"REMOTE_PORT"`
	//[13] FAWRY_BILLER_ID                                NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	FAWRY_BILLER_ID *string `gorm:"column:FAWRY_BILLER_ID;type:NVARCHAR;size:200;" json:"FAWRY_BILLER_ID" db:"FAWRY_BILLER_ID"`
	//[14] TERMINAL_ID                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TERMINAL_ID *string `gorm:"column:TERMINAL_ID;type:NVARCHAR;size:200;" json:"TERMINAL_ID" db:"TERMINAL_ID"`
	//[15] PYMT_METHOD                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	PYMT_METHOD *string `gorm:"column:PYMT_METHOD;type:NVARCHAR;size:200;" json:"PYMT_METHOD" db:"PYMT_METHOD"`
	//[16] FAWRY_TRANS_ID                                 NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	FAWRY_TRANS_ID *string `gorm:"column:FAWRY_TRANS_ID;type:NVARCHAR;size:200;" json:"FAWRY_TRANS_ID" db:"FAWRY_TRANS_ID"`
	//[17] ERROR_MSG                                      NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	ERROR_MSG *string `gorm:"column:ERROR_MSG;type:NVARCHAR;size:512;" json:"ERROR_MSG" db:"ERROR_MSG"`
	//[18] ERROR_NO                                       NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	ERROR_NO *string `gorm:"column:ERROR_NO;type:NVARCHAR;size:200;" json:"ERROR_NO" db:"ERROR_NO"`
	//[19] IS_RETRY                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_RETRY *int64 `gorm:"column:IS_RETRY;type:INT;" json:"IS_RETRY" db:"IS_RETRY"`
}

// TableName sets the insert table name for this struct type
func (f *FAWRY_TRANS) TableName() string {
	return "FAWRY_TRANS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FAWRY_TRANS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FAWRY_TRANS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FAWRY_TRANS) Validate(action Action) error {
	return nil
}
