package binarytree

// PostorderTraversa 二叉树的后序遍历
func PostorderTraversa(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	res = append(res, PostorderTraversa(root.Left)...)
	res = append(res, PostorderTraversa(root.Right)...)
	res = append(res, root.Val)
	return res
}

// MyPostorderTraversa 二叉树的后序遍历
func MyPostorderTraversa(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	post(root, &res)
	return res
}

func post(root *TreeNode, res *[]int) *[]int {
	if root == nil {
		return res
	}
	post(root.Left, res)
	post(root.Right, res)
	*res = append(*res, root.Val)
	return res
}

func printstack(a []*TreeNode) []int {
	res := []int{}
	for _, value := range a {
		res = append(res, value.Val)
		b := value.Left
		c := value.Right
		if b != nil {
			res = append(res, b.Val)
		}
		if c != nil {
			res = append(res, c.Val)
		}

	}
	return res

}

// PostorderTraversaWithFor  二叉树的后序遍历非递归
// 时间复杂度太高
func PostorderTraversaWithFor(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	pre := &TreeNode{}
	stack = append(stack, root)
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		if (cur.Left == nil && cur.Right == nil) || pre == cur.Right {
			res = append(res, cur.Val)
			pre = cur
			stack = stack[:len(stack)-1]
		} else {
			if cur.Right != nil {
				stack = append(stack, cur.Right)
			}
			if cur.Left != nil {
				stack = append(stack, cur.Left)
			}
		}
	}
	return res
}

// Postorder 二叉树的后序遍历非递归
// 广度优先遍历后  逆序数组
func Postorder(root *TreeNode) []int {
	res := []int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		res = append(res, node.Val)
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
	}

	return reverse(res)
}

func reverse(res []int) []int {
	if len(res) == 0 || len(res) == 1 {
		return res
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

// PostorderTraversaWithFor 二叉树的后序遍历非递归
// func PostorderTraversaWithFor(root *TreeNode) []int {
// 	res := []int{}
// 	stack := []*TreeNode{root}
// 	var preVisited *TreeNode
// 	for len(stack) > 0 {
// 		if (root.Left == nil && root.Right == nil) || (root.Right == preVisited) {
// 			res = append(res, root.Val)
// 			preVisited = root
// 			printstack(stack)
// 			root = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 		} else {
// 			if root.Right != nil {
// 				stack = append(stack, root.Right)

// 			}
// 			if root.Left != nil {
// 				stack = append(stack, root.Left)
// 			}
// 			preVisited = root
// 			root = stack[len(stack)-1]
// 			stack = stack[:len(stack)-1]
// 		}
// 	}
// 	return res
// }
