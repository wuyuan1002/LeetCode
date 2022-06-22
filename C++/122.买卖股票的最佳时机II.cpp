
// 122. 买卖股票的最佳时机 II
//
// 给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
//
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 1. 等价于每天都买卖 -- 前一天比后一天小,买卖; 前一天比后一天大,不买卖 -- 所有上涨交易日都买卖（赚到所有利润）, 所有下降交易日都不买卖（永不亏钱）
    // 2. 动态规划 -- 类似714
    // 一天共有两种状态:
    // 0: 持有股票
    // 1: 未持有股票
    //
    // dp[i][j]表示第i天状态j时所持有的现金数量(越大说明越有钱)
    int maxProfit(std::vector<int>& prices)
    {
        if (prices.size() < 2) {
            return 0;
        }

        int profit = 0; // 利润
        for (int i = 1; i < prices.size(); i++) {
            if (prices[i] > prices[i - 1]) {
                profit += prices[i] - prices[i - 1];
            }
        }

        return profit;
    }
};
