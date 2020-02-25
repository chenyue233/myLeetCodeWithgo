package strstr

import (
	"testing"
)

func teststrStr(t *testing.T) {
	testArray := [][]string{
		{"aaab", "aa"},
		{"bbbc", "d"},
		{"ssfddass", "dd"},
		{"", "a"},
	}
	expected := []int{0, -1, 3, -1}
	for index, value := range testArray {
		resOfMyFunc := strStr(value[0], value[1])
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
		}
	}
}
