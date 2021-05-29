type MyStack struct {
	Val []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Val = append(this.Val, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	l := len(this.Val)
	var temp = this.Val[l-1]
	this.Val = this.Val[0 : l-1]
	return temp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Val[len(this.Val)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.Val) == 0
}
