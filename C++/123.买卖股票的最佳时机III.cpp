
// 123. 买卖股票的最佳时机 III
//
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划
    //
    // 一天一共就有五个状态：
    // 0. 没有操作
    // 1. 第一次买入
    // 2. 第一次卖出
    // 3. 第二次买入
    // 4. 第二次卖出
    // dp[i][j]中 i表示第i天，j为 [0 - 4] 五个状态，dp[i][j]表示第i天状态j时所持有的现金数量(越大说明越有钱)
    int maxProfit(std::vector<int>& prices) {
        std::vector<std::vector<int>> dp(prices.size(), std::vector<int>(5));

        // 初始化dp
        dp[0][0] = 0;
        dp[0][1] = -prices[0];
        dp[0][2] = 0;
        dp[0][3] = -prices[0];
        dp[0][4] = 0;

        // 开始dp
        for (int i = 1; i < prices.size(); i++) {
            dp[i][0] = dp[i - 1][0];                                      // 保持前一天的未操作状态
            dp[i][1] = std::max(dp[i - 1][0] - prices[i], dp[i - 1][1]);  // 保持前一天的买入状态或在前一天没有买入时进行买入
            dp[i][2] = std::max(dp[i - 1][1] + prices[i], dp[i - 1][2]);  // 保持前一天的卖出状态或在前一天没有卖出时卖出
            dp[i][3] = std::max(dp[i - 1][2] - prices[i], dp[i - 1][3]);  // 保持前一天的买入状态或在前一天没有买入时进行买入
            dp[i][4] = std::max(dp[i - 1][3] + prices[i], dp[i - 1][4]);  // 保持前一天的卖出状态或在前一天没有卖出时卖出
        }

        return dp[prices.size() - 1][4];
    }
};
