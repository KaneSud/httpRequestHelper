package requests

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func parseInt64(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseInt(itemStr, 10, 64)
	if err != nil {
		return err
	}
	fieldVal.SetInt(curItem)
	return nil
}

func parseUint64(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseUint(itemStr, 10, 64)
	if err != nil {
		return err
	}
	fieldVal.SetUint(curItem)
	return nil
}

func parseUint32(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseUint(itemStr, 10, 32)
	if err != nil {
		return err
	}
	fieldVal.SetUint(curItem)
	return nil
}

func parseFloat64(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseFloat(itemStr, 64)
	if err != nil {
		return err
	}
	fieldVal.SetFloat(curItem)
	return nil
}

func parseBool(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseBool(itemStr)
	if err != nil {
		return err
	}
	fieldVal.SetBool(curItem)
	return nil
}

func parseBoolPointer(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseBool(itemStr)
	if err != nil {
		return err
	}
	newValue := reflect.ValueOf(&curItem)
	fieldVal.Set(newValue)
	return nil
}

func parseInt64Pointer(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseInt(itemStr, 10, 64)
	if err != nil {
		return err
	}
	newValue := reflect.ValueOf(&curItem)
	fieldVal.Set(newValue)
	return nil
}

func parseComplex128(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	if len(itemStr) == 0 {
		return nil
	}
	curItem, err := strconv.ParseComplex(itemStr, 128)
	if err != nil {
		return err
	}
	fieldVal.SetComplex(curItem)
	return nil
}

func parseString(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemStr := r.FormValue(jsonTag)
	fieldVal.SetString(itemStr)
	return nil
}

func parseInt64Slice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedInt []int64
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseInt(itemsSplit[idx], 10, 64)
		if err != nil {
			return err
		}
		parsedInt = append(parsedInt, curItem)
	}
	if isNil(parsedInt) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedInt)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseUint64Slice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedInt []uint64
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseUint(itemsSplit[idx], 10, 64)
		if err != nil {
			return err
		}
		parsedInt = append(parsedInt, curItem)
	}
	if isNil(parsedInt) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedInt)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseUint32Slice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedInt []uint32
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseUint(itemsSplit[idx], 10, 32)
		if err != nil {
			return err
		}
		parsedInt = append(parsedInt, uint32(curItem))
	}
	if isNil(parsedInt) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedInt)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseFloat64Slice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedFloat []float64
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseFloat(itemsSplit[idx], 64)
		if err != nil {
			return err
		}
		parsedFloat = append(parsedFloat, curItem)
	}
	if isNil(parsedFloat) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedFloat)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseBoolSlice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedBool []bool
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseBool(itemsSplit[idx])
		if err != nil {
			return err
		}
		parsedBool = append(parsedBool, curItem)
	}
	if isNil(parsedBool) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedBool)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseComplex128Slice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	var parsedComplex []complex128
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	for idx, _ := range itemsSplit {
		curItem, err := strconv.ParseComplex(itemsSplit[idx], 128)
		if err != nil {
			return err
		}
		parsedComplex = append(parsedComplex, curItem)
	}
	if isNil(parsedComplex) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(parsedComplex)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

func parseStringSlice(r *http.Request, fieldVal reflect.Value, jsonTag string) error {
	itemsStr := r.FormValue(jsonTag)
	if len(itemsStr) == 0 {
		return nil
	}
	itemsSplit := strings.Split(itemsStr, ",")
	if isNil(itemsSplit) {
		fieldVal.Set(reflect.Zero(fieldVal.Type()))
	} else {
		reflectSlice := reflect.ValueOf(itemsSplit)
		fieldVal.Set(reflectSlice)
	}
	return nil
}

var parserMap map[reflect.Type]func(*http.Request, reflect.Value, string) error

func init() {
	parserMap = map[reflect.Type]func(*http.Request, reflect.Value, string) error{
		reflect.TypeOf((*int64)(nil)):       parseInt64Pointer,
		reflect.TypeOf((*bool)(nil)):        parseBoolPointer,
		reflect.TypeOf((int64)(0)):          parseInt64,
		reflect.TypeOf((uint64)(0)):         parseUint64,
		reflect.TypeOf((uint32)(0)):         parseUint32,
		reflect.TypeOf((float64)(0)):        parseFloat64,
		reflect.TypeOf((bool)(false)):       parseBool,
		reflect.TypeOf((complex128)(0)):     parseComplex128,
		reflect.TypeOf((string)("")):        parseString,
		reflect.TypeOf(([]int64)(nil)):      parseInt64Slice,
		reflect.TypeOf(([]uint64)(nil)):     parseUint64Slice,
		reflect.TypeOf(([]uint32)(nil)):     parseUint32Slice,
		reflect.TypeOf(([]float64)(nil)):    parseFloat64Slice,
		reflect.TypeOf(([]bool)(nil)):       parseBoolSlice,
		reflect.TypeOf(([]complex128)(nil)): parseComplex128Slice,
		reflect.TypeOf(([]string)(nil)):     parseStringSlice,
	}
}
