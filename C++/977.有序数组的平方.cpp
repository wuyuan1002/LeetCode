// 977. 有序数组的平方
// 给你一个按非递减顺序排序的整数数组nums, 返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

#include <math.h>

#include <vector>

class Solution {
public:
    // 1. 遍历数组对每个数平方, 然后进行排序
    // 2. 双指针, 从两边进行遍历数组, 由于数组本身是由小到大排序的, 平方的最大的值由数组的最大或最小值贡献
    std::vector<int> sortedSquares(std::vector<int>& nums) {
        std::vector<int> result(nums.size(), 0);
        int l = 0, r = nums.size() - 1;
        for (int i = result.size() - 1; i >= 0; i--) {
            if (std::abs(nums[l]) > std::abs(nums[r])) {
                result[i] = nums[l] * nums[l];
                l++;
            } else {
                result[i] = nums[r] * nums[r];
                r--;
            }
        }
        return result;
    }
};

int main() {
    std::vector<int> list = {-6, -3, -2, 0, 4, 6, 9};

    auto res = Solution().sortedSquares(list);
    for (auto i : res) {
        printf("%d\n", i);
    }
    return 0;
}