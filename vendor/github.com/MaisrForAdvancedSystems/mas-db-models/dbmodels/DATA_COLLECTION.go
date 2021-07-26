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

type DATA_COLLECTION struct {
	//[ 0] CUSTKEY                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	CUSTKEY string `gorm:"primary_key;column:CUSTKEY;type:NVARCHAR;size:60;" json:"CUSTKEY" db:"CUSTKEY"`
	//[ 1] DEVICE_ID                                      NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:60;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 2] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[ 3] C_TYPE                                         NVARCHAR(20)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 20      default: []
	C_TYPE *string `gorm:"column:C_TYPE;type:NVARCHAR;size:20;" json:"C_TYPE" db:"C_TYPE"`
	//[ 4] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[ 5] BRANCH_ID                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	BRANCH_ID *int64 `gorm:"column:BRANCH_ID;type:INT;" json:"BRANCH_ID" db:"BRANCH_ID"`
	//[ 6] BRANCH_NAME                                    NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	BRANCH_NAME *string `gorm:"column:BRANCH_NAME;type:NVARCHAR;size:100;" json:"BRANCH_NAME" db:"BRANCH_NAME"`
	//[ 7] BILLGROUP                                      NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BILLGROUP *string `gorm:"column:BILLGROUP;type:NCHAR;size:20;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 8] BOOK_NO                                        NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	BOOK_NO *string `gorm:"column:BOOK_NO;type:NCHAR;size:20;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 9] WALK_NO                                        NCHAR(20)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 20      default: []
	WALK_NO *string `gorm:"column:WALK_NO;type:NCHAR;size:20;" json:"WALK_NO" db:"WALK_NO"`
	//[10] CYCLE_ID                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID *int64 `gorm:"column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[11] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[12] COLLETION_DATE                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	COLLETION_DATE *time.Time `gorm:"column:COLLETION_DATE;type:DATETIME;" json:"COLLETION_DATE" db:"COLLETION_DATE"`
	//[13] READING_DATE                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	READING_DATE *time.Time `gorm:"column:READING_DATE;type:DATETIME;" json:"READING_DATE" db:"READING_DATE"`
	//[14] IS_CURRENT_STATION                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_CURRENT_STATION *bool `gorm:"column:IS_CURRENT_STATION;type:BIT;" json:"IS_CURRENT_STATION" db:"IS_CURRENT_STATION"`
	//[15] LAT_R                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_R *float64 `gorm:"column:LAT_R;type:FLOAT;" json:"LAT_R" db:"LAT_R"`
	//[16] LAT_C                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_C *float64 `gorm:"column:LAT_C;type:FLOAT;" json:"LAT_C" db:"LAT_C"`
	//[17] LNG_R                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_R *float64 `gorm:"column:LNG_R;type:FLOAT;" json:"LNG_R" db:"LNG_R"`
	//[18] LNG_C                                          FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_C *float64 `gorm:"column:LNG_C;type:FLOAT;" json:"LNG_C" db:"LNG_C"`
	//[19] CL_BLNCE                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CL_BLNCE *float64 `gorm:"column:CL_BLNCE;type:FLOAT;" json:"CL_BLNCE" db:"CL_BLNCE"`
	//[20] CONSUMPTION                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CONSUMPTION *int64 `gorm:"column:CONSUMPTION;type:INT;" json:"CONSUMPTION" db:"CONSUMPTION"`
	//[21] CR_READING                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CR_READING *int64 `gorm:"column:CR_READING;type:INT;" json:"CR_READING" db:"CR_READING"`
	//[22] PERM_ID                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PERM_ID *int64 `gorm:"column:PERM_ID;type:INT;" json:"PERM_ID" db:"PERM_ID"`
	//[23] READING_DESCRPENCY                             INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	READING_DESCRPENCY *int64 `gorm:"column:READING_DESCRPENCY;type:INT;" json:"READING_DESCRPENCY" db:"READING_DESCRPENCY"`
	//[24] READING_NOTE                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	READING_NOTE *int64 `gorm:"column:READING_NOTE;type:INT;" json:"READING_NOTE" db:"READING_NOTE"`
	//[25] COLLECTION_NOTE                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COLLECTION_NOTE *int64 `gorm:"column:COLLECTION_NOTE;type:INT;" json:"COLLECTION_NOTE" db:"COLLECTION_NOTE"`
	//[26] IS_NOTIFIED                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_NOTIFIED *bool `gorm:"column:IS_NOTIFIED;type:BIT;" json:"IS_NOTIFIED" db:"IS_NOTIFIED"`
	//[27] NOTIFICATION_DATE                              DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	NOTIFICATION_DATE *time.Time `gorm:"column:NOTIFICATION_DATE;type:DATETIME;" json:"NOTIFICATION_DATE" db:"NOTIFICATION_DATE"`
	//[28] SOURCE_TYPE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SOURCE_TYPE *int64 `gorm:"column:SOURCE_TYPE;type:INT;" json:"SOURCE_TYPE" db:"SOURCE_TYPE"`
	//[29] CONNECTION_ADDRESS                             NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CONNECTION_ADDRESS *string `gorm:"column:CONNECTION_ADDRESS;type:NVARCHAR;size:100;" json:"CONNECTION_ADDRESS" db:"CONNECTION_ADDRESS"`
	//[30] SOURCE_NAME                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	SOURCE_NAME *string `gorm:"column:SOURCE_NAME;type:NVARCHAR;size:200;" json:"SOURCE_NAME" db:"SOURCE_NAME"`
	//[31] DELIVERY_ST                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DELIVERY_ST *int64 `gorm:"column:DELIVERY_ST;type:INT;" json:"DELIVERY_ST" db:"DELIVERY_ST"`
	//[32] ACCURECY_C                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY_C *float64 `gorm:"column:ACCURECY_C;type:FLOAT;" json:"ACCURECY_C" db:"ACCURECY_C"`
	//[33] ACCURECY_R                                     FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	ACCURECY_R *float64 `gorm:"column:ACCURECY_R;type:FLOAT;" json:"ACCURECY_R" db:"ACCURECY_R"`
}

// TableName sets the insert table name for this struct type
func (d *DATA_COLLECTION) TableName() string {
	return "DATA_COLLECTION"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DATA_COLLECTION) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DATA_COLLECTION) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DATA_COLLECTION) Validate(action Action) error {
	return nil
}
