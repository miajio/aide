package excel

import (
	"encoding/json"
	"reflect"
)

func readExcelObj(obj any) ([]ExcelVal, error) {
	result := []ExcelVal{}

	valueOf := reflect.ValueOf(obj)
	typeOf := reflect.TypeOf(obj)
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
		typeOf = typeOf.Elem()
	}
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Field(i)
		tagVal := typeOf.Field(i).Tag.Get("excel")
		fieldName := typeOf.Field(i).Name
		if tagVal != "" && tagVal != "-" && field.IsValid() && field.CanInterface() {
			excelTag, err := ParseExcelTag(fieldName, tagVal)
			if err != nil {
				return nil, err
			}

			res := ExcelVal{ExcelTag: excelTag}
			// 如果是结构体，直接json解析
			if field.Kind() == reflect.Struct {
				jsonStr, _ := json.Marshal(field.Interface())
				res.Val = jsonStr
				continue
			}
			res.Val = field.Interface()
			result = append(result, res)
		}
	}
	return result, nil
}
