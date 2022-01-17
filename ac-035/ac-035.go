package ac_035

type Node struct {
	Val int

	Next *Node

	Random *Node
}

//解法一：两次循环。第一次循环，复制节点，用map存储映射关系，第二次循环，将复制出来的新节点连接上
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node)

	//第一遍循环 用map记录新节点的映射关系
	for curr := head; curr != nil; {
		m[curr] = &Node{
			Val: curr.Val,
		}
		curr = curr.Next
	}

	//第二遍循环 从map中获取复制的节点 并添加映射关系
	for curr := head; curr != nil; {
		m[curr].Next = m[curr.Next]
		m[curr].Random = m[curr.Random]
		curr = curr.Next
	}
	return m[head]
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 1. 新生成copy节点放到原节点的后面，例如1->2->3这样的链表就变成了这样1->1'->2->2'->3->3'
	// 此时拷贝节点的random指针都指向nil
	var copy *Node = nil
	for cur := head; cur != nil; cur = cur.Next.Next { // 因为是加上了copy节点，所以每次要移动2步
		copy = &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
	}

	// 2. 将copy节点的random安排上，思路：
	// cur.Next指向的替身替身1，cur.Random指向的原节点的random，那么cur.Random.Next指向的是谁？答案是cur.Random的替身2
	// 那么将 替身1的random 指向 替身2，就完成了copy节点的random指向，非常巧妙
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil { // random指针可能为空
			cur.Next.Random = cur.Random.Next // cur.Next是替身1，cur.Random.Next是替身2
		}
	}

	// 3. 拆分原节点 和 copy节点
	var temp *Node = nil
	newHead := head.Next
	for cur := head; cur != nil && cur.Next != nil; {
		temp = cur.Next      // copy节点
		cur.Next = temp.Next // 原链表
		cur = temp
	}
	return newHead
}
