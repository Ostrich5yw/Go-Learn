package leetcodeProject

import (
	"strconv"
	"strings"
)

type Codec struct {
	str string
}

func Constructor7() Codec {
	return Codec{str: ""}
}

/**
采用DFS深度优先遍历，实现树的序列化，为空则置为"n"
*/
func rSerial(root *TreeNode) string {
	if root == nil {
		return "n"
	}
	left := "," + rSerial(root.Left)
	right := "," + rSerial(root.Right)
	sel := strconv.Itoa(root.Val)
	return sel + left + right
}

// Serializes a tree to a single string.
func (cc *Codec) Serialize(root *TreeNode) string {
	return rSerial(root)
}

/**
通过字符串恢复树，同样采用DFS深度优先遍历，实现树的反序列化
可以采用两种方式，一种是通过返回值给root.Right root.Left赋值，一种是通过指针给root.Right root.Left赋值
*/
func rDeserial(datas []string, index int) (int, *TreeNode) { //index记录当前遍历到的datas的位置
	var root *TreeNode
	word := datas[index]
	if word != "n" {
		Iword, _ := strconv.Atoi(word)
		root = &TreeNode{Val: Iword, Left: nil, Right: nil} //给当前结点赋值
		index, root.Left = rDeserial(datas, index+1)        //深度搜索继续遍历其左右子节点
		index, root.Right = rDeserial(datas, index)
	} else {
		index++
		root = nil
	}
	return index, root
}

// Deserializes your encoded data to tree.
func (cc *Codec) Deserialize(data string) *TreeNode {
	datas := strings.Split(data, ",")
	if datas[0] == "n" {
		return nil
	}
	_, root := rDeserial(datas, 0)
	return root
}

/**
func rDeserial(datas []string, index int, pre, root *TreeNode) int { //pre负责记录root的父节点，在root对应"n"时，将pre.left或者pre.right置为nil
	word := datas[index]
	if word != "n" {
		Iword, _ := strconv.Atoi(word)
		newroot := TreeNode{Val: Iword, Left: &TreeNode{}, Right: &TreeNode{}} //创建当前结点，通过赋值给指针指向位置的方式，实现对参数中的root赋值
		*root = newroot
		index = rDeserial(datas, index+1, root, root.Left) //index记录当前遍历的位置
		index = rDeserial(datas, index, root, root.Right)
	} else {
		if pre.Left == root {	//判断root是pre的左子树还是右子树
			pre.Left = nil
		} else {
			pre.Right = nil
		}
		index++
	}
	return index
}

// Deserializes your encoded data to tree.
func (cc *Codec) Deserialize(data string) *TreeNode {
	root := &TreeNode{}
	datas := strings.Split(data, ",")
	if datas[0] == "n" {
		return nil
	}
	rDeserial(datas, 0, root, root)
	return root
}
*/
