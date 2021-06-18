package typeconv

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Int16MaxValue() int16 {
	return 32767
}

func Int16MinValue() int16 {
	return -32768
}

func Int32MaxValue() int32 {
	return 2147483647
}

func Int32MinValue() int32 {
	return -2147483648
}

func Int64MaxValue() int64 {
	return 9223372036854775807
}

func Int64MinValue() int64 {
	return -9223372036854775808
}

func Float32MaxValue() float32 {
	return 3.40282347e+38
}

func Float32MinValue() float32 {
	return -3.40282347e+38
}

func Float64MaxValue() float64 {
	return 1.7976931348623157e+308
}

func Float64MinValue() float64 {
	return -1.7976931348623157e+308
}

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

func ChangeType(v interface{}, conversionType reflect.Type, format string) (interface{}, error) {
	conv := GetTypeConverter(v)
	if conv == nil {
		return nil, errors.New("invalid cast : TypeConverter not found")
	}
	if v == nil {
		return nil, nil
	}
	if conversionType.Name() == reflect.TypeOf(time.Now()).Name() {
		return conv.ToTime(v, format)
	}
	switch conversionType.Kind() {
	case reflect.Bool:
		return conv.ToBool(v)
	case reflect.Int:
		return conv.ToInt(v)
	case reflect.Int16:
		return conv.ToInt16(v)
	case reflect.Int32:
		return conv.ToInt32(v)
	case reflect.Int64:
		return conv.ToInt64(v)
	case reflect.Float32:
		return conv.ToFloat32(v)
	case reflect.Float64:
		return conv.ToFloat64(v)
	case reflect.String:
		return conv.ToString(v, format)
	default:
		return nil, fmt.Errorf("invalid cast: %T could not cast to float64", v)
	}
}

func ToTimef(v interface{}, format string) (*time.Time, error) {
	r, err := ChangeType(v, reflect.TypeOf(time.Now()), format)
	if err != nil {
		return nil, err
	}
	return r.(*time.Time), nil
}

func ToStringf(v interface{}, format string) (*string, error) {
	r, err := ChangeType(v, reflect.TypeOf(""), format)
	if err != nil {
		return nil, err
	}
	return r.(*string), nil
}

func ToBool(v interface{}) (*bool, error) {
	r, err := ChangeType(v, reflect.TypeOf(false), "")
	if err != nil {
		return nil, err
	}
	return r.(*bool), nil
}

func ToTime(v interface{}) (*time.Time, error) {
	return ToTimef(v, "")
}

func ToInt(v interface{}) (*int, error) {
	r, err := ChangeType(v, reflect.TypeOf(int(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*int), nil
}

func ToInt16(v interface{}) (*int16, error) {
	r, err := ChangeType(v, reflect.TypeOf(int16(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*int16), nil
}

func ToInt32(v interface{}) (*int32, error) {
	r, err := ChangeType(v, reflect.TypeOf(int32(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*int32), nil
}

func ToInt64(v interface{}) (*int64, error) {
	r, err := ChangeType(v, reflect.TypeOf(int64(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*int64), nil
}

func ToFloat32(v interface{}) (*float32, error) {
	r, err := ChangeType(v, reflect.TypeOf(float32(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*float32), nil
}

func ToFloat64(v interface{}) (*float64, error) {
	r, err := ChangeType(v, reflect.TypeOf(float64(0)), "")
	if err != nil {
		return nil, err
	}
	return r.(*float64), nil
}

func ToString(v interface{}) (*string, error) {
	return ToStringf(v, "")
}
