package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

//gocyclo:ignore
func Any2String(v any) (string, error) {
	switch s := v.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case json.Number:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case template.HTML:
		return string(s), nil
	case template.HTMLAttr:
		return string(s), nil
	case template.URL:
		return string(s), nil
	case template.JS:
		return string(s), nil
	case template.JSStr:
		return string(s), nil
	case template.CSS:
		return string(s), nil
	case template.Srcset:
		return string(s), nil
	case nil:
		return "", nil
	case fmt.Stringer:
		return s.String(), nil
	case error:
		return s.Error(), nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", v, v)
	}
}

func Any2IntMust(v any) int {
	value, err := Any2Int(v)
	if err != nil {
		panic(err)
	}
	return value
}

func Any2Int(v any) (int, error) {
	switch s := v.(type) {
	case string:
		return strconv.Atoi(s)
	case int:
		return s, nil
	case float64:
		return int(s), nil
	case int64:
		return int(s), nil
	default:
		return 0, errors.New("cannot parse to int")
	}
}

func Any2Float(v any) (float64, error) {
	// 입력값의 타입을 확인
	valueType := reflect.TypeOf(v)

	// 입력값이 숫자 형식인지 확인
	switch valueType.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(v.(int64)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(v.(uint64)), nil
	case reflect.Float32, reflect.Float64:
		return v.(float64), nil
	}

	// 변환이 불가능한 경우 에러 반환
	return 0, fmt.Errorf("지원되지 않는 데이터 타입: %s", valueType.String())
}

func Any2FloatMust(v any) float64 {
	float, err := Any2Float(v)
	if err != nil {
		panic(err)
	}

	return float
}
