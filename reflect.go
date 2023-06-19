package utils

import "reflect"

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
