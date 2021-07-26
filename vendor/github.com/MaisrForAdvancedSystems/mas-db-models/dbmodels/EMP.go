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

type EMP struct {
	//[ 0] BRANCH_ID                                      INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	BRANCH_ID int32 `gorm:"primary_key;column:BRANCH_ID;type:INT;" json:"BRANCH_ID" db:"BRANCH_ID"`
	//[ 1] ID                                             INT                  null: false  primary: true   isArray: false  auto: false  col: INT             len: -1      default: []
	ID int32 `gorm:"primary_key;column:ID;type:INT;" json:"ID" db:"ID"`
	//[ 2] USER_NAME                                      NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	USER_NAME *string `gorm:"column:USER_NAME;type:NVARCHAR;size:510;" json:"USER_NAME" db:"USER_NAME"`
	//[ 3] PASSWORD                                       NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	PASSWORD *string `gorm:"column:PASSWORD;type:NVARCHAR;size:510;" json:"PASSWORD" db:"PASSWORD"`
	//[ 4] FULL_NAME                                      NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	FULL_NAME *string `gorm:"column:FULL_NAME;type:NVARCHAR;size:510;" json:"FULL_NAME" db:"FULL_NAME"`
	//[ 5] ADDRESS                                        NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	ADDRESS *string `gorm:"column:ADDRESS;type:NVARCHAR;size:510;" json:"ADDRESS" db:"ADDRESS"`
	//[ 6] TEL                                            NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	TEL *string `gorm:"column:TEL;type:NVARCHAR;size:510;" json:"TEL" db:"TEL"`
	//[ 7] DEVICE_ID                                      NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	DEVICE_ID *string `gorm:"column:DEVICE_ID;type:NVARCHAR;size:510;" json:"DEVICE_ID" db:"DEVICE_ID"`
	//[ 8] NOTES                                          NVARCHAR(510)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 510     default: []
	NOTES *string `gorm:"column:NOTES;type:NVARCHAR;size:510;" json:"NOTES" db:"NOTES"`
	//[ 9] DISABLED                                       BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	DISABLED *bool `gorm:"column:DISABLED;type:BIT;" json:"DISABLED" db:"DISABLED"`
	//[10] EMP_ID                                         INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_ID *int64 `gorm:"column:EMP_ID;type:INT;" json:"EMP_ID" db:"EMP_ID"`
	//[11] READING                                        BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	READING *bool `gorm:"column:READING;type:BIT;" json:"READING" db:"READING"`
	//[12] COLLECTION                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	COLLECTION *bool `gorm:"column:COLLECTION;type:BIT;" json:"COLLECTION" db:"COLLECTION"`
	//[13] SYS_ADMIN                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	SYS_ADMIN *bool `gorm:"column:SYS_ADMIN;type:BIT;" json:"SYS_ADMIN" db:"SYS_ADMIN"`
	//[14] DATA_ADMIN                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	DATA_ADMIN *bool `gorm:"column:DATA_ADMIN;type:BIT;" json:"DATA_ADMIN" db:"DATA_ADMIN"`
	//[15] DATA_ENTRY                                     BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	DATA_ENTRY *bool `gorm:"column:DATA_ENTRY;type:BIT;" json:"DATA_ENTRY" db:"DATA_ENTRY"`
	//[16] REPORTING_VIEWER                               BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	REPORTING_VIEWER *bool `gorm:"column:REPORTING_VIEWER;type:BIT;" json:"REPORTING_VIEWER" db:"REPORTING_VIEWER"`
	//[17] CASHER                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	CASHER *bool `gorm:"column:CASHER;type:BIT;" json:"CASHER" db:"CASHER"`
	//[18] MOBILE1                                        NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	MOBILE1 *string `gorm:"column:MOBILE1;type:NVARCHAR;size:100;" json:"MOBILE1" db:"MOBILE1"`
	//[19] MOBILE2                                        NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	MOBILE2 *string `gorm:"column:MOBILE2;type:NVARCHAR;size:100;" json:"MOBILE2" db:"MOBILE2"`
	//[20] NID                                            NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	NID *string `gorm:"column:NID;type:NVARCHAR;size:100;" json:"NID" db:"NID"`
	//[21] EMAIL                                          NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	EMAIL *string `gorm:"column:EMAIL;type:NVARCHAR;size:100;" json:"EMAIL" db:"EMAIL"`
	//[22] ALOOW_CANCEL                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALOOW_CANCEL *bool `gorm:"column:ALOOW_CANCEL;type:BIT;" json:"ALOOW_CANCEL" db:"ALOOW_CANCEL"`
	//[23] ALLOW_MODIFY_READING                           BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALLOW_MODIFY_READING *bool `gorm:"column:ALLOW_MODIFY_READING;type:BIT;" json:"ALLOW_MODIFY_READING" db:"ALLOW_MODIFY_READING"`
	//[24] FLAGE1                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE1 *bool `gorm:"column:FLAGE1;type:BIT;" json:"FLAGE1" db:"FLAGE1"`
	//[25] FLAGE2                                         BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	FLAGE2 *bool `gorm:"column:FLAGE2;type:BIT;" json:"FLAGE2" db:"FLAGE2"`
	//[26] VALUE2                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	VALUE2 *string `gorm:"column:VALUE2;type:NVARCHAR;size:100;" json:"VALUE2" db:"VALUE2"`
	//[27] VALUE1                                         NVARCHAR(100)        null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 100     default: []
	VALUE1 *string `gorm:"column:VALUE1;type:NVARCHAR;size:100;" json:"VALUE1" db:"VALUE1"`
	//[28] ALlOW_CANCEL                                   BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	ALlOW_CANCEL *bool `gorm:"column:ALlOW_CANCEL;type:BIT;" json:"ALlOW_CANCEL" db:"ALlOW_CANCEL"`
	//[29] ORIGINAL_STATION                               NCHAR(200)           null: true   primary: false  isArray: false  auto: false  col: NCHAR           len: 200     default: []
	ORIGINAL_STATION *string `gorm:"column:ORIGINAL_STATION;type:NCHAR;size:200;" json:"ORIGINAL_STATION" db:"ORIGINAL_STATION"`
	//[30] EMP_TYPE                                       INT                  null: true   primary: false  isArray: false  auto: false  col: INT             len: -1      default: []
	EMP_TYPE *int64 `gorm:"column:EMP_TYPE;type:INT;" json:"EMP_TYPE" db:"EMP_TYPE"`
	//[31] BILNG_EMP_ID                                   NVARCHAR(60)         null: true   primary: false  isArray: false  auto: false  col: NVARCHAR        len: 60      default: []
	BILNG_EMP_ID *string `gorm:"column:BILNG_EMP_ID;type:NVARCHAR;size:60;" json:"BILNG_EMP_ID" db:"BILNG_EMP_ID"`
	//[32] MARKETING                                      BIT                  null: true   primary: false  isArray: false  auto: false  col: BIT             len: -1      default: []
	//MARKETING *bool `gorm:"column:MARKETING;type:BIT;" json:"MARKETING" db:"MARKETING"`
}

// TableName sets the insert table name for this struct type
func (e *EMP) TableName() string {
	return "EMP"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (e *EMP) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (e *EMP) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (e *EMP) Validate(action Action) error {
	return nil
}
