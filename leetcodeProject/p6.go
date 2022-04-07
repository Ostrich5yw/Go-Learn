package leetcodeProject

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	res := make([]string, numRows)
	point, mark := 0, 0 //point表示字符存入第几个数组，mark表示目前是向下(0, Z的直边部分)还是向上(1, Z的斜边部分)
	for i := 0; i < len(s); i++ {
		if point <= 0 && mark == 1 {
			point = 0
			mark = 0
			i -= 1
			continue //此处要continue，因为numRows==2时，可能会出现point-=2后，point已经指向0的情况，所以需要直接再转化为mark=0
		}
		if point >= numRows && mark == 0 {
			point -= 2
			mark = 1
			i -= 1
			continue
		}
		if mark == 0 {
			res[point] += string(s[i])
			point += 1
		} else {
			res[point] += string(s[i])
			point -= 1
		}
	}
	tt := ""
	for _, t := range res {
		tt += t
	}
	return tt
}
