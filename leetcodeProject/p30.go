package leetcodeProject

func FindSubstring(s string, words []string) []int {
	res := []int{}
	wMap := make(map[string]int)
	wLen, wsLen := len(words[0]), 0
	for _, w := range words {
		wMap[w] += 1
		wsLen += len(w)
	}
	point := 0
LOOP:
	for ; point < len(s)-wsLen+1; point += 1 {
		tempMap := make(map[string]int)
		for k, v := range wMap {
			tempMap[k] = v
		}
		for i := point; i < point+wsLen; i += wLen {
			if v, bol := tempMap[s[i:i+wLen]]; bol && v > 0 {
				tempMap[s[i:i+wLen]] -= 1
			} else {
				break
			}
		}
		for _, v := range tempMap {
			if v != 0 {
				point += 1
				goto LOOP
			}
		}
		res = append(res, point)
	}
	return res
}
