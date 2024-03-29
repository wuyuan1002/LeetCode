
// 309. 最佳买卖股票时机含冷冻期
//
// 给定一个整数数组，其中第i个元素代表了第i天的股票价格 。
//
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
//
// 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划
    //
    // 一天共有四种状态：
    // 0：买入股票状态（今天买入股票，或者是之前就买入了股票然后没有操作）
    // 卖出股票状态，这里就有两种卖出股票状态
    // 		1：两天前就卖出了股票，度过了冷冻期，一直没操作，今天保持卖出股票状态
    // 		2：今天卖出了股票
    // 3：今天为冷冻期状态，但冷冻期状态不可持续，只有一天！

    // dp[i][j]表示第i天状态j时所持有的现金数量(越大说明越有钱)
    // 递推公式:
    // dp[i][0] = max(dp[i - 1][0], max(dp[i - 1][3], dp[i - 1][1]) - prices[i];
    // dp[i][1] = max(dp[i - 1][1], dp[i - 1][3]);
    // dp[i][2] = dp[i - 1][0] + prices[i];
    // dp[i][3] = dp[i - 1][2];
    int maxProfit(std::vector<int>& prices) {
        std::vector<std::vector<int>> dp(prices.size(), std::vector<int>(4, 0));

        // 初始化dp
        dp[0][0] = -prices[0];
        dp[0][1] = 0;
        dp[0][2] = 0;
        dp[0][3] = 0;

        // 开始dp
        for (int i = 1; i < prices.size(); i++) {
            dp[i][0] = std::max(dp[i - 1][0], std::max(dp[i - 1][3], dp[i - 1][1]) - prices[i]);
            dp[i][1] = std::max(dp[i - 1][1], dp[i - 1][3]);
            dp[i][2] = dp[i - 1][0] + prices[i];
            dp[i][3] = dp[i - 1][2];
        }

        return std::max(dp[prices.size() - 1][3], std::max(dp[prices.size() - 1][1], dp[prices.size() - 1][2]));
    }
};
