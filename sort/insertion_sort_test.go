package sort

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var(
	testArray = []int{1,2,31,2,14,42,11}
	resArray = []int{1,2,2,11,14,31,42}
)
func TestInsertionSort(t *testing.T)  {
	Convey("插入排序",t, func() {
		res := InsertionSort(testArray)
		for i,v := range res{
			So(v,ShouldEqual,resArray[i])
		}
		
	})
}

func TestInsertionSorts(t *testing.T) {
	testArray := [][]int{
		{1, 2, 1, 15, 25, 111, 2344},
		{111, 2311, 12, 441111, 12111, 121},
		{12452, 422, 122, 22},
		{1231, 2, 3},
		{0, 9, 5, 2, 5, 4, 3},
		{123},
	}
	expected := [][]int{
		{1, 1, 2, 15, 25, 111, 2344},
		{12, 111, 121, 2311, 12111, 441111},
		{22, 122, 422, 12452},
		{2, 3, 1231},
		{0, 2, 3, 4, 5, 5, 9},
		{123},
	}
	for idx,arr := range testArray {
		res := InsertionSort(arr)
		fmt.Println(res)
		for i:=0;i<len(res);i++  {
			if res[i] != expected[idx][i]{
				t.Errorf("expect:%d,but got:%d",res[i],expected[idx][i])
			}
		}
	}
}