package devide

import "testing"

func TestDevidewithSub(t *testing.T) {
	testArray := [][]int{
		{10, 3},
		{9, 2},
		{-5, 3},
		{15, 2},
		{17, 1},
		{-10, -3},
		{5, -1},
	}
	expected := []int{3, 4, -1, 7, 17, 3, -5}
	for index, value := range testArray {
		resOfMyFunc := DevidewithSub(value[0], value[1])
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
		}
	}
}

func TestDivide(t *testing.T) {
	testArray := [][]int{
		{10, 3},
		{9, 2},
		{-5, 3},
		{15, 2},
		{17, 1},
		{-10, -3},
		{5, -1},
	}
	expected := []int{3, 4, -1, 7, 17, 3, -5}
	for index, value := range testArray {
		resOfMyFunc := Divide(value[0], value[1])
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
		}
	}
}

func BenchmarkDevidewithSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DevidewithSub(-1000000, 2)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(17, 1)
	}
}
