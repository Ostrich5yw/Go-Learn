package leetcodeProject

func IsUgly(n int) bool {
	if n < 1 { //负数或0肯定不是丑数
		return false
	}

	for n%5 == 0 { //把因数5,3,2除完，如果为1说明只包含这三个因数，否则还有其他因数
		n /= 5
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}
