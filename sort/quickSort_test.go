package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestQuickSort(t *testing.T)  {
	Convey("快速排序",t, func() {
		res := QuickSort(testArray,0,6)
		for i,v := range res{
			So(v,ShouldEqual,resArray[i])
		}
		
	})
}