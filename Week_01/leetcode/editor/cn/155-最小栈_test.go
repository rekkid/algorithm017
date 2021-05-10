//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
//
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
//
//
//
//
// 示例:
//
// 输入：
//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]
//
//输出：
//[null,null,null,null,-3,null,0,-2]
//
//解释：
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> 返回 -3.
//minStack.pop();
//minStack.top();      --> 返回 0.
//minStack.getMin();   --> 返回 -2.
//
//
//
//
// 提示：
//
//
// pop、top 和 getMin 操作总是在 非空栈 上调用。
//
// Related Topics 栈 设计
// 👍 895 👎 0

package main

import (
	"math"
	"testing"
)

func TestMinStack(t *testing.T) {
	obj := Constructor()
	obj.Push(-2)
}

//leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	normal []int
	min    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.normal = append(this.normal, val)
	// !!这里容易错,开始时min长度为0
	if len(this.min) == 0 || val < this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[len(this.min)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.normal) != 0 {
		this.normal = this.normal[:len(this.normal)-1]
		this.min = this.min[:len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.normal) == 0 {
		return math.MinInt64
	} else {
		return this.normal[len(this.normal)-1]
	}
}

func (this *MinStack) GetMin() int {
	if len(this.min) != 0 {
		return this.min[len(this.min)-1]
	} else {
		return math.MinInt64
	}

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
