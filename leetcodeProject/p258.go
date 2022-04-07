package leetcodeProject

import "strconv"

func AddDigits(num int) int {
	for {
		snum := strconv.Itoa(num)
		num = 0
		for _, c := range snum {
			num += (int)(c - 48)
		}
		if num < 10 {
			return num
		}
	}
}
