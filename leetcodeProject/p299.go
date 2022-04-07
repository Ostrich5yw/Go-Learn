package leetcodeProject

import "strconv"

func GetHint(secret string, guess string) string {
	nMap, sLen, gLen, aNum, bNum := make(map[int]int), len(secret), len(guess), 0, 0	
	for x, y := sLen-1, gLen-1; x >= 0 && y >= 0; {		//注意猜测的数(guess)和实际的数(secret)有可能位数不相等
		/**
			aNum记录两个数中，位数和值都相同的数量
			bNum记录两个数中，值相同但是位数不同的数量
			1807 7810
			aNum = 1     bNum = 3
		*/
		if x < 0 {			//如果secret已经走完了，但是guess没走完，继续看guess是否还包含secret中未匹配的值
			for ; y >= 0; y-- {
				tempy := int(guess[y] - '0')
				if nMap[tempy] < 0 {
					bNum++
				}
				nMap[tempy]++
			}
		}
		if y < 0 {		//如果guess已经走完了，但是secret没走完，继续看secret是否还包含guess中未匹配的值
			for ; x >= 0; x-- {
				tempx := int(secret[x] - '0')
				if nMap[tempx] < 0 {
					bNum++
				}
				nMap[tempx]++
			}
		}
		if secret[x] == guess[y] {
			aNum++
		} else {
			tempx, tempy := int(secret[x]-'0'), int(guess[y]-'0')		//如果temp是guess中的数，则map--，否则map++
			if nMap[tempx] < 0 {										//当temp是secret中，且map[temp]<0	说明匹配到了
				bNum++													//当temp是guess中，且map[temp]>0	说明匹配到了
			}
			nMap[tempx]++
			if nMap[tempy] > 0 {
				bNum++								//bNum不能等最后才加，因为当map[temp]=-1  我们不知道是加了一次减了两次 还是 减了一次
			}
			nMap[tempy]--
		}
		x--
		y--
	}
	return strconv.Itoa(aNum) + "A" + strconv.Itoa(bNum) + "B"
}
