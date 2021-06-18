package typeconv

import (
	"fmt"
	"time"
)

type TypeTimeConverter struct {
}

func (conv *TypeTimeConverter) GetName() string {
	return "typeconv.TypeTimeConverter"
}
func (conv *TypeTimeConverter) IsConvertible(v interface{}) bool {
	if _, ok := v.(time.Time); ok {
		return true
	}
	return false
}

func (conv *TypeTimeConverter) getTime(v interface{}) time.Time {
	return v.(time.Time)
}

func (conv *TypeTimeConverter) ToBool(v interface{}) (bool, error) {
	return false, fmt.Errorf("invalid cast: %T could not cast to bool")
}
func (conv *TypeTimeConverter) ToTime(v interface{}, format string) (time.Time, error) {
	return conv.getTime(v), nil
}
func (conv *TypeTimeConverter) ToInt(v interface{}) (int, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to int", v)
}
func (conv *TypeTimeConverter) ToInt16(v interface{}) (int16, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to int16", v)
}
func (conv *TypeTimeConverter) ToInt32(v interface{}) (int32, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to int32", v)
}
func (conv *TypeTimeConverter) ToInt64(v interface{}) (int64, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to int64", v)
}
func (conv *TypeTimeConverter) ToFloat32(v interface{}) (float32, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to float32", v)
}
func (conv *TypeTimeConverter) ToFloat64(v interface{}) (float64, error) {
	return 0, fmt.Errorf("invalid cast: %T could not cast to float64", v)
}
func (conv *TypeTimeConverter) ToString(v interface{}, format string) (string, error) {
	return conv.getTime(v).Format(format), nil
}
