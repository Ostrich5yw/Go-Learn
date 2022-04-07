package leetcodeProject

func ReverseBits(num uint32) uint32 {
	var u uint32 = 0
	for i := 0; i < 32; i++ {
		u = u << 1     // 左移一位，空出最低位
		u += num & 1   // 与1做与运算可以得到num的最低位，相加即是把最低位赋给u
		num = num >> 1 // num右移一位
	}
	return u
}

// func reverseBits(num uint32) uint32 {
// 	var snum string = ""
// 	var inum int = int(num)
// 	for inum != 0 {
// 		snum = snum + strconv.Itoa(inum%2)
// 		inum = inum / 2
// 	}
// 	for len(snum) < 32 {
// 		snum = snum + "0"
// 	}
// 	u64num, _ := strconv.ParseUint(snum, 2, 32)
// 	return uint32(u64num)
// }
