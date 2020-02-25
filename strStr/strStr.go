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

// MYstrStr 自己实现
func MYstrStr(haystack string, needle string) int {
	needlelen := len(needle)
	haylen := len(haystack)
	j := 0
	i := 0
	if needlelen == 0 {
		return 0
	}

	for i < haylen && j < needlelen {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == needlelen {
		return i - j
	}
	return -1

}
