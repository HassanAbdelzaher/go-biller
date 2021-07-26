package tools

import (
	"math"
	"strconv"
	"strings"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func RoundTo(n float64, decimals uint32) float64 {
	return math.Round(n*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}

func ToFloatPointer(fl float64) *float64 {
	return &fl
}

func FloatPtrToFloat(fl *float64) float64 {
	if fl == nil {
		return 0
	}
	return *fl
}

func FloatPtrToInt32Ptr(fl *float64) *int32 {
	if fl == nil {
		return nil
	}
	i := FloatPtrToInt32(fl)
	return &i
}

func FloatPtrToInt32(fl *float64) int32 {
	if fl == nil {
		return 0
	}
	return int32(*fl)
}

func Int32PtrPtrToInt32(fl *int32) int32 {
	if fl == nil {
		return 0
	}
	return *fl
}

func Int32PtrToInt64Ptr(fl *int32) *int64 {
	if fl == nil {
		return nil
	}
	return ToIntPointer(int64(*fl))
}

func Int32ToInt32Ptr(fl int32) *int32 {
	return &fl
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

func StringToInt32(f *string) *int32 {
	if f == nil {
		return nil
	}
	i, err := strconv.Atoi(*f)
	if err != nil {
		return nil
	}
	r := int32(i)
	return &r
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

func ToTimePrt(m time.Time) *time.Time {
	return &m
}
