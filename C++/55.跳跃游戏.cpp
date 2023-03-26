
// 55. 跳跃游戏
//
// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。

#include <algorithm>
#include <vector>

class Solution {
public:
    bool canJump(std::vector<int>& nums) {
        int farthest = 0;  // 能跳到的最远距离
        for (int i = 0; i < nums.size(); ++i) {
            if (i > farthest) {
                return false;
            }
            farthest = std::max(farthest, i + nums[i]);
        }
        return true;
    }
};
