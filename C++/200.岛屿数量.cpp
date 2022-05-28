
// 200. 岛屿数量
//
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。

#include <vector>

class Solution {
public:
    int numIslands(std::vector<std::vector<char>>& grid)
    {
        int count = 0;
        for (int i = 0; i < grid.size(); ++i) {
            for (int j = 0; j < grid[0].size(); ++j) {
                // 若当前节点已被访问过则是*, 水是0, 只有当前节点是陆地时是1
                if (grid[i][j] == '1') {
                    ++count;
                    dfs(grid, i, j);
                }
            }
        }

        return count;
    }

    void dfs(std::vector<std::vector<char>>& grid, int i, int j)
    {
        // 若当前位置不是岛屿或已被访问过，直接返回
        if (i < 0 || i >= grid.size() || j < 0 || j >= grid[0].size() || grid[i][j] != '1') {
            return;
        }

        // 标记当前位置已被访问
        grid[i][j] = '*';

        // 遍历前后左右
        dfs(grid, i + 1, j);
        dfs(grid, i - 1, j);
        dfs(grid, i, j + 1);
        dfs(grid, i, j - 1);
    }
};
