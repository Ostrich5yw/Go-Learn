package leetcodeProject

func TwoSum(numbers []int, target int) []int {
	res := []int{0, 0}
	for i := 0; i < len(numbers); i++ {
		nowTarget := target - numbers[i]
		left := 0
		right := len(numbers) - 1
		for left <= right {
			mid := (left + right) / 2
			if numbers[mid] == nowTarget {
				res[0] = i + 1
				res[1] = mid + 2
				if mid != i {
					res[1] -= 1
				}
				return res
			} else if numbers[mid] < nowTarget {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return res
}
