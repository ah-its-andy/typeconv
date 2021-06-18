package typeconv

import (
	"time"
)

type TypeConverter interface {
	GetName() string
	IsConvertible(v interface{}) bool
	ToBool(v interface{}) (bool, error)
	ToTime(v interface{}, format string) (time.Time, error)
	ToInt(v interface{}) (int, error)
	ToInt16(v interface{}) (int16, error)
	ToInt32(v interface{}) (int32, error)
	ToInt64(v interface{}) (int64, error)
	ToFloat32(v interface{}) (float32, error)
	ToFloat64(v interface{}) (float64, error)
	ToString(v interface{}, format string) (string, error)
	ChangeType(v interface{}, format string) (interface{}, error)
}
