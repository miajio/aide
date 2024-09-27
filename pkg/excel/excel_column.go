package excel

// ExcelVal 单元格值
type ExcelVal struct {
	*ExcelTag
	Val any // 值
}

// ToColumnName 将数字转换为列名
func ToColumnName(n int) string {
	if n <= 0 {
		return ""
	}
	var result []byte
	for n > 0 {
		n--
		result = append(result, byte('A'+n%26))
		n /= 26
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
