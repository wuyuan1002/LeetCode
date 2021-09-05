package main

// 1049. 最后一块石头的重量 II

// 有一堆石头，用整数数组stones 表示。其中stones[i] 表示第 i 块石头的重量。
//
// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为x 和y，且x <= y。那么粉碎的可能结果如下：
//
// 如果x == y，那么两块石头都会被完全粉碎；
// 如果x != y，那么重量为x的石头将会完全粉碎，而重量为y的石头新重量为y-x。
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

func main() {

}

// 动态规划 -- 01背包
// 类似hoot100 416
// 尽量让石头分成重量相同的两堆，相撞之后剩下的石头最小，这样就化解成01背包问题了
func lastStoneWeightII(stones []int) int {
	if stones == nil || len(stones) == 0 {
		return 0
	}

	sum := 0
	for _, n := range stones {
		sum += n
	}

	target := sum / 2           // 目标和向下取整，接下来计算容量为target的背包所能背的最大重量
	dp := make([]int, target+1) // dp[i]表示容量为i的背包所能盛放的最大重量
	dp[0] = 1
	for _, stone := range stones { // 遍历物品
		for i := target; i >= stone; i-- { // 遍历背包容量
			dp[i] = max(dp[i], dp[i-stone]+stone)
		}
	}

	return sum - 2*dp[target]
}
