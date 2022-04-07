package leetcodeProject

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	smap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if v, ok := smap[s[i]]; ok {
			smap[s[i]] = v + 1			//s中出现给该字符加1
		} else {
			smap[s[i]] = 1
		}
		if v, ok := smap[t[i]]; ok {
			smap[t[i]] = v - 1			//t中出现给该字符减1
		} else {
			smap[t[i]] = -1
		}
	}
	for _, v := range smap {
		if v != 0 {
			return false
		}
	}
	return true
}
