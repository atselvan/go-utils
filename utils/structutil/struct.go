package structutil

import (
	"reflect"
	"strings"
)

const (
	StructTagRequired     = "required"
	StructTagMapstructure = "mapstructure"
	StructTagJson         = "json"
	StructTagYaml         = "yaml"
)

// GetFieldTagValue returns the value of the mapstructure field tag for a struct field.
func GetFieldTagValue(structPointer any, fieldPointer any) string {
	var tagValue string
	structReflect := reflect.ValueOf(structPointer).Elem()
	fieldReflect := reflect.ValueOf(fieldPointer).Elem()
	for i := 0; i < structReflect.NumField(); i++ {
		fieldValue := structReflect.Field(i)
		if fieldValue.Addr().Interface() == fieldReflect.Addr().Interface() {
			tagValue = structReflect.Type().Field(i).Tag.Get(StructTagMapstructure)
		}
	}
	return strings.Split(tagValue, ",")[0]
}

// GetJsonFieldTagValue returns the value of the json field tag for a struct field.
func GetJsonFieldTagValue(structPointer any, fieldPointer any) string {
	var tagValue string
	structReflect := reflect.ValueOf(structPointer).Elem()
	fieldReflect := reflect.ValueOf(fieldPointer).Elem()
	for i := 0; i < structReflect.NumField(); i++ {
		fieldValue := structReflect.Field(i)
		if fieldValue.Addr().Interface() == fieldReflect.Addr().Interface() {
			tagValue = structReflect.Type().Field(i).Tag.Get(StructTagJson)
		}
	}
	return strings.Split(tagValue, ",")[0]
}

// GetYamlFieldTagValue returns the value of the json field tag for a struct field.
func GetYamlFieldTagValue(structPointer any, fieldPointer any) string {
	var tagValue string
	structReflect := reflect.ValueOf(structPointer).Elem()
	fieldReflect := reflect.ValueOf(fieldPointer).Elem()
	for i := 0; i < structReflect.NumField(); i++ {
		fieldValue := structReflect.Field(i)
		if fieldValue.Addr().Interface() == fieldReflect.Addr().Interface() {
			tagValue = structReflect.Type().Field(i).Tag.Get(StructTagYaml)
		}
	}
	return strings.Split(tagValue, ",")[0]
}
