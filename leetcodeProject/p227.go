package leetcodeProject

import (
	d "DataStructure"
	"strconv"
)

// func main() {

// 	fmt.Println(calculate("2*3-4"))
// }

func findNum(nowPos int, s string) (int, int) { //寻找数字："1+ 211  /     332"	发现数字或者乘除符号后，需要将整个数字字符串搜索完毕，并转为int
	start := -1
	for ; nowPos < len(s); nowPos++ {
		if start == -1 && s[nowPos] >= '0' && s[nowPos] <= '9' { // start记录数字起始位置(因为可能有空格)
			start = nowPos
		}
		if start != -1 && (s[nowPos] < '0' || s[nowPos] > '9') { // 已经记录了起始位置，并且当前又一次发现非数字符号，停止遍历
			break
		}
	}
	nowNum, _ := strconv.Atoi(s[start:nowPos]) //将string转为int
	return nowNum, nowPos - 1                  //遍历时，比如"234/1" nowPos会指向/，而下面循环中又会执行i++，导致switch跳过'/'号，所以需要nowPos-1
}

/**
"-5 + 11 * 20 / 1"
——> opt变为-1
——> 5 * opt后，将-5压栈
——> opt变为1
——> 11 * opt后，将11压栈
——> 遇到*，弹出11，找到20，11*20后，将220压栈
——> 遇到/，弹出220，找到1,220/1后，将220压栈
——> 结束，遍历栈，结果为[-5,220] 相加得215
*/
func Calculate2(s string) int {
	res := 0
	stack := d.NewStack()
	nowOpt := 1
	preNum, nowNum := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			nowOpt = 1
			continue
		case '-':
			nowOpt = -1
			continue
		case '*':
			preNum = (*stack.Pop()).(int) //遇到乘除号，弹出上一个数字，即将乘除号前后的两数字做乘除，并将结果再次压栈
			nowNum, i = findNum(i, s)
			stack.Push(preNum * nowNum)
			continue
		case '/':
			preNum = (*stack.Pop()).(int)
			nowNum, i = findNum(i, s)
			stack.Push(preNum / nowNum)
			continue
		default:
			nowNum, i = findNum(i, s)
			stack.Push(nowOpt * nowNum)
			continue
		}
	}
	for !stack.Empty() {
		res += (*stack.Pop()).(int)
	}
	return res
}
