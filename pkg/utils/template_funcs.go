package utils

import (
	"fmt"
	"html/template"
	"reflect"
)

type Ms = map[string]interface{}

var FuncMapCommon = template.FuncMap{
	"map":      TemplateFuncMap,
	"hasField": HasField,
	"getf":     Getf,
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

func Getf(obj interface{}, fieldName string, def interface{}) interface{} {
	rv := reflect.ValueOf(obj)
	// It might be a pointer to a pointer to a pointer ...
	// Who knows ¯\_(ツ)_/¯
	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	fmt.Println(rv.Kind())
	switch rv.Kind() {
	case reflect.Map:
		fmt.Println(obj)
		m, ok := obj.(Ms)
		fmt.Println(m, ok)
		if !ok {
			return def
		}
		v, ok := m[fieldName]
		if !ok {
			return def
		}
		return v
	case reflect.Struct:
		if rv.FieldByName(fieldName).IsValid() {
			return reflect.Indirect(rv).FieldByName(fieldName)
		}
		return def
	default:
		return def
	}
}
