package binarytree

// 力扣98 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*TreeNode
	var res []int
	for len(stack) > 0 || root != nil {
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
	if !isAddSlice(res) {
		return false
	}
	return true
}

func isAddSlice(res []int) bool {
	for i := 0; i < len(res)-2; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func isValidBSTs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	Min := -1 << 63
	Max := 1<<63 - 1
	res := helper(Min, Max, root)
	return res
}

func helper(min int, max int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := min < root.Val && root.Val < max
	ress := helper(min, root.Val, root.Left)
	resss := helper(root.Val, max, root.Right)
	return res && ress && resss
}
