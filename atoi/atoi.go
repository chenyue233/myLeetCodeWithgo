package atoi

import (
	"math"
	"strings"
)

// MyAtoi 力扣第8题
// 字符串转化为整数
func MyAtoi(str string) int {
	// 通过无符号整数来计算最大最小值
	// maxInt := int32(^uint32(0) >> 1)
	// minInt := ^maxInt
	res := 0
	flags := 1
	i := 0
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}
	if str[i] == '-' {
		i++
		flags = -1
	} else if str[i] == '+' {
		i++
		flags = 1
	}

	for ; i < len(str); i++ {
		if res*flags >= math.MaxInt32 {
			return math.MaxInt32
		} else if res*flags <= math.MinInt32 {
			return math.MinInt32
		}
		if str[i] < '0' || string(str[i]) > "9" {
			return res * flags
		}
		// int(str[i])string转化int中是转化为ascii码，需要减去0的ascii码 48
		res = res*10 + int(str[i]) - '0'
	}
	if flags*res >= math.MaxInt32 {
		return math.MaxInt32
	} else if flags*res <= math.MinInt32 {
		return math.MinInt32
	}

	return flags * res
}
