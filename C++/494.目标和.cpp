
// 494. 目标和
//
// 给你一个整数数组 nums 和一个整数 target 。
//
// 向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
//
// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，
// 然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同表达式的数目。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 1. 回溯法 -- 超时
    // 2. 动态规划 -- 01背包问题 - 组合问题
    // 类似377、416、518
    //
    // left组合 - right组合 = target
    // left + right等于sum，而sum是固定的
    // left - (sum - left) = target -> left = (target + sum)/2 -> 问题转换成了组合总和
    int findTargetSumWays(std::vector<int>& nums, int target) {
        int sum = 0;
        std::for_each(nums.begin(), nums.end(), [&sum](int n) { sum += n; });

        if (target > sum || (sum + target) % 2 == 1) {
            return 0;
        }

        // 计算目标值
        int t = (sum + target) / 2;
        if (t < 0) {  // 防止nums = [100], target = -200
            return 0;
        }

        // 进行dp
        std::vector<int> dp(t + 1, 0);  // dp[i] += dp[i - num]
        dp[0] = 1;

        for (int num : nums) {
            for (int i = t; i >= num; i--) {
                dp[i] += dp[i - num];
            }
        }

        return dp[t];
    }
};
