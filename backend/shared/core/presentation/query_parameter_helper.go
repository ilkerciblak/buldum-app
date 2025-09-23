package corepresentation

import (
	"net/http"
	"reflect"
)

func PathValuesMapper(r *http.Request, target any) map[string]string {
	valMap := make(map[string]string)
	v := reflect.ValueOf(target)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("path")
		if tag == "" {
			continue
		}

		val := r.PathValue(tag)
		if val == "" {
			continue
		}

		valMap[tag] = val

	}

	return valMap
}

func QueryParametersMapper(r *http.Request, target any) map[string]interface{} {
	valMap := make(map[string]interface{})

	v := reflect.ValueOf(target)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}
		val := r.URL.Query()[tag]
		if len(val) == 0 || val[0] == "" {
			continue
		}
		if field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Array {

			valMap[tag] = val
		} else {
			valMap[tag] = val[0]
		}

	}
	return valMap
}
