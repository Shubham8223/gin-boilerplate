package utils

import (
	"reflect"
)

func StructToMapUpdate(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(input).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		jsonKey := fieldType.Tag.Get("json")

		if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() {
			result[jsonKey] = fieldVal.Elem().Interface()
		}
	}
	return result
}


func StructToMapCreate(input interface{}, output interface{}) {
	inVal := reflect.ValueOf(input).Elem()
	outVal := reflect.ValueOf(output).Elem()

	for i := 0; i < inVal.NumField(); i++ {
		inField := inVal.Field(i)
		inFieldType := inVal.Type().Field(i)

		outField := outVal.FieldByName(inFieldType.Name)
		if outField.IsValid() && outField.CanSet() && inField.Type() == outField.Type() {
			outField.Set(inField)
		}
	}
}
