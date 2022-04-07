package leetcodeProject

func findK2(root, p, q *TreeNode) *TreeNode {
	if root == nil { //叶节点返回nil
		return nil
	}
	/**
			7
		5		8
	  1	  6       9

	  找6,9最近父节点，会找到7直接返回，不用担心找到5或8
	*/
	if (p.Val <= root.Val && q.Val >= root.Val) || (p.Val >= root.Val && q.Val <= root.Val) {
		return root //如果p，q分别在自己的左右，则说明自己就是最近父节点，直接返回自己
	}
	left := findK2(root.Left, p, q) //分别查看左右树
	right := findK2(root.Right, p, q)
	if left != nil { //左右树返回不为空，说明最近父节点在左右子树上，直接返回
		return left
	} else if right != nil {
		return right
	}
	return nil //不在左右子树，说明这个分支不存在最近父节点，返回nil
}
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	return findK2(root, p, q)
}
