package binarytree

import "testing"

func TestLevelOrder(t *testing.T) {
	testNode := &TreeNode{1,
		&TreeNode{2, &TreeNode{2, nil, nil}, nil},
		&TreeNode{3, &TreeNode{1, nil, nil}, nil},
	}
	expected := [][]int{{1}, {2, 3}, {2, 1}}
	res := levelOrder(testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (&expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}
	// 防止数组越界
	ress := res[:len(expected)]
	for index, value := range expected {
		for i, value := range value {
			if value != ress[index][i] {
				t.Errorf("expected '%d',but got '%d'", expected, res)
			}
		}
	}
}
