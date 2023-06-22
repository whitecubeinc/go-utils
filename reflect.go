package utils

import (
	"reflect"
)

func GetNotEmptyFields(v any) (fields []string) {
	reflectType := reflect.TypeOf(v)
	reflectValue := reflect.ValueOf(v)

	for i := 0; i < reflectType.NumField(); i++ {
		if reflectType.Field(i).Type.Kind() == reflect.Ptr {
			if !reflectValue.Field(i).IsNil() {
				fields = append(fields, reflectType.Field(i).Name)
			}
		} else {
			if !reflectValue.Field(i).IsZero() {
				fields = append(fields, reflectType.Field(i).Name)
			}
		}
	}

	return
}

// NewStringEnum strings.ToUpper, strings.ToLower 및 strcase.ToXXX 사용
func NewStringEnum[T any](keyCaseFunction func(string) string, valueCaseFunction func(string) string) T {
	// FIXME add test code
	dest := new(T)

	destTypeOf := Deref(reflect.TypeOf(dest))
	value := reflect.ValueOf(dest)

	direct := reflect.Indirect(value)
	base := Deref(value.Type())

	vp := reflect.New(base)
	v := reflect.Indirect(vp)

	valuesIdx := -1
	values := make(map[string]string)

	for i := 0; i < destTypeOf.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			valuePtr := reflect.New(v.Field(i).Type())
			fieldValue := reflect.Indirect(valuePtr)

			value := destTypeOf.Field(i).Name
			if valueCaseFunction != nil {
				value = valueCaseFunction(destTypeOf.Field(i).Name)
			}

			key := destTypeOf.Field(i).Name
			if keyCaseFunction != nil {
				key = keyCaseFunction(destTypeOf.Field(i).Name)
			}
			values[key] = value

			fieldValue.SetString(value)
			v.Field(i).Set(fieldValue)
		case reflect.Map:
			if destTypeOf.Field(i).Name == "ValueMap" {
				valuesIdx = i
			}
		}
	}

	if valuesIdx >= 0 {
		v.Field(valuesIdx).Set(reflect.ValueOf(values))
	}

	direct.Set(v)

	return *dest
}

// NewConstantFromTag tag key = value , string only
func NewConstantFromTag[T any](keyCaseFunction func(string) string) T {
	// FIXME add test code
	dest := new(T)

	destTypeOf := Deref(reflect.TypeOf(dest))

	value := reflect.ValueOf(dest)

	direct := reflect.Indirect(value)
	base := Deref(value.Type())

	vp := reflect.New(base)
	v := reflect.Indirect(vp)

	valuesIdx := -1
	values := make(map[string]string)

	for i := 0; i < destTypeOf.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.String:
			valuePtr := reflect.New(v.Field(i).Type())
			fieldValue := reflect.Indirect(valuePtr)

			value := destTypeOf.Field(i).Tag.Get("value")
			key := destTypeOf.Field(i).Name
			if keyCaseFunction != nil {
				key = keyCaseFunction(destTypeOf.Field(i).Name)
			}
			values[key] = value

			fieldValue.SetString(value)
			v.Field(i).Set(fieldValue)
		case reflect.Map:
			if destTypeOf.Field(i).Name == "ValueMap" {
				valuesIdx = i
			}
		}
	}

	if valuesIdx >= 0 {
		v.Field(valuesIdx).Set(reflect.ValueOf(values))
	}

	direct.Set(v)

	return *dest
}

func Deref(t reflect.Type) reflect.Type {
	// FIXME add test code
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}
