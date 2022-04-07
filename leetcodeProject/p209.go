package leetcodeProject

import "math"

func MinSubArrayLen(target int, nums []int) int {
	res := 0
	minLen := math.MaxInt32
	slow := 0			//slow ,quick 双指针滑动窗口
	for quick := 0; quick < len(nums); quick++ {
		res += nums[quick]		//quick每次向后一个，加入新元素
		for res >= target {		//因为quick++导致窗口总和超出target，则通过slow++减少窗口总和
			if minLen > quick-slow+1 {			//如果当前长度小于minLen，则替换
				minLen = quick - slow + 1
			}
			res -= nums[slow]
			slow++
		}

	}
	if minLen == math.MaxInt32 {	//如果无法达到target，返回0
		return 0
	}
	return minLen
}

// func minSubArrayLen(target int, nums []int) int {
//     res := math.MaxInt32
//     for i:=0;i < len(nums);i ++{
//         lennum := 1
//         total := nums[i]
//         for j:=i + 1;j < len(nums);j ++{
//             if total >= target{
//                 break
//             }
//             total += nums[j]
//             lennum ++
//         }
//         if total >= target && lennum < res{
//             res = lennum
//         }
//     }
//     if res == math.MaxInt32{
//         return 0
//     }
//     return res
// }
