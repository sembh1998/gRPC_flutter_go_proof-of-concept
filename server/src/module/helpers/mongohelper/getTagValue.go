package mongohelper

import (
	"log"
	"reflect"
)

func GetTagValue(fieldtype interface{}, enclosingStruct interface{}, tag string) string {
	log.Printf("fieldtype: %v", fieldtype)
	log.Printf("enclosingStruct: %v", enclosingStruct)
	log.Printf("tag: %v", tag)

	structType := reflect.TypeOf(enclosingStruct)
	log.Printf("structType: %v", structType)

	fieldName := reflect.TypeOf(fieldtype).Name()
	log.Printf("fieldName: %v", fieldName)

	nstructType := reflect.TypeOf(enclosingStruct)
	if structType.Kind() == reflect.Ptr {
		nstructType = structType.Elem()
		log.Printf("nstructType: %v", nstructType)
	}
	nfieldType := reflect.TypeOf(fieldtype)
	if nfieldType.Kind() == reflect.Ptr {
		nfieldType = nfieldType.Elem()
		log.Printf("nfieldType: %v", nfieldType)
	}

	if fieldName == "" {
		log.Printf("nstructType: %v", nstructType)
		log.Printf("nfieldType: %v", nfieldType)
		for i := 0; i < nstructType.NumField(); i++ {
			field := nstructType.Field(i)
			log.Printf("for %v field: %v, nfieldType: %v", i, field, nfieldType)
			if field.Type == nfieldType {
				fieldName = field.Name
				log.Printf("fieldName: %v", fieldName)
				break
			}
		}
	}

	field, _ := structType.FieldByName(fieldName)
	log.Printf("field: %v", field)

	tagreturn := field.Tag.Get(tag)
	log.Printf("tagreturn: %v", tagreturn)

	return tagreturn
}
