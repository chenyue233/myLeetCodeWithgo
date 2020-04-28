package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHeadSort(t *testing.T)  {
	Convey("堆排序",t, func() {
		HeadSort(testArray)
		for i,v := range testArray{
			So(v,ShouldEqual,resArray[i])
		}
		
	})
}
