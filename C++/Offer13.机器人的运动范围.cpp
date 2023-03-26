
// Offer 13. 机器人的运动范围
//
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。
// 一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），
// 也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，
// 因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

#include <functional>
#include <vector>

class Solution {
public:
    int movingCount(int m, int n, int k) {
        std::vector<std::vector<bool>> visited(m, std::vector<bool>(n, false));
        return dfs(0, 0, m, n, k, visited);
    }

    int dfs(int i, int j, int m, int n, int k, std::vector<std::vector<bool>>& visited) {
        std::function<bool()> can_move = [=]() {
            std::function<int(int)> sum = [](int num) {
                int count = 0;
                while (num > 0) {
                    count += num % 10;
                    num /= 10;
                }
                return count;
            };
            return i >= 0 && i < m && j >= 0 && j < n && !visited[i][j] && sum(i) + sum(j) <= k;
        };

        int count = 0;
        if (can_move()) {
            visited[i][j] = true;
            count = 1 + dfs(i + 1, j, m, n, k, visited) + dfs(i - 1, j, m, n, k, visited) + dfs(i, j + 1, m, n, k, visited) + dfs(i, j - 1, m, n, k, visited);
        }
        return count;
    }
};

int main() {
    Solution().movingCount(2, 3, 4);
    return 0;
}
