
// 40. 组合总和 II
//
// 给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
// candidates中的每个数字在每个组合中只能使用一次。
//
// 注意：解集不能包含重复的组合。

#include <algorithm>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> combinationSum2(std::vector<int>& candidates, int target)
    {
        // 先排序方便剪枝 -- 见39
        // 也可在dfs时使用visited记录已访问的数字来进行剪枝去重
        quick_sort(candidates, 0, candidates.size() - 1);
        // std::sort(candidates.begin(), candidates.end());

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

            if (i > start && candidates[i] == candidates[i - 1]) {
                // 跳过重复元素
                continue;
            }

            res.push_back(candidates[i]);
            dfs(candidates, i + 1, current_sum + candidates[i], target, res, result);
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
