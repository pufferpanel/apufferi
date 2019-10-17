package apufferi

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"time"
)

//Converts the val parameter to the same type as the target
func Convert(val interface{}, target interface{}) (interface{}, error) {
	switch target.(type) {
	case string:
		return cast.ToStringE(val)
	case int:
		return cast.ToIntE(val)
	case int8:
		return cast.ToInt8E(val)
	case int16:
		return cast.ToInt16E(val)
	case int32:
		return cast.ToInt32E(val)
	case int64:
		return cast.ToInt64E(val)
	case uint:
		return cast.ToUintE(val)
	case uint8:
		return cast.ToUint8E(val)
	case uint16:
		return cast.ToUint16E(val)
	case uint32:
		return cast.ToUint32E(val)
	case uint64:
		return cast.ToUint64E(val)
	case bool:
		return cast.ToBoolE(val)
	case time.Duration:
		return cast.ToDurationE(val)
	case time.Time:
		return cast.ToTimeE(val)
	case float32:
		return cast.ToFloat64E(val)
	case float64:
		return cast.ToFloat64E(val)
	case map[string]string:
		return cast.ToStringMapStringE(val)
	case map[string][]string:
		return cast.ToStringMapStringSliceE(val)
	case map[string]bool:
		return cast.ToStringMapBoolE(val)
	case map[string]interface{}:
		return cast.ToStringMapE(val)
	case map[string]int:
		return cast.ToStringMapIntE(val)
	case map[string]int64:
		return cast.ToStringMapInt64E(val)
	case []interface{}:
		return cast.ToSliceE(val)
	case []bool:
		return cast.ToBoolSliceE(val)
	case []string:
		return cast.ToStringSliceE(val)
	case []int:
		return cast.ToIntSliceE(val)
	case []time.Duration:
		return cast.ToDurationSliceE(val)
	}

	return nil, errors.New(fmt.Sprintf("cannot convert %s to %s", reflect.TypeOf(val), reflect.TypeOf(target)))
}
