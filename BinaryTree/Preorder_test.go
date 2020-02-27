package binarytree

import (
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{1, 2, 3}
	res := PreorderTraversal(&testNode)
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

func TestMyPreorderTraversal(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{1, 2, 3}
	res := MyPreorderTraversal(&testNode)
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

func TestPreorderTraversalWithFor(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{1, 2, 3}
	res := PreorderTraversalWithFor(&testNode)
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
