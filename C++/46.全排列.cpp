
// 46. 全排列
//
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

#include <algorithm>
#include <unordered_set>
#include <vector>

class Solution {
public:
    // 回溯法
    std::vector<std::vector<int>> permute(std::vector<int>& nums) {
        std::vector<std::vector<int>> result;
        dfs(nums, 0, result);
        return result;
    }

    void dfs(std::vector<int>& nums, int index, std::vector<std::vector<int>>& result) {
        if (index == nums.size()) {
            result.push_back(nums);
            return;
        }

        // 由于本题nums中不包含重复数字, 所以可以不用visited记录已访问的数字, 而47题就需要
        std::unordered_set<int> visited;

        for (int i = index; i < nums.size(); ++i) {
            if (visited.find(nums[i]) != visited.end()) {
                continue;
            }

            visited.insert(nums[i]);

            std::swap(nums[index], nums[i]);
            dfs(nums, index + 1, result);
            std::swap(nums[index], nums[i]);
        }
    }
};
