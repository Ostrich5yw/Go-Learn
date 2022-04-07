package leetcodeProject

import (
	"math"
	"strconv"
)

func FractionToDecimal(numerator int, denominator int) string {
	hashmap := make(map[int]int)
	var res = ""
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		res = "-"
	}
	numerator = int(math.Abs(float64(numerator)))
	denominator = int(math.Abs(float64(denominator)))
	res += strconv.Itoa(numerator / denominator)
	var decNum = numerator % denominator
	if decNum == 0 {
		return res
	}
	res += "."
	var mark = len(res)
	for decNum != 0 {
		v, ok := hashmap[decNum]
		if ok {
			res = res[:v] + "(" + res[v:] + ")"
			break
		} else {
			hashmap[decNum] = mark
			mark += 1
		}
		if decNum*10 < denominator {
			res = res + "0"
		} else {
			res = res + strconv.Itoa((decNum*10)/denominator)
		}
		decNum = decNum * 10 % denominator
	}
	return res
}
