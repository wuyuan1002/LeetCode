
// 64. 最小路径和
//
// 给定一个包含非负整数的 m x n 网格 grid ，
// 请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 动态规划 -- 礼物的最大价值 -- 当前位置的最小和取决于我的上面和左面位置的最小和
    int minPathSum(std::vector<std::vector<int>>& grid)
    {
        std::vector<int> dp(grid[0].size(), 0); // 一维dp数组
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[0].size(); ++j) {
                if (j == 0) { // 第一列
                    dp[j] = grid[i][j] + dp[j];
                } else if (i == 0) { // 第一行
                    dp[j] = grid[i][j] + dp[j - 1];
                } else {
                    dp[j] = grid[i][j] + std::min(dp[j], dp[j - 1]);
                }
            }
        }

        return dp[grid[0].size() - 1];
    }
};

int main()
{
    std::vector<std::vector<int>> grid = { { 1, 3, 1 }, { 1, 5, 1 }, { 4, 2, 1 } };
    Solution().minPathSum(grid);
}