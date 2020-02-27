package binarytree

// 力扣94 中序遍历

// TreeNode 定义二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	res = inorderTraversal(root.Left)
	res = append(res, root.Val)
	// ... 语法糖  切片打散传入
	res = append(res, inorderTraversal(root.Right)...)
	return res

}

func myinorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) *[]int {
	if node == nil {
		return res
	}
	dfs(node.Left, res)
	*res = append(*res, node.Val)
	dfs(node.Right, res)
	return res
}

func inorderTraversalWithFor(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
