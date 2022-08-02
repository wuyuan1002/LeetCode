
// 121. 买卖股票的最佳时机
//
// 给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。
// 设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划
    //
    // 找出在哪天进行买入、在哪天进行卖出可以获得最大利润 -- 只能买卖一次
    // 左取最小值, 右取最大值
    int maxProfit(std::vector<int>& prices)
    {
        if (prices.size() < 2) {
            return 0;
        }

        int min_cost = prices[0]; // 最小成本
        int max_profit = 0; // 最大利润

        for (int cost : prices) {
            if (cost < min_cost) {
                // 更新最小成本
                min_cost = cost;
            } else {
                // 更新最大利润
                max_profit = std::max(max_profit, cost - min_cost);
            }
        }

        return max_profit;
    }
};
