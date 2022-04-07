package leetcodeProject

func dfs(flag []int, que [][]int, lenque []int, num int) bool {
	if flag[num] == -1 { //如果为-1，说明成环
		return false
	} else if flag[num] == 1 { //如果为1，说明该节点是安全节点
		return true
	}
	flag[num] = -1                     //标记为-1，表示待定，在该次深度遍历中已经出现过一次，如果再次出现这说明成环
	for i := 0; i < lenque[num]; i++ { //循环遍历指向num节点的所有节点，如果都是安全的，则num是安全的（如下将其置为1）
		if !dfs(flag, que, lenque, que[num][i]) {
			return false
		}
	}
	flag[num] = 1
	return true
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
	var que [][]int
	for x := 0; x < numCourses; x++ {
		ar := make([]int, numCourses)
		que = append(que, ar)
	}
	flag := make([]int, numCourses)   //flag 标记该节点是否为安全节点，1——安全，-1——不安全(待定)，0——还未检测
	lenque := make([]int, numCourses) //记录每个que长度
	for i := 0; i < len(prerequisites); i++ {
		temp := prerequisites[i][1] //que[i][] 记录指向i的所有节点     lenque 记录每一个que[i][]的长度
		que[temp][lenque[temp]] = prerequisites[i][0]
		lenque[temp]++

		// temp := prerequisites[i][0] //结果一样的
		// que[temp][lenque[temp]] = prerequisites[i][1]
		// lenque[temp]++
	}
	// fmt.Print(lenque)
	for i := 0; i < numCourses; i++ { //保障所有节点都被遍历过，非-1即1，保障没有0
		if !dfs(flag, que, lenque, i) {
			return false
		}
	}
	return true
}
