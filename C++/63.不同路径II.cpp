
// 63. 不同路径 II

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

#include <vector>

class Solution {
public:
    // 动态规划 -- 到达某一个位置的走法 == 它上面位置的走法个数 + 左面位置的走法个数
    int uniquePathsWithObstacles(std::vector<std::vector<int>>& obstacleGrid)
    {
        std::vector<std::vector<int>> dp(obstacleGrid.size(), std::vector<int>(obstacleGrid[0].size(), 0)); // dp数组
        // 初始化第一行和第一列的走法个数
        for (int i = 0; i < obstacleGrid.size(); ++i) {
            if (obstacleGrid[i][0] == 1) {
                break;
            }
            dp[i][0] = 1;
        }
        for (int j = 0; j < obstacleGrid[0].size(); ++j) {
            if (obstacleGrid[0][j] == 1) {
                break;
            }
            dp[0][j] = 1;
        }

        // 开始dp
        for (int i = 1; i < obstacleGrid.size(); ++i) {
            for (int j = 1; j < obstacleGrid[0].size(); ++j) {
                if (obstacleGrid[i][j] == 1) {
                    dp[i][j] = 0;
                } else {
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                }
            }
        }

        return dp[obstacleGrid.size() - 1][obstacleGrid[0].size() - 1];
    }
};
