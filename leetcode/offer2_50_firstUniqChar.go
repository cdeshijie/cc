/*package leetcode

func firstUniqChar(s string) byte { //用map，map遍历是无序的
	if len(s) == 0 {
		return ' '
	}
	m := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		_, ok := m[s[i]]
		if ok == true {
			m[s[i]]++
		} else {
			m[s[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if v == 1 && ok {
			return s[i]
		}
	}
	return ' '
}*/
