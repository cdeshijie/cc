/*package leetcode

func validateStackSequences(pushed []int, popped []int) bool {//栈顶和出栈元素作比较，相同则出，不同就试着入栈，入不了栈说明早就入栈且不是栈顶，此时报错
    stack:=make([]int,len(pushed))
    pushedpointer:=0
    find:=func(x int)int{//寻找出栈元素在入栈元素的位置
        for i:=0;i<len(pushed);i++{
            if pushed[i]==x{
                return i
            }
        }
        return -1
    }
    for i:=0;i<len(popped);i++{//遍历出栈元素
        if pushedpointer<=find(popped[i]){//如果出栈元素位置大于已经入栈的元素位置，则继续入栈
            for pushedpointer<=find(popped[i]){
                stack=append(stack,pushed[pushedpointer])
                pushedpointer++
            }
        }
        if popped[i]==stack[len(stack)-1]{//入栈后这里一定会相等，如果上面没有入栈，说明将要出栈的元素早就在栈内，且不在栈顶，此时报错
            stack=stack[:len(stack)-1]
        }else{
            return false
        }
    }
    return true
}
func validateStackSequences2(pushed []int, popped []int) bool {//每次入栈都与出栈元素作比较，如果相等则出栈，且继续比较，直到栈为空，或者不相等，继续入栈
    stack:=[]int{}
    poppedpointer:=0
    for i:=0;i<len(pushed);i++{
        stack=append(stack,pushed[i])
        for len(stack)!=0&&stack[len(stack)-1]==popped[poppedpointer]{
            stack=stack[:len(stack)-1]
            poppedpointer++
        }
    }
    if len(stack)==0{//所有元素都入栈后，如果都出栈了则长度为0
        return true
    }else{
        return false
    }
}
*/