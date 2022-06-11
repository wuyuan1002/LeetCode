
// 75. 颜色分类
//
// 给定一个包含红色、白色和蓝色，一共n 个元素的数组，
// 原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

#include <vector>

class Solution {
public:
    // 1. 插入排序
    // 2. 双指针
    // [0, p) -> 0
    // [p、q] -> 1
    // (q, nums.size() - 1] -> 2
    void sortColors(std::vector<int>& nums)
    {
        int p = 0, q = nums.size() - 1;
        for (int i = 0; i <= q;) {
            if (nums[i] == 0) {
                std::swap(nums[i++], nums[p++]);
            } else if (nums[i] == 1) {
                ++i;
            } else {
                std::swap(nums[i], nums[q--]);
            }
        }
    }
};
