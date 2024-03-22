/*package leetcode//回溯问题难点就在于去重，第二种方法暂时没有悟透，有空回来再悟
func permutation(s string) []string {//第一种方法，回溯，这里是找到全排列然后将不同的结果放入字符串切片中
    re:=[]string{}
    t:=""
    vis:=make([]bool,len(s))//判断是否访问过这个字母
    var dfs func(t string)
    dfs=func(t string){
        if len(t)==len(s){//因为是一个一个字母添加，说明此时到了叶子结点
            for _,v:=range re{//添加结果与之前结果是否重复
                if v==t{
                    return
                }
            }
            re=append(re,t)
            return
        }
        for i:=range s{//本质是循环套循环，在循环过程中没被访问则加入字符串，在回退过程中再撤销加入的操作
            if vis[i]==false{
                t+=string(s[i])
                vis[i]=true
                dfs(t)
                t=t[:len(t)-1]
                vis[i]=false
            }
        }
    }
    dfs(t)
    return re
}
func permutation(s string) []string {//第二种使用了hashmap来优化查询次数，hashmap同样会记录所有字母的出现次数，
    res := []string{}   //返回值列表
    Hmap := make(map[byte]int)  //aab在第一种需要循环三次，这里因为哈希表只用循环两次，节省了时间
    ele := ""   //待构造的字符串
    for i:=0; i<len(s); i++ {
        Hmap[s[i]]++    //计数
    }
    var dfs func(start int)
    dfs = func(start int) {
        if start == len(s) {
            res = append(res, ele)  //字符串构造完毕 添加进返回值列表
            return
        }
        for k, _ := range Hmap {//这里aab的哈希表循环两次，但是下次循环因为a的值在表中为2，所以还可以再用一次
            if Hmap[k] != 0 {   //次数不为0说明可用
                ele += string(k)
                Hmap[k]--
                dfs(start+1)    //从新的点继续构造字符串
                ele = ele[:len(ele)-1]  //回溯
                Hmap[k]++
            }
        }
    }
    dfs(0)
    return res
}
*/