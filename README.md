# aide
[![Go Reference](https://pkg.go.dev/badge/github.com/miajio/aide.svg)](https://pkg.go.dev/github.com/miajio/aide)

golang common util aide

# Install

```
go get github.com/miajio/aide
```

# Usage

see package package test file

# Package

stream: [github.com/miajio/aide/pkg/stream](./pkg/stream)

system: [github.com/miajio/aide/pkg/system](./pkg/system)

utils: [github.com/miajio/aide/pkg/utils](./pkg/utils)

web: [github.com/miajio/aide/pkg/web](./pkg/web)

excel: [github.com/miajio/aide/pkg/excel](./pkg/excel)

# Excel

basic: [github.com/xuri/excelize/v2](https://github.com/xuri/excelize)

Encapsulation based on excelize v2

add excel tag

``` go
// excel tag
type ExcelTag struct {
	Column   int    // column index
	Name     string // column name
	Default  string // column default name
	Format   string // column need format function name
	Color    string // column text color
	BgColor  string // column background color
	Box      bool   // column need box
	BoxColor string // column box color
}

// struct tag example
type User struct {
	ID    int    `excel:"column:1;name:ID;default:0;box:true;boxColor:000000"`
	Name  string `excel:"column:2;name:Name;default:Miajio;box:true;boxColor:000000"`
	Age   int    `excel:"column:3;name:Age;default:0;box:true;boxColor:000000"`
	Sex   int    `excel:"column:4;name:Sex;default:Man;format:SexFormat;box:true;boxColor:000000"`
	Berth int64  `excel:"column:5;name:Berth; format:UnixDateFormatYMD;box:true;boxColor:000000"`
}

func main() {
	b2, _ := time.Parse("2006-01-02", "1995-11-02")
	us := stream.NewSlice()
	us.Add(User{ID: 1, Age: 18, Sex: 1, Berth: time.Now().Unix()})
	us.Add(User{ID: 2, Name: "Kevin", Age: 19, Sex: 1, Berth: b2.Unix()})
	us.Add(User{ID: 3, Name: "Chris", Age: 20, Sex: 2, Berth: b2.Unix()})

	f := excel.New()
    defer f.Close()
	f.AddFormatFunc("UnixDateFormatYMD", excel.UnixDateFormatYMD)
	f.AddFormatFunc("SexToString", func(val any) string {
		if utils.IsZero(val) || val == 1 {
			return "MAN"
		}
		return "WOMAN"
	})
	f.CreateSheet("sheet")
	f.WriteSlice("sheet", 0, us)
	f.SaveAs("test.xlsx")
}
```

# License
Apache License 2.0