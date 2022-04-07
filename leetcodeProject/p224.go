package leetcodeProject

import (
	d "DataStructure"
	"strconv"
)

// func main() {
// 	fmt.Println(calculate("2+111"))
// }

/**
符号之所以要入栈，是因为 1 - (2+3) - 5的这种情况，第一个‘-’只作用于(2+3) 而不作用于5,所以遇到')'就将'-'出栈
1 - (2 + (3 - 4) - 1)
遇到第一个(括号 - 入栈  作用范围(2 + (3 - 4) - 1)
遇到第二个(括号 + 入栈  作用范围(3 - 4)
遇到第一个)括号 + 出栈
遇到第二个)括号 - 出栈
*/
func Calculate(s string) int {
	res := 0
	stack := d.NewStack()
	stack.Push(1)
	opt := 1
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			opt = 1 * (*stack.Top()).(int)
			i++
		case '-':
			opt = -1 * (*stack.Top()).(int)
			i++
		case '(':
			stack.Push(opt)
			i++
		case ')':
			stack.Pop()
			i++
		default:
			start := i
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
			}
			temp, _ := strconv.Atoi(s[start:i])
			res += opt * temp
		}
	}
	return res
}
