/*package leetcode

type stack struct {
	top   int
	value []int
}

type CQueue struct {
	stack1 stack
	stack2 stack
}

func Constructor() CQueue {
	c := CQueue{
		stack1: stack{
			top:   -1,
			value: []int{},
		},
		stack2: stack{
			top:   -1,
			value: []int{},
		},
	}
	return c
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.top++
	this.stack1.value = append(this.stack1.value, value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.top != -1 {
		t := this.stack2.value[this.stack2.top]
		this.stack2.value = this.stack2.value[:this.stack2.top]
		this.stack2.top--
		return t
	}
	if this.stack1.top == -1 && this.stack2.top == -1 {
		return -1
	}
	for this.stack1.top != -1 {
		this.stack2.top++
		this.stack2.value = append(this.stack2.value, this.stack1.value[this.stack1.top])
		this.stack1.top--
	}
	this.stack1.value = []int{}
	t := this.stack2.value[this.stack2.top]
	this.stack2.value = this.stack2.value[:this.stack2.top]
	this.stack2.top--
	return t
}*/
