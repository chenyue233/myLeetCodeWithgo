package binarytree

import "testing"

func TestIsValidBST(t *testing.T) {
	testNode := TreeNode{1, &TreeNode{1, nil, nil}, nil}
	expected := true
	res := isValidBST(&testNode)

	if res != expected {
		t.Errorf("expected '%v',but got '%v'", expected, res)
	}

}

func testIsAddSlice(t *testing.T) {
	testArray := []int{1, 1}
	expected := false
	res := isAddSlice(testArray)
	if res != expected {
		t.Errorf("expected '%v',but got '%v'", expected, res)
	}
}
