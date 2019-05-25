package utils

import "reflect"

func StructName(v interface{}) string {
	name := reflect.TypeOf(v).String()
	if reflect.ValueOf(v).Type().Kind() == reflect.Ptr {
		return name[1:]
	}
	return name
}
