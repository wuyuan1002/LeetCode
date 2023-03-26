
// 216. 组合总和 III
//
// 找出所有相加之和为n 的k个数的组合。组合中只允许含有 1 -9 的正整数，并且每种组合中不存在重复的数字。
//
// 说明：
// 所有数字都是正整数。
// 解集不能包含重复的组合。

#include <algorithm>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> combinationSum3(int k, int n) {
        std::vector<int> res;
        std::vector<std::vector<int>> result;
        dfs(1, k, 0, n, res, result);
        return result;
    }

    void dfs(int start, int k, int current_sum, int target, std::vector<int>& res, std::vector<std::vector<int>>& result) {
        if (current_sum == target && res.size() == k) {
            result.push_back(res);
            return;
        }

        for (int i = start; i <= 9; ++i) {
            if (current_sum + i > target) {
                // 剪枝
                return;
            }

            res.push_back(i);
            dfs(i + 1, k, current_sum + i, target, res, result);
            res.pop_back();
        }
    }
};
