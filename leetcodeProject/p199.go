package leetcodeProject

import (
	d "DataStructure"
	"container/list"
)

// func main() {
// 	root := d.NewTreeNode(0)
// 	root.Left = d.NewTreeNode(1)
// 	root.Right = d.NewTreeNode(2)
// 	fmt.Println(rightSideView(root))
// }
func RightSideView(root *d.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int = make([]int, 100)
	var num int = 0
	var arraylist = list.New()
	arraylist.PushBack(root)
	temp := arraylist.Front() //temp负责遍历
	last := arraylist.Back()  //last负责指向每层的末尾元素
	for temp != last.Next() {
		if temp.Value.(*d.TreeNode).Left != nil { //分别将左右子节点加入数组
			arraylist.PushBack(temp.Value.(*d.TreeNode).Left)
		}
		if temp.Value.(*d.TreeNode).Right != nil {
			arraylist.PushBack(temp.Value.(*d.TreeNode).Right)
		}
		if temp == last { //temp == last 说明该层遍历完成
			last = arraylist.Back() //last重新指向下一层末尾元素
			res[num] = temp.Value.(*d.TreeNode).Val
			num++
		}
		temp = temp.Next() //temp继续向后遍历
	}
	return res
}
