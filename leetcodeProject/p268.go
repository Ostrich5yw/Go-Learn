package leetcodeProject

func MissingNumber(nums []int) int {
	res := 0
	for i := 0; i <= len(nums); i++ {	//将整个数组都遍历一遍，最终结果我们希望求得 0,1,2,3..n相异或的结果，和该数组相异或的结果进行异或操作，得到的就是res
		res ^= i						//因为这样的话，其他数字都出现两次，而缺少的数字只出现一次
		if i < len(nums) {
			res ^= nums[i]
		}
	}
	return res
}
