package excel_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/miajio/aide/pkg/excel"
)

type Person struct {
	Name string `excel:"1; name :姓名;default :未知;format  : 大写 ;color: 红色;bgColor:蓝色; box: true;boxColor:绿色"`
	Age  int    `excel:"2;name:年龄;default:0"`
	City string `excel:"column:3;name:城市"`
}

func TestParseExcelTag(et *testing.T) {
	p := Person{"John", 30, "New York"}
	t := reflect.TypeOf(p)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldName := field.Name
		excelTag, err := excel.ParseExcelTag(fieldName, field.Tag.Get("excel"))
		if err != nil {
			et.Fatal(err)
		}

		fmt.Printf("Field: %s, Column: %d, Name: %s, Default: %s, Format: %s, Color: %s, BgColor: %s, Box: %v, BoxColor: %s\n", field.Name, excelTag.Column, excelTag.Name, excelTag.Default, excelTag.Format, excelTag.Color, excelTag.BgColor, excelTag.Box, excelTag.BoxColor)

	}
}
