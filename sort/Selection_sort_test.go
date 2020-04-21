package sort

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelection(t *testing.T) {
	Convey("选择排序",t, func() {
			res := Selection(testArray)
		for i,v := range res{
			So(v,ShouldEqual,resArray[i])
		}
	})
}
