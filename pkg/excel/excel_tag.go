package excel

import (
	"fmt"
	"strconv"
	"strings"
)

type ExcelTag struct {
	Column   int    // 列号
	Name     string // 列名
	Default  string // 默认值
	Format   string // 格式化方法名
	Color    string // 字体颜色
	BgColor  string // 背景颜色
	Box      bool   // 是否有边框
	BoxColor string // 边框颜色
}

func ParseExcelTag(fieldName, tag string) (*ExcelTag, error) {
	excelTag := &ExcelTag{}
	tag = strings.ReplaceAll(tag, "=", ":")

	parts := strings.Split(tag, ";")
	for _, part := range parts {
		keyValue := strings.Split(part, ":")
		if len(keyValue) == 2 {
			key := keyValue[0]
			value := keyValue[1]

			// 去除前后的空格
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)

			switch key {
			case "column":
				column, err := strconv.Atoi(value)
				if err == nil {
					excelTag.Column = column
				}
			case "name":
				excelTag.Name = value
			case "default":
				excelTag.Default = value
			case "format":
				excelTag.Format = value
			case "color":
				excelTag.Color = value
			case "bgColor":
				excelTag.BgColor = value
			case "box":
				// 如果没有指定值，默认为false
				if strings.ToUpper(value) == "TRUE" {
					excelTag.Box = true
				}
			case "boxColor":
				excelTag.BoxColor = value
			}
		}
	}
	if excelTag.Column == 0 {
		column, _ := strconv.Atoi(parts[0])
		excelTag.Column = column
		if excelTag.Column == 0 {
			return nil, fmt.Errorf("invalid excel tag: %s", tag)
		}
	}
	return excelTag, nil
}
