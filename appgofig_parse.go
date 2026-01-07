package appgofig

import (
	"fmt"
	"reflect"
	"strconv"
)

// stringToValue takes rawInput and converts it to configType if possible, errors out otherwise. Uses strconv methods.
func stringToValue(rawInput string, configType reflect.Kind) (any, error) {
	switch configType {
	case reflect.String:
		return rawInput, nil
	case reflect.Bool:
		return strconv.ParseBool(rawInput)
	case reflect.Int64:
		return strconv.ParseInt(rawInput, 0, 64)
	case reflect.Float64:
		return strconv.ParseFloat(rawInput, 64)
	default:
		return nil, fmt.Errorf("unsupported type %q - allowed are string, bool, int64, float64", configType)
	}
}

// valueToString returns a string representation of value for supported types within AppGofig
func valueToString(value reflect.Value) string {
	switch value.Kind() {
	case reflect.String:
		return value.String()
	case reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	default:
		return " - unsupported type " + value.Kind().String() + " - "
	}
}
