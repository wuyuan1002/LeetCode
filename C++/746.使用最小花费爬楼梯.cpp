
// 746. 使用最小花费爬楼梯
//
// 给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
// 一旦你支付此费用，即可选择向上爬一个或者两个台阶。
//
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
//
// 请你计算并返回达到楼梯顶部的最低花费。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划
    // dp[i]: 第i个台阶的最小花费
    // dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
    int minCostClimbingStairs(std::vector<int>& cost)
    {
        std::vector<int> dp(cost.size(), 0);
        dp[0] = cost[0];
        dp[1] = cost[1];

        for (int i = 2; i < cost.size(); i++) {
            dp[i] = std::min(dp[i - 1], dp[i - 2]) + cost[i];
        }

        return std::min(dp[cost.size() - 1], dp[cost.size() - 2]);
    }
};
