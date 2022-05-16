// 977. 有序数组的平方
// 给你一个按非递减顺序排序的整数数组nums, 返回每个数字的平方组成的新数组，要求也按非递减顺序排序。

#include <vector>

class Solution {
public:
    std::vector<int> sortedSquares(std::vector<int>& nums)
    {
        std::vector<int> result(nums.size(), 0);
        int l = 0, r = nums.size() - 1;
        for (int i = result.size() - 1; i >= 0; --i) {
            if (std::abs(nums[l]) > std::abs(nums[r])) {
                result[i] = (nums[l] * nums[l]);
                ++l;
            } else {
                result[i] = (nums[r] * nums[r]);
                --r;
            }
        }
        return result;
    }
};

int main()
{
    std::vector<int> list = { -6, -3, -2, 0, 4, 6, 9 };

    auto res = Solution().sortedSquares(list);
    for (auto i : res) {
        printf("%d\n", i);
    }
    return 0;
}