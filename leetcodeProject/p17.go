package leetcodeProject

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := []string{""}
	sMap := make(map[int32]string, 8)
	sMap[2], sMap[3], sMap[4], sMap[5], sMap[6] = "abc", "def", "ghi", "jkl", "mno"
	sMap[7], sMap[8], sMap[9] = "pqrs", "tuv", "wxyz"

	for _, c := range digits {
		c = c - '0'
		digit := sMap[c]
		tempRes := []string{}
		for _, s := range res {
			for _, tc := range digit {
				tempRes = append(tempRes, s+string(tc))
			}
		}
		res = []string{}
		res = append(res, tempRes...)
	}
	return res
}
