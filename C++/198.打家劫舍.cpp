
// 198. 打家劫舍
//
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，
// 影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，
// 如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
//
// 给定一个代表每个房屋存放金额的非负整数数组，
// 计算你不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 同 213、337
    // 动态规划
    // 偷到某个房屋的最大价值 = max(当前房屋价值+前两个房屋最大价值, 前一个房屋的最大价值)
    int rob(std::vector<int>& nums)
    {
        if (nums.size() == 0) {
            return 0;
        } else if (nums.size() == 1) {
            return nums[0];
        }

        int pre_two = nums[0]; // 前两个房屋的最大价值
        int pre_one = std::max(nums[0], nums[1]); // 前一个房屋的最大价值

        for (int i = 2; i < nums.size(); i++) {
            int price = std::max(nums[i] + pre_two, pre_one); // 偷当前房屋的最大价值

            pre_two = pre_one;
            pre_one = price;
        }

        return pre_one;
    }
};
