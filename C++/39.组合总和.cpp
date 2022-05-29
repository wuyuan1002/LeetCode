
// 39. 组合总和
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
// 找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，
// 并以列表形式返回。你可以按 任意顺序 返回这些组合。
//
// candidates 中的 同一个 数字可以 无限制重复被选取 。
// 如果至少一个数字的被选数量不同，则两种组合是不同的。 
//
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    // 1. 回溯法
    // 2. 动态规划 -- 完全背包问题
    std::vector<std::vector<int>> combinationSum(std::vector<int>& candidates, int target)
    {
        // 此题说明了数组中无重复元素, 所以不用先排序进行剪枝 -- 见40
        // quick_sort(candidates, 0, candidates.size() - 1);

        std::vector<int> res;
        std::vector<std::vector<int>> result;
        dfs(candidates, 0, 0, target, res, result);
        return result;
    }

    void dfs(std::vector<int>& candidates, int start, int current_sum, int target, std::vector<int>& res, std::vector<std::vector<int>>& result)
    {
        if (current_sum == target) {
            result.push_back(res);
            return;
        }

        for (int i = start; i < candidates.size(); ++i) {
            if (current_sum + candidates[i] > target) {
                // 剪枝
                return;
            }

            // 此题说明了数组中无重复元素, 所以不用跳过重复元素
            // if (i > start && candidates[i] == candidates[i - 1]) {
            //     // 跳过重复元素
            //     continue;
            // }

            res.push_back(candidates[i]);
            dfs(candidates, i, current_sum + candidates[i], target, res, result);
            res.pop_back();
        }
    }

    void quick_sort(std::vector<int>& nums, int left, int right)
    {
        if (nums.empty() || left >= right) {
            return;
        }

        int l = left, r = right;
        int pivot = nums[left];
        while (l < r) {
            for (; l < r && nums[r] >= pivot; --r) { }
            for (; l < r && nums[l] <= pivot; ++l) { }
            if (l < r) {
                std::swap(nums[l], nums[r]);
            }
        }
        std::swap(nums[l], nums[left]);

        quick_sort(nums, left, l - 1);
        quick_sort(nums, r + 1, right);
    }
};
