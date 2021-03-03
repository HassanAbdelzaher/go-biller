package tools

import (
	"strconv"
	"strings"
)

func ToFloatPointer(fl float64) *float64 {
	return &fl
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


func Float64ToString(f *float64) *string{
	if f==nil{
		return nil
	}
	str:=strconv.FormatFloat(*f,'f',-1,64)
	return &str
}


func Float32ToString(f *float32) *string{
	if f==nil{
		return nil
	}
	f64:=float64(*f)
	str:=strconv.FormatFloat(f64,'f',-1,32)
	return &str
}


func Int64ToString(f *int64) *string{
	if f==nil{
		return nil
	}
	str:=strconv.FormatInt(*f,10)
	return &str
}


func Int32ToString(f *int32) *string{
	if f==nil{
		return nil
	}
	i:=int64(*f)
	str:=strconv.FormatInt(i,10)
	return &str
}

func StringCompare(s1 string,s2 string)  bool{
	return strings.TrimSpace(s1)==strings.TrimSpace(s2)
}


func StringComparePointer(s1 *string,s2 *string)  bool{
	if s1==nil || s2==nil{
		return false
	}
	return StringCompare(*s1,*s2)
}

func StringComparePointer2(s1 string,s2 *string)  bool{
	if s2==nil{
		return false
	}
	return StringCompare(s1,*s2)
}

func BoolToString(b *bool) *string{
	if b==nil{
		return nil
	}
	one:="1"
	zero:="0"
	if *b {
		return &one
	}
	return &zero
}