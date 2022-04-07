package leetcodeProject

func MaxProfit(prices []int) int {
	len := len(prices)
	buy, sell, cool := make([]int, len), make([]int, len), make([]int, len) //分别表示i之前最后一个操作为买入，卖出，冷冻
	for i, val := range prices {
		if i == 0 {
			sell[i] = 0
			buy[i] = -val
			cool[i] = 0
			continue
		}
		sell[i] = max(buy[i-1]+val, sell[i-1]) //卖出之前的操作必然是买入，也有一种特殊情况是第i步没有进行操作
		buy[i] = max(cool[i-1]-val, buy[i-1])  //买入之前的操作必然是冷冻，也有一种特殊情况是第i步没有进行操作
		cool[i] = max(cool[i-1], sell[i-1])    //冷冻操作的上一步必然是卖出，也有一种特殊情况是第i步没有进行操作，所以第i步之前最后一个操作为cool，则等于第i-1步之前最后一个操作为cool
	}
	return max(max(sell[len-1], buy[len-1]), cool[len-1])
}
  