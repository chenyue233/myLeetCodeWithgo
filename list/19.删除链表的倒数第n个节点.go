package list

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := new(ListNode)
	dummyNode.Next = head
	fast, slow := dummyNode, dummyNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}

// @lc code=end
