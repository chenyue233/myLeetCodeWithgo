package binarytree

// 力扣 103 二叉树的锯齿形遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	Level := 0
	var deque []*TreeNode
	deque = append(deque, root)
	for len(deque) > 0 {
		var levelNodes []*TreeNode
		var curLevel []int
		for _, value := range deque {
			if Level%2 == 0 {
				curLevel = append(curLevel, value.Val)
			} else {
				curLevel = append([]int{value.Val}, curLevel...)
			}
			if value.Left != nil {
				levelNodes = append(levelNodes, value.Left)
			}
			if value.Right != nil {
				levelNodes = append(levelNodes, value.Right)
			}
		}
		res = append(res, curLevel)
		deque = levelNodes
		Level++
	}
	return res
}
