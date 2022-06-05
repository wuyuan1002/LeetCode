
// 42. 接雨水
//
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

#include <algorithm>
#include <stack>
#include <vector>

class Solution {
public:
    // 1. 双指针 -- 类似 11
    int trap(std::vector<int>& height)
    {
        int lmax = 0, rmax = 0; // 左右最高的列
        int res = 0; // 盛水量

        for (int l = 0, r = height.size() - 1; l < r;) {
            // 更新左右最大值
            lmax = std::max(lmax, height[l]);
            rmax = std::max(rmax, height[r]);

            // 将当前列的盛水量加到总结果
            res += lmax <= rmax ? lmax - height[l++] : rmax - height[r--];
        }

        return res;
    }

    // 2. 单调栈 -- 类似 84、239、739、496、503
    int trap(std::vector<int>& height)
    {
        int res = 0; // 盛水量
        std::stack<int> stack; // 单调递减栈
        for (int i = 0; i < height.size(); i++) {
            // 若栈不为空且当前元素大于栈顶元素，依次将栈内大于当前元素的值出栈，计算它能接的雨水数量
            while (!stack.empty() && height[i] > height[stack.top()]) {
                // 获取当前栈顶元素，并出栈
                int top = height[stack.top()];
                stack.pop();

                // 若当前栈顶元素为栈内第一个元素，直接退出循环然后将当前元素入栈，因为左面没有比它高的元素，构不成凹形无法盛接雨水
                if (stack.empty()) {
                    break;
                }

                int left = height[stack.top()]; // 左侧高列
                int right = height[i]; // 右侧高列
                int width = i - stack.top() - 1; // 左右侧高列间距
                int height = std::min(left, right) - top; // 当前列的高度差

                // 将当前列的盛水量加到总结果
                res += width * height;
            }

            // 若栈为空或当前元素小于栈顶元素了，将当前元素入栈
            stack.push(i);
        }

        return res;
    }
};
