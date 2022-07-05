
// 322. 零钱兑换
//
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回-1。
//
// 你可以认为每种硬币的数量是无限的。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 类似Hot100 518, 416，39(与回溯法的区别)
    // 完全背包问题
    // 动态规划 -- dp[i]表示i元的最少硬币数 -- dp[i] = min(dp[i], dp[i-num]+1)
    int coinChange(std::vector<int>& coins, int amount)
    {
        if (coins.size() == 0 || amount < 0) {
            return -1;
        }

        std::vector<int> dp(amount + 1, amount + 1); // 初始化dp数组 -- dp[i]最大的可能性也只能是amount, 因此用amount+1表示该金额无法被组成
        dp[0] = 0;

        for (int coin : coins) { // 遍历每一个物品
            for (int i = 1; i <= amount; i++) { // 遍历背包容量 -- 求选择这个物品后构成目标值的最优解 -- 每个物品只能被使用1次时应该是倒序遍历，如Hot100 416、494
                if (i >= coin) {
                    dp[i] = std::min(dp[i - coin] + 1, dp[i]);
                }
            }
        }

        return dp[amount] <= amount ? dp[amount] : -1;
    }
};
