package math

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testArray := []int{121, 1111111111111, 2222, -123, 0, 1, 12345}
	expected := []bool{true, true, true, false, true, true, false}
	for index, value := range testArray {
		resOfMyFunc := isPalindrome(value)
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%t',but got '%t'", expected[index], resOfMyFunc)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome(12221)
	}
}
