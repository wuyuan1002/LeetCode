
// 77. 组合
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
//
// 你可以按 任何顺序 返回答案。

#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> combine(int n, int k) {
        std::vector<int> res;
        std::vector<std::vector<int>> result;
        dfs(n, k, 1, res, result);
        return result;
    }

    void dfs(int n, int k, int start, std::vector<int>& res, std::vector<std::vector<int>>& result) {
        if (res.size() == k) {
            result.push_back(res);
            return;
        }

        for (int i = start; i <= n; ++i) {
            if (res.size() + n - start + 1 < k) {
                // 剪枝 -- 若已有元素+剩余元素 < k 时可以直接跳出，因为无论如何都不会满足总数等于k了
                return;
            }

            res.push_back(i);
            dfs(n, k, i + 1, res, result);
            res.pop_back();
        }
    }
};
