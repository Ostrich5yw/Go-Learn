package leetcodeProject

type NumArray struct {
	ttt []int
}
// 题目强调需要进行无数次SumRange，所以要在Constructor阶段存放0到所有位置的和
func Constructor8(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]		//nums[i]记录0-i的数字和
	}
	return NumArray{ttt: nums}
}

func (array *NumArray) SumRange(left int, right int) int {
	if left == 0 {					//如果left=0，刚好就是num[right]
		return array.ttt[right]
	}
	return array.ttt[right] - array.ttt[left-1]		//0-right之和减0-left之和等于left-right之和
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
