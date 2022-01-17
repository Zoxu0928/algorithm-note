package ac_006

type ListNode struct {
	Val  int
	Next *ListNode
}

//两次遍历解法
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	//数组元素前后交换位置
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
