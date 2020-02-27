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

	inorderTraversal(root.Left)
	res = append(res, root.Val)
	// ... 语法糖  切片打散传入
	res = append(res, inorderTraversal(root.Right)...)
	return res

}
