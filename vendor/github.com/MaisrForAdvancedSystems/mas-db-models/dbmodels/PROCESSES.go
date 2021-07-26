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

type PROCESSES struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: true   col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] DESCRIPTION                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 2] BILLGROUP                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NVARCHAR;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 3] BOOK_NO                                        NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NVARCHAR;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 4] WALK_NO                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WALK_NO *int64 `gorm:"column:WALK_NO;type:INT;" json:"WALK_NO" db:"WALK_NO"`
	//[ 5] START_TIME                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	START_TIME *time.Time `gorm:"column:START_TIME;type:DATETIME;" json:"START_TIME" db:"START_TIME"`
	//[ 6] END_TIME                                       DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	END_TIME *time.Time `gorm:"column:END_TIME;type:DATETIME;" json:"END_TIME" db:"END_TIME"`
	//[ 7] NO_PHASES                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NO_PHASES *int64 `gorm:"column:NO_PHASES;type:INT;" json:"NO_PHASES" db:"NO_PHASES"`
	//[ 8] IS_RUNNING                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_RUNNING *int64 `gorm:"column:IS_RUNNING;type:INT;" json:"IS_RUNNING" db:"IS_RUNNING"`
	//[ 9] IS_COMPLETED                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_COMPLETED *int64 `gorm:"column:IS_COMPLETED;type:INT;" json:"IS_COMPLETED" db:"IS_COMPLETED"`
	//[10] IS_ERROR                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_ERROR *int64 `gorm:"column:IS_ERROR;type:INT;" json:"IS_ERROR" db:"IS_ERROR"`
	//[11] IS_WARRNING                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_WARRNING *int64 `gorm:"column:IS_WARRNING;type:INT;" json:"IS_WARRNING" db:"IS_WARRNING"`
	//[12] IS_SUCCSSED                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_SUCCSSED *int64 `gorm:"column:IS_SUCCSSED;type:INT;" json:"IS_SUCCSSED" db:"IS_SUCCSSED"`
	//[13] IS_READING                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_READING *int64 `gorm:"column:IS_READING;type:INT;" json:"IS_READING" db:"IS_READING"`
	//[14] IS_COLLECTION                                  INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_COLLECTION *int64 `gorm:"column:IS_COLLECTION;type:INT;" json:"IS_COLLECTION" db:"IS_COLLECTION"`
	//[15] IS_SUPPORT_CANCELL                             INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_SUPPORT_CANCELL *int64 `gorm:"column:IS_SUPPORT_CANCELL;type:INT;" json:"IS_SUPPORT_CANCELL" db:"IS_SUPPORT_CANCELL"`
	//[16] ERROR_DESCRIPTION                              NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	ERROR_DESCRIPTION *string `gorm:"column:ERROR_DESCRIPTION;type:NVARCHAR;" json:"ERROR_DESCRIPTION" db:"ERROR_DESCRIPTION"`
	//[17] PROGRESS                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PROGRESS *int64 `gorm:"column:PROGRESS;type:INT;" json:"PROGRESS" db:"PROGRESS"`
	//[18] CURRENT_PHASE                                  NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CURRENT_PHASE *string `gorm:"column:CURRENT_PHASE;type:NVARCHAR;" json:"CURRENT_PHASE" db:"CURRENT_PHASE"`
	//[19] PHASE_PROGRESS                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PHASE_PROGRESS *int64 `gorm:"column:PHASE_PROGRESS;type:INT;" json:"PHASE_PROGRESS" db:"PHASE_PROGRESS"`
	//[20] STAMP_USER                                     NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;" json:"STAMP_USER" db:"STAMP_USER"`
	//[21] ROWS_COUNT                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ROWS_COUNT *int64 `gorm:"column:ROWS_COUNT;type:INT;" json:"ROWS_COUNT" db:"ROWS_COUNT"`
	//[22] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[23] CYCLE_ID                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID *int64 `gorm:"column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[24] PROCESS_TYPE                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PROCESS_TYPE *int64 `gorm:"column:PROCESS_TYPE;type:INT;" json:"PROCESS_TYPE" db:"PROCESS_TYPE"`
	//[25] PROCESS_GUID                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PROCESS_GUID *int64 `gorm:"column:PROCESS_GUID;type:INT;" json:"PROCESS_GUID" db:"PROCESS_GUID"`
}

// TableName sets the insert table name for this struct type
func (p *PROCESSES) TableName() string {
	return "PROCESSES"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *PROCESSES) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *PROCESSES) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *PROCESSES) Validate(action Action) error {
	return nil
}
