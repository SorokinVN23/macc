package utils

import (
	"errors"
	"fmt"
)

func Sum(v1 interface{}, v2 interface{}) (interface{}, error) {
	switch v1.(type) {
	case int64:
		res := v1.(int64) + v2.(int64)
		return interface{}(res), nil
	case float64:
		res := v1.(float64) + v2.(float64)
		return interface{}(res), nil
	}

	return nil, errors.New("sum operation is not supported")
}

func ConvertToInt64(val interface{}) (int64, error) {
	switch v := val.(type) {
	case int64:
		return v, nil
	case float64:
		return int64(v), nil
	}

	return 0, fmt.Errorf("invalid val %+v for ConvertToInt64", val)
}

func ConvertToFloat64(val interface{}) (float64, error) {
	switch v := val.(type) {
	case int64:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	}

	return 0, fmt.Errorf("invalid val %+v for ConvertToFloat64", val)
}
