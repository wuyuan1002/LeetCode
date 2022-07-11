
// 279. 完全平方数
//

// 给定正整数n，找到若干个完全平方数（比如1, 4, 9, 16, ...）使得它们的和等于 n。
// 你需要让组成和的完全平方数的个数最少。
//
// 给你一个整数 n ，返回和为 n 的完全平方数的 最少数量 。
// 完全平方数 是一个整数，其值等于另一个整数的平方；换句话说，其值等于一个整数自乘的积。
// 例如，1、4、9 和 16 都是完全平方数，而 3 和 11 不是。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划 -- 完全背包问题
    // 类似322, 只是物品不是零钱而是1、4、9这样的完全平方数了
    // dp[i] = min(dp[i], dp[i-j*j]+1)
    int numSquares(int n)
    {
        if (n <= 0) {
            return 0;
        }

        std::vector<int> dp(n + 1, n + 1); // -- dp[i]最大的可能性也只能是n, 因此用n+1表示该金额无法被组成
        dp[0] = 0;

        for (int i = 1; i <= n; i++) { // 遍历背包容量
            for (int j = 1; j * j <= i; j++) { // 遍历物品
                dp[i] = std::min(dp[i], dp[i - j * j] + 1);
            }
        }

        return dp[n] > n ? 0 : dp[n];
    }
};
