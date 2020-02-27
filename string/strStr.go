package string

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
	// 边界值
	if needlelen == 0 {
		return 0
	}
	// i或者j只要其中一个超过了字符串的长度立刻就结束循环
	for i < haylen && j < needlelen {
		// 两个字符串的从0开始，只要某个字符相等就比较下一个字符
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			// 如果不相等，haystack下一个字符与needle第一个字符串对比，直到相等
			// 如果对比到一半忽然不相等了，从新开始对比，需要用i-j=0
			i = i - j + 1
			j = 0
		}
	}
	// j是haylen与neelen对比成功的次数，如果成功的次数和neelen的长度相等就返回：总循环的次数i减去成功的次数，就是开始成功的地方
	if j == needlelen {
		return i - j
	}
	return -1

}
