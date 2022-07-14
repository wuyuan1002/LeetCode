
// 343. 整数拆分
//
// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
//
// 返回 你可以获得的最大乘积 。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 同剪绳子
    int integerBreak(int n)
    {
        std::vector<int> dp(n + 1, 0); // 拆分数字i所能得到的最大乘积 -- dp[i] = max(dp[i], dp[j] * dp[i-j])
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 3;

        for (int i = 3; i <= n; i++) {
            for (int j = 1; j <= i / 2; j++) {
                dp[i] = std::max(dp[i], dp[j] * dp[i - j]);
            }
        }

        return dp[n];
    }
};
