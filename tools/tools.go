package tools

// 字符串逆序
func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 字符串逆序（仅限ascii字符，无中文等）
func ReverseStrASCII(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

// 是否有效成对括号
func IsBracketsPair(left, right byte) bool {
	result := false
	switch left {
	case '(':
		result = right == ')'
	case '[':
		result = right == ']'
	case '{':
		result = right == '}'
	default:
	}
	return result
}
