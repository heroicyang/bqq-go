// A Tencent Business QQ client in go.
// https://github.com/heroicyang/bqq-go
//
// @author Herioc Yang <me@heroicyang.com>

package bqq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// API call result
type Result map[string]interface{}

// Make a result from tencent API response.
func MakeResult(jsonBytes []byte) (res Result, err error) {
	err = json.Unmarshal(jsonBytes, &res)

	if err != nil {
		fmt.Errorf("cannot format tencent response. %v", err)
		return
	}

	return
}

// Gets a field.
//
// Field can be a dot separated string.
// If field name is "a.b.c", it will try to return value of res["a"]["b"]["c"].
//
// To access array items, use index value in field.
// For instance, field "a.0.c" means to read res["a"][0]["c"].
//
// It doesn't work with Result which has a key contains dot. Use GetField in this case.
//
// Returns nil if field doesn't exist.
func (res *Result) Get(field string) interface{} {
	if field == "" {
		return res
	}

	fields := strings.Split(field, ".")
	return res.get(fields)
}

// Get a field.
// Arguments are treated as keys to access value in Result.
// If arguments are "a","b","c", it will try to return value of res["a"]["b"]["c"].
//
// To access array items, use index value as a string.
// For instance, args of "a", "0", "c" means to read res["a"][0]["c"].
//
// Returns nil if field doesn't exist.
func (res Result) GetField(fields ...string) interface{} {
	if len(fields) == 0 {
		return res
	}

	return res.get(fields)
}

func (res Result) get(fields []string) interface{} {
	v, ok := res[fields[0]]

	if !ok || v == nil {
		return nil
	}

	if len(fields) == 1 {
		return v
	}

	value := getValueField(reflect.ValueOf(v), fields[1:])

	if !value.IsValid() {
		return nil
	}

	return value.Interface()
}

func getValueField(value reflect.Value, fields []string) reflect.Value {
	valueType := value.Type()
	kind := valueType.Kind()
	field := fields[0]

	switch kind {
	case reflect.Array, reflect.Slice:
		n, err := strconv.ParseUint(field, 0, 10)

		if err != nil {
			return reflect.Value{}
		}

		if n >= uint64(value.Len()) {
			return reflect.Value{}
		}

		value = reflect.ValueOf(value.Index(int(n)).Interface())

	case reflect.Map:
		v := value.MapIndex(reflect.ValueOf(field))

		if !v.IsValid() {
			return v
		}

		value = reflect.ValueOf(v.Interface())

	default:
		value = reflect.Value{}
	}

	if len(fields) == 1 {
		return value
	}

	return getValueField(value, fields[1:])
}
