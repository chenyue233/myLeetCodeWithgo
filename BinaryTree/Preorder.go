package binarytree

// PreorderTraversal 二叉树的前序遍历
func PreorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, PreorderTraversal(root.Left)...)
	res = append(res, PreorderTraversal(root.Right)...)
	return res
}

// MyPreorderTraversal 二叉树的前序遍历
func MyPreorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	dfsss(root, &res)
	return res

}

func dfsss(root *TreeNode, res *[]int) *[]int {
	if root == nil {
		return res
	}
	*res = append(*res, root.Val)
	dfsss(root.Left, res)
	dfsss(root.Right, res)
	return res
}

// PreorderTraversalWithFor 二叉树的前序遍历
func PreorderTraversalWithFor(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	if root == nil {
		return res
	}
	for len(stack) > 0 || root != nil {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
