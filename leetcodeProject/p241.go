package leetcodeProject

import "strconv"

// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func DiffWaysToCompute(expression string) []int {
	if isDigit(expression) {
		tmp, _ := strconv.Atoi(expression)
		return []int{tmp}
	}
	var res []int
	for index, c := range expression {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" { //每个递归，都会遍历所有符号位，对符号位两边进行相加并将结果加入res数组
			left := DiffWaysToCompute(expression[:index])
			right := DiffWaysToCompute(expression[index+1:])
			for _, leftNum := range left {			//左侧共有len(left)种结果，右侧共有len(right)种结果，将其组合得到所有组合结果
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}
	return res
}
