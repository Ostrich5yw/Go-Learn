package leetcodeProject

import (
	"sync"
	"time"
)

func findK(root *TreeNode, k int, m *int, wait *sync.WaitGroup) { //中序遍历
	if root == nil {
		return
	}
	findK(root.Left, k, m, wait)
	*m = *m + 1  //m记录整体的序数
	if *m == k { //当m==k，停止这个goroutine，m中存入相应的Val值
		*m = root.Val
		wait.Done()                 //wait.Done()减一后，主函数向下执行到return需要一段时间
		time.Sleep(3 * time.Second) //所以这里停止3s，给主函数向下执行的时间
	}
	findK(root.Right, k, m, wait)

}
func KthSmallest(root *TreeNode, k int) int {
	var m int = 0
	wait := sync.WaitGroup{}
	wait.Add(1)
	go findK(root, k, &m, &wait)
	wait.Wait() //找到m后，主程序继续运行返回m
	return m
}

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// func main() {
// 	t1 := &TreeNode{1, nil, nil}
// 	t222 := &TreeNode{6, nil, nil}
// 	t22 := &TreeNode{5, nil, t222}
// 	t2 := &TreeNode{4, nil, t22}
// 	t3 := &TreeNode{2, t1, nil}
// 	t4 := &TreeNode{3, t3, t2}
// 	fmt.Println(kthSmallest(t4, 3))
// }
