//package leetcode

/*type MinStack struct { //用两个栈，第一个栈正常栈，第二个栈记录第一个栈里的元素在作为栈顶时的最小值
	stack1 []int //所以最小值就是当前第二个栈的栈顶元素
	stack2 []int //记录对应元素在作为栈顶时的最小值
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) { //入栈，栈空则说明两个栈都为空，所以两个都入相同元素，不空栈一直接入，第二个栈入当前最小值和要入栈的元素两个之间的最小值
	if len(this.stack1) == 0 {
		this.stack1 = append(this.stack1, x)
		this.stack2 = append(this.stack2, x)
		return
	}
	if x < this.stack2[len(this.stack2)-1] {
		this.stack1 = append(this.stack1, x)
		this.stack2 = append(this.stack2, x)
	} else {
		this.stack1 = append(this.stack1, x)
		this.stack2 = append(this.stack2, this.stack2[len(this.stack2)-1])
	}
	return
}

func (this *MinStack) Pop() { //两个栈直接都减少一个
	if len(this.stack1) == 0 {
		return
	} else {
		this.stack1 = this.stack1[:len(this.stack1)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack1[len(this.stack1)-1]
}
func (this *MinStack) Min() int {
	return this.stack2[len(this.stack2)-1]
}*/

/*type MinStack struct {//第二种方法，差分法，每次记录入栈值与最小值的差，公式为data=x-min，data写入栈是差值，x是传入的值，min是记录的最小值，这个公式要注意处理第一个元素，因为第一个元素前没有min
    stack []int
    min int
}

func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    if len(this.stack)==0{//第一次传入，特殊处理传入0，认为min和data相等,这样处理可以让所有元素都符合data=x-min这个公式
        this.stack=append(this.stack,0)
        this.min=x
    }else{
        t:=x-this.min//其他情况，传入差值，如果差小于0说明最小值应发生改变
        this.stack=append(this.stack,t)
        if t<0{
            this.min=x
        }
    }
}


func (this *MinStack) Pop()  {//这里没有处理第一个元素是因为无论怎样都会删除，并且删除后即使min改变了，由于此时len为0，下次入栈会自己修改min
    if this.stack[len(this.stack)-1]>=0{//大于等于0说明入栈时的x大于等于当前最小值min，所以此时直接出就行
        this.stack=this.stack[:len(this.stack)-1]
    }else{
        this.min=this.min-this.stack[len(this.stack)-1]//小于0，说明出栈后最小值要发生改变，就需要找到更早的最小值，根据公式data=x-min，在data小于0时变为data=newmin-oldmin，所以oldmin=newmin-data
        this.stack=this.stack[:len(this.stack)-1]//newmin是当前最小值，oldmin是要算的那个最小值
    }
}


func (this *MinStack) Top() int {
    //data=x-min
    if this.stack[len(this.stack)-1]>0{
        return this.min+this.stack[len(this.stack)-1]
    }else{
        return this.min
    }
}


func (this *MinStack) Min() int {
    return this.min
}*/
