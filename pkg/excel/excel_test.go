package excel_test

import (
	"testing"
	"time"

	"github.com/miajio/aide/pkg/excel"
	"github.com/miajio/aide/pkg/stream"
	"github.com/miajio/aide/pkg/utils"
)

type User struct {
	ID    int    `excel:"column:1;name:ID;default:0;box:true;boxColor:000000"`
	Name  string `excel:"column:2;name:Name;default:Miajio;box:true;boxColor:000000"`
	Age   int    `excel:"column:3;name:Age;default:0;box:true;boxColor:000000"`
	Sex   int    `excel:"column:4;name:Sex;default:Man;format:SexFormat;box:true;boxColor:000000"`
	Berth int64  `excel:"column:5;name:Berth; format:UnixDateFormatYMD;box:true;boxColor:000000"`
}

func TestExcel(t *testing.T) {

	b2, _ := time.Parse("2006-01-02", "1995-11-02")

	us := stream.NewSlice()
	us.Add(User{ID: 1, Age: 18, Sex: 1, Berth: time.Now().Unix()})
	us.Add(User{ID: 2, Name: "Kevin", Age: 19, Sex: 1, Berth: b2.Unix()})
	us.Add(User{ID: 3, Name: "Chris", Age: 20, Sex: 2, Berth: b2.Unix()})

	f := excel.New()
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
