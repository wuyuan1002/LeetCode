
// 416. 分割等和子集
//
// 给你一个 只包含正整数 的 非空 数组 nums 。
// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

#include <vector>

class Solution {
public:
    // Hot100 322, 518，39(与回溯法的区别)
    // 01背包问题
    // 背包问题 - 存在问题 -- dp[i] = dp[i] || dp[i-num]
    bool canPartition(std::vector<int>& nums) {
        int sum = 0;
        std::for_each(nums.begin(), nums.end(), [&sum](int n) { sum += n; });
        if (sum & 1 != 0) {
            // 总和为奇数，无法分成相等的两部分
            return false;
        }

        int target = sum / 2;                     // 背包容量
        std::vector<bool> dp(target + 1, false);  // dp[i]表示是否存在子集和为i
        dp[0] = true;                             // target=0不需要选择任何元素，所以是可以实现的

        for (int num : nums) {                     // 遍历选择列表 -- 所有零钱
            for (int i = target; i >= num; i--) {  // 遍历目标值 -- 目标值target -- 每个物品只能被使用1次时就应该倒序遍历
                dp[i] = dp[i] || dp[i - num];
            }
        }

        return dp[target];
    }
};
