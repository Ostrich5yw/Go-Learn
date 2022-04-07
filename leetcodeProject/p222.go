package leetcodeProject

import (
	d "DataStructure"
)

func getDepth(root *d.TreeNode) int { //计算当前结点的最大深度。由于是完全二叉树，所以最大深度等于最大左子树深度
	res := 0
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

func CountNodes(root *d.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := getDepth(root.Left) //得到左右子树深度
	right := getDepth(root.Right)
	/**
		左右子树深度一致，说明以root=1为根的左子树一定是完全二叉树，右子树不一定。所以总结点数=完全二叉左子树 + 根节点 + （右子树）
			1          1
		  2	  3      2   3
	     2 2 3      2 2 3 3
		左右子树深度不一致，说明以root=1为根的右子树一定是完全二叉树，左子树一定不是。所以总结点数=完全二叉右子树 + 根节点 + （左子树）
	   	   1
		2	   3
	  2   2  3   3
	 2 2
	*/
	if left == right { //如果深度一致，向右子树继续遍历，同时加上完全二叉的左子树与根节点
		return CountNodes(root.Right) + (1<<left - 1) + 1
	} else { //如果深度不一致，向左子树继续遍历，同时加上完全二叉的右子树与根节点
		return CountNodes(root.Left) + (1<<right - 1) + 1
	}
}
