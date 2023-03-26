
// 152. 乘积最大子数组
//
// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），
// 并返回该子数组所对应的乘积。

#include <vector>

class Solution {
public:
    // 动态规划
    // 以某个值结尾的最大乘积 = max(当前元素, max(以上一个元素结尾的的最小乘积 * 当前元素, 以上一个元素结尾的的最大乘积 * 当前元素))
    int maxProduct(std::vector<int>& nums) {
        if (nums.size() < 2) {
            return nums[0];
        }

        int max_product = nums[0];  // 最大乘积

        // 以前一个元素结尾的子数组乘积的最小值、最大值
        int min_pre = nums[0], max_pre = nums[0];

        for (int i = 1; i < nums.size(); i++) {
            // 更新当前元素相乘后的最大最小值
            int max = max_pre, min = min_pre;
            max_pre = std::max(max * nums[i], std::max(nums[i], min * nums[i]));
            min_pre = std::min(min * nums[i], std::min(nums[i], max * nums[i]));

            // 更新总结果
            max_product = std::max(max_product, max_pre);
        }

        return max_product;
    }
};
