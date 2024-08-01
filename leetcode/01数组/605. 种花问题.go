package main

// 605. 种花问题

// 假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。
// 另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。

// canPlaceFlowers .
// 若要在flowerbed[i]处种花，需要满足flowerbed[i-1]、flowerbed[i]、flowerbed[i+1]三处都为0
// 为简化首尾的判断逻辑，防止下标溢出的情况，可以在数组首尾各插入一个0
func canPlaceFlowers(flowerbed []int, n int) bool {
	// 在首尾分别插入0
	flowerbed = append(append([]int{0}, flowerbed...), 0)

	// 在能种花的位置进行种花并递减花的数量
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1 // 进行种花
			n--              // 递减花的数量

			// 若已经种够了则跳出，不继续进行种花
			if n <= 0 {
				break
			}
		}
	}

	return n <= 0
}
