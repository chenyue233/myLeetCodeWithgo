package strstr

import (
	"strings"
)

// 力扣第28题,实现strStr
// strings包的函数实现
func strStr(haystack string, needle string) int {
	res := strings.Index(haystack, needle)
	return res
}

// 自己实现
func mYstrStr(haystack string, needle string) int {
	needlelen := len(needle)
	if needlelen == 0 {
		return 0
	}
	for index, _ := range haystack {
		if haystack[index:needlelen] == needle {
			return index
		}
	}
}
