package leetcodeProject

func isBadVersion(a int) bool { return false }

func FirstBadVersion(n int) int {
	res := 1
	for slow, quick := 1, n; slow <= quick; {
		mid := (slow + quick) / 2
		if isBadVersion(mid) {
			quick = mid - 1
			res = mid
		} else {
			slow = mid + 1
		}
	}
	return res
}
