
// 300. 最长递增子序列
//
// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 类似 128、674
    // 动态规划 -- dp[i]表示以第i个数字结尾的最长递增序列的长度 -- dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)

    int lengthOfLIS(std::vector<int>& nums) {
        std::vector<int> dp(nums.size(), 1);

        int max_len = 1;
        for (int i = 1; i < nums.size(); i++) {
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = std::max(dp[i], dp[j] + 1);
                }
            }
            max_len = std::max(max_len, dp[i]);
        }

        return max_len;
    }
};
