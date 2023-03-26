
// 90. 子集 II
//
// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> subsetsWithDup(std::vector<int>& nums) {
        // 因为有重复元素, 所以先排序方便剪枝
        quick_sort(nums, 0, nums.size() - 1);

        std::vector<int> res;
        std::vector<std::vector<int>> result;
        for (int i = 0; i <= nums.size(); ++i) {
            dfs(nums, 0, i, res, result);
        }
        return result;
    }

    void dfs(std::vector<int>& nums, int start, int len, std::vector<int>& res, std::vector<std::vector<int>>& result) {
        if (res.size() == len) {
            result.push_back(res);
            return;
        }

        for (int i = start; i < nums.size(); ++i) {
            if (i > start && nums[i] == nums[i - 1]) {
                // 剪枝
                continue;
            }

            res.push_back(nums[i]);
            dfs(nums, i + 1, len, res, result);
            res.pop_back();
        }
    }

    void quick_sort(std::vector<int>& nums, int left, int right) {
        if (nums.empty() || left >= right) {
            return;
        }

        int l = left, r = right;
        int pivot = nums[left];
        while (l < r) {
            for (; l < r && nums[r] >= pivot; --r) {
            }
            for (; l < r && nums[l] <= pivot; ++l) {
            }
            if (l < r) {
                std::swap(nums[l], nums[r]);
            }
        }
        std::swap(nums[l], nums[left]);

        quick_sort(nums, left, l - 1);
        quick_sort(nums, r + 1, right);
    }
};
