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

type STATM_DEPOSITS struct {
	//[ 0] DEPOSIT_ID                                     BIGINT               null: false  primary: true   isArray: false  auto: false  col: BIGINT          len: -1      default: []
	DEPOSIT_ID int64 `gorm:"primary_key;column:DEPOSIT_ID;type:BIGINT;" json:"DEPOSIT_ID" db:"DEPOSIT_ID"`
	//[ 1] RECEIPT_NO                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	RECEIPT_NO *string `gorm:"column:RECEIPT_NO;type:NVARCHAR;" json:"RECEIPT_NO" db:"RECEIPT_NO"`
	//[ 2] AMOUNT                                         DECIMAL              null: false  primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	AMOUNT float64 `gorm:"column:AMOUNT;type:DECIMAL;" json:"AMOUNT" db:"AMOUNT"`
	//[ 3] COUNT                                          INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COUNT int32 `gorm:"column:COUNT;type:INT;" json:"COUNT" db:"COUNT"`
	//[ 4] DELIVERY_MAN                                   NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DELIVERY_MAN *string `gorm:"column:DELIVERY_MAN;type:NVARCHAR;" json:"DELIVERY_MAN" db:"DELIVERY_MAN"`
	//[ 5] DELIVERY_DATE                                  DATETIME             null: false  primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	DELIVERY_DATE *time.Time `gorm:"column:DELIVERY_DATE;type:DATETIME;" json:"DELIVERY_DATE" db:"DELIVERY_DATE"`
	//[ 6] STAMP_USER                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;" json:"STAMP_USER" db:"STAMP_USER"`
	//[ 7] STAMP_DATE                                     DATETIME             null: false  primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 8] BILNG_DATE                                     DATETIME             null: false  primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATETIME;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 9] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[10] CANCEL_AMOUNT                                  FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CANCEL_AMOUNT *float64 `gorm:"column:CANCEL_AMOUNT;type:FLOAT;" json:"CANCEL_AMOUNT" db:"CANCEL_AMOUNT"`
	//[11] CANCEL_BY                                      FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CANCEL_BY *float64 `gorm:"column:CANCEL_BY;type:FLOAT;" json:"CANCEL_BY" db:"CANCEL_BY"`
	//[12] CANCEL_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CANCEL_DATE *time.Time `gorm:"column:CANCEL_DATE;type:DATETIME;" json:"CANCEL_DATE" db:"CANCEL_DATE"`
	//[13] NET_AMOUNT                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	NET_AMOUNT *float64 `gorm:"column:NET_AMOUNT;type:FLOAT;" json:"NET_AMOUNT" db:"NET_AMOUNT"`
	//[14] FALGE1                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FALGE1 *bool `gorm:"column:FALGE1;type:BIT;" json:"FALGE1" db:"FALGE1"`
	//[15] FALGE2                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FALGE2 *bool `gorm:"column:FALGE2;type:BIT;" json:"FALGE2" db:"FALGE2"`
	//[16] BILNG_DATE2                                    DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE2 *time.Time `gorm:"column:BILNG_DATE2;type:DATE;" json:"BILNG_DATE2" db:"BILNG_DATE2"`
	//[17] BILLING_DEPOSITID                              INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	BILLING_DEPOSITID *int64 `gorm:"column:BILLING_DEPOSITID;type:INT;" json:"BILLING_DEPOSITID" db:"BILLING_DEPOSITID"`
	//[18] BANK_RECPT_NO                                  NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	BANK_RECPT_NO *string `gorm:"column:BANK_RECPT_NO;type:NCHAR;size:100;" json:"BANK_RECPT_NO" db:"BANK_RECPT_NO"`
	//[19] BANK_ACC_NO                                    NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	BANK_ACC_NO *string `gorm:"column:BANK_ACC_NO;type:NCHAR;size:100;" json:"BANK_ACC_NO" db:"BANK_ACC_NO"`
	//[20] BANK_NAME                                      NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	BANK_NAME *string `gorm:"column:BANK_NAME;type:NCHAR;size:100;" json:"BANK_NAME" db:"BANK_NAME"`
	//[21] BANK_POST_DATE                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BANK_POST_DATE *time.Time `gorm:"column:BANK_POST_DATE;type:DATETIME;" json:"BANK_POST_DATE" db:"BANK_POST_DATE"`
	//[22] DEDUCATIONS_AMOUNT                             FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DEDUCATIONS_AMOUNT *float64 `gorm:"column:DEDUCATIONS_AMOUNT;type:FLOAT;" json:"DEDUCATIONS_AMOUNT" db:"DEDUCATIONS_AMOUNT"`
	//[23] CHQ_NO                                         NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	CHQ_NO *string `gorm:"column:CHQ_NO;type:NCHAR;size:100;" json:"CHQ_NO" db:"CHQ_NO"`
	//[24] CHQ_BANK                                       NCHAR(100)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 100     default: []
	CHQ_BANK *string `gorm:"column:CHQ_BANK;type:NCHAR;size:100;" json:"CHQ_BANK" db:"CHQ_BANK"`
	//[25] IS_POSTED                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_POSTED *bool `gorm:"column:IS_POSTED;type:BIT;" json:"IS_POSTED" db:"IS_POSTED"`
	//[26] BILNG_DEPOSIT_ID                               BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	BILNG_DEPOSIT_ID *int64 `gorm:"column:BILNG_DEPOSIT_ID;type:BIGINT;" json:"BILNG_DEPOSIT_ID" db:"BILNG_DEPOSIT_ID"`
	//[27] POST_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	POST_DATE *time.Time `gorm:"column:POST_DATE;type:DATETIME;" json:"POST_DATE" db:"POST_DATE"`
	//[28] STATION_NO                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int64 `gorm:"column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[29] FROM_BILNG                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FROM_BILNG *bool `gorm:"column:FROM_BILNG;type:BIT;" json:"FROM_BILNG" db:"FROM_BILNG"`
	//[30] deposit_id_t                                   BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	Deposit_id_t *int64 `gorm:"column:deposit_id_t;type:BIGINT;" json:"deposit_id_t" db:"deposit_id_t"`
	//[31] deposit_id_t2                                  BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	Deposit_id_t2 *int64 `gorm:"column:deposit_id_t2;type:BIGINT;" json:"deposit_id_t2" db:"deposit_id_t2"`
	//[32] USER_ID                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	USER_ID *int64 `gorm:"column:USER_ID;type:INT;" json:"USER_ID" db:"USER_ID"`
}

// TableName sets the insert table name for this struct type
func (s *STATM_DEPOSITS) TableName() string {
	return "STATM_DEPOSITS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *STATM_DEPOSITS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *STATM_DEPOSITS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *STATM_DEPOSITS) Validate(action Action) error {
	return nil
}
