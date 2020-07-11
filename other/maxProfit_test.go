package other

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var v = 5

func TestInsertionSort(t *testing.T) {
	Convey("插入排序", t, func() {
		res := maxProfit([]int{1, 2, 1, 2, 4, 5, 6, 1, 2})

		So(v, ShouldEqual, res)
		fmt.Println(res)

	})
}
