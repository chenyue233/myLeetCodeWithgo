package other

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSelection(t *testing.T) {
	Convey("选择排序", t, func() {
		res := pigai(6)
		fmt.Println(res)
	})
}
