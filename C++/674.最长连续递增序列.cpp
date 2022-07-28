
// 674. 最长连续递增序列
//
// 给定一个未经排序的整数数组，找到最长且 连续递增的子序列，并返回该序列的长度。
//
// 连续递增的子序列 可以由两个下标 l 和 r（l < r）确定，如果对于每个 l <= i < r，
// 都有 nums[i] < nums[i + 1] ，那么子序列 [nums[l], nums[l + 1], ...,
// nums[r - 1], nums[r]] 就是连续递增子序列。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 类似 128、300
    // 动态规划 -- dp[i]表示以下标i处数字结尾的连续递增序列长度
    int findLengthOfLCIS(std::vector<int>& nums)
    {
        std::vector<int> dp(nums.size(), 1);

        int max_len = 1;
        for (int i = 1; i < nums.size(); i++) {
            if (nums[i] > nums[i - 1]) {
                dp[i] = dp[i - 1] + 1;
            }
            max_len = std::max(max_len, dp[i]);
        }

        return max_len;
    }
};
