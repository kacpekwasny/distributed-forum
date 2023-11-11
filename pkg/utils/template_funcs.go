package utils

import "html/template"

var FuncMapCommon = template.FuncMap{
	"map": TemplateFuncMap,
}

func TemplateFuncMap(els ...any) map[any]any {
	m := make(map[any]any)
	for i := 0; i <= len(m)/2; i++ {
		idx := 2 * i
		m[els[idx]] = els[idx+1]
	}
	return m
}
