
// 62. 不同路径
//
// 一个机器人位于一个 m x n网格的左上角 （起始点在下图中标记为 “Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
// 问总共有多少条不同的路径？

#include <vector>

class Solution {
public:
    // 动态规划 -- 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
    int uniquePaths(int m, int n) {
        std::vector<int> dp(n);  // dp数组
        dp[0] = 1;               //第一个位置的走法个数为1

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (j == 0) {
                    dp[j] = dp[j];  // 第一列
                } else {
                    dp[j] = dp[j] + dp[j - 1];
                }
            }
        }

        return dp[n - 1];
    }
};
