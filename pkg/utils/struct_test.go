package utils_test

import (
	"fmt"
	"testing"

	"github.com/miajio/aide/pkg/utils"
)

type O1 struct {
	Id int64 `json:"id"`
	O2 O2    `json:"o2"`
}

type O2 struct {
	Name string `json:"-"`
	Age  int    `json:"age"`
}

func TestStructToMap(t *testing.T) {
	res := utils.StructToMap(O1{
		Id: 1,
		O2: O2{
			Name: "test",
			Age:  18,
		},
	}, "json")
	fmt.Println(res)

	res = utils.StructToMap(10, "json")
	fmt.Println(res)

}
