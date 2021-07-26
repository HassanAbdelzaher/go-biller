package tools

import (
	"strconv"
	"strings"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func ToInt32Pointer(fl int32) *int32 {
	return &fl
}

func IntPtrToInt(fl *int32) int32 {
	if fl == nil {
		return 0
	}
	return *fl
}

func ToDatePointer(fl time.Time) *time.Time {
	return &fl
}

func DatePtrToDate(fl *time.Time) time.Time {
	if fl == nil {
		return time.Now()
	}
	return *fl
}

func ToFloatPointer(fl float64) *float64 {
	return &fl
}
func ToFloat32Pointer(fl float32) *float32 {
	return &fl
}

func FloatPtrToFloat(fl *float64) float64 {
	if fl == nil {
		return 0
	}
	return *fl
}

func IntPtrPtrToInt(fl *int64) int64 {
	if fl == nil {
		return 0
	}
	return *fl
}

func ToIntPointer(int642 int64) *int64 {
	return &int642
}

func ToStringPointer(s string) *string {
	return &s
}

func ToBoolPointer(s bool) *bool {
	return &s
}

func Float64ToString(f *float64) *string {
	if f == nil {
		return nil
	}
	str := strconv.FormatFloat(*f, 'f', -1, 64)
	return &str
}

func Float32ToString(f *float32) *string {
	if f == nil {
		return nil
	}
	f64 := float64(*f)
	str := strconv.FormatFloat(f64, 'f', -1, 32)
	return &str
}

func Int64ToString(f *int64) *string {
	if f == nil {
		return nil
	}
	str := strconv.FormatInt(*f, 10)
	return &str
}

func Int32ToString(f *int32) *string {
	if f == nil {
		return nil
	}
	i := int64(*f)
	str := strconv.FormatInt(i, 10)
	return &str
}

func StringCompare(s1 string, s2 string) bool {
	return strings.TrimSpace(s1) == strings.TrimSpace(s2)
}

func StringComparePointer(s1 *string, s2 *string) bool {
	if s1 == nil || s2 == nil {
		return false
	}
	return StringCompare(*s1, *s2)
}

func StringComparePointer2(s1 string, s2 *string) bool {
	if s2 == nil {
		return false
	}
	return StringCompare(s1, *s2)
}

func BoolToString(b *bool) *string {
	if b == nil {
		return nil
	}
	one := "1"
	zero := "0"
	if *b {
		return &one
	}
	return &zero
}

func Sum(floats ...*float64) float64 {
	var sm float64 = 0
	for id := range floats {
		if floats[id] != nil {
			sm = sm + *floats[id]
		}
	}
	return sm
}

func SumF(floats ...*float64) *float64 {
	var sm =Sum(floats...)
	return &sm
}

func Max(floats ...*float64) *float64 {
	var mx *float64 = nil
	for id := range floats {
		if floats[id] != nil {
			if mx == nil || *mx < *floats[id] {
				mx = floats[id]
			}
		}
	}
	return mx
}

func Min(floats ...*float64) *float64 {
	var mx *float64 = nil
	for id := range floats {
		if floats[id] != nil {
			if mx == nil || *mx > *floats[id] {
				mx = floats[id]
			}
		}
	}
	return mx
}

func Divide(m *float64, n *float64) *float64 {
	if m == nil || n == nil {
		return nil
	}
	if *n == 0 {
		return nil
	}
	dv := (*m) / (*n)
	return &dv
}

func Multiply(m *float64, n *float64) *float64 {
	if m == nil || n == nil {
		return nil
	}
	rs := (*m) * (*n)
	return &rs
}

func DefaultF(m *float64, n float64) float64 {
	if m == nil {
		return n
	}
	return *m
}

func DefaultI(m *int64, n int64) int64 {
	if m == nil {
		return n
	}
	return *m
}

func DefaultTime(m *time.Time, n time.Time) time.Time {
	if m == nil {
		return n
	}
	return *m
}

func DefaultTimeStamp(m *timestamp.Timestamp, n time.Time) time.Time {
	if m == nil {
		return n
	}
	return m.AsTime()
}

func MaxDate(dates []*time.Time, limit *time.Time) *time.Time {
	if dates == nil || len(dates) == 0 {
		return nil
	}
	var maxDate *time.Time = nil
	for id := range dates {
		dt := dates[id]
		if dt == nil {
			continue
		}
		if maxDate == nil {
			if limit == nil {
				maxDate = dt
			} else {
				if dt.After(*limit) {
					continue
				} else {
					maxDate = dt
				}
			}
		}
		if limit == nil {
			if dt.After(*maxDate) {
				maxDate = dt
			}

		} else {
			if dt.After(*maxDate) && (*dt == *limit || dt.Before(*limit)) {
				maxDate = dt
			}
		}
	}
	return maxDate
}
