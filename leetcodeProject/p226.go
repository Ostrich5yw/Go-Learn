package leetcodeProject

func recurInvert(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {	//对于空节点或者叶节点，直接返回
		return root
	}
	root.Left = recurInvert(root.Left)		//翻转左子树
	root.Right = recurInvert(root.Right)	//翻转右子树
	temp := root.Left			//翻转自己（即对调自己的左右子树）
	root.Left = root.Right
	root.Right = temp
	return root
}
func InvertTree(root *TreeNode) *TreeNode {
	return recurInvert(root)
}
