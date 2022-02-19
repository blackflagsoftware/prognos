package util

import (
	"reflect"
)

func GetTypeCount(i interface{}) int {
	switch reflect.ValueOf(i).Kind() {
	case reflect.Map:
		return reflect.ValueOf(i).Len()
	case reflect.Array:
		return reflect.ValueOf(i).Len()
	case reflect.Slice:
		return reflect.ValueOf(i).Len()
	default:
		return 1
	}
}
