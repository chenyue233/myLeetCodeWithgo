package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMergeSort(t *testing.T) {
	Convey("归并排序",t, func() {
		res := Selection(testArray)
		for i,v := range res{
			So(v,ShouldEqual,resArray[i])
		}
	})
}
