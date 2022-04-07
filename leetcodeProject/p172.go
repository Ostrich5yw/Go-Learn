package leetcodeProject

func TrailingZeroes(n int) int {
	var num int = 0
	n = int(n/5) * 5
	for n != 0 {
		temp := n
		for temp%5 == 0 {
			temp = temp / 5
			num++
		}
		n = n - 5
	}
	return num
}
