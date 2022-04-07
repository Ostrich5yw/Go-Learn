package leetcodeProject

func CountPrimes(n int) int {
	var num int = 0
	mark := make([]bool, n)
	for i := 2; i < n; i++ {
		if mark[i] { //不是质数直接跳过
			continue
		}
		num++                      //是质数则计数器加1，并将他所有倍数标记为素数
		for j := 2; j*i < n; j++ { //如果2是质数，那么2*1 2*2 2*3均不是质数
			mark[j*i] = true
		}
	}
	return num
}
