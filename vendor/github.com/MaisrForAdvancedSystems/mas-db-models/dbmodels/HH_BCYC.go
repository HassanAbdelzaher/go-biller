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

type HH_BCYC struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] BILLGROUP                                      NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILLGROUP string `gorm:"primary_key;column:BILLGROUP;type:NVARCHAR;size:60;" json:"BILLGROUP" db:"BILLGROUP"`
	//[ 2] BOOK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BOOK_NO string `gorm:"primary_key;column:BOOK_NO;type:NVARCHAR;size:60;" json:"BOOK_NO" db:"BOOK_NO"`
	//[ 3] WALK_NO                                        NVARCHAR(60)         null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	WALK_NO string `gorm:"primary_key;column:WALK_NO;type:NVARCHAR;size:60;" json:"WALK_NO" db:"WALK_NO"`
	//[ 4] CYCLE_ID                                       INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	CYCLE_ID int32 `gorm:"primary_key;column:CYCLE_ID;type:INT;" json:"CYCLE_ID" db:"CYCLE_ID"`
	//[ 5] IS_COLLECTION                                  BIT                  null: false  primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_COLLECTION bool `gorm:"column:IS_COLLECTION;type:BIT;" json:"IS_COLLECTION" db:"IS_COLLECTION"`
	//[ 6] IS_READING                                     BIT                  null: false  primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_READING bool `gorm:"column:IS_READING;type:BIT;" json:"IS_READING" db:"IS_READING"`
	//[ 7] BRANCH                                         NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	BRANCH *string `gorm:"column:BRANCH;type:NVARCHAR;size:200;" json:"BRANCH" db:"BRANCH"`
	//[ 8] BILNG_DATE                                     DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	BILNG_DATE *time.Time `gorm:"column:BILNG_DATE;type:DATE;" json:"BILNG_DATE" db:"BILNG_DATE"`
	//[ 9] BDB_CDB_C                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	BDB_CDB_C *int64 `gorm:"column:BDB_CDB_C;type:INT;" json:"BDB_CDB_C" db:"BDB_CDB_C"`
	//[10] BDB_CDB_DATE_C                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BDB_CDB_DATE_C *time.Time `gorm:"column:BDB_CDB_DATE_C;type:DATETIME;" json:"BDB_CDB_DATE_C" db:"BDB_CDB_DATE_C"`
	//[11] BDB_CDB_USER_C                                 NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	BDB_CDB_USER_C *string `gorm:"column:BDB_CDB_USER_C;type:NVARCHAR;size:200;" json:"BDB_CDB_USER_C" db:"BDB_CDB_USER_C"`
	//[12] BDB_CDB_R                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	BDB_CDB_R *int64 `gorm:"column:BDB_CDB_R;type:INT;" json:"BDB_CDB_R" db:"BDB_CDB_R"`
	//[13] BDB_CDB_DATE_R                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	BDB_CDB_DATE_R *time.Time `gorm:"column:BDB_CDB_DATE_R;type:DATETIME;" json:"BDB_CDB_DATE_R" db:"BDB_CDB_DATE_R"`
	//[14] BDB_CDB_USER_R                                 NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	BDB_CDB_USER_R *string `gorm:"column:BDB_CDB_USER_R;type:NVARCHAR;size:200;" json:"BDB_CDB_USER_R" db:"BDB_CDB_USER_R"`
	//[15] CDB_HH_R                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_HH_R *int64 `gorm:"column:CDB_HH_R;type:INT;" json:"CDB_HH_R" db:"CDB_HH_R"`
	//[16] CDB_HH_DATE_R                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_HH_DATE_R *time.Time `gorm:"column:CDB_HH_DATE_R;type:DATETIME;" json:"CDB_HH_DATE_R" db:"CDB_HH_DATE_R"`
	//[17] CDB_HH_USER_R                                  NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CDB_HH_USER_R *string `gorm:"column:CDB_HH_USER_R;type:NVARCHAR;size:200;" json:"CDB_HH_USER_R" db:"CDB_HH_USER_R"`
	//[18] CDB_HH_C                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_HH_C *int64 `gorm:"column:CDB_HH_C;type:INT;" json:"CDB_HH_C" db:"CDB_HH_C"`
	//[19] CDB_HH_DATE_C                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_HH_DATE_C *time.Time `gorm:"column:CDB_HH_DATE_C;type:DATETIME;" json:"CDB_HH_DATE_C" db:"CDB_HH_DATE_C"`
	//[20] CDB_HH_USER_C                                  NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CDB_HH_USER_C *string `gorm:"column:CDB_HH_USER_C;type:NVARCHAR;size:200;" json:"CDB_HH_USER_C" db:"CDB_HH_USER_C"`
	//[21] HH_CDB_C                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	HH_CDB_C *int64 `gorm:"column:HH_CDB_C;type:INT;" json:"HH_CDB_C" db:"HH_CDB_C"`
	//[22] HH_CDB_DATE_C                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HH_CDB_DATE_C *time.Time `gorm:"column:HH_CDB_DATE_C;type:DATETIME;" json:"HH_CDB_DATE_C" db:"HH_CDB_DATE_C"`
	//[23] HH_CDB_USER_C                                  NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	HH_CDB_USER_C *string `gorm:"column:HH_CDB_USER_C;type:NVARCHAR;size:200;" json:"HH_CDB_USER_C" db:"HH_CDB_USER_C"`
	//[24] HH_CDB_R                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	HH_CDB_R *int64 `gorm:"column:HH_CDB_R;type:INT;" json:"HH_CDB_R" db:"HH_CDB_R"`
	//[25] HH_CDB_DATE_R                                  DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	HH_CDB_DATE_R *time.Time `gorm:"column:HH_CDB_DATE_R;type:DATETIME;" json:"HH_CDB_DATE_R" db:"HH_CDB_DATE_R"`
	//[26] HH_CDB_USER_R                                  NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	HH_CDB_USER_R *string `gorm:"column:HH_CDB_USER_R;type:NVARCHAR;size:200;" json:"HH_CDB_USER_R" db:"HH_CDB_USER_R"`
	//[27] CDB_BDB_R                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_BDB_R *int64 `gorm:"column:CDB_BDB_R;type:INT;" json:"CDB_BDB_R" db:"CDB_BDB_R"`
	//[28] CDB_BDB_DATE_R                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_BDB_DATE_R *time.Time `gorm:"column:CDB_BDB_DATE_R;type:DATETIME;" json:"CDB_BDB_DATE_R" db:"CDB_BDB_DATE_R"`
	//[29] CDB_BDB_USER_R                                 NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CDB_BDB_USER_R *string `gorm:"column:CDB_BDB_USER_R;type:NVARCHAR;size:200;" json:"CDB_BDB_USER_R" db:"CDB_BDB_USER_R"`
	//[30] CDB_BDB_C                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CDB_BDB_C *int64 `gorm:"column:CDB_BDB_C;type:INT;" json:"CDB_BDB_C" db:"CDB_BDB_C"`
	//[31] CDB_BDB_DATE_C                                 DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CDB_BDB_DATE_C *time.Time `gorm:"column:CDB_BDB_DATE_C;type:DATETIME;" json:"CDB_BDB_DATE_C" db:"CDB_BDB_DATE_C"`
	//[32] CDB_BDB_USER_C                                 NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CDB_BDB_USER_C *string `gorm:"column:CDB_BDB_USER_C;type:NVARCHAR;size:200;" json:"CDB_BDB_USER_C" db:"CDB_BDB_USER_C"`
	//[33] ISCYCLE_COMPLETED_C                            INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ISCYCLE_COMPLETED_C *int64 `gorm:"column:ISCYCLE_COMPLETED_C;type:INT;" json:"ISCYCLE_COMPLETED_C" db:"ISCYCLE_COMPLETED_C"`
	//[34] ISCYCLE_COMPLETED_R                            INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ISCYCLE_COMPLETED_R *int64 `gorm:"column:ISCYCLE_COMPLETED_R;type:INT;" json:"ISCYCLE_COMPLETED_R" db:"ISCYCLE_COMPLETED_R"`
	//[35] CLOSE_DATE_C                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CLOSE_DATE_C *time.Time `gorm:"column:CLOSE_DATE_C;type:DATETIME;" json:"CLOSE_DATE_C" db:"CLOSE_DATE_C"`
	//[36] CLOSE_BY_C                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CLOSE_BY_C *string `gorm:"column:CLOSE_BY_C;type:NVARCHAR;size:200;" json:"CLOSE_BY_C" db:"CLOSE_BY_C"`
	//[37] CLOSE_DATE_R                                   DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	CLOSE_DATE_R *time.Time `gorm:"column:CLOSE_DATE_R;type:DATETIME;" json:"CLOSE_DATE_R" db:"CLOSE_DATE_R"`
	//[38] CLOSE_BY_R                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	CLOSE_BY_R *string `gorm:"column:CLOSE_BY_R;type:NVARCHAR;size:200;" json:"CLOSE_BY_R" db:"CLOSE_BY_R"`
	//[39] DEVICEID_R                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	DEVICEID_R *string `gorm:"column:DEVICEID_R;type:NVARCHAR;size:200;" json:"DEVICEID_R" db:"DEVICEID_R"`
	//[40] DEVICEID_C                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	DEVICEID_C *string `gorm:"column:DEVICEID_C;type:NVARCHAR;size:200;" json:"DEVICEID_C" db:"DEVICEID_C"`
	//[41] EMPID_C                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMPID_C *int64 `gorm:"column:EMPID_C;type:INT;" json:"EMPID_C" db:"EMPID_C"`
	//[42] EMPID_R                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMPID_R *int64 `gorm:"column:EMPID_R;type:INT;" json:"EMPID_R" db:"EMPID_R"`
	//[43] COUNT_C                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COUNT_C *int64 `gorm:"column:COUNT_C;type:INT;" json:"COUNT_C" db:"COUNT_C"`
	//[44] COUNT_R                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COUNT_R *int64 `gorm:"column:COUNT_R;type:INT;" json:"COUNT_R" db:"COUNT_R"`
	//[45] DESCRIPTION                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:200;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[46] IS_ALLOWED_C                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_ALLOWED_C *bool `gorm:"column:IS_ALLOWED_C;type:BIT;" json:"IS_ALLOWED_C" db:"IS_ALLOWED_C"`
	//[47] IS_ALLOWED_R                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_ALLOWED_R *bool `gorm:"column:IS_ALLOWED_R;type:BIT;" json:"IS_ALLOWED_R" db:"IS_ALLOWED_R"`
	//[48] CL_BLNCE                                       FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	CL_BLNCE *float64 `gorm:"column:CL_BLNCE;type:FLOAT;" json:"CL_BLNCE" db:"CL_BLNCE"`
	//[49] COLLECTED_AMOUNT                               FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	COLLECTED_AMOUNT *float64 `gorm:"column:COLLECTED_AMOUNT;type:FLOAT;" json:"COLLECTED_AMOUNT" db:"COLLECTED_AMOUNT"`
	//[50] COLLECTED_COUNT                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	COLLECTED_COUNT *int64 `gorm:"column:COLLECTED_COUNT;type:INT;" json:"COLLECTED_COUNT" db:"COLLECTED_COUNT"`
	//[51] READED_COUNT                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	READED_COUNT *int64 `gorm:"column:READED_COUNT;type:INT;" json:"READED_COUNT" db:"READED_COUNT"`
	//[52] REACHED_C                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	REACHED_C *int64 `gorm:"column:REACHED_C;type:INT;" json:"REACHED_C" db:"REACHED_C"`
	//[53] REACHED_R                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	REACHED_R *int64 `gorm:"column:REACHED_R;type:INT;" json:"REACHED_R" db:"REACHED_R"`
	//[54] POSTED_C                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	POSTED_C *int64 `gorm:"column:POSTED_C;type:INT;" json:"POSTED_C" db:"POSTED_C"`
	//[55] POSTED_R                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	POSTED_R *int64 `gorm:"column:POSTED_R;type:INT;" json:"POSTED_R" db:"POSTED_R"`
	//[56] ISCLOSED_INDEVICE_C                            BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ISCLOSED_INDEVICE_C *bool `gorm:"column:ISCLOSED_INDEVICE_C;type:BIT;" json:"ISCLOSED_INDEVICE_C" db:"ISCLOSED_INDEVICE_C"`
	//[57] ISCLOSED_INDEVICE_R                            BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ISCLOSED_INDEVICE_R *bool `gorm:"column:ISCLOSED_INDEVICE_R;type:BIT;" json:"ISCLOSED_INDEVICE_R" db:"ISCLOSED_INDEVICE_R"`
	//[58] IS_METER_WALK                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_METER_WALK *bool `gorm:"column:IS_METER_WALK;type:BIT;" json:"IS_METER_WALK" db:"IS_METER_WALK"`
	//[59] IS_CUSTOMER_WALK                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_CUSTOMER_WALK *bool `gorm:"column:IS_CUSTOMER_WALK;type:BIT;" json:"IS_CUSTOMER_WALK" db:"IS_CUSTOMER_WALK"`
	//[60] CLOSE_INDEVICE_COLlECTED                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CLOSE_INDEVICE_COLlECTED *int64 `gorm:"column:CLOSE_INDEVICE_COLlECTED;type:INT;" json:"CLOSE_INDEVICE_COLlECTED" db:"CLOSE_INDEVICE_COLlECTED"`
	//[61] CLOSE_INDEVICE_READED                          INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CLOSE_INDEVICE_READED *int64 `gorm:"column:CLOSE_INDEVICE_READED;type:INT;" json:"CLOSE_INDEVICE_READED" db:"CLOSE_INDEVICE_READED"`
	//[62] CLOSE_INDEVICE_DATE_R                          DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CLOSE_INDEVICE_DATE_R *time.Time `gorm:"column:CLOSE_INDEVICE_DATE_R;type:DATE;" json:"CLOSE_INDEVICE_DATE_R" db:"CLOSE_INDEVICE_DATE_R"`
	//[63] CLOSE_INDEVICE_DATE_C                          DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	CLOSE_INDEVICE_DATE_C *time.Time `gorm:"column:CLOSE_INDEVICE_DATE_C;type:DATE;" json:"CLOSE_INDEVICE_DATE_C" db:"CLOSE_INDEVICE_DATE_C"`
	//[64] IS_REVIRSE_C                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_REVIRSE_C *bool `gorm:"column:IS_REVIRSE_C;type:BIT;" json:"IS_REVIRSE_C" db:"IS_REVIRSE_C"`
	//[65] IS_REVIRSE_R                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_REVIRSE_R *bool `gorm:"column:IS_REVIRSE_R;type:BIT;" json:"IS_REVIRSE_R" db:"IS_REVIRSE_R"`
	//[66] FLAGE1                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE1 *bool `gorm:"column:FLAGE1;type:BIT;" json:"FLAGE1" db:"FLAGE1"`
	//[67] FLAGE2                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE2 *bool `gorm:"column:FLAGE2;type:BIT;" json:"FLAGE2" db:"FLAGE2"`
	//[68] FLAGE3                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE3 *bool `gorm:"column:FLAGE3;type:BIT;" json:"FLAGE3" db:"FLAGE3"`
	//[69] NOTE_C                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	NOTE_C *string `gorm:"column:NOTE_C;type:NVARCHAR;size:100;" json:"NOTE_C" db:"NOTE_C"`
	//[70] NOTE_R                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	NOTE_R *string `gorm:"column:NOTE_R;type:NVARCHAR;size:100;" json:"NOTE_R" db:"NOTE_R"`
	//[71] NOTE_DATE_C                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	NOTE_DATE_C *time.Time `gorm:"column:NOTE_DATE_C;type:DATETIME;" json:"NOTE_DATE_C" db:"NOTE_DATE_C"`
	//[72] NOTE_DATE_R                                    DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	NOTE_DATE_R *time.Time `gorm:"column:NOTE_DATE_R;type:DATETIME;" json:"NOTE_DATE_R" db:"NOTE_DATE_R"`
	//[73] IS_NOTIFIED                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_NOTIFIED *bool `gorm:"column:IS_NOTIFIED;type:BIT;" json:"IS_NOTIFIED" db:"IS_NOTIFIED"`
	//[74] NOTIFICATION_DATE                              DATE                 null: true   primary: false  isArray: false  auto: false  col: DATE            len: -1      default: []
	NOTIFICATION_DATE *time.Time `gorm:"column:NOTIFICATION_DATE;type:DATE;" json:"NOTIFICATION_DATE" db:"NOTIFICATION_DATE"`
	//[75] LAT_MIN                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_MIN *float64 `gorm:"column:LAT_MIN;type:FLOAT;" json:"LAT_MIN" db:"LAT_MIN"`
	//[76] LAT_MAX                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LAT_MAX *float64 `gorm:"column:LAT_MAX;type:FLOAT;" json:"LAT_MAX" db:"LAT_MAX"`
	//[77] LNG_MIN                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_MIN *float64 `gorm:"column:LNG_MIN;type:FLOAT;" json:"LNG_MIN" db:"LNG_MIN"`
	//[78] LNG_MAX                                        FLOAT                null: true   primary: false  isArray: false  auto: false  col: FLOAT           len: -1      default: []
	LNG_MAX *float64 `gorm:"column:LNG_MAX;type:FLOAT;" json:"LNG_MAX" db:"LNG_MAX"`
	//[79] APPLY_REF                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	APPLY_REF *int64 `gorm:"column:APPLY_REF;type:INT;" json:"APPLY_REF" db:"APPLY_REF"`
	//[80] CLIENT_VERSION                                 NCHAR(40)            null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 40      default: []
	CLIENT_VERSION *string `gorm:"column:CLIENT_VERSION;type:NCHAR;size:40;" json:"CLIENT_VERSION" db:"CLIENT_VERSION"`
	//[81] NEXT_SEQ                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	NEXT_SEQ *int64 `gorm:"column:NEXT_SEQ;type:INT;" json:"NEXT_SEQ" db:"NEXT_SEQ"`
	//[82] CLOSE_INDEVICE_BY_C                            NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CLOSE_INDEVICE_BY_C *string `gorm:"column:CLOSE_INDEVICE_BY_C;type:NVARCHAR;size:100;" json:"CLOSE_INDEVICE_BY_C" db:"CLOSE_INDEVICE_BY_C"`
	//[83] CLOSE_INDEVICE_BY_R                            NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	CLOSE_INDEVICE_BY_R *string `gorm:"column:CLOSE_INDEVICE_BY_R;type:NVARCHAR;size:100;" json:"CLOSE_INDEVICE_BY_R" db:"CLOSE_INDEVICE_BY_R"`
	//[84] GARD                                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	GARD *bool `gorm:"column:GARD;type:BIT;" json:"GARD" db:"GARD"`
	//[85] GARD_OK                                        BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	GARD_OK *bool `gorm:"column:GARD_OK;type:BIT;" json:"GARD_OK" db:"GARD_OK"`
	//[86] EDAMS_CYCLE_C                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_CYCLE_C *bool `gorm:"column:EDAMS_CYCLE_C;type:BIT;" json:"EDAMS_CYCLE_C" db:"EDAMS_CYCLE_C"`
	//[87] EDAMS_CYCLE_R                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_CYCLE_R *bool `gorm:"column:EDAMS_CYCLE_R;type:BIT;" json:"EDAMS_CYCLE_R" db:"EDAMS_CYCLE_R"`
	//[88] ERROR_MESSAGE                                  NVARCHAR(512)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 512     default: []
	ERROR_MESSAGE *string `gorm:"column:ERROR_MESSAGE;type:NVARCHAR;size:512;" json:"ERROR_MESSAGE" db:"ERROR_MESSAGE"`
	//[89] SECTOR_CODE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SECTOR_CODE *int64 `gorm:"column:SECTOR_CODE;type:INT;" json:"SECTOR_CODE" db:"SECTOR_CODE"`
	//[90] IS_BILL_PREPEARED                              BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_BILL_PREPEARED *bool `gorm:"column:IS_BILL_PREPEARED;type:BIT;" json:"IS_BILL_PREPEARED" db:"IS_BILL_PREPEARED"`
	//[91] PREV_OWNER_R                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PREV_OWNER_R *int64 `gorm:"column:PREV_OWNER_R;type:INT;" json:"PREV_OWNER_R" db:"PREV_OWNER_R"`
	//[92] TRANSFEER_DATE_R                               DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	TRANSFEER_DATE_R *time.Time `gorm:"column:TRANSFEER_DATE_R;type:DATETIME;" json:"TRANSFEER_DATE_R" db:"TRANSFEER_DATE_R"`
	//[93] PREV_OWNER_C                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PREV_OWNER_C *int64 `gorm:"column:PREV_OWNER_C;type:INT;" json:"PREV_OWNER_C" db:"PREV_OWNER_C"`
	//[94] TRANSFEER_DATE_C                               DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	TRANSFEER_DATE_C *time.Time `gorm:"column:TRANSFEER_DATE_C;type:DATETIME;" json:"TRANSFEER_DATE_C" db:"TRANSFEER_DATE_C"`
	//[95] ALLOW_FAWRY                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ALLOW_FAWRY *int64 `gorm:"column:ALLOW_FAWRY;type:INT;" json:"ALLOW_FAWRY" db:"ALLOW_FAWRY"`
	//[96] FAWRY_OPEN_BY                                  VARCHAR(128)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
	FAWRY_OPEN_BY *string `gorm:"column:FAWRY_OPEN_BY;type:VARCHAR;size:128;" json:"FAWRY_OPEN_BY" db:"FAWRY_OPEN_BY"`
	//[97] FAWRY_CLOSE_BY                                 VARCHAR(128)         null: true   primary: false  isArray: false  auto: false  col: VARCHAR         len: 128     default: []
	FAWRY_CLOSE_BY *string `gorm:"column:FAWRY_CLOSE_BY;type:VARCHAR;size:128;" json:"FAWRY_CLOSE_BY" db:"FAWRY_CLOSE_BY"`
	//[98] FAWRY_OPEN_DATE                                DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	FAWRY_OPEN_DATE *time.Time `gorm:"column:FAWRY_OPEN_DATE;type:DATETIME;" json:"FAWRY_OPEN_DATE" db:"FAWRY_OPEN_DATE"`
	//[99] FAWRY_CLOSE_DATE                               DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	FAWRY_CLOSE_DATE *time.Time `gorm:"column:FAWRY_CLOSE_DATE;type:DATETIME;" json:"FAWRY_CLOSE_DATE" db:"FAWRY_CLOSE_DATE"`
}

// TableName sets the insert table name for this struct type
func (h *HH_BCYC) TableName() string {
	return "HH_BCYC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HH_BCYC) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HH_BCYC) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HH_BCYC) Validate(action Action) error {
	return nil
}

func (h *HH_BCYC) IsCycleComplatedC() bool {
	if h.ISCYCLE_COMPLETED_C == nil || *h.ISCYCLE_COMPLETED_C == 0 {
		return false
	}
	return true
}

func (h *HH_BCYC) IsCycleComplatedR() bool {
	if h.ISCYCLE_COMPLETED_R == nil || *h.ISCYCLE_COMPLETED_R == 0 {
		return false
	}
	return true
}
