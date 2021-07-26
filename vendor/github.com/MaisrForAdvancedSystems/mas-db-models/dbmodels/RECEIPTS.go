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

type RECEIPTS struct {
	//[ 0] RECEIPT_NO                                     NVARCHAR(120)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	RECEIPT_NO string `gorm:"primary_key;column:RECEIPT_NO;type:NVARCHAR;size:120;" json:"RECEIPT_NO" db:"RECEIPT_NO"`
	//[ 1] CUSTKEY                                        NCHAR(120)           null: false  primary: false  isArray: false  auto: false  col: NCHAR           len: 120     default: []
	CUSTKEY string `gorm:"column:CUSTKEY;type:NCHAR;size:120;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 2] PAYMENT_NO                                     NCHAR(120)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 120     default: []
	PAYMENT_NO *string `gorm:"column:PAYMENT_NO;type:NCHAR;size:120;" json:"PAYMENT_NO" db:"PAYMENT_NO"`
	//[ 3] AMOUNT                                         FLOAT                null: false  primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	AMOUNT float64 `gorm:"column:AMOUNT;type:FLOAT;" json:"AMOUNT" db:"AMOUNT"`
	//[ 4] COLLECTION_DATE                                DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	COLLECTION_DATE *time.Time `gorm:"column:COLLECTION_DATE;type:DATETIME;" json:"COLLECTION_DATE" db:"COLLECTION_DATE"`
	//[ 5] COLLECTION_TYPE                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COLLECTION_TYPE *int64 `gorm:"column:COLLECTION_TYPE;type:INT;" json:"COLLECTION_TYPE" db:"COLLECTION_TYPE"`
	//[ 6] COLLECTION_METHOD                              INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COLLECTION_METHOD *int64 `gorm:"column:COLLECTION_METHOD;type:INT;" json:"COLLECTION_METHOD" db:"COLLECTION_METHOD"`
	//[ 7] PAYMENT_METHOD                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PAYMENT_METHOD *int64 `gorm:"column:PAYMENT_METHOD;type:INT;" json:"PAYMENT_METHOD" db:"PAYMENT_METHOD"`
	//[ 8] CHEQ_NO                                        NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	CHEQ_NO *string `gorm:"column:CHEQ_NO;type:NVARCHAR;size:60;" json:"CHEQ_NO" db:"CHEQ_NO"`
	//[ 9] CHEQ_BANK                                      NVARCHAR(120)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	CHEQ_BANK *string `gorm:"column:CHEQ_BANK;type:NVARCHAR;size:120;" json:"CHEQ_BANK" db:"CHEQ_BANK"`
	//[10] DISCOUNT_AMOUNT                                FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	DISCOUNT_AMOUNT *float64 `gorm:"column:DISCOUNT_AMOUNT;type:FLOAT;" json:"DISCOUNT_AMOUNT" db:"DISCOUNT_AMOUNT"`
	//[11] DISCOUNT_DOCUMENT_NO                           NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	DISCOUNT_DOCUMENT_NO *string `gorm:"column:DISCOUNT_DOCUMENT_NO;type:NVARCHAR;size:60;" json:"DISCOUNT_DOCUMENT_NO" db:"DISCOUNT_DOCUMENT_NO"`
	//[12] DOCUMENT_NO                                    NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	DOCUMENT_NO *string `gorm:"column:DOCUMENT_NO;type:NVARCHAR;size:60;" json:"DOCUMENT_NO" db:"DOCUMENT_NO"`
	//[13] CANCELLED                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CANCELLED bool `gorm:"column:CANCELLED;type:BIT;" json:"CANCELLED" db:"CANCELLED"`
	//[14] CANCELLED_BY                                   NVARCHAR(120)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	CANCELLED_BY *string `gorm:"column:CANCELLED_BY;type:NVARCHAR;size:120;" json:"CANCELLED_BY" db:"CANCELLED_BY"`
	//[15] CANCELLED_DATE                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CANCELLED_DATE *time.Time `gorm:"column:CANCELLED_DATE;type:DATETIME;" json:"CANCELLED_DATE" db:"CANCELLED_DATE"`
	//[16] CANCELLED_REASON                               NVARCHAR(520)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 520     default: []
	CANCELLED_REASON *string `gorm:"column:CANCELLED_REASON;type:NVARCHAR;size:520;" json:"CANCELLED_REASON" db:"CANCELLED_REASON"`
	//[17] LAT                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT *float64 `gorm:"column:LAT;type:FLOAT;" json:"LAT" db:"LAT"`
	//[18] LNG                                            FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG *float64 `gorm:"column:LNG;type:FLOAT;" json:"LNG" db:"LNG"`
	//[19] ACCURECY                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY *float64 `gorm:"column:ACCURECY;type:FLOAT;" json:"ACCURECY" db:"ACCURECY"`
	//[20] DEVICE_ID                                      NVARCHAR(120)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:120;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[21] STAMP_USER                                     NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;size:60;" json:"STAMP_USER" db:"STAMP_USER"`
	//[22] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[23] DEPOSIT_ID                                     BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	DEPOSIT_ID *int64 `gorm:"column:DEPOSIT_ID;type:BIGINT;" json:"DEPOSIT_ID" db:"DEPOSIT_ID"`
	//[24] IS_POSTED                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_POSTED *bool `gorm:"column:IS_POSTED;type:BIT;" json:"IS_POSTED" db:"IS_POSTED"`
	//[25] POST_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	POST_DATE *time.Time `gorm:"column:POST_DATE;type:DATETIME;" json:"POST_DATE" db:"POST_DATE"`
	//[26] POST_BY                                        NVARCHAR(120)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	POST_BY *string `gorm:"column:POST_BY;type:NVARCHAR;size:120;" json:"POST_BY" db:"POST_BY"`
	//[27] TRANS_NO                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	TRANS_NO *int64 `gorm:"column:TRANS_NO;type:INT;" json:"TRANS_NO" db:"TRANS_NO"`
	//[28] CYCLE_ID                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID *int32 `gorm:"column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[29] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[30] INS_CYCLE_ID                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	INS_CYCLE_ID *int64 `gorm:"column:INS_CYCLE_ID;type:INT;" json:"INS_CYCLE_ID" db:"INS_CYCLE_ID"`
	//[31] INS_BILNG_DATE                                 DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	INS_BILNG_DATE *time.Time `gorm:"column:INS_BILNG_DATE;type:DATE;" json:"INS_BILNG_DATE" db:"INS_BILNG_DATE"`
	//[32] IS_HAFZA_PRINTED                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_HAFZA_PRINTED *bool `gorm:"column:IS_HAFZA_PRINTED;type:BIT;" json:"IS_HAFZA_PRINTED" db:"IS_HAFZA_PRINTED"`
	//[33] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[34] STATION_NO                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int64 `gorm:"column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[35] BILLGROUP                                      NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[36] BOOK_NO                                        NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;size:60;" json:"BOOK_NO" db:"BOOK_NO"`
	//[37] WALK_NO                                        NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NVARCHAR;size:60;" json:"WALK_NO" db:"WALK_NO"`
	//[38] STATM_NO                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATM_NO *int64 `gorm:"column:STATM_NO;type:INT;" json:"STATM_NO" db:"STATM_NO"`
	//[39] CANCELLED_AMOUNT                               FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CANCELLED_AMOUNT *float64 `gorm:"column:CANCELLED_AMOUNT;type:FLOAT;" json:"CANCELLED_AMOUNT" db:"CANCELLED_AMOUNT"`
	//[40] RECEIPT_CHARGE1                                FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE1 *float64 `gorm:"column:RECEIPT_CHARGE1;type:FLOAT;" json:"RECEIPT_CHARGE1" db:"RECEIPT_CHARGE1"`
	//[41] RECEIPT_CHARGE2                                FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE2 *float64 `gorm:"column:RECEIPT_CHARGE2;type:FLOAT;" json:"RECEIPT_CHARGE2" db:"RECEIPT_CHARGE2"`
	//[42] RECEIPT_CHARGE3                                FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	RECEIPT_CHARGE3 *float64 `gorm:"column:RECEIPT_CHARGE3;type:FLOAT;" json:"RECEIPT_CHARGE3" db:"RECEIPT_CHARGE3"`
	//[43] HAFZA_PRINT_COUNT                              INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	HAFZA_PRINT_COUNT *int64 `gorm:"column:HAFZA_PRINT_COUNT;type:INT;" json:"HAFZA_PRINT_COUNT" db:"HAFZA_PRINT_COUNT"`
	//[44] HAFZA_PRINT_DATE                               DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HAFZA_PRINT_DATE *time.Time `gorm:"column:HAFZA_PRINT_DATE;type:DATETIME;" json:"HAFZA_PRINT_DATE" db:"HAFZA_PRINT_DATE"`
	//[45] COLLECTION_ID                                  BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	COLLECTION_ID *int64 `gorm:"column:COLLECTION_ID;type:BIGINT;" json:"COLLECTION_ID" db:"COLLECTION_ID"`
	//[46] tmp_id                                         BIGINT               null: true   primary: false  isArray: false  auto: false  col: BIGINT          len: -1      default: []
	Tmp_id *int64 `gorm:"column:tmp_id;type:BIGINT;" json:"tmp_id" db:"tmp_id"`
	//[47] FAWRY_TRANS_NO                                 VARCHAR(1000)        null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 1000    default: []
	FAWRY_TRANS_NO *string `gorm:"column:FAWRY_TRANS_NO;type:VARCHAR;size:1000;" json:"FAWRY_TRANS_NO" db:"FAWRY_TRANS_NO"`
	//[48] CANCELLED_RECIEPT_NO                           VARCHAR(256)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 256     default: []
	CANCELLED_RECIEPT_NO *string `gorm:"column:CANCELLED_RECIEPT_NO;type:VARCHAR;size:256;" json:"CANCELLED_RECIEPT_NO" db:"CANCELLED_RECIEPT_NO"`
	//[49] USER_ID                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	USER_ID *int64 `gorm:"column:USER_ID;type:INT;" json:"USER_ID" db:"USER_ID"`
	//[50] FPTN                                           NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	FPTN *string `gorm:"column:FPTN;type:NVARCHAR;size:510;" json:"FPTN" db:"FPTN"`
	//[51] BLRPTN                                         NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	BLRPTN *string `gorm:"column:BLRPTN;type:NVARCHAR;size:510;" json:"BLRPTN" db:"BLRPTN"`
	//[52] FCRN                                           NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	FCRN *string `gorm:"column:FCRN;type:NVARCHAR;size:510;" json:"FCRN" db:"FCRN"`
	//[53] COMMENT                                        VARCHAR(255)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 255     default: []
	COMMENT *string `gorm:"column:COMMENT;type:VARCHAR;size:255;" json:"COMMENT" db:"COMMENT"`

	RECEIPT_TYPE *int64 `gorm:"column:RECEIPT_TYPE;type:INT;" json:"RECEIPT_TYPE" db:"RECEIPT_TYPE"`

	//IS_HH_NOTIFIED bool `gorm:"column:IS_HH_NOTIFIED;type:bit;" json:"IS_HH_NOTIFIED" db:"IS_HH_NOTIFIED"`
}

// TableName sets the insert table name for this struct type
func (r *RECEIPTS) TableName() string {
	return "RECEIPTS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RECEIPTS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RECEIPTS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RECEIPTS) Validate(action Action) error {
	return nil
}
