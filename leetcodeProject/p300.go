package leetcodeProject

func LengthOfLIS(nums []int) int {
	nLen := len(nums)
	dp := make([]int, nLen)
	nMax := 1
	for i := 0; i < nLen; i++ { //状态转移方程 dp[i] = max(dp[j]) + 1      (dp[j]<dp[i]   0<j<i)
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		nMax = max(nMax, dp[i])
	}
	return nMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
