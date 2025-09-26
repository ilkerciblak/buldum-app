package corepresentation

import (
	"maps"
	"net/http"
	"reflect"
	"strconv"
)

func QueryAndPathParametersMapper(r *http.Request, target any) map[string]interface{} {
	valMap := make(map[string]interface{})

	for k, val := range QueryParametersMapper(r, target) {
		valMap[k] = val
	}

	for k, val := range PathValuesMapper(r, target) {
		valMap[k] = val
	}

	return valMap

}

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
		fieldValue := v.Field(i)
		if field.Type.Kind() == reflect.Struct {
			nestedMap := QueryParametersMapper(r, fieldValue.Interface())
			maps.Copy(valMap, nestedMap)
		}

		tag := field.Tag.Get("query")
		if tag == "" {
			continue
		}
		val := r.URL.Query()[tag]
		if len(val) == 0 || val[0] == "" {
			continue
		}
		if field.Type.Kind() == reflect.Slice || field.Type.Kind() == reflect.Array {
			valMap[tag] = convertSliceValues(val, field.Type.Elem())
		} else {
			valMap[tag] = convertSingleValue(val[0], field.Type)
		}

	}

	return valMap
}

func convertSingleValue(value string, targetType reflect.Type) interface{} {
	switch targetType.Kind() {
	case reflect.String:
		return value
	case reflect.Bool:
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	case reflect.Int:
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	case reflect.Int8:
		if parsed, err := strconv.ParseInt(value, 10, 8); err == nil {
			return int8(parsed)
		}
	case reflect.Int16:
		if parsed, err := strconv.ParseInt(value, 10, 16); err == nil {
			return int16(parsed)
		}
	case reflect.Int32:
		if parsed, err := strconv.ParseInt(value, 10, 32); err == nil {
			return int32(parsed)
		}
	case reflect.Int64:
		if parsed, err := strconv.ParseInt(value, 10, 64); err == nil {
			return parsed
		}
	case reflect.Uint:
		if parsed, err := strconv.ParseUint(value, 10, 0); err == nil {
			return uint(parsed)
		}
	case reflect.Uint8:
		if parsed, err := strconv.ParseUint(value, 10, 8); err == nil {
			return uint8(parsed)
		}
	case reflect.Uint16:
		if parsed, err := strconv.ParseUint(value, 10, 16); err == nil {
			return uint16(parsed)
		}
	case reflect.Uint32:
		if parsed, err := strconv.ParseUint(value, 10, 32); err == nil {
			return uint32(parsed)
		}
	case reflect.Uint64:
		if parsed, err := strconv.ParseUint(value, 10, 64); err == nil {
			return parsed
		}
	case reflect.Float32:
		if parsed, err := strconv.ParseFloat(value, 32); err == nil {
			return float32(parsed)
		}
	case reflect.Float64:
		if parsed, err := strconv.ParseFloat(value, 64); err == nil {
			return parsed
		}
	}
	// Fallback to string if conversion fails
	return value
}

// convertSliceValues converts string slice to match the exact slice element type
func convertSliceValues(values []string, elemType reflect.Type) interface{} {
	switch elemType.Kind() {
	case reflect.String:
		return values
	case reflect.Bool:
		result := make([]bool, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseBool(v); err == nil {
				result = append(result, parsed)
			}
		}
		return result
	case reflect.Int:
		result := make([]int, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.Atoi(v); err == nil {
				result = append(result, parsed)
			}
		}
		return result
	case reflect.Int8:
		result := make([]int8, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseInt(v, 10, 8); err == nil {
				result = append(result, int8(parsed))
			}
		}
		return result
	case reflect.Int16:
		result := make([]int16, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseInt(v, 10, 16); err == nil {
				result = append(result, int16(parsed))
			}
		}
		return result
	case reflect.Int32:
		result := make([]int32, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseInt(v, 10, 32); err == nil {
				result = append(result, int32(parsed))
			}
		}
		return result
	case reflect.Int64:
		result := make([]int64, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseInt(v, 10, 64); err == nil {
				result = append(result, parsed)
			}
		}
		return result
	case reflect.Uint:
		result := make([]uint, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseUint(v, 10, 0); err == nil {
				result = append(result, uint(parsed))
			}
		}
		return result
	case reflect.Uint8:
		result := make([]uint8, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseUint(v, 10, 8); err == nil {
				result = append(result, uint8(parsed))
			}
		}
		return result
	case reflect.Uint16:
		result := make([]uint16, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseUint(v, 10, 16); err == nil {
				result = append(result, uint16(parsed))
			}
		}
		return result
	case reflect.Uint32:
		result := make([]uint32, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseUint(v, 10, 32); err == nil {
				result = append(result, uint32(parsed))
			}
		}
		return result
	case reflect.Uint64:
		result := make([]uint64, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseUint(v, 10, 64); err == nil {
				result = append(result, parsed)
			}
		}
		return result
	case reflect.Float32:
		result := make([]float32, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseFloat(v, 32); err == nil {
				result = append(result, float32(parsed))
			}
		}
		return result
	case reflect.Float64:
		result := make([]float64, 0, len(values))
		for _, v := range values {
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				result = append(result, parsed)
			}
		}
		return result
	}
	// Fallback to string slice
	return values
}
