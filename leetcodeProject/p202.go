package leetcodeProject

func IsHappy(n int) bool {
	hashmap := make(map[int]bool)
	for n != 1 {
		temp := 0
		for n != 0 { //计算这个数各位平方和
			temp += (n % 10) * (n % 10)
			n = n / 10
		}
		_, ok := hashmap[temp]
		if ok { //如果有重复说明有无限循环
			return false
		}
		hashmap[temp] = true
		n = temp
	}
	return true
}
