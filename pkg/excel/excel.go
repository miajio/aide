package excel

import (
	"fmt"

	"github.com/miajio/aide/pkg/stream"
	"github.com/miajio/aide/pkg/utils"
	"github.com/xuri/excelize/v2"
)

// File 文件
// 基于github.com/xuri/excelize/v2实现
type File struct {
	*excelize.File
	formatMap map[string]FormatFunc
}

// New 创建文件
func New(opts ...excelize.Options) *File {
	return &File{excelize.NewFile(opts...), make(map[string]FormatFunc)}
}

// OpenFile 打开文件
func OpenFile(name string, opts ...excelize.Options) (*File, error) {
	file, err := excelize.OpenFile(name, opts...)
	if err != nil {
		return nil, err
	}
	return &File{file, make(map[string]FormatFunc)}, nil
}

// CreateSheet 创建Sheet
func (f *File) CreateSheet(name string) (int, error) {
	i, err := f.File.NewSheet(name)
	if err != nil {
		return 0, err
	}
	f.File.SetActiveSheet(i)
	return i, nil
}

func (f *File) AddFormatFunc(format string, fn FormatFunc) {
	f.formatMap[format] = fn
}

// WriteSlice 写入Slice
// 当start <= 1 时，写入表头
// sheet: Sheet名称
// start: 开始行
// datas: 数据
func (f *File) WriteSlice(sheet string, start int, datas *stream.Slice) error {
	if nil == datas || datas.IsEmpty() {
		return nil
	}
	row := start

	_, err := datas.ItemForEachError(func(i int, data any) error {
		realRow := row + i
		excelVals, err := readExcelObj(data)
		if err != nil {
			return err
		}

		if 1 >= realRow {
			realRow = 1
			for _, excelVal := range excelVals {
				cell := fmt.Sprintf("%s%d", ToColumnName(excelVal.Column), realRow)

				if err := f.styleSet(sheet, cell, excelVal); err != nil {
					return err
				}

				if err := f.SetCellValue(sheet, cell, excelVal.Name); err != nil {
					return err
				}
			}
			row = 2
			realRow = 2
		}

		for _, excelVal := range excelVals {
			cell := fmt.Sprintf("%s%d", ToColumnName(excelVal.Column), realRow)
			val := excelVal.Val
			if excelVal.Default != "" && utils.IsZero(val) {
				val = excelVal.Default
			}

			if excelVal.Format != "" {
				if fn, ok := f.formatMap[excelVal.Format]; ok {
					val = fn(val)
				}
			}

			if err := f.styleSet(sheet, cell, excelVal); err != nil {
				return err
			}

			if err := f.SetCellValue(sheet, cell, val); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// styleSet 设置样式
func (f *File) styleSet(sheet, cell string, excelVal ExcelVal) error {
	style := &excelize.Style{}
	useStyle := false

	if excelVal.Color != "" {
		useStyle = true
		style.Font = &excelize.Font{Color: excelVal.Color}
	}

	if excelVal.BgColor != "" {
		useStyle = true
		style.Fill = excelize.Fill{Type: "pattern", Pattern: 1, Color: []string{excelVal.BgColor}}
	}

	if excelVal.Box {
		useStyle = true
		if excelVal.BoxColor == "" {
			style.Border = []excelize.Border{
				{Type: "left", Color: "000000", Style: 1},
				{Type: "top", Color: "000000", Style: 1},
				{Type: "right", Color: "000000", Style: 1},
				{Type: "bottom", Color: "000000", Style: 1},
			}
		} else {
			style.Border = []excelize.Border{
				{Type: "left", Color: excelVal.BoxColor, Style: 1},
				{Type: "top", Color: excelVal.BoxColor, Style: 1},
				{Type: "right", Color: excelVal.BoxColor, Style: 1},
				{Type: "bottom", Color: excelVal.BoxColor, Style: 1},
			}
		}
	}

	if useStyle {
		styleID, err := f.NewStyle(style)
		if err != nil {
			return err
		}
		f.SetCellStyle(sheet, cell, cell, styleID)
	}
	return nil
}

// SaveAs 保存为文件
func (f *File) SaveAs(name string, opts ...excelize.Options) error {
	return f.File.SaveAs(name, opts...)
}
