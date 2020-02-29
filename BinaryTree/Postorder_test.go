package binarytree

import "testing"

func TestPostorderTraversa(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{2, 3, 1}
	res := PostorderTraversa(&testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (&expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}
	// 防止数组越界
	ress := res[:len(expected)]
	for index, value := range expected {
		if value != ress[index] {
			t.Errorf("expected '%d',but got '%d'", expected, res)
		}
	}
}

func TestMyPostorderTraversa(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{2, 3, 1}
	res := MyPostorderTraversa(&testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (&expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}
	// 防止数组越界
	ress := res[:len(expected)]
	for index, value := range expected {
		if value != ress[index] {
			t.Errorf("expected '%d',but got '%d'", expected, res)
		}
	}
}

func TestPostorder(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{2, 3, 1}
	res := Postorder(&testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (&expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}
	// 防止数组越界
	ress := res[:len(expected)]
	for index, value := range expected {
		if value != ress[index] {
			t.Errorf("expected '%d',but got '%d'", expected, res)
		}
	}
}

func TestReverse(t *testing.T) {
	testNode := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	res := reverse(testNode)
	if len(res) != len(expected) {
		t.Errorf("len(res)!= len(expected),expected '%d',but got '%d'", expected, res)
	}
	if (expected == nil) != (res == nil) {
		t.Errorf("if (&expected == nil) != (res == nil),expected '%d',but got '%d'", expected, res)
	}

	for index, value := range expected {
		if value != res[index] {
			t.Errorf("expected '%d',but got '%d'", expected, res)
		}
	}
}
