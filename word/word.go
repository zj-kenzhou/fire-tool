package word

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// UnderLineToCamelCaseTitle 下划线转驼峰首字母大写
func UnderLineToCamelCaseTitle(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = cases.Title(language.Und).String(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderLineToCamelCase 下划线转驼峰
func UnderLineToCamelCase(s string) string {
	s = UnderLineToCamelCaseTitle(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CenterLineToCamelCase 中划线转驼峰
func CenterLineToCamelCase(s string) string {
	s = strings.Replace(s, "-", " ", -1)
	s = cases.Title(language.Und).String(s)
	s = strings.Replace(s, " ", "", -1)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToCenterLine 驼峰转中划线
func CamelCaseToCenterLine(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '-')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// CamelCaseToUnderLine 驼峰转下划线
func CamelCaseToUnderLine(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
