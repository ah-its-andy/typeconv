package typeconv

import (
	"strconv"
	"strings"
	"time"
)

type StringTypeConverter struct {
}

func (conv *StringTypeConverter) GetName() string {
	return "typeconv.StringTypeConverter"
}

func (conv *StringTypeConverter) IsConvertible(v interface{}) bool {
	if _, ok := v.(string); ok {
		return true
	}
	return false
}

func (conv *StringTypeConverter) getStr(v interface{}) string {
	return v.(string)
}

func (conv *StringTypeConverter) ToBool(v interface{}) (bool, error) {
	s := conv.getStr(v)
	r := strings.EqualFold(s, "true")
	return r, nil
}

func (conv *StringTypeConverter) ToTime(v interface{}, format string) (time.Time, error) {
	s := conv.getStr(v)
	r, err := time.Parse(format, s)
	if err != nil {
		return time.Now(), err
	}
	return r, nil
}

func (conv *StringTypeConverter) ToInt(v interface{}) (int, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(r), nil
}
func (conv *StringTypeConverter) ToInt16(v interface{}) (int16, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(r), nil
}
func (conv *StringTypeConverter) ToInt32(v interface{}) (int32, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(r), nil
}
func (conv *StringTypeConverter) ToInt64(v interface{}) (int64, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return r, nil
}
func (conv *StringTypeConverter) ToFloat32(v interface{}) (float32, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0.0, err
	}
	return float32(r), nil
}
func (conv *StringTypeConverter) ToFloat64(v interface{}) (float64, error) {
	s := conv.getStr(v)
	r, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return r, nil
}
func (conv *StringTypeConverter) ToString(v interface{}, format string) (string, error) {
	return v.(string), nil
}
