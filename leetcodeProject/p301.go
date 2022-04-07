package leetcodeProject

func isValid(str string) bool { //判断符号是否符合规则
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

func RemoveInvalidParentheses(s string) (ans []string) {
	curSet := map[string]struct{}{s: {}} //将初始字符串存入set
	for {
		for str := range curSet {
			if isValid(str) {
				ans = append(ans, str)
			}
		}
		//当curSet中第一次出现符合条件的字符串，立刻返回，ans即是删除最少括号得到的结果
		if len(ans) > 0 {
			return
		}
		nextSet := map[string]struct{}{} //做下一轮迭代
		for str := range curSet {        //对curSet中所有的结果进行下一轮迭代
			for i, ch := range str { //按位删除符号，并将结果串存入下一轮结果集中
				//针对())))))这种情况，其实从i=1到i=6得到的结果是一致的，所以跳过
				if i > 0 && byte(ch) == str[i-1] {
					continue
				}
				if ch == '(' || ch == ')' { //如果是其他字符不管，比如(a)
					nextSet[str[:i]+str[i+1:]] = struct{}{}
				}
			}
		}
		curSet = nextSet
	}
}
