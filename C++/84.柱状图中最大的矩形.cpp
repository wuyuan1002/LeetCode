
// 84. 柱状图中最大的矩形
//
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

#include <algorithm>
#include <stack>
#include <string>
#include <vector>

class Solution {
public:
    int largestRectangleArea(std::vector<int>& heights, int w)
    {
        heights.push_back(-1); // 往heights末尾插入一个比所有元素都小的值, 确保所有元素都被计算到
        int res = 0; // 最大面积
        std::stack<int> stack; // 单调递增栈 -- 存元素下标，栈内每个元素的前一个元素都是它左面第一个比它低的元素
        for (int i = 0; i < heights.size(); ++i) {
            while (!stack.empty() && heights[i] < heights[stack.top()]) {
                // 获取栈顶元素并出栈
                int top = heights[stack.top()];
                stack.pop();

                // 获取左侧比当前位置低的柱子下标
                int left = stack.empty() ? -1 : stack.top();

                // 计算最大面积
                res = std::max(res, top * (i - left - 1));
            }
            stack.push(i);
        }
        heights.pop_back();
        return res;
    }
};

int main()
{
    std::vector<int> heights = { 5, 3, 2 };
    Solution().largestRectangleArea(heights);
}
