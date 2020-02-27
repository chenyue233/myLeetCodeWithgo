package math

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	testStrings := [...]string{"2000000000000", "-2000000000000", "42", "-42", "-000001", "0000001", "q-1", "a--2", "    1"}
	maxInt := math.MaxInt32
	minInt := math.MinInt32
	expected := [...]int{maxInt, minInt, 42, -42, -1, 1, 0, 0, 1}
	for index, value := range testStrings {
		resOfMyFunc := MyAtoi(value)
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
		}
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyAtoi("-000001")
	}
}
