
// 15. 三数之和
//
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。

#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> threeSum(std::vector<int>& nums)
    {
        // std::sort(nums.begin(), nums.end());
        quick_sort(nums, 0, nums.size() - 1);

        std::vector<std::vector<int>> result;
        for (int i = 0; i < nums.size(); ++i) {
            if (nums[i] > 0) {
                // 剪枝
                break;
            } else if (i > 0 && nums[i] == nums[i - 1]) {
                // 跳过重复元素
                continue;
            }

            int l = i + 1, r = nums.size() - 1;
            while (l < r) {
                int sum = nums[i] + nums[l] + nums[r];
                if (sum > 0) {
                    --r;
                } else if (sum < 0) {
                    ++l;
                } else {
                    result.push_back({ nums[i], nums[l], nums[r] });

                    // 跳过重复元素
                    for (++l; l < r && nums[l] == nums[l - 1]; ++l) { }
                    for (--r; l < r && nums[r] == nums[r + 1]; --r) { }
                }
            }
        }

        return result;
    }

    void quick_sort(std::vector<int>& nums, int left, int right)
    {
        if (nums.empty() || left >= right) {
            return;
        }

        int l = left, r = right;
        while (l != r) {
            for (; l < r && nums[r] >= nums[left]; --r) { }
            for (; l < r && nums[l] <= nums[left]; ++l) { }
            if (l < r) {
                std::swap(nums[l], nums[r]);
            }
        }
        std::swap(nums[left], nums[l]);

        quick_sort(nums, left, l - 1);
        quick_sort(nums, r + 1, right);
    }
};
