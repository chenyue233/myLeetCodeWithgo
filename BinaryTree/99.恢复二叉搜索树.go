package binarytree

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
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
func recoverTree(root *TreeNode) {
	var res []int
	helpers(root, &res)
	sort.Ints(res)
	cur := 0
	inorderFill(root, res, &cur)

}

func helpers(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	helpers(root.Left, res)
	*res = append(*res, root.Val)
	helpers(root.Right, res)
}

func inorderFill(root *TreeNode, array []int, current *int) {
	if root == nil {
		return
	}
	inorderFill(root.Left, array, current)
	root.Val = array[*current]
	*current++
	inorderFill(root.Right, array, current)
}

func recoverTrees(root *TreeNode) {
	cur := root
	var pre *TreeNode
	var firstNode *TreeNode
	var secondNode *TreeNode
	for cur != nil {
		if cur.Left == nil {
			if pre != nil && pre.Val > cur.Val {
				if firstNode == nil {
					firstNode = pre
				}
				secondNode = cur
			}
			pre = cur
			cur = cur.Right
		} else {
			leftRight := cur.Left
			for leftRight.Right != nil && leftRight.Right != cur {
				leftRight = leftRight.Right
			}
			if leftRight.Right != cur {
				leftRight.Right = cur
				cur = cur.Left
			} else {
				leftRight.Right = nil
				if pre != nil && pre.Val > cur.Val {
					if firstNode == nil {
						firstNode = pre
					}
					secondNode = cur
				}
				pre = cur
				cur = cur.Right
			}
		}

	}
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
}
