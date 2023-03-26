
// 518. 零钱兑换 II
//
// 给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。
// 请你计算并返回可以凑成总金额的硬币组合数。如果任何硬币组合都无法凑出总金额，返回 0 。
//
// 假设每一种面额的硬币有无限个。
//
// 题目数据保证结果符合 32 位带符号整数。

#include <vector>

class Solution {
public:
    // 见Hot100 322, 416，39(与回溯法的区别)
    // 同样类似于跳楼梯，只不过跳的楼梯数不是1和2，而是每一个金额数了
    // 背包问题 - 组合问题-- dp[i] += dp[i-num]
    // 同Hot100 494
    int change(int amount, std::vector<int>& coins) {
        if (coins.empty()) {
            return 0;
        }

        std::vector<int> dp(amount + 1, 0);
        dp[0] = 1;

        for (int coin : coins) {
            for (int i = 1; i <= amount; i++) {
                if (i >= coin) {
                    dp[i] += dp[i - coin];
                }
            }
        }

        return dp[amount];
    }
};
