package models

import (
	billing "MaisrForAdvancedSystems/go-biller/proto"
	"time"
)

type Transaction struct {
	Code string
	EffectDate time.Time
	BilngDate time.Time
	Amount float64
	TaxAmount float64
	DiscountAmount float64
	Ctype string
	PropRef string
	ServiceType billing.SERVICE_TYPE
	NoUnits int64
}

type MeasuredTransaction struct {
	Transaction
	CrReading *float64
	PrReading *float64
	Consump float64
	ReadType billing.READING_TYPE
	MeterType *string
	MeterRef *string
}
