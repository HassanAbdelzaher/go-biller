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

type USERS struct {
	//[ 0] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 1] STATION_NO                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	STATION_NO *int64 `gorm:"column:STATION_NO;type:INT;" json:"STATION_NO" db:"STATION_NO"`
	//[ 2] FULL_NAME                                      NVARCHAR(300)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 300     default: []
	FULL_NAME *string `gorm:"column:FULL_NAME;type:NVARCHAR;size:300;" json:"FULL_NAME" db:"FULL_NAME"`
	//[ 3] USER_NAME                                      NVARCHAR(40)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 40      default: []
	USER_NAME *string `gorm:"column:USER_NAME;type:NVARCHAR;size:40;" json:"USER_NAME" db:"USER_NAME"`
	//[ 4] PASSWORD                                       NVARCHAR(120)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 120     default: []
	PASSWORD string `gorm:"column:PASSWORD;type:NVARCHAR;size:120;" json:"PASSWORD" db:"PASSWORD"`
	//[ 5] SYS_ADMIN                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	SYS_ADMIN *bool `gorm:"column:SYS_ADMIN;type:BIT;" json:"SYS_ADMIN" db:"SYS_ADMIN"`
	//[ 6] READING_ADMIN                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	READING_ADMIN *bool `gorm:"column:READING_ADMIN;type:BIT;" json:"READING_ADMIN" db:"READING_ADMIN"`
	//[ 7] COLLECTION_ADMIN                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	COLLECTION_ADMIN *bool `gorm:"column:COLLECTION_ADMIN;type:BIT;" json:"COLLECTION_ADMIN" db:"COLLECTION_ADMIN"`
	//[ 8] REPORING_ADMIN                                 BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	REPORING_ADMIN *bool `gorm:"column:REPORING_ADMIN;type:BIT;" json:"REPORING_ADMIN" db:"REPORING_ADMIN"`
	//[ 9] HH_MONITOR                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	HH_MONITOR *bool `gorm:"column:HH_MONITOR;type:BIT;" json:"HH_MONITOR" db:"HH_MONITOR"`
	//[10] USER_MANAGEMENT                                BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	USER_MANAGEMENT *bool `gorm:"column:USER_MANAGEMENT;type:BIT;" json:"USER_MANAGEMENT" db:"USER_MANAGEMENT"`
	//[11] DATA_ADMIN                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	DATA_ADMIN *bool `gorm:"column:DATA_ADMIN;type:BIT;" json:"DATA_ADMIN" db:"DATA_ADMIN"`
	//[12] SYSTEM_MNTINANCE                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	SYSTEM_MNTINANCE *bool `gorm:"column:SYSTEM_MNTINANCE;type:BIT;" json:"SYSTEM_MNTINANCE" db:"SYSTEM_MNTINANCE"`
	//[13] NID                                            NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	NID *string `gorm:"column:NID;type:NVARCHAR;size:100;" json:"NID" db:"NID"`
	//[14] EMAIL                                          NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	EMAIL *string `gorm:"column:EMAIL;type:NVARCHAR;size:100;" json:"EMAIL" db:"EMAIL"`
	//[15] ALLOW_COLLECTION                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_COLLECTION *bool `gorm:"column:ALLOW_COLLECTION;type:BIT;" json:"ALLOW_COLLECTION" db:"ALLOW_COLLECTION"`
	//[16] ALLOW_MODIFY_READING                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_MODIFY_READING *bool `gorm:"column:ALLOW_MODIFY_READING;type:BIT;" json:"ALLOW_MODIFY_READING" db:"ALLOW_MODIFY_READING"`
	//[17] ALLOW_DEPOSIT                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_DEPOSIT *bool `gorm:"column:ALLOW_DEPOSIT;type:BIT;" json:"ALLOW_DEPOSIT" db:"ALLOW_DEPOSIT"`
	//[18] ALOOW_CANCEL                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALOOW_CANCEL *bool `gorm:"column:ALOOW_CANCEL;type:BIT;" json:"ALOOW_CANCEL" db:"ALOOW_CANCEL"`
	//[19] ALLOW_POST_C                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_POST_C *bool `gorm:"column:ALLOW_POST_C;type:BIT;" json:"ALLOW_POST_C" db:"ALLOW_POST_C"`
	//[20] ALLOW_POST_R                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_POST_R *bool `gorm:"column:ALLOW_POST_R;type:BIT;" json:"ALLOW_POST_R" db:"ALLOW_POST_R"`
	//[21] VALUE1                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	VALUE1 *string `gorm:"column:VALUE1;type:NVARCHAR;size:100;" json:"VALUE1" db:"VALUE1"`
	//[22] VALUE2                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	VALUE2 *string `gorm:"column:VALUE2;type:NVARCHAR;size:100;" json:"VALUE2" db:"VALUE2"`
	//[23] FLAGE1                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE1 *bool `gorm:"column:FLAGE1;type:BIT;" json:"FLAGE1" db:"FLAGE1"`
	//[24] FLAGE2                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE2 *bool `gorm:"column:FLAGE2;type:BIT;" json:"FLAGE2" db:"FLAGE2"`
	//[25] FLAGE3                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE3 *bool `gorm:"column:FLAGE3;type:BIT;" json:"FLAGE3" db:"FLAGE3"`
	//[26] FLAGE4                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE4 *bool `gorm:"column:FLAGE4;type:BIT;" json:"FLAGE4" db:"FLAGE4"`
	//[27] FLAGE5                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE5 *bool `gorm:"column:FLAGE5;type:BIT;" json:"FLAGE5" db:"FLAGE5"`
	//[28] FLAGE6                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE6 *bool `gorm:"column:FLAGE6;type:BIT;" json:"FLAGE6" db:"FLAGE6"`
	//[29] FLAGE7                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE7 *bool `gorm:"column:FLAGE7;type:BIT;" json:"FLAGE7" db:"FLAGE7"`
	//[30] FLAGE8                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE8 *bool `gorm:"column:FLAGE8;type:BIT;" json:"FLAGE8" db:"FLAGE8"`
	//[31] FLAGE9                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE9 *bool `gorm:"column:FLAGE9;type:BIT;" json:"FLAGE9" db:"FLAGE9"`
	//[32] FLAGE10                                        BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE10 *bool `gorm:"column:FLAGE10;type:BIT;" json:"FLAGE10" db:"FLAGE10"`
	//[33] IS_ENC                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_ENC *bool `gorm:"column:IS_ENC;type:BIT;" json:"IS_ENC" db:"IS_ENC"`
	//[34] ALLOW_COMPENSATION                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_COMPENSATION *bool `gorm:"column:ALLOW_COMPENSATION;type:BIT;" json:"ALLOW_COMPENSATION" db:"ALLOW_COMPENSATION"`
	//[35] ALLOW_CANCEL                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_CANCEL *bool `gorm:"column:ALLOW_CANCEL;type:BIT;" json:"ALLOW_CANCEL" db:"ALLOW_CANCEL"`
	//[36] MAP_READING                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_READING *bool `gorm:"column:MAP_READING;type:BIT;" json:"MAP_READING" db:"MAP_READING"`
	//[37] MAP_COLLECTION                                 BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_COLLECTION *bool `gorm:"column:MAP_COLLECTION;type:BIT;" json:"MAP_COLLECTION" db:"MAP_COLLECTION"`
	//[38] MAP_LOCATION                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_LOCATION *bool `gorm:"column:MAP_LOCATION;type:BIT;" json:"MAP_LOCATION" db:"MAP_LOCATION"`
	//[39] MAP_PATH                                       BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_PATH *bool `gorm:"column:MAP_PATH;type:BIT;" json:"MAP_PATH" db:"MAP_PATH"`
	//[40] MAP_COMPLAINTS                                 BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_COMPLAINTS *bool `gorm:"column:MAP_COMPLAINTS;type:BIT;" json:"MAP_COMPLAINTS" db:"MAP_COMPLAINTS"`
	//[41] MAP_BAD_CONN                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MAP_BAD_CONN *bool `gorm:"column:MAP_BAD_CONN;type:BIT;" json:"MAP_BAD_CONN" db:"MAP_BAD_CONN"`
	//[42] WALK_ARRANGE_POST                              BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	WALK_ARRANGE_POST *bool `gorm:"column:WALK_ARRANGE_POST;type:BIT;" json:"WALK_ARRANGE_POST" db:"WALK_ARRANGE_POST"`
	//[43] UPDATE_CUSTOMER_LOCATION                       BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	UPDATE_CUSTOMER_LOCATION *bool `gorm:"column:UPDATE_CUSTOMER_LOCATION;type:BIT;" json:"UPDATE_CUSTOMER_LOCATION" db:"UPDATE_CUSTOMER_LOCATION"`
	//[44] ALLOW_CANCEL_DEP_COL                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_CANCEL_DEP_COL *bool `gorm:"column:ALLOW_CANCEL_DEP_COL;type:BIT;" json:"ALLOW_CANCEL_DEP_COL" db:"ALLOW_CANCEL_DEP_COL"`
	//[45] ALLOW_RECOL_CANCEL                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_RECOL_CANCEL *bool `gorm:"column:ALLOW_RECOL_CANCEL;type:BIT;" json:"ALLOW_RECOL_CANCEL" db:"ALLOW_RECOL_CANCEL"`
	//[46] ALLOW_MAPS                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_MAPS *bool `gorm:"column:ALLOW_MAPS;type:BIT;" json:"ALLOW_MAPS" db:"ALLOW_MAPS"`
	//[47] ALLOW_LOADDATA_C                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_LOADDATA_C *bool `gorm:"column:ALLOW_LOADDATA_C;type:BIT;" json:"ALLOW_LOADDATA_C" db:"ALLOW_LOADDATA_C"`
	//[48] ALLOW_LOADDATA_R                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_LOADDATA_R *bool `gorm:"column:ALLOW_LOADDATA_R;type:BIT;" json:"ALLOW_LOADDATA_R" db:"ALLOW_LOADDATA_R"`
	//[49] IS_CASHIER                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_CASHIER *bool `gorm:"column:IS_CASHIER;type:BIT;" json:"IS_CASHIER" db:"IS_CASHIER"`
	//[50] CASHIER_ID                                     INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	CASHIER_ID *int64 `gorm:"column:CASHIER_ID;type:INT;" json:"CASHIER_ID" db:"CASHIER_ID"`
	//[51] IS_DESK_CAHIER                                 BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	IS_DESK_CAHIER *bool `gorm:"column:IS_DESK_CAHIER;type:BIT;" json:"IS_DESK_CAHIER" db:"IS_DESK_CAHIER"`
	//[52] EDAMS_CLEAR_READINGS                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_CLEAR_READINGS *bool `gorm:"column:EDAMS_CLEAR_READINGS;type:BIT;" json:"EDAMS_CLEAR_READINGS" db:"EDAMS_CLEAR_READINGS"`
	//[53] EDAMS_RECALC                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_RECALC *bool `gorm:"column:EDAMS_RECALC;type:BIT;" json:"EDAMS_RECALC" db:"EDAMS_RECALC"`
	//[54] EDAMS_CS                                       BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_CS *bool `gorm:"column:EDAMS_CS;type:BIT;" json:"EDAMS_CS" db:"EDAMS_CS"`
	//[55] ALLOW_METER_MODIFY                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_METER_MODIFY *bool `gorm:"column:ALLOW_METER_MODIFY;type:BIT;" json:"ALLOW_METER_MODIFY" db:"ALLOW_METER_MODIFY"`
	//[56] ALLOW_GARD                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_GARD *bool `gorm:"column:ALLOW_GARD;type:BIT;" json:"ALLOW_GARD" db:"ALLOW_GARD"`
	//[57] RESEND_COLLECTION                              BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	RESEND_COLLECTION *bool `gorm:"column:RESEND_COLLECTION;type:BIT;" json:"RESEND_COLLECTION" db:"RESEND_COLLECTION"`
	//[58] RESEND_READING                                 BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	RESEND_READING *bool `gorm:"column:RESEND_READING;type:BIT;" json:"RESEND_READING" db:"RESEND_READING"`
	//[59] CLOSE_COLLECTION                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CLOSE_COLLECTION *bool `gorm:"column:CLOSE_COLLECTION;type:BIT;" json:"CLOSE_COLLECTION" db:"CLOSE_COLLECTION"`
	//[60] CLOSE_READING                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CLOSE_READING *bool `gorm:"column:CLOSE_READING;type:BIT;" json:"CLOSE_READING" db:"CLOSE_READING"`
	//[61] COLLECTION_DISCOUNT                            BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	COLLECTION_DISCOUNT *bool `gorm:"column:COLLECTION_DISCOUNT;type:BIT;" json:"COLLECTION_DISCOUNT" db:"COLLECTION_DISCOUNT"`
	//[62] REFRESH_BILLS_DATA                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	REFRESH_BILLS_DATA *bool `gorm:"column:REFRESH_BILLS_DATA;type:BIT;" json:"REFRESH_BILLS_DATA" db:"REFRESH_BILLS_DATA"`
	//[63] CALCULATE_DUE_AMOUNT                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CALCULATE_DUE_AMOUNT *bool `gorm:"column:CALCULATE_DUE_AMOUNT;type:BIT;" json:"CALCULATE_DUE_AMOUNT" db:"CALCULATE_DUE_AMOUNT"`
	//[64] COLLECT_BILL_SINGL_UINT                        BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	COLLECT_BILL_SINGL_UINT *bool `gorm:"column:COLLECT_BILL_SINGL_UINT;type:BIT;" json:"COLLECT_BILL_SINGL_UINT" db:"COLLECT_BILL_SINGL_UINT"`
	//[65] PRINT_PAYMENT_REQUEST                          BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	PRINT_PAYMENT_REQUEST *bool `gorm:"column:PRINT_PAYMENT_REQUEST;type:BIT;" json:"PRINT_PAYMENT_REQUEST" db:"PRINT_PAYMENT_REQUEST"`
	//[66] PARTIAL_COLLECTION                             BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	PARTIAL_COLLECTION *bool `gorm:"column:PARTIAL_COLLECTION;type:BIT;" json:"PARTIAL_COLLECTION" db:"PARTIAL_COLLECTION"`
	//[67] ALONE_PANEL                                    BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALONE_PANEL *bool `gorm:"column:ALONE_PANEL;type:BIT;" json:"ALONE_PANEL" db:"ALONE_PANEL"`
	//[68] REFRESH_CUSTOMER_BILLS                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	REFRESH_CUSTOMER_BILLS *bool `gorm:"column:REFRESH_CUSTOMER_BILLS;type:BIT;" json:"REFRESH_CUSTOMER_BILLS" db:"REFRESH_CUSTOMER_BILLS"`
	//[69] EDAMS_RECALC_NEW                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	EDAMS_RECALC_NEW *bool `gorm:"column:EDAMS_RECALC_NEW;type:BIT;" json:"EDAMS_RECALC_NEW" db:"EDAMS_RECALC_NEW"`
	//[70] ALLOW_FAWRY_OPEN                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_FAWRY_OPEN *bool `gorm:"column:ALLOW_FAWRY_OPEN;type:BIT;" json:"ALLOW_FAWRY_OPEN" db:"ALLOW_FAWRY_OPEN"`
	//[71] ALLOW_FAWRY_CLOSE                              BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_FAWRY_CLOSE *bool `gorm:"column:ALLOW_FAWRY_CLOSE;type:BIT;" json:"ALLOW_FAWRY_CLOSE" db:"ALLOW_FAWRY_CLOSE"`
	//[72] ALLOW_MODIFY_PREV_READING                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_MODIFY_PREV_READING *bool `gorm:"column:ALLOW_MODIFY_PREV_READING;type:BIT;" json:"ALLOW_MODIFY_PREV_READING" db:"ALLOW_MODIFY_PREV_READING"`
	//[73] LIST_COLLECTION                                BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	LIST_COLLECTION *bool `gorm:"column:LIST_COLLECTION;type:BIT;" json:"LIST_COLLECTION" db:"LIST_COLLECTION"`
	//[74] MARKETING                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	MARKETING *bool `gorm:"column:MARKETING;type:BIT;" json:"MARKETING" db:"MARKETING"`
	//[75] PREPEAR_HAFZA                                  BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	PREPEAR_HAFZA *bool `gorm:"column:PREPEAR_HAFZA;type:BIT;" json:"PREPEAR_HAFZA" db:"PREPEAR_HAFZA"`
}

// TableName sets the insert table name for this struct type
func (u *USERS) TableName() string {
	return "USERS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USERS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USERS) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USERS) Validate(action Action) error {
	return nil
}
