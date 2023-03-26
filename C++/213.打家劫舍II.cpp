
// 213. 打家劫舍 II
//
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。
// 这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
// 同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
//
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，
// 今晚能够偷窃到的最高金额。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 同 198、337
    // 房屋首尾相连，偷了第一个就不能偷最后一个，可以把环形房屋拆成两个单排房屋分别求最大值

    // 环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，
    // 因此可以把此环状排列房间问题约化为两个单排排列房间子问题：
    //
    // 在不偷窃第一个房子的情况下（即 nums[1:]），最大金额是 p1
    // 在不偷窃最后一个房子的情况下（即 nums[:n-1]），最大金额是 p2
    // 综合偷窃最大金额： 为以上两种情况的较大值，即 max(p1,p2)max(p1,p2) 。
    int rob(std::vector<int>& nums) {
        if (nums.size() == 0) {
            return 0;
        } else if (nums.size() == 1) {
            return nums[0];
        } else if (nums.size() == 2) {
            return std::max(nums[0], nums[1]);
        }

        return std::max(rob(nums, 1, nums.size()), rob(nums, 0, nums.size() - 1));
    }

    int rob(std::vector<int>& nums, int start, int end) {
        int pre_two = nums[start];                             // 前两个房屋的最大价值
        int pre_one = std::max(nums[start], nums[start + 1]);  // 前一个房屋的最大价值

        for (int i = start + 2; i < end; i++) {
            int price = std::max(nums[i] + pre_two, pre_one);
            pre_two = pre_one;
            pre_one = price;
        }

        return pre_one;
    }
};
