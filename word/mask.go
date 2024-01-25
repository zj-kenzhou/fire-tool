package word

import "strings"

// MaskString 脱敏数据
func MaskString(input string, start *int32, end *int32) string {
	if start == nil && end == nil {
		return input
	}
	inputLen := len(input)
	if start == nil && end != nil {
		if inputLen > int(*end) {
			return strings.Repeat("*", int(*end)) + input[int(*end):]
		}
	}
	if end == nil && start != nil {
		if inputLen > int(*start) {
			return input[:int(*start)] + strings.Repeat("*", inputLen-int(*start))
		}
	}
	if start != nil && end != nil && *start < *end {
		if inputLen > int(*start) && inputLen <= int(*end) {
			return input[:int(*start)] + strings.Repeat("*", inputLen-int(*start))
		}
		if inputLen > int(*end) {
			return input[:int(*start)] + strings.Repeat("*", int(*end)-int(*start)) + input[int(*end):]
		}
	}
	return input
}
