package typeconv

import (
	"reflect"
	"time"
)

type TypeConverter interface {
	GetName() string
	IsConvertible(v interface{}) bool
	ToBool(v interface{}, format FormatProvider) (*bool, error)
	ToByte(v interface{}, format FormatProvider) (*byte, error)
	ToTime(v interface{}, format FormatProvider) (*time.Time, error)
	ToInt(v interface{}, format FormatProvider) (*int, error)
	ToInt16(v interface{}, format FormatProvider) (*int16, error)
	ToInt32(v interface{}, format FormatProvider) (*int32, error)
	ToInt64(v interface{}, format FormatProvider) (*int64, error)
	ToFloat32(v interface{}, format FormatProvider) (*float32, error)
	ToFloat64(v interface{}, format FormatProvider) (*float64, error)
	ToString(v interface{}, format FormatProvider) (*string, error)
	ChangeType(v interface{}, conversionType reflect.Type, format FormatProvider) (interface{}, error)
}
