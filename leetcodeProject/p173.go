package leetcodeProject

import (
	d "DataStructure"
	"math"
)

type BSTIterator struct {
	Nodes   [100000]*d.TreeNode
	nowNode int
}

func midTraverse(root *d.TreeNode, Nodes []*d.TreeNode, num int) int {
	// fmt.Println(root.Val)
	if root == nil {
		return num
	}
	if root.Left != nil {
		Nodes[num] = root.Left
		num = midTraverse(root.Left, Nodes, num)
	}
	Nodes[num] = root
	num += 1
	if root.Right != nil {
		Nodes[num] = root.Right
		num = midTraverse(root.Right, Nodes, num)
	}
	return num
}

func Constructor(root *d.TreeNode) BSTIterator {
	bs := BSTIterator{}

	midTraverse(root, bs.Nodes[:], 1)
	for root.Left != nil {
		root = root.Left
	}
	root.Left = &d.TreeNode{Val: math.MinInt64, Left: nil, Right: nil}
	bs.Nodes[0] = root.Left
	bs.nowNode = 0
	return bs
}

func (p *BSTIterator) Next() int {
	p.nowNode = p.nowNode + 1
	return p.Nodes[p.nowNode].Val
}

func (p *BSTIterator) HasNext() bool {
	// fmt.Println(this.Nodes[this.nowNode])
	if p.Nodes[p.nowNode+1] != nil {
		return true
	} else {
		return false
	}
}
