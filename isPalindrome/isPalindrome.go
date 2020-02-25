package ispalindrome

// IsPalindrome 力扣第九题回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	myx := x
	res := 0
	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}
	return myx == res
}
