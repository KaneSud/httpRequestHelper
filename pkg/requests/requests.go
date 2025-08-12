package requests

import (
	"errors"
	"net/http"
	"reflect"
)

var (
	parsedError            = errors.New("cant_parse_value")
	overflowError          = errors.New("value_overflow")
	setValueError          = errors.New("cant_set_value")
	reflectConversionError = errors.New("cant_convert_internal")
	nilProvided            = errors.New("nil_provided")
	unsupportedType        = errors.New("unsupported_type")
)

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	default:
	}
	return false
}

func ParseQueryStruct(body interface{}, r *http.Request) error {
	if isNil(body) || r == nil {
		return nilProvided
	}
	ts := reflect.ValueOf(body)
	t := ts.Elem()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			fieldVal := t.Type().Field(i)
			jsonTag := fieldVal.Tag.Get("json")
			if len(jsonTag) == 0 {
				continue
			}
			val := t.FieldByName(fieldVal.Name)
			parser, ok := parserMap[val.Type()]
			if !ok {
				return unsupportedType
			}
			if !val.IsValid() || !val.CanSet() {
				return setValueError
			}
			err := parser(r, val, jsonTag)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
