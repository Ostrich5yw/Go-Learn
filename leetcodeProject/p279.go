package leetcodeProject

func NumSquares(n int) int {		//动态规划，对于一个整数i,最少平方和个数为：1 + min(nums[i-j*j])  1<j*j<i,即在上一个数的基础上，再加j*j得到
	nums := make([]int, n+1)		//nums的序数是目标和，nums的值是所需的最少平方和个数
	for i := 1; i <= n; i++ {
		minNum := 1000000
		for j := 1; j*j <= i; j++ {
			minNum = min(minNum, nums[i-j*j])
		}
		nums[i] = minNum + 1
	}
	return nums[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
