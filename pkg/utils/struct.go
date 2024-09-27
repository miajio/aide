package utils

import "reflect"

// StructToMap 将结构体转换为map
func StructToMap(obj any, tag string) map[string]any {
	result := make(map[string]interface{})
	valueOf := reflect.ValueOf(obj)
	typeOf := reflect.TypeOf(obj)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
	}
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		tagVal := typeOf.Field(i).Tag.Get(tag)
		if tagVal != "" && tagVal != "-" && field.IsValid() && field.CanInterface() {
			// 如果是结构体，递归调用
			if field.Kind() == reflect.Struct {
				result[tagVal] = StructToMap(field.Interface(), tag)
				continue
			}
			result[tagVal] = field.Interface()
		}
	}
	return result
}

// IsZero 判断是否为空
func IsZero(v any) bool {
	return v == nil || reflect.ValueOf(v).IsZero()
}
