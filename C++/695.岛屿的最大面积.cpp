
// 695. 岛屿的最大面积
//
//
// 给定一个包含了一些 0 和 1 的非空二维数组grid 。
//
// 一个岛屿是由一些相邻的1(代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。
// 你可以假设grid 的四个边缘都被 0（代表水）包围着。
//
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)

#include <algorithm>
#include <vector>

class Solution {
public:
    // 等于在所有机器人的运动的范围中最大的那个
    int maxAreaOfIsland(std::vector<std::vector<int>>& grid) {
        int max_area = 0;
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[i].size(); ++j) {
                // 若当前节点已被访问过则是-1, 水是0, 只有当前节点是陆地时是1
                if (grid[i][j] == 1) {
                    max_area = std::max(max_area, dfs(grid, i, j));
                }
            }
        }
        return max_area;
    }

    int dfs(std::vector<std::vector<int>>& grid, int i, int j) {
        // 若当前位置不是岛屿或已被访问过，直接返回
        if (i < 0 || i >= grid.size() || j < 0 || j >= grid[0].size() || grid[i][j] != 1) {
            return 0;
        }

        // 标记当前位置已被访问
        grid[i][j] = -1;

        // 前后左右求当前岛屿的面积
        return 1 + dfs(grid, i + 1, j) + dfs(grid, i - 1, j) + dfs(grid, i, j + 1) + dfs(grid, i, j - 1);
    }
};
