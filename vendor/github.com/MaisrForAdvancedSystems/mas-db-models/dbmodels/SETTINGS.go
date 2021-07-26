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

type SETTINGS struct {
	//[ 0] KEY_WORD                                       NVARCHAR(200)        null: false  primary: true   isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	KEY_WORD string `gorm:"primary_key;column:KEY_WORD;type:NVARCHAR;size:200;" json:"KEY_WORD" db:"KEY_WORD"`
	//[ 1] KEY_VALUE                                      NVARCHAR(200)        null: false  primary: false  isArray: false  auto: false  col: NVARCHAR        len: 200     default: []
	KEY_VALUE string `gorm:"column:KEY_VALUE;type:NVARCHAR;size:200;" json:"KEY_VALUE" db:"KEY_VALUE"`
	//[ 2] DESCRIPTION                                    NCHAR(510)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 510     default: []
	DESCRIPTION *string `gorm:"column:DESCRIPTION;type:NCHAR;size:510;" json:"DESCRIPTION" db:"DESCRIPTION"`
}

// TableName sets the insert table name for this struct type
func (s *SETTINGS) TableName() string {
	return "SETTINGS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SETTINGS) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SETTINGS) Prepare() {
}

type SettingKey string

const SETTING_INVOICE_REPORT SettingKey = "INVOICE_REPORT"
const SETTING_INFROM_REPORT SettingKey = "INFROM_REPORT"
const SETTING_PARTIAL_REPORT SettingKey = "PARTIAL_REPORT"
const SETTING_LIST_INVOICE_REPORT SettingKey = "LIST_INVOICE_REPORT"
const SETTING_ENABLE_WEBSOCKETS SettingKey = "ENABLE_WEBSOCKETS"
const SETTING_ENABLE_PARTIAL_COLLECTION SettingKey = "ENABLE_PARTIAL_COLLECTION"
const SETTING_STABLE_LOCATION_TIMEOUT SettingKey = "STABLE_LOCATION_TIMEOUT"
const SETTING_ENABLE_HH_PRINTING SettingKey = "ENABLE_HH_PRINTING"
const SETTING_DEVICE_LOCATION_INTERVAL SettingKey = "DEVICE_LOCATION_INTERVAL"
const SETTING_PING_INTERVAL SettingKey = "PING_INTERVAL"
const SETTING_MIN_CHARGE_VALUE SettingKey = "MIN_CHARGE_VALUE"
const SETTING_MIN_CHARGE_RATIO SettingKey = "MIN_CHARGE_RATIO"
const SETTING_MIN_CHARGE_OP SettingKey = "MIN_CHARGE_OP"
const SETTING_DISABLE_COLLECTION SettingKey = "DisableCollection"
const SETTING_MAX_CHARGE SettingKey = "MAX_CHARGE"
const SETTING_SURNAME_MAX_SIZE SettingKey = "SURNAME_MAX_SIZE"
const SETTING_UA_ADRESS_MAX_SIZE SettingKey = "UA_ADRESS_MAX_SIZE"
const SETTING_DISTRIBUTE_STATMENT_CTYPES SettingKey = "DISTRIBUTE_STATMENT_CTYPES"
const SETTING_MASGATE_API_KEY SettingKey = "MASGATE_API_KEY"
const SETTING_MASGATE_SERVER SettingKey = "MASGATE_SERVER"
const SETTING_MinCLBlnce SettingKey = "MinCLBlnce"
const SETTING_DisableCollection SettingKey = "DisableCollection"
const SETTING_CheckBefourHHClose SettingKey = "CheckBefourHHClose"
const SETTING_MUST_GARD SettingKey = "MUST_GARD"
const SETTING_MIN_CHARGE_CURCHARGE SettingKey = "MIN_CHARGE_CURCHARGE"
const SETTING_MAX_ROWS_PER_HH_REQUEST SettingKey = "MAX_ROWS_PER_HH_REQUEST"
const SETTING_GARD_PAYMENT_NO_STYLE SettingKey = "GARD_PAYMENT_NO_STYLE"
const SETTING_MATCH_RECEIPT_PAYMENT SettingKey = "MATCH_RECEIPT_PAYMENT"
const SETTING_PULL_COLLECTED_STATMENTS_BILLING SettingKey = "PULL_COL_STM_BILLING"
const SETTING_TRANSFEER_METER_CONDITION SettingKey = "TRANSFEER_METER_CONDITION"
const SETTING_ENABLE_COMPRESSION SettingKey = "ENABLE_COMPRESSION"                           //0 disable ,1 enable server only ,2 enable server and mobile
const SETTING_COVERT_NEGATIVE_BALANCE_TOZERO SettingKey = "CONVERT_NEGATIVE_CL_BLNCE_TOZERO" //0 disable ,1 enable server only
const SETTING_MARKETING SettingKey = "MARKETING"                                             //0 disable ,1 enable server only
const SETTING_EDAMS_MRECEIPT SettingKey = "EDAMS_MRECEIPT"                                   //0 disable ,1 enable server only
const SETTING_COLLECTION_DIRECTION SettingKey = "COLLECTION_DIRECTION"                       //0 down,1 up
const SETTING_MARKETING_TRANS SettingKey = "MARKETING_TRANS"
const SETTING_EDAMS_PARTIAL_TRANS_STYPE SettingKey = "EDAMS_PARTIAL_TRANS_STYPE"
const SETTING_READING_DATE_MODE SettingKey = "READING_DATE_MODE"
const SETTING_FIX_EDAMS_PR_DATE1 SettingKey = "FIX_EDAMS_PR_DATE1"
const SETTING_FAWRY_IP SettingKey = "FAWRY_IP"
const SETTING_COMPANY_NAME SettingKey = "COMPANY_NAME"
const SETTING_COMPANY_LOGO SettingKey = "COMPANY_LOGO"
const SETTING_HH_SECTOR_NAME SettingKey = "HH_SECTOR_NAME"
const SETTING_CUSTKEY_TITLE SettingKey = "CUSTKEY_TITLE"
const SETTING_PROPREF_TITLE SettingKey = "PROPREF_TITLE"
const SETTING_OLDKEY_TITLE SettingKey = "OLDKEY_TITLE"
const SETTING_HHCONTROLLER_VERSION SettingKey = "HHCONTROLLER_VERSION"
