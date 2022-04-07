package leetcodeProject

func ConvertToTitle(columnNumber int) string {
	res := ""
	// temp := int(math.Pow(26, ？(digit)))
	for columnNumber != 0 {
		columnNumber = columnNumber - 1
		localNum := columnNumber%26 + 'A'
		columnNumber = columnNumber / 26
		res = res + string(rune(localNum))
	}
	return res
}
