package ac_024

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针法
func reverseList(head *ListNode) *ListNode {
	//第一个指针
	var prev *ListNode
	//第二个指针
	curr := head
	for curr != nil {
		next := curr.Next
		//当前节点指向前一个节点
		curr.Next = prev
		//移动第一个指针
		prev = curr
		//移动第二个指针
		curr = next
	}
	return prev
}
