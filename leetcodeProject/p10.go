package leetcodeProject

/**
dp[i][j] = dp[i - 1][j - 1], if p[j - 1] != '*' && (s[i - 1] == p[j - 1] || p[j - 1] == '.');
dp[i][j] = dp[i][j - 2], if p[j - 1] == '*' and the pattern repeats for 0 time;
dp[i][j] = dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'), if p[j - 1] == '*' and the pattern repeats for at least 1 time.

*/
func IsMatch(s string, p string) bool {
	sLen, pLen := len(s), len(p)
	var dp [][]bool
	for x := 0; x <= sLen; x++ {		//建立一个存储状态的数组
		ar := make([]bool, pLen+1)		//其中dp[i][j]=true表示s[0, i)与p[0, j)模式匹配
		dp = append(dp, ar)
	}
	dp[0][0] = true //边界条件, 如果从0开始则边界为01，因此我们将结果存放在1——(Len+1)，所以要注意dp比s和p多一位
	for i := 0; i <= sLen; i++ {
		for j := 1; j <= pLen; j++ {
			if p[j-1] == '*' {		//dp[i][j-2]表示'*'代表的次数为0，后面表示'*'代表的次数为1
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				dp[i][j] = i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') && dp[i-1][j-1]
			}
		}
	}
	return dp[sLen][pLen]
}
