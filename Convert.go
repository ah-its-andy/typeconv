package typeconv

import (
	"errors"
	"reflect"
	"time"
)

var typeConverters []TypeConverter = make([]TypeConverter, 0)

func SetTypeConverters(tc []TypeConverter) {
	for _, e := range tc {
		hasElement := false
		for _, t := range typeConverters {
			if t.GetName() == e.GetName() {
				hasElement = true
				break
			}
		}
		if !hasElement {
			typeConverters = append(typeConverters, e)
		}
	}
}

func GetTypeConverter(v interface{}) TypeConverter {
	var typeConv TypeConverter
	for _, e := range typeConverters {
		if e.IsConvertible(v) {
			typeConv = e
			break
		}
	}
	return typeConv
}

func ChangeType(v interface{}, conversionType reflect.Type, format FormatProvider) (interface{}, error) {
	typeConverter := GetTypeConverter(v)
	if typeConverter == nil {
		return nil, errors.New("invalid cast : TypeConverter not found")
	}
	return typeConverter.ChangeType(v, conversionType, format)
}

func ToBoolFormat(v interface{}, format FormatProvider) (*bool, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*bool), nil
}
func ToByteFormat(v interface{}, format FormatProvider) (*byte, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*byte), nil
}
func ToTimeFormat(v interface{}, format FormatProvider) (*time.Time, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*time.Time), nil
}
func ToIntFormat(v interface{}, format FormatProvider) (*int, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*int), nil
}
func ToInt16Format(v interface{}, format FormatProvider) (*int16, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*int16), nil
}
func ToInt32Format(v interface{}, format FormatProvider) (*int32, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*int32), nil
}
func ToInt64Format(v interface{}, format FormatProvider) (*int64, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*int64), nil
}
func ToFloat32Format(v interface{}, format FormatProvider) (*float32, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*float32), nil
}
func ToFloat64Format(v interface{}, format FormatProvider) (*float64, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*float64), nil
}
func ToStringFormat(v interface{}, format FormatProvider) (*string, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), format)
	if err != nil {
		return nil, err
	}
	return r.(*string), nil
}

func ToTimef(v interface{}, format string) (*time.Time, error) {
	formatProvider, err := GetTimeFormat(format)
	if err != nil {
		return nil, err
	}
	return ToTimeFormat(v, formatProvider)
}

func ToFloat32f(v interface{}, format string) (*float32, error) {
	formatProvider, err := GetTimeFormat(format)
	if err != nil {
		return nil, err
	}
	return ToFloat32Format(v, formatProvider)
}

func ToFloat64f(v interface{}, format string) (*float64, error) {
	formatProvider, err := GetTimeFormat(format)
	if err != nil {
		return nil, err
	}
	return ToFloat64Format(v, formatProvider)
}

func ToStringf(v interface{}, format string) (*string, error) {
	formatProvider, err := GetTimeFormat(format)
	if err != nil {
		return nil, err
	}
	return ToStringFormat(v, formatProvider)
}

func ToBool(v interface{}) (*bool, error) {
	return ToBoolFormat(v, nil)
}
func ToByte(v interface{}) (*byte, error) {
	return ToByteFormat(v, nil)
}
func ToTime(v interface{}) (*time.Time, error) {
	return ToTimeFormat(v, nil)
}
func ToInt(v interface{}) (*int, error) {
	return ToIntFormat(v, nil)
}
func ToInt16(v interface{}) (*int16, error) {
	return ToInt16Format(v, nil)
}
func ToInt32(v interface{}) (*int32, error) {
	return ToInt32Format(v, nil)
}
func ToInt64(v interface{}) (*int64, error) {
	return ToInt64Format(v, nil)
}
func ToFloat32(v interface{}) (*float32, error) {
	return ToFloat32Format(v, nil)
}
func ToFloat64(v interface{}) (*float64, error) {
	return ToFloat64Format(v, nil)
}
func ToString(v interface{}) (*string, error) {
	return ToStringFormat(v, nil)
}
