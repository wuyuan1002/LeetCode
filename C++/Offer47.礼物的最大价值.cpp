
// Offer 47. 礼物的最大价值
// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划 -- 同64
    int maxValue(std::vector<std::vector<int>>& grid) {
        std::vector<int> dp(grid[0].size(), 0);

        for (int i = 0; i < grid.size(); i++) {
            for (int j = 0; j < grid[0].size(); j++) {
                if (j == 0) {  // 第一列
                    dp[j] = grid[i][j] + dp[j];
                } else if (i == 0) {  // 第一行
                    dp[j] = grid[i][j] + dp[j - 1];
                } else {
                    dp[j] = grid[i][j] + std::min(dp[j], dp[j - 1]);
                }
            }
        }

        return dp[grid[0].size() - 1];
    }
};
