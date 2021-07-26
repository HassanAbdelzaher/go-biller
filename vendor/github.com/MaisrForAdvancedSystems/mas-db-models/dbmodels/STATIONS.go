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

type STATIONS struct {
	//[ 0] STATION_NO                                     INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO int32 `gorm:"primary_key;column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 1] NAME                                           NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	NAME *string `gorm:"column:NAME;type:NVARCHAR;size:200;" json:"NAME" db:"NAME"`
	//[ 2] DESCRIPTION                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NVARCHAR;size:200;" json:"DESCRIPTION" db:"DESCRIPTION"`
	//[ 3] ISCURRENT                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	ISCURRENT *int64 `gorm:"column:ISCURRENT;type:INT;" json:"ISCURRENT" db:"ISCURRENT"`
	//[ 4] DB_STANDALONE                                  INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DB_STANDALONE *int64 `gorm:"column:DB_STANDALONE;type:INT;" json:"DB_STANDALONE" db:"DB_STANDALONE"`
	//[ 5] DB_ACCESS                                      INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DB_ACCESS *int64 `gorm:"column:DB_ACCESS;type:INT;" json:"DB_ACCESS" db:"DB_ACCESS"`
	//[ 6] IS_HEADQUARTERS                                INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_HEADQUARTERS *int64 `gorm:"column:IS_HEADQUARTERS;type:INT;" json:"IS_HEADQUARTERS" db:"IS_HEADQUARTERS"`
	//[ 7] IS_MAINTSITE                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_MAINTSITE *int64 `gorm:"column:IS_MAINTSITE;type:INT;" json:"IS_MAINTSITE" db:"IS_MAINTSITE"`
	//[ 8] IS_MRREADING                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_MRREADING *int64 `gorm:"column:IS_MRREADING;type:INT;" json:"IS_MRREADING" db:"IS_MRREADING"`
	//[ 9] IS_CSERVICES                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_CSERVICES *int64 `gorm:"column:IS_CSERVICES;type:INT;" json:"IS_CSERVICES" db:"IS_CSERVICES"`
	//[10] IS_RECEIPTING                                  INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	IS_RECEIPTING *int64 `gorm:"column:IS_RECEIPTING;type:INT;" json:"IS_RECEIPTING" db:"IS_RECEIPTING"`
	//[11] TEL_WORK                                       NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TEL_WORK *string `gorm:"column:TEL_WORK;type:NVARCHAR;size:200;" json:"TEL_WORK" db:"TEL_WORK"`
	//[12] TEL_FAX                                        NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TEL_FAX *string `gorm:"column:TEL_FAX;type:NVARCHAR;size:200;" json:"TEL_FAX" db:"TEL_FAX"`
	//[13] TEL_EMAIL                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TEL_EMAIL *string `gorm:"column:TEL_EMAIL;type:NVARCHAR;size:200;" json:"TEL_EMAIL" db:"TEL_EMAIL"`
	//[14] TEL_AHOURS                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TEL_AHOURS *string `gorm:"column:TEL_AHOURS;type:NVARCHAR;size:200;" json:"TEL_AHOURS" db:"TEL_AHOURS"`
	//[15] TEL_COMPL                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	TEL_COMPL *string `gorm:"column:TEL_COMPL;type:NVARCHAR;size:200;" json:"TEL_COMPL" db:"TEL_COMPL"`
	//[16] LN_ADDRESS1                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	LN_ADDRESS1 *string `gorm:"column:LN_ADDRESS1;type:NVARCHAR;size:200;" json:"LN_ADDRESS1" db:"LN_ADDRESS1"`
	//[17] LN_ADDRESS2                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	LN_ADDRESS2 *string `gorm:"column:LN_ADDRESS2;type:NVARCHAR;size:200;" json:"LN_ADDRESS2" db:"LN_ADDRESS2"`
	//[18] LN_ADDRESS3                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	LN_ADDRESS3 *string `gorm:"column:LN_ADDRESS3;type:NVARCHAR;size:200;" json:"LN_ADDRESS3" db:"LN_ADDRESS3"`
	//[19] LN_CITY                                        NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	LN_CITY *string `gorm:"column:LN_CITY;type:NVARCHAR;size:200;" json:"LN_CITY" db:"LN_CITY"`
	//[20] LN_POSTAL                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	LN_POSTAL *string `gorm:"column:LN_POSTAL;type:NVARCHAR;size:200;" json:"LN_POSTAL" db:"LN_POSTAL"`
	//[21] WORKING_MON                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_MON *int64 `gorm:"column:WORKING_MON;type:INT;" json:"WORKING_MON" db:"WORKING_MON"`
	//[22] WORKING_TUE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_TUE *int64 `gorm:"column:WORKING_TUE;type:INT;" json:"WORKING_TUE" db:"WORKING_TUE"`
	//[23] WORKING_WED                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_WED *int64 `gorm:"column:WORKING_WED;type:INT;" json:"WORKING_WED" db:"WORKING_WED"`
	//[24] WORKING_THU                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_THU *int64 `gorm:"column:WORKING_THU;type:INT;" json:"WORKING_THU" db:"WORKING_THU"`
	//[25] WORKING_FRI                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_FRI *int64 `gorm:"column:WORKING_FRI;type:INT;" json:"WORKING_FRI" db:"WORKING_FRI"`
	//[26] WORKING_SAT                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_SAT *int64 `gorm:"column:WORKING_SAT;type:INT;" json:"WORKING_SAT" db:"WORKING_SAT"`
	//[27] WORKING_SUN                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	WORKING_SUN *int64 `gorm:"column:WORKING_SUN;type:INT;" json:"WORKING_SUN" db:"WORKING_SUN"`
	//[28] OVERTIME_MON                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_MON *int64 `gorm:"column:OVERTIME_MON;type:INT;" json:"OVERTIME_MON" db:"OVERTIME_MON"`
	//[29] OVERTIME_TUE                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_TUE *int64 `gorm:"column:OVERTIME_TUE;type:INT;" json:"OVERTIME_TUE" db:"OVERTIME_TUE"`
	//[30] OVERTIME_WED                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_WED *int64 `gorm:"column:OVERTIME_WED;type:INT;" json:"OVERTIME_WED" db:"OVERTIME_WED"`
	//[31] OVERTIME_THU                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_THU *int64 `gorm:"column:OVERTIME_THU;type:INT;" json:"OVERTIME_THU" db:"OVERTIME_THU"`
	//[32] OVERTIME_FRI                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_FRI *int64 `gorm:"column:OVERTIME_FRI;type:INT;" json:"OVERTIME_FRI" db:"OVERTIME_FRI"`
	//[33] OVERTIME_SAT                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_SAT *int64 `gorm:"column:OVERTIME_SAT;type:INT;" json:"OVERTIME_SAT" db:"OVERTIME_SAT"`
	//[34] OVERTIME_SUN                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	OVERTIME_SUN *int64 `gorm:"column:OVERTIME_SUN;type:INT;" json:"OVERTIME_SUN" db:"OVERTIME_SUN"`
	//[35] SITE_MEMO                                      NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	SITE_MEMO *string `gorm:"column:SITE_MEMO;type:NVARCHAR;size:200;" json:"SITE_MEMO" db:"SITE_MEMO"`
	//[36] STAMP_DATE                                     DATETIME             null: true   primary: false  isArray: false  auto: false  col: DATETIME        len: -1      default: []
	STAMP_DATE *time.Time `gorm:"column:STAMP_DATE;type:DATETIME;" json:"STAMP_DATE" db:"STAMP_DATE"`
	//[37] STAMP_USER                                     NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	STAMP_USER *string `gorm:"column:STAMP_USER;type:NVARCHAR;size:200;" json:"STAMP_USER" db:"STAMP_USER"`
	//[38] MAINTSITE_CODE                                 INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	MAINTSITE_CODE *int64 `gorm:"column:MAINTSITE_CODE;type:INT;" json:"MAINTSITE_CODE" db:"MAINTSITE_CODE"`
	//[39] DEFAULTWAREHOUSE                               INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DEFAULTWAREHOUSE *int64 `gorm:"column:DEFAULTWAREHOUSE;type:INT;" json:"DEFAULTWAREHOUSE" db:"DEFAULTWAREHOUSE"`
	//[40] DEFAULTDEPOT                                   INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	DEFAULTDEPOT *int64 `gorm:"column:DEFAULTDEPOT;type:INT;" json:"DEFAULTDEPOT" db:"DEFAULTDEPOT"`
	//[41] WAREHOUSESETTING                               SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	WAREHOUSESETTING *int64 `gorm:"column:WAREHOUSESETTING;type:SMALLINT;" json:"WAREHOUSESETTING" db:"WAREHOUSESETTING"`
	//[42] DEPOTDSETTING                                  SMALLINT             null: true   primary: false  isArray: false  auto: false  col: SMALLINT        len: -1      default: []
	DEPOTDSETTING *int64 `gorm:"column:DEPOTDSETTING;type:SMALLINT;" json:"DEPOTDSETTING" db:"DEPOTDSETTING"`
	//[43] IP_ADDRESS                                     NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	IP_ADDRESS *string `gorm:"column:IP_ADDRESS;type:NVARCHAR;size:40;" json:"IP_ADDRESS" db:"IP_ADDRESS"`
	//[44] PORT_NO                                        INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	PORT_NO *int64 `gorm:"column:PORT_NO;type:INT;" json:"PORT_NO" db:"PORT_NO"`
	//[45] SECTOR_CODE                                    INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	SECTOR_CODE *int64 `gorm:"column:SECTOR_CODE;type:INT;" json:"SECTOR_CODE" db:"SECTOR_CODE"`
	//[46] SECTOR_NAME                                    NVARCHAR(200)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	SECTOR_NAME *string `gorm:"column:SECTOR_NAME;type:NVARCHAR;size:200;" json:"SECTOR_NAME" db:"SECTOR_NAME"`
}

// TableName sets the insert table name for this struct type
func (s *STATIONS) TableName() string {
	return "STATIONS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *STATIONS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *STATIONS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *STATIONS) Validate(action Action) error {
	return nil
}
