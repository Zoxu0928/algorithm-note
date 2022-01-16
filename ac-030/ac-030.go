package ac_030

import "math"

type MinStack struct {
	stack, minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	//数据栈正常存入
	this.stack = append(this.stack, x)
	//获得最小栈当前存储的最小值
	top := this.minStack[len(this.minStack)-1]
	//比较一下当前最小值和要插入的值
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	stackLen := len(this.stack) - 1
	this.stack = this.stack[:stackLen]
	minStackLen := len(this.minStack) - 1
	this.minStack = this.minStack[:minStackLen]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
