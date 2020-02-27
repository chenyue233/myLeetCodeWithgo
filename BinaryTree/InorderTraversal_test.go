package binarytree

import "testing"

type question struct {
	para
	ans
}

// para 是参数
type para struct {
	pre []int
	in  []int
}

// ans 是答案
type ans struct {
	one []int
}

func TestinorderTraversal(t *testing.T) {
	testArray := []question{
		question{
			para{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
			},
			ans{
				[]int{1, 3, 2}
			},
		},
	}
	for _, value := range testArray {
		resOfMyFunc := inorderTraversal(value[0])
		if resOfMyFunc != expected[index] {
			t.Errorf("expected '%d',but got '%d'", expected[index], resOfMyFunc)
			
		}

	}

}
