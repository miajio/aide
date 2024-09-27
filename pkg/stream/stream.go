package stream

// Slice 切片
// 流式切片中所有操作均会操作原切片, 若需要复制, 请使用Copy方法将原切片进行备份
type Slice struct {
	datas  []any
	size   int
	equals func(a, b any) bool // 相等函数-默认使用 ==
}

// KV 键值对
type KV struct {
	Key   any
	Value any
}

// NewSlice 创建Slice
// 若datas非any, 可能出现异常 cannot use us (variable of type []XXX) as []any value in argument to NewSlice
// 不建议直接使用NewSlice进行初始化结构体数组
func NewSlice(datas ...any) *Slice {
	s := &Slice{datas: datas, size: len(datas)}
	s.equals = s.defaultEquals
	return s
}

// Add 添加
func (s *Slice) Add(data any) *Slice {
	s.datas = append(s.datas, data)
	s.size++
	return s
}

// AddAll 添加多个
func (s *Slice) AddAll(datas ...any) *Slice {
	s.datas = append(s.datas, datas...)
	s.size += len(datas)
	return s
}

// Get 获取
func (s *Slice) Get(index int) any {
	if index < 0 || index >= len(s.datas) {
		return nil
	}
	return s.datas[index]
}

// Set 设置
// 若index超出范围, 则会在末尾添加
// 若index小于0, 则会在开头添加
func (s *Slice) Set(index int, data any) *Slice {
	if index < 0 {
		index = 0
	}
	if index >= len(s.datas) {
		index = len(s.datas)
	}
	s.size++
	s.datas = append(s.datas[:index], append([]any{data}, s.datas[index:]...)...)
	return s
}

// Remove 删除
// 边界情况, 不进行任何处理
func (s *Slice) Remove(index int) *Slice {
	// 边界情况 不进行任何处理
	if index < 0 || index >= len(s.datas) {
		return s
	}
	s.datas = append(s.datas[:index], s.datas[index+1:]...)
	s.size--
	return s
}

// Clear 清空
func (s *Slice) Clear() *Slice {
	s = NewSlice()
	return s
}

// ForEach 遍历
func (s *Slice) ForEach(do func(data any)) *Slice {
	for i := 0; i < len(s.datas); i++ {
		do(s.datas[i])
	}
	return s
}

// ForEachError 遍历
// 若有错误, 则会返回错误
func (s *Slice) ForEachError(do func(data any) error) (*Slice, error) {
	cpDatas := s.Copy()
	for i := 0; i < len(s.datas); i++ {
		if err := do(s.datas[i]); err != nil {
			return cpDatas, err
		}
	}
	return s, nil
}

// ItemForEach 遍历
func (s *Slice) ItemForEach(do func(i int, data any)) *Slice {
	for i := 0; i < len(s.datas); i++ {
		do(i, s.datas[i])
	}
	return s
}

// ItemForEachError 遍历
// 若有错误, 则会返回错误
func (s *Slice) ItemForEachError(do func(i int, data any) error) (*Slice, error) {
	cpDatas := s.Copy()
	for i := 0; i < len(s.datas); i++ {
		if err := do(i, s.datas[i]); err != nil {
			return cpDatas, err
		}
	}
	return s, nil
}

// Sort 排序
func (s *Slice) Sort(sort func(i, j any) bool) *Slice {
	for i := 0; i < len(s.datas); i++ {
		for j := i + 1; j < len(s.datas); j++ {
			if sort(s.datas[i], s.datas[j]) {
				s.datas[i], s.datas[j] = s.datas[j], s.datas[i]
			}
		}
	}
	return s
}

// Filter 过滤
func (s *Slice) Filter(filter func(data any) bool) *Slice {
	for i := len(s.datas) - 1; i >= 0; i-- {
		if !filter(s.datas[i]) {
			s.Remove(i)
		}
	}
	return s
}

// ToMap 映射
func (s *Slice) ToMap(mapper func(data any) KV) map[any]any {
	res := make(map[any]any)
	for i := 0; i < len(s.datas); i++ {
		kv := mapper(s.datas[i])
		res[kv.Key] = kv.Value
	}
	return res
}

// ToList 转换为列表
func (s *Slice) ToList(to func(data any) any) []any {
	if to != nil {
		res := make([]any, 0)
		for i := len(s.datas) - 1; i >= 0; i-- {
			res = append(res, to(s.datas[i]))
		}
		return res
	}
	return s.datas
}

// Copy 复制
func (s *Slice) Copy() *Slice {
	n := NewSlice()
	n.datas = s.datas
	n.size = len(n.datas)
	return n
}

// Size 获取长度
func (s *Slice) Size() int {
	return s.size
}

// IsEmpty 是否为空
func (s *Slice) IsEmpty() bool {
	return s.size == 0
}

// SetEquals 设置相等函数
func (s *Slice) SetEquals(equals func(a, b any) bool) *Slice {
	s.equals = equals
	return s
}

// defaultEquals 默认相等函数
func (s *Slice) defaultEquals(a, b any) bool {
	return a == b
}

// Contains 是否包含
func (s *Slice) Contains(data any) bool {
	return s.IndexOf(data) >= 0
}

// IndexOf 获取索引
func (s *Slice) IndexOf(data any) int {
	return s.LastIndexOfRange(data, 0, s.size)
}

// LastIndexOfRange 范围内获取最后一个索引
func (s *Slice) LastIndexOfRange(data any, start, end int) int {
	es := s.Copy().ToList(nil)
	if data == nil {
		for i := end - 1; i >= start; i-- {
			if es[i] == nil {
				return i
			}
		}
	} else {
		eq := s.defaultEquals
		if s.equals != nil {
			eq = s.equals
		}
		for i := end - 1; i >= start; i-- {
			if eq(data, es[i]) {
				return i
			}
		}
	}
	return -1
}
