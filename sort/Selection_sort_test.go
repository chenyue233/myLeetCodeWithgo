package sort

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelection(t *testing.T) {
	Convey("插入排序",t, func() {
			res := Selection(testArray)
			fmt.Println(res)
		for i,v := range res{
			So(v,ShouldEqual,resArray[i])
		}
	})
}
