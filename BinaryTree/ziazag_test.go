package binarytree

import "testing"

func TestZigzagLevelOrder(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	expected := [][]int{
		{1},
		{3, 2},
		{4, 5, 6},
	}
	res := zigzagLevelOrder(&testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (&expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}
	// 防止数组越界
	ress := res[:len(expected)]
	for index, value := range expected {
		for i, v := range value {
			if v != ress[index][i] {
				t.Errorf("expected '%d',but got '%d'", expected, res)
			}
		}

	}
}
