package utils

import (
	"html/template"
	"reflect"
)

var FuncMapCommon = template.FuncMap{
	"map":      TemplateFuncMap,
	"hasField": HasField,
}

func TemplateFuncMap(els ...any) map[any]any {
	m := make(map[any]any)
	for i := 0; i <= len(m)/2; i++ {
		idx := 2 * i
		m[els[idx]] = els[idx+1]
	}
	return m
}

func HasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}
