package leetcodeProject

type NumArray1 struct {
	trees []int
	nLen  int
}

/**
	偶数个：1,2,3,4
		10
	3		7
   1 2     3 4
	奇数个：1,2,3
		6
	  1   5
	     2 3
	可以发现叶节点比内部节点多1，所以n-(2n-1)存放叶节点，1-(n-1)存放内部节点
*/
func Constructor11(nums []int) NumArray1 {
	nLen := len(nums)
	trees := make([]int, 2*nLen)           //构建2*n的空间来存放线段树
	for i, j := nLen, 0; i < 2*nLen; i++ { //n 到 2n-1用来存放原数组（叶节点）
		trees[i] = nums[j]
		j++
	}
	for i := nLen - 1; i > 0; i-- { // n-1 到 0 用来存放左右节点和
		trees[i] = trees[2*i] + trees[2*i+1]
	}
	return NumArray1{trees: trees, nLen: nLen}
}

func (tt *NumArray1) Update(index int, val int) {
	index = index + tt.nLen //index存放在index + n位置上
	tt.trees[index] = val
	for index > 0 { //依次更新所有父节点的值
		right, left := index, index
		if index%2 == 0 { //为零说明index在左子树，right++，否则left--;left和right分别在index/2的左右两边
			right = index + 1
		} else {
			left = index - 1
		}
		tt.trees[index/2] = tt.trees[right] + tt.trees[left]
		index /= 2
	}
}

func (tt *NumArray1) SumRange(left int, right int) int {
	// get leaf with value 'l'
	left += tt.nLen
	// get leaf with value 'r'
	right += tt.nLen
	sum := 0
	/**
	计算1,2,3,4,5,6,7,8  索引2-6的和
	第一轮：left到7，right到11，sum+right（7）=7
	第二轮：left到26，right到10，sum+left（7）+right（11）=25
	left在right右侧，遍历停止
			  36
		 10		   26
	   3   7     11  15
	  1 2 3 4   5 6 7  8
	*/
	for left <= right {
		if (left % 2) == 1 { //如果left是右子树，则只加当前left，不再由这个left向其父节点延伸
			sum += tt.trees[left]
			left++
		}
		if (right % 2) == 0 { //如果right是左子树，则只加当前right，不再由这个right向其父节点延伸
			sum += tt.trees[right]
			right--
		}
		left /= 2
		right /= 2
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
