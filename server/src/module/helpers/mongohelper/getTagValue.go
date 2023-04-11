package mongohelper

import (
	"reflect"
)

func GetTagValue(enclosingStruct interface{}, fieldName, tag string) string {
	structType := reflect.TypeOf(enclosingStruct)
	field, _ := structType.FieldByName(fieldName)
	return field.Tag.Get(tag)
}
