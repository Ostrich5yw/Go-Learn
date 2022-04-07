package leetcodeProject

import (
	"strings"
)

func FindRepeatedDnaSequences(s string) []string {
	var args string = ""
	selfmap := make(map[string]bool, 100)
	res := make([]string, 100)
	num := 0
	for i := 0; i < len(s)-10; i++ {
		args = s[i : i+10]
		temps := s[i+1:]
		if strings.Contains(temps, args) && !selfmap[args] {
			selfmap[args] = false
		}
	}
	for k := range selfmap {
		res[num] = k
		num++
	}
	return res[0:num]
}
