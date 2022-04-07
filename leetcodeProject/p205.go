package leetcodeProject

func IsIsomorphic(s string, t string) bool {
	if s == "" && t == "" {
		return true
	}
	var len int = len(s)
	mark1 := make(map[uint8]uint8, 1000)
	mark2 := make(map[uint8]uint8, 1000)
	for i := 0; i < len; i++ {
		v1, ok1 := mark1[s[i]]
		v2, ok2 := mark2[t[i]]
		if (ok1 && v1 != t[i]) || (ok2 && v2 != s[i]) {
			return false
		}
		mark1[s[i]] = t[i]
		mark2[t[i]] = s[i]
	}

	return true
}
