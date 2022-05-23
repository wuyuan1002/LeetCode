
// 47. 全排列 II
//
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

#include <algorithm>
#include <unordered_set>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> permuteUnique(std::vector<int>& nums)
    {
        std::vector<std::vector<int>> result;
        dfs(nums, 0, result);
        return result;
    }

    void dfs(std::vector<int>& nums, int index, std::vector<std::vector<int>>& result)
    {
        if (index == nums.size()) {
            result.push_back(nums);
            return;
        }

        // 由于本题nums中包含重复数字, 所以需要visited记录已访问的数字, 而46题就不需要
        // 还可以使用先在permuteUnique中给nums排序, 之后在dfs中判断当前数字和前一个数字是否相等来去重
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
