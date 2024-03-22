/*package leetcode  单调队列法

type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor() MaxQueue {
	return MaxQueue{}
}

func (this *MaxQueue) Max_value() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.max) != 0 && this.max[len(this.max)-1] < value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	if this.queue[0] == this.max[0] {
		this.max = this.max[1:]
	}
	t := this.queue[0]
	this.queue = this.queue[1:]
	return t
}*/
