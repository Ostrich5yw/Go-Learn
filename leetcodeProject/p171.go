package leetcodeProject

import (
	"math"
)

func TitleToNumber(columnTitle string) int {
	res := 0
	j := len(columnTitle) - 1
	for i := 0; i < len(columnTitle); i++ {
		res += int(math.Pow(26, float64(j))) * (int)(columnTitle[i]-64)
		j--
	}
	return res
}
