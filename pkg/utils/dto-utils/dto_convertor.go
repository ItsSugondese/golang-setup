package dto_utils

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
)

func ConvertSlice[S any, D any](source []S) []D {
	var dest []D
	for _, srcItem := range source {
		var destItem D
		err := mapstructure.Decode(srcItem, &destItem)
		if err != nil {
			panic(err)
		}
		dest = append(dest, destItem)
	}
	return dest
}

func NullAwareMapDtoConvertor(src interface{}, dest interface{}) {
	srcValue := reflect.ValueOf(src)
	destValue := reflect.ValueOf(dest).Elem()

	if srcValue.Kind() != reflect.Struct || destValue.Kind() != reflect.Struct {
		panic("src and dest must be structs")
	}

	for i := 0; i < srcValue.NumField(); i++ {
		srcField := srcValue.Field(i)
		destField := destValue.FieldByName(srcValue.Type().Field(i).Name)

		if !destField.IsValid() {
			continue
		}

		if srcField.Kind() == reflect.Ptr {
			if srcField.IsNil() {
				// If srcField is nil, do not update destField
				continue
			}
		}

		if srcField.CanSet() {
			if destField.Kind() == reflect.Ptr {
				// Create a new pointer if srcField is non-nil
				if !srcField.IsNil() {
					destField.Set(reflect.New(srcField.Type().Elem()))
					destField.Elem().Set(srcField.Elem())
				}
			} else {
				destField.Set(srcField)
			}
		}
	}
}
