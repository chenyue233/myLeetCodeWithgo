package string

import (
	"testing"
)

func TestStrStr(t *testing.T) {
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

func TestMYstrStr(t *testing.T) {
	testArray := [][]string{
		{"aaab", "aa"},
		{"bbbc", "d"},
		{"ssfddass", "dd"},
		{"", "a"},
	}
	expected := []int{0, -1, 3, -1}
	for index, value := range testArray {
		resOfMyFunc := MYstrStr(value[0], value[1])
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
		}
	}
}

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr("he11llll", "ll")
	}
}

func BenchmarkMYstrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr("he11llll", "ll")
	}
}
