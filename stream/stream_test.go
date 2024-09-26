package stream_test

import (
	"testing"

	"github.com/miajio/aide/aide/stream"
)

// 测试 NewSlice 函数的正常情况
func TestNewSlice(t *testing.T) {
	// 定义一些测试数据
	datas1 := []any{1, 2, 3}
	datas2 := []any{"a", "b", "c"}

	// 创建切片实例
	s1 := stream.NewSlice(datas1...)
	s2 := stream.NewSlice(datas2...)

	// 验证切片的属性
	if s1.Size() != len(datas1) {
		t.Errorf("Expected size %d, got %d", len(datas1), s1.Size())
	}
	if s2.Size() != len(datas2) {
		t.Errorf("Expected size %d, got %d", len(datas2), s2.Size())
	}
}

// 测试 NewSlice 函数的边界情况
func TestNewSliceEmpty(t *testing.T) {
	// 创建空切片实例
	s := stream.NewSlice()

	// 验证切片的属性
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

// TestAdd 测试 Add 方法
func TestAdd(t *testing.T) {
	s := &stream.Slice{}

	// 测试添加一个元素
	s.Add(1)
	if len(s.ToList(nil)) != 1 || s.Size() != 1 {
		t.Errorf("添加一个元素失败")
	}

	// 测试添加多个元素
	s.Add(2)
	s.Add(3)
	if len(s.ToList(nil)) != 3 || s.Size() != 3 {
		t.Errorf("添加多个元素失败")
	}

	// 测试添加 nil 元素
	s.Add(nil)

	// 验证切片的属性
	if s.Size() != 4 {
		t.Errorf("Expected size 4, got %d", s.Size())
	}

	// 验证切片的内容
	if s.Get(0) != 1 || s.Get(1) != 2 || s.Get(2) != 3 || s.Get(3) != nil {
		t.Errorf("切片的内容错误")
	}

	if !s.Contains(nil) {
		t.Errorf("切片中不包含 nil 元素")
	}
}

func TestSlice_AddAll(t *testing.T) {
	// 创建一个测试用的 Slice 实例
	s := &stream.Slice{}

	// 测试空数据添加
	s.AddAll()
	if len(s.ToList(nil)) != 0 || s.Size() != 0 {
		t.Errorf("添加空数据后，长度或大小不应改变")
	}

	// 测试添加单个数据
	s.AddAll(1)
	if len(s.ToList(nil)) != 1 || s.Size() != 1 {
		t.Errorf("添加单个数据后，长度或大小不正确")
	}

	// 测试添加多个数据
	s.AddAll(2, 3, 4)
	if len(s.ToList(nil)) != 4 || s.Size() != 4 {
		t.Errorf("添加多个数据后，长度或大小不正确")
	}
}

func TestSlice_Get(t *testing.T) {
	s := stream.NewSlice(1, 2, 3, 4, 5)

	// 正常情况
	val := s.Get(0)
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	// 边界情况：索引超出范围
	val = s.Get(len(s.ToList(nil)))
	if val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}

func TestSet(t *testing.T) {
	// 创建测试用的 Slice 实例
	slice := stream.NewSlice()

	// 边界情况：索引超出范围
	slice = slice.Set(-1, "invalid")
	if slice == nil {
		t.Errorf("Expected error for invalid index, got nil")
	}

	// 正常情况：设置数据
	slice.Set(0, "data1")
	if slice.Get(0) != "data1" {
		t.Errorf("Expected data1, got %v", slice.Get(0))
	}

	slice.Set(1, "data2")
	if slice.Get(1) != "data2" {
		t.Errorf("Expected data2, got %v", slice.Get(1))
	}

	slice = slice.Set(slice.Size(), "invalid2")
	if slice == nil {
		t.Errorf("Expected error for invalid index, got nil")
	}
}

// 测试正常删除中间元素
func TestRemoveMiddle(t *testing.T) {
	s := stream.NewSlice(1, 2, 3, 4, 5)
	result := s.Remove(2)
	expected := stream.NewSlice(1, 2, 4, 5)
	if !sliceEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试删除头部元素
func TestRemoveHead(t *testing.T) {
	s := stream.NewSlice(1, 2, 3, 4, 5)
	result := s.Remove(0)
	expected := stream.NewSlice(2, 3, 4, 5)
	if !sliceEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试删除尾部元素
func TestRemoveTail(t *testing.T) {
	s := stream.NewSlice(1, 2, 3, 4, 5)
	result := s.Remove(4)
	expected := stream.NewSlice(1, 2, 3, 4)
	if !sliceEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 测试删除越界索引
func TestRemoveOutOfBounds(t *testing.T) {
	s := stream.NewSlice(1, 2, 3, 4, 5)
	result := s.Remove(10)
	expected := s
	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// 判断两个切片是否相等的辅助函数
func sliceEqual(s1, s2 *stream.Slice) bool {
	if s1.Size() != s2.Size() {
		return false
	}
	data1 := s1.ToList(nil)
	data2 := s2.ToList(nil)

	for i := range data1 {
		if data1[i] != data2[i] {
			return false
		}
	}
	return true
}

// 测试 IsEmpty 函数
func TestIsEmpty(t *testing.T) {
	// 创建一个空的切片实例
	slice := stream.NewSlice()

	// 验证空切片的 IsEmpty 结果应该为 true
	if !slice.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true for an empty slice")
	}

	// 添加一个元素到切片
	slice.Add(1)

	// 验证非空切片的 IsEmpty 结果应该为 false
	if slice.IsEmpty() {
		t.Errorf("Expected IsEmpty to be false for a non-empty slice")
	}
}

// Contains 函数的测试用例
func TestContains(t *testing.T) {
	// 创建示例切片
	s := stream.NewSlice(1, 2, 3, 4, 5)

	// 测试包含的数据
	if !s.Contains(3) {
		t.Errorf("Expected true, got false")
	}

	// 测试不包含的数据
	if s.Contains(6) {
		t.Errorf("Expected false, got true")
	}
}

// 测试 IndexOf 函数
func TestIndexOf(t *testing.T) {
	// 创建 Slice 实例
	slice := stream.NewSlice()

	// 测试空切片
	index := slice.IndexOf(1)
	if index != -1 {
		t.Errorf("Expected -1 for empty slice, got %d", index)
	}

	// 添加元素
	slice.Add(1)
	slice.Add(2)
	slice.Add(3)

	// 测试存在的元素
	index = slice.IndexOf(2)
	if index != 1 {
		t.Errorf("Expected index 1 for element 2, got %d", index)
	}

	// 测试不存在的元素
	index = slice.IndexOf(4)
	if index != -1 {
		t.Errorf("Expected -1 for non-existing element, got %d", index)
	}
}

func TestSlice(t *testing.T) {
	// 创建一个测试切片
	slice := stream.NewSlice(1, 2, 3, 4, 5)
	slice.ForEach(func(data any) {
		if data.(int)%2 == 0 {
			slice.Remove(slice.IndexOf(data))
		}
	}).Filter(func(data any) bool {
		return data.(int) > 3
	}).Add(1).Sort(func(i, j any) bool {
		return i.(int) > j.(int)
	})

	if !slice.Contains(1) && !slice.Contains(5) {
		t.Errorf("Expected true, got false")
	}

	m := slice.ToMap(func(data any) stream.KV {
		var res stream.KV
		if data.(int) == 1 {
			res.Key = "a"
			res.Value = 1
		} else {
			res.Key = "b"
			res.Value = 5
		}
		return res
	})
	if m["a"] != 1 && m["b"] != 5 {
		t.Errorf("Expected true, got false")
	}

}
