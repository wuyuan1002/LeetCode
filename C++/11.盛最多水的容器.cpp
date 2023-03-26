
// 11. 盛最多水的容器
//
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。
// 在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 双指针 -- 两个指针向中间移动，每次都是移动数字小的那个
    // 类似 42
    int maxArea(std::vector<int>& height) {
        int l = 0, r = height.size() - 1;  // 左右指针
        int res = 0;                       // 最大盛水量

        while (l < r) {
            res = height[l] < height[r] ? std::max(res, (r - l) * height[l++]) : std::max(res, (r - l) * height[r--]);
        }
        return res;
    }
};

int main() {
    std::vector<int> list = {1, 8, 6, 2, 5, 4, 8, 3, 7};
    Solution().maxArea(list);
}