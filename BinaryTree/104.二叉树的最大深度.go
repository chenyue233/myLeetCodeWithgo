package binarytree

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftint := maxDepth(root.Left)
	Rightint := maxDepth(root.Right)
	maxint := func(left, right int) int {
		if left > right {
			return left + 1
		}
		return right + 1
	}
	return maxint(leftint, Rightint)

}

// @lc code=end
func maxDepthh(root *TreeNode) int {
	var stack []*TreeNode
	stack = append(stack, root)
	res := 0
	for len(stack) > 0 {
		for i := 0; i < len(stack); i++ {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		res++

	}
	return res
}
