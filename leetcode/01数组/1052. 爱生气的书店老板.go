package main

// 1052. 爱生气的书店老板

// 有一个书店老板，他的书店开了 n 分钟。每分钟都有一些顾客进入这家商店。给定一个长度为 n 的整数数组 customers ，
// 其中 customers[i] 是在第 i 分钟开始时进入商店的顾客数量，所有这些顾客在第 i 分钟结束后离开。
// 在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。
// 当书店老板生气时，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的。
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 minutes 分钟不生气，但却只能使用一次。
//
// 请你返回 这一天营业下来，最多有多少客户能够感到满意 。

// maxSatisfied .
// leetcode 1004. 最大连续1的个数 III
// leetcode 424. 替换后的最长重复字符
//
// 滑动窗口
//
// grumpy[i]=0说明顾客在这一分钟一定会满意因为老板一定不会生气，grumpy[i]=1说明顾客可能会不满意，因为这一分钟老板会生气，
// 但是因为老板可以在连续minutes分钟内保持不生气，也就是若当前grumpy[i]正好在minutes分钟内则可以将grumpy[i]变为0使得顾客满意
//
// 因此先将grumpy[i]=0的顾客数量求出来，这些是一定会满意的顾客，因为老板一定不会生气，
// 然后使用一个长度为minutes的滑动窗口，在这个窗口内grumpy[i]为1的可以变成0，也就是将原本因为老板生气而不满意的顾客变为满意，
// 窗口滑动的过程中更新最多能有多少顾客会因为老板变得不生气而变得满意
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// 求出一定会满意的顾客数量 -- 因为在这一分钟老板一定没有生气
	mustSatisfy := 0
	for i := range customers {
		if grumpy[i] == 0 {
			mustSatisfy += customers[i]
		}
	}

	// 最大的使变满意的客户数量、当前窗口内能使变满意的客户数量
	maxChange, currentChange := 0, 0
	// 构造一个长度为minutes的窗口并累加窗口内所有生气变为不生气的顾客数量
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			currentChange += customers[i]
		}
	}
	maxChange = currentChange

	// 不断向后移动窗口，更新窗口内所有生气变为不生气的顾客数量的最大值
	for i := minutes; i < len(customers); i++ {
		// 累加新的滑入窗口的原来是不满意的顾客数量
		if grumpy[i] == 1 {
			currentChange += customers[i]
		}
		// 减去滑出窗口的原来是不满意的顾客数量
		if grumpy[i-minutes] == 1 {
			currentChange -= customers[i-minutes]
		}

		// 统计因为老板由生气变为不生气而变得满意的顾客数量最大值
		maxChange = max(maxChange, currentChange)
	}

	// 最终 一定满意的顾客 + 由于变得不生气而变得满意 的顾客数量就是最多能够使满意的顾客数量
	return mustSatisfy + maxChange
}
