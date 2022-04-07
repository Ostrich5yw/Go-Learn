package leetcodeProject

func dfs2(rec [][]int, lenRec []int, flag []int, res *[]int, nowNum int) bool { //思路与p207一致，加入15,16行记录遍历顺序
	if flag[nowNum] == -1 {
		return false
	} else if flag[nowNum] == 1 {
		return true
	}
	flag[nowNum] = -1
	for i := 0; i < lenRec[nowNum]; i++ {
		if !dfs2(rec, lenRec, flag, res, rec[nowNum][i]) {
			return false
		}
	}
	t1 := *res                //t1为地址res中的值	//如果dfs2()中的参数是res []int，则append之后返回的新数组与参数地址不同，则无法通过将res放在参数中来修改它
	*res = append(t1, nowNum) //所以这里存放的是res *[]int，将res修改完成后，把地址res下存放的值，替换为append返回值
	flag[nowNum] = 1
	return true
}
func FindOrder(numCourses int, prerequisites [][]int) []int {
	var rec [][]int
	for i := 0; i < 10; i++ {
		t := make([]int, numCourses)
		rec = append(rec, t)
	}
	lenRec := make([]int, numCourses)
	flag := make([]int, numCourses)
	res := &[]int{}
	for i := 0; i < len(prerequisites); i++ {
		temp := prerequisites[i][0]
		rec[temp][lenRec[temp]] = prerequisites[i][1]
		lenRec[temp]++
	}
	for i := 0; i < numCourses; i++ {
		if flag[i] == 0 {
			if !dfs2(rec, lenRec, flag, res, i) {
				return []int{} //如果有环，返回空数组
			}
		}
	}
	return *res //返回地址res下存放的值
}
