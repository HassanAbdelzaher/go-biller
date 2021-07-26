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

type ST_BCYC struct {
	//[ 0] BOOK_NO                                        NVARCHAR(40)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NVARCHAR;size:40;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 1] WALK_NO                                        NVARCHAR(40)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	WALK_NO string `gorm:"primary_key;column:WALK_NO;type:NVARCHAR;size:40;" json:"WALK_NO" db:"WALK_NO"`
	//[ 2] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 3] BRANCH                                         NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BRANCH *string `gorm:"column:BRANCH;type:NVARCHAR;" json:"BRANCH" db:"BRANCH"`
	//[ 4] BILLGROUP                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NVARCHAR;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 5] BILNG_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATETIME;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 6] BDB_CDB                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	BDB_CDB *int64 `gorm:"column:BDB_CDB;type:INT;" json:"BDB_CDB" db:"BDB_CDB"`
	//[ 7] BDB_CDB_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BDB_CDB_DATE *time.Time `gorm:"column:BDB_CDB_DATE;type:DATETIME;" json:"BDB_CDB_DATE" db:"BDB_CDB_DATE"`
	//[ 8] BDB_CDB_USER                                   NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	BDB_CDB_USER *string `gorm:"column:BDB_CDB_USER;type:NVARCHAR;" json:"BDB_CDB_USER" db:"BDB_CDB_USER"`
	//[ 9] CDB_HH                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_HH *int64 `gorm:"column:CDB_HH;type:INT;" json:"CDB_HH" db:"CDB_HH"`
	//[10] CDB_HH_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_HH_DATE *time.Time `gorm:"column:CDB_HH_DATE;type:DATETIME;" json:"CDB_HH_DATE" db:"CDB_HH_DATE"`
	//[11] CDB_HH_USER                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CDB_HH_USER *string `gorm:"column:CDB_HH_USER;type:NVARCHAR;" json:"CDB_HH_USER" db:"CDB_HH_USER"`
	//[12] HH_CDB                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	HH_CDB *int64 `gorm:"column:HH_CDB;type:INT;" json:"HH_CDB" db:"HH_CDB"`
	//[13] HH_CDB_DATE                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HH_CDB_DATE *time.Time `gorm:"column:HH_CDB_DATE;type:DATETIME;" json:"HH_CDB_DATE" db:"HH_CDB_DATE"`
	//[14] HH_CDB_USER                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	HH_CDB_USER *string `gorm:"column:HH_CDB_USER;type:NVARCHAR;" json:"HH_CDB_USER" db:"HH_CDB_USER"`
	//[15] U_HH                                           INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	U_HH *int64 `gorm:"column:U_HH;type:INT;" json:"U_HH" db:"U_HH"`
	//[16] U_HH_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	U_HH_DATE *time.Time `gorm:"column:U_HH_DATE;type:DATETIME;" json:"U_HH_DATE" db:"U_HH_DATE"`
	//[17] U_HH_USER                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	U_HH_USER *string `gorm:"column:U_HH_USER;type:NVARCHAR;" json:"U_HH_USER" db:"U_HH_USER"`
	//[18] HH_U                                           INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	HH_U *int64 `gorm:"column:HH_U;type:INT;" json:"HH_U" db:"HH_U"`
	//[19] HH_U_DATE                                      DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HH_U_DATE *time.Time `gorm:"column:HH_U_DATE;type:DATETIME;" json:"HH_U_DATE" db:"HH_U_DATE"`
	//[20] HH_U_USER                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	HH_U_USER *string `gorm:"column:HH_U_USER;type:NVARCHAR;" json:"HH_U_USER" db:"HH_U_USER"`
	//[21] CDB_BDB                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_BDB *int64 `gorm:"column:CDB_BDB;type:INT;" json:"CDB_BDB" db:"CDB_BDB"`
	//[22] CDB_BDB_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_BDB_DATE *time.Time `gorm:"column:CDB_BDB_DATE;type:DATETIME;" json:"CDB_BDB_DATE" db:"CDB_BDB_DATE"`
	//[23] CDB_BDB_USER                                   NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CDB_BDB_USER *string `gorm:"column:CDB_BDB_USER;type:NVARCHAR;" json:"CDB_BDB_USER" db:"CDB_BDB_USER"`
	//[24] ISCYCLE_COMPLETED                              INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ISCYCLE_COMPLETED *int64 `gorm:"column:ISCYCLE_COMPLETED;type:INT;" json:"ISCYCLE_COMPLETED" db:"ISCYCLE_COMPLETED"`
	//[25] CLOSE_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CLOSE_DATE *time.Time `gorm:"column:CLOSE_DATE;type:DATETIME;" json:"CLOSE_DATE" db:"CLOSE_DATE"`
	//[26] CLOSE_BY                                       NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	CLOSE_BY *string `gorm:"column:CLOSE_BY;type:NVARCHAR;" json:"CLOSE_BY" db:"CLOSE_BY"`
	//[27] DEVICE_ID                                      NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[28] DESCRIPTION                                    NVARCHAR             null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: -1      default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[29] IsReady                                        BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IsReady *bool `gorm:"column:IsReady;type:BIT;" json:"IsReady" db:"IsReady"`
}

// TableName sets the insert table name for this struct type
func (s *ST_BCYC) TableName() string {
	return "ST_BCYC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *ST_BCYC) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *ST_BCYC) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *ST_BCYC) Validate(action Action) error {
	return nil
}
