package typeconv

import (
	"reflect"
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

func (conv *StringTypeConverter) ToBool(v interface{}, format FormatProvider) (*bool, error) {
	r := strings.EqualFold(v.(string), "true")
	return &r, nil
}

func (conv *StringTypeConverter) ToTime(v interface{}, format FormatProvider) (*time.Time, error) {

}

func (conv *StringTypeConverter) ToInt(v interface{}, format FormatProvider) (*int, error)
func (conv *StringTypeConverter) ToInt16(v interface{}, format FormatProvider) (*int16, error)
func (conv *StringTypeConverter) ToInt32(v interface{}, format FormatProvider) (*int32, error)
func (conv *StringTypeConverter) ToInt64(v interface{}, format FormatProvider) (*int64, error)
func (conv *StringTypeConverter) ToFloat32(v interface{}, format FormatProvider) (*float32, error)
func (conv *StringTypeConverter) ToFloat64(v interface{}, format FormatProvider) (*float64, error)
func (conv *StringTypeConverter) ToString(v interface{}, format FormatProvider) (*string, error)
func (conv *StringTypeConverter) ChangeType(v interface{}, conversionType reflect.Type, format FormatProvider) (interface{}, error)
