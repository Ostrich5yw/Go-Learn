package leetcodeProject

import "strconv"

func isLegal(a, b, c int) bool {
	return c-a == b
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func IsAdditiveNumber(num string) bool {
	numA, numB, numC, lenA, lenB, lenC := 0, 0, 0, 0, 0, 0
	var err error
	/**
	本题最关键是找到前两个数，用num[0:i]和num[i:j+1]分别表示前两个数
	lenA, lenB记录这两个数长度，则numC的长度lenC必然等于max(lenA,lenB)或者max(lenA,lenB)+1
	*/
	for i := 1; i <= len(num); i++ {
		if numA, err = strconv.Atoi(num[0:i]); err != nil || (i != 1 && num[0] == '0') {	//求numA
			continue
		}
		tt := numA
		for j := i + 1; j <= len(num); j++ {
			if numB, err = strconv.Atoi(num[i:j]); err != nil || (j != i+1 && num[i] == '0') {//求numB
				continue
			}
			numA, lenA, lenB = tt, i-0, j-i
			lenC = max1(lenA, lenB)
			start, end := j, j+lenC		//记录numA，numB的长度，同时numC只可能在num[j:j+max(lenA,lenB)] 与 num[j:j+max(lenA,lenB)+1] 之中
			for end <= len(num) {	
				//**********************************************************************************************************
				if numC, err = strconv.Atoi(num[start:end]); err != nil || (end != start+1 && num[start] == '0') {	//如果num[start:end]大于MAXINT，直接break
					break
				}
				lenC = end - start
				if !isLegal(numA, numB, numC) {	//num[start:end]!=numA+numB  则继续判断num[start:end+1]
					if end+1 > len(num) {
						break
					}
					if numC, err = strconv.Atoi(num[start : end+1]); err != nil || (end+1 != start+1 && num[start] == '0') {
						break
					}
					if !isLegal(numA, numB, numC) {
						break
					}
					lenC = lenC + 1
					end = end + 1		//num[start:end+1]==numA+numB   则更新lenC和end的大小
				}
				//**********************************************************************************************************
				if end == len(num) {		//end等于len(num)且num[start:end]==numA+numB，说明该string满足条件
					return true
				}
				numA, numB, lenA, lenB = numB, numC, lenB, lenC		//numA,numB整体向后，lenC，start，end重新记录新的numC可能的取值范围
				lenC = max1(lenA, lenB)
				start, end = end, end+lenC
			}
		}
	}
	return false
}
