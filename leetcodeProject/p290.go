package leetcodeProject

import "strings"

/**
如果只有map[byte]string  则如下出错
	abba  cat cat cat cat
如果只有map[string]byte  则如下出错
	abba  dog cat cat rabbit
*/
func WordPattern(pattern string, s string) bool { //需要pattern和s互相对应，所以需要分别建立两个map进行记录
	pLen, words, pMap, sMap := len(pattern), strings.Split(s, " "), make(map[byte]string), make(map[string]byte)
	if pLen != len(words) {
		return false
	}
	for index, word := range words {
		if val, ok := sMap[word]; ok {
			if val != pattern[index] {
				return false
			}
		}
		if val, ok := pMap[pattern[index]]; ok {
			if val != word {
				return false
			}
			continue
		}
		pMap[pattern[index]] = word
		sMap[word] = pattern[index]
	}
	return true
}
