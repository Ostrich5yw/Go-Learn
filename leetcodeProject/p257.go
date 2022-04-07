package leetcodeProject

import "strconv"

func findWay(root *TreeNode) (res []string) {
	if root == nil {
		return []string{}
	}
	if root.Left == nil && root.Right == nil {		//找到子节点，返回他本身            son
		s := strconv.Itoa(root.Val)
		return []string{s}
	}
	s := strconv.Itoa(root.Val) + "->"
	s1 := findWay(root.Left)		//左节点的字符串数组				
	s2 := findWay(root.Right)		//右节点的字符串数组
	for i := 0; i < len(s1); i++ {	//分别为左右节点数组中的每个字符串加上当前结点			root->leftson    root->rightson 
		res = append(res, s+s1[i])													//leftson rightson分别为左右子树的包含所有路径的数组
	}
	for i := 0; i < len(s2); i++ {
		res = append(res, s+s2[i])
	}
	return res						//返回当前结点的字符串数组
}
func BinaryTreePaths(root *TreeNode) []string {
	return findWay(root)
}
