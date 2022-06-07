
// 53. 最大子序和
//
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
//
// 子数组 是数组中的一个连续部分。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 1.
    int maxSubArray(std::vector<int>& nums)
    {
        if (nums.empty()) {
            return 0;
        } else if (nums.size() == 1) {
            return nums[0];
        }

        int sum = nums[0], max_sum = nums[0]; // 当前和、最大和
        for (int i = 1; i < nums.size(); i++) {
            sum = std::max(sum + nums[i], nums[i]);
            max_sum = std::max(sum, max_sum);
        }

        return max_sum;
    }

    // 2. 动态规划
    int maxSubArray(std::vector<int>& nums)
    {
        if (nums.empty()) {
            return 0;
        }

        int max_sum = nums[0]; // 最大和

        std::vector<int> dp(nums.size()); // dp[i]表示以第i个数结尾的最大和
        dp[0] = nums[0];

        for (int i = 1; i < nums.size(); i++) {
            dp[i] = std::max(dp[i - 1] + nums[i], nums[i]);
            max_sum = std::max(dp[i], max_sum);
        }

        return max_sum;
    }
};
