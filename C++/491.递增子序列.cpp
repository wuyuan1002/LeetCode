
// 491. 递增子序列
//
// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。
//
// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

#include <unordered_set>
#include <vector>

class Solution {
public:
    // 同 90子集, 相当于只求递增的子集 -- 如[4,4,3,2,1]的递增子集只有[4,4], 而子集却包括[[],[2],[3],[4],[2,3],[2,4],[3,4],[4,4],[2,3,4],[2,4,4],[3,4,4],[2,3,4,4]]
    std::vector<std::vector<int>> findSubsequences(std::vector<int>& nums)
    {
        // 由于只求在原数组顺序中的递增子集, 所以要保证数组顺序, 所以不能通过对数组排序来去重

        std::vector<int> res;
        std::vector<std::vector<int>> result;
        for (int i = 2; i <= nums.size(); ++i) {
            dfs(nums, 0, i, res, result);
        }
        return result;
    }

    void dfs(std::vector<int>& nums, int start, int len, std::vector<int>& res, std::vector<std::vector<int>>& result)
    {
        if (res.size() == len) {
            result.push_back(res);
            return;
        }

        std::unordered_set<int> visited; // 用来记录数字是否已在当前层使用过
        for (int i = start; i < nums.size(); ++i) {
            if (res.size() > 0 && nums[i] < res.back()) {
                // 剪枝
                continue;
            }

            if (visited.find(nums[i]) != visited.end()) {
                // 剪枝
                continue;
            }
            visited.insert(nums[i]);

            res.push_back(nums[i]);
            dfs(nums, i + 1, len, res, result);
            res.pop_back();
        }
    }
};
