package support

import (
	"regexp"
	"strings"
)

// 从原字符串中获取 CJK 字符串，删除非 CJK 字符串
func GetCJK(s string) string {
	return regexp.MustCompile("[^\u4e00-\u9fff]").ReplaceAllString(s, "")
}

func HTTP2HTTPS(s string) string {
	return strings.Replace(s, "http://", "https://", -1)
}

func SnakeCased(s string) string {
	newStr := make([]rune, 0)

	for idx, chr := range s {
		if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
			if idx > 0 {
				newStr = append(newStr, '_')
			}
			chr -= 'A' - 'a'
		}
		newStr = append(newStr, chr)
	}

	return string(newStr)
}
