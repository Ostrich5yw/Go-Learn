package leetcodeProject

func recursion(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = recursion(root.Left, res)
	res = recursion(root.Right, res)
	return res
}
func PreorderTraversal(root *TreeNode) []int {
	var res []int
	return recursion(root, res)
}
