package leetcodeProject

func LongestValidParentheses(s string) int {
	sLen, res := len(s), 0
	leftNum, rightNum := 0, 0
	for i := 0; i < sLen; i++ {
		if s[i] == '(' {
			leftNum++
		}
		if s[i] == ')' {
			rightNum++
		}
		if leftNum == rightNum {
			res = max(res, leftNum+rightNum)
		} else if leftNum < rightNum {
			leftNum, rightNum = 0, 0
		}
	}
	leftNum, rightNum = 0, 0
	for i := sLen - 1; i >= 0; i-- {
		if s[i] == '(' {
			leftNum++
		}
		if s[i] == ')' {
			rightNum++
		}
		if leftNum == rightNum {
			res = max(res, leftNum+rightNum)
		} else if leftNum > rightNum {
			leftNum, rightNum = 0, 0
		}
	}
	return res
}
