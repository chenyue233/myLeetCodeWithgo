package devide

// DevidewithSub 用减法做两数相除
// 力扣29 两数相除

func DevidewithSub(dividend int, divisor int) int {
	var flags bool
	switch {
	case divisor > 0 && dividend > 0:
		flags = true
	case divisor > 0 && dividend < 0:
		dividend = -dividend
		flags = false
	case divisor < 0 && dividend < 0:
		divisor = -divisor
		dividend = -dividend
		flags = true
	case divisor < 0 && dividend > 0:
		divisor = -divisor
		flags = false
	default:
		flags = true
	}
	result := 0
	for dividend >= divisor {
		dividend -= divisor
		result++
	}
	if !flags {
		result = -result
	}
	if !(result >= int(-0x80000000) && result <= int(0x7FFFFFFF)) {
		return int(0x7FFFFFFF)
	}
	return result

}

// Divide 两数相除
func Divide(dividend int, divisor int) int {
	res, c := 0, 1
	var flag bool
	switch {
	case dividend > 0 && divisor < 0:
		{
			divisor = -divisor
			flag = false
		}
	case dividend < 0 && divisor > 0:
		{

			dividend = -dividend
			flag = false

		}
	case dividend < 0 && divisor < 0:
		{

			dividend = -dividend
			divisor = -divisor
			flag = true

		}
	default:
		flag = true
	}
	for dividend >= divisor<<1 {
		divisor <<= 1
		c <<= 1
	}
	for c > 0 && dividend > 0 {
		if dividend >= divisor {
			dividend -= divisor
			res += c
		}
		c >>= 1
		divisor >>= 1
	}
	if !flag {
		res = -res
	}
	if !(res >= int(-0x80000000) && res <= int(0x7FFFFFFF)) {
		return int(0x7FFFFFFF)
	}
	return res
}
