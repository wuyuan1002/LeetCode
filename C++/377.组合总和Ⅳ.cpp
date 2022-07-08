
// 377. 组合总和 Ⅳ
// 给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。
// 请你从 nums 中找出并返回总和为 target 的元素组合的个数。

// 题目数据保证答案符合 32 位整数范围。

#include <iostream>
#include <stdint.h>
#include <vector>

class Solution {
public:
    // 动态规划 -- 完全背包问题(数字可以被多次使用)
    int combinationSum4(std::vector<int>& nums, int target)
    {
        std::vector<int> dp(target + 1, 0); // 总和为i的组合个数 -- dp[i] += dp[i-num]
        dp[0] = 1;

        for (int i = 1; i <= target; i++) {
            for (int num : nums) {
                if (i >= num && dp[i] < INT32_MAX - dp[i - num]) { // c++会出现两个数组值相加溢出的情况, 这里做一下特殊判断
                    dp[i] += dp[i - num];
                }
            }
        }

        return dp[target];
    }
};

int main()
{
    std::vector<int> nums = { 1, 2, 3 };
    std::cout << Solution().combinationSum4(nums, 4);
}