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

type METR_RDGS struct {
	//[ 0] BILLGROUP                                      NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:256;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 1] METER_TYPE                                     NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	METER_TYPE string `gorm:"primary_key;column:METER_TYPE;type:NVARCHAR;size:256;" json:"METER_TYPE" db:"METER_TYPE"`
	//[ 2] METER_REF                                      NVARCHAR(256)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 256     default: []
	METER_REF string `gorm:"primary_key;column:METER_REF;type:NVARCHAR;size:256;" json:"METER_REF" db:"METER_REF"`
	//[ 3] READ_NO                                        INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	READ_NO int32 `gorm:"primary_key;column:READ_NO;type:INT;" json:"READ_NO" db:"READ_NO"`
	//[ 4] BOOK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 5] WALK_NO                                        INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WALK_NO int32 `gorm:"column:WALK_NO;type:INT;" json:"WALK_NO" db:"WALK_NO"`
	//[ 6] SEQNCE_NO                                      INT                  null: false  primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SEQNCE_NO int32 `gorm:"column:SEQNCE_NO;type:INT;" json:"SEQNCE_NO" db:"SEQNCE_NO"`
	//[ 7] CUSTKEY                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CUSTKEY *string `gorm:"column:CUSTKEY;type:NVARCHAR;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 8] CR_DATE                                        DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CR_DATE *time.Time `gorm:"column:CR_DATE;type:DATETIME;" json:"CR_DATE" db:"CR_DATE"`
	//[ 9] CR_TIME                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CR_TIME *string `gorm:"column:CR_TIME;type:NVARCHAR;" json:"CR_TIME" db:"CR_TIME"`
	//[10] CR_READING                                     DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	CR_READING *float64 `gorm:"column:CR_READING;type:DECIMAL;" json:"CR_READING" db:"CR_READING"`
	//[11] CONSUMP                                        DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	CONSUMP *float64 `gorm:"column:CONSUMP;type:DECIMAL;" json:"CONSUMP" db:"CONSUMP"`
	//[12] AVRG_CONS                                      DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	AVRG_CONS *float64 `gorm:"column:AVRG_CONS;type:DECIMAL;" json:"AVRG_CONS" db:"AVRG_CONS"`
	//[13] PR_DATE1                                       DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	PR_DATE1 *time.Time `gorm:"column:PR_DATE1;type:DATETIME;" json:"PR_DATE1" db:"PR_DATE1"`
	//[14] PR_READ1                                       DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	PR_READ1 *float64 `gorm:"column:PR_READ1;type:DECIMAL;" json:"PR_READ1" db:"PR_READ1"`
	//[15] PR_CONS1                                       DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	PR_CONS1 *float64 `gorm:"column:PR_CONS1;type:DECIMAL;" json:"PR_CONS1" db:"PR_CONS1"`
	//[16] PR_FLOW1                                       DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	PR_FLOW1 *float64 `gorm:"column:PR_FLOW1;type:DECIMAL;" json:"PR_FLOW1" db:"PR_FLOW1"`
	//[17] PR_MREAD1                                      DECIMAL              null: true   primary: false  isArray: false  auto: false  col: DECIMAL         len: -1      default: []
	PR_MREAD1 *float64 `gorm:"column:PR_MREAD1;type:DECIMAL;" json:"PR_MREAD1" db:"PR_MREAD1"`
	//[18] PR_RDGN1                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PR_RDGN1 *int64 `gorm:"column:PR_RDGN1;type:INT;" json:"PR_RDGN1" db:"PR_RDGN1"`
}

// TableName sets the insert table name for this struct type
func (m *METR_RDGS) TableName() string {
	return "METR_RDGS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *METR_RDGS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *METR_RDGS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *METR_RDGS) Validate(action Action) error {
	return nil
}
