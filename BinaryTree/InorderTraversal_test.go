package binarytree

import (
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{2, 1, 3}
	res := inorderTraversal(&testNode)
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

func TestMyInorderTraversal(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{1, nil, nil}, nil}
	expected := [...]int{1, 1}
	res := myinorderTraversal(&testNode)
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

func TestInorderTraversalWithFor(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	expected := [...]int{2, 1, 3}
	res := inorderTraversalWithFor(&testNode)
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
