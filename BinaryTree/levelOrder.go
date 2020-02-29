package binarytree

// levelOrder 二叉树的层次遍历
// 力扣102
// BFS
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		var stacks []*TreeNode
		var curLevel []int
		for _, value := range stack {
			curLevel = append(curLevel, value.Val)
			if value.Left != nil {
				stacks = append(stacks, value.Left)
			}
			if value.Right != nil {
				stacks = append(stacks, value.Right)
			}
		}
		res = append(res, curLevel)
		stack = stacks
	}
	return res
}

// 递归
func levelOrders(root *TreeNode) [][]int {
	return nil
}
