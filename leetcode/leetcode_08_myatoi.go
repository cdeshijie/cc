/*package leetcode

func myAtoi(s string) int {
    if len(s)==0{
        return 0
    }
    i:=0
    for i<len(s)&&s[i]==' '{
        i++
    }
    if i==len(s){
        return 0
    }
    flag:=1
    if s[i]=='-'{
        flag=-1
        i++
    }else if s[i]=='+'{
        i++
    }
    result:=0
    for i<len(s)&&s[i]>='0'&&s[i]<='9'{
        result=result*10+int(s[i]-'0')
        if result>math.MaxInt32{
            if flag>0{
                return math.MaxInt32
            }else{
                return flag*(math.MaxInt32+1)
            }
        }
        i++
    }
    return flag*result
}*/