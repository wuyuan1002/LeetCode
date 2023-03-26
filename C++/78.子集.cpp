
// 78. 子集
//
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> subsets(std::vector<int>& nums) {
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
            res.push_back(nums[i]);
            dfs(nums, i + 1, len, res, result);
            res.pop_back();
        }
    }
};
