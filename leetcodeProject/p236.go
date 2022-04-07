package leetcodeProject

import (
	"sync"
	"time"
)

func findK1(res, root, p, q *TreeNode, wait *sync.WaitGroup) int {
	var t1, t2 int
	if root == nil { //搜索到叶节点，返回0
		return 0
	}
	if root == p || root == q { //自己就是搜索的节点，则还需要判断子节点中还有无另一个搜索节点，如果有，直接程序终止，返回自己，否则返回1，表示这个分支有一个搜索节点
		if findK1(res, root.Left, p, q, wait) == 1 || findK1(res, root.Right, p, q, wait) == 1 {
			*res = *root
			wait.Done()
			time.Sleep(1 * time.Second)
		}
		return 1
	}
	t1 = findK1(res, root.Left, p, q, wait)
	t2 = findK1(res, root.Right, p, q, wait)
	if t1 == 1 && t2 == 1 { //左右子树都有搜索节点的第一个节点，证明自己就是他们的最近父节点（直接wait.Done()返回自己）
		*res = *root
		wait.Done()
		time.Sleep(1 * time.Second)
	} else if t1 == 1 || t2 == 1 { //左右子树一边有搜索节点，返回1，证明该分支有一个搜索节点
		return 1
	}
	return 0 //返回0表示这个分支无搜索节点
}
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res TreeNode
	wait := sync.WaitGroup{}
	wait.Add(1)
	go findK1(&res, root, p, q, &wait)
	wait.Wait()
	return &res
}
