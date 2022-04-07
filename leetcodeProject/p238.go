package leetcodeProject

func ProductExceptSelf(nums []int) []int {
	lennum, left, right := len(nums), 1, 1
	res := make([]int, lennum)
	for i := 0; i < lennum; i++ {
		res[i] = 1
	}
	for i := 0; i < lennum; i++ {			//left, right 分别记录左右累乘，遍历完成，每个位置相当于被赋值了两次
		res[i] = res[i] * left
		left = left * nums[i]

		res[lennum-i-1] = res[lennum-i-1] * right
		right = right * nums[lennum-i-1]
	}
	/**
	相当于
	for(int i=0;i<len;i++){
        output[i] = left;
        left *= nums[i];
    }
    for(int j=len-1;j>=0;j--){
        output[j] *= right;
        right *= nums[j];
    }
	*/
	return res
}
