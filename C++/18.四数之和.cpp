
// 18. 四数之和
//
// 给定一个包含n 个整数的数组nums和一个目标值target，
// 判断nums中是否存在四个元素 a，b，c和 d，使得a + b + c + d的值与target相等？
// 找出所有满足条件且不重复的四元组。
//
// 注意：答案中不可以包含重复的四元组。

#include <algorithm>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> fourSum(std::vector<int>& nums, int target) {
        quick_sort(nums, 0, nums.size() - 1);

        std::vector<std::vector<int>> result;
        for (int i = 0; i < nums.size(); ++i) {  // 先固定一个值
            if (i > 0 && nums[i] == nums[i - 1]) {
                // 跳过重复元素
                continue;
            }

            for (int j = i + 1; j < nums.size(); ++j) {  // 进行三数之和
                if (j > i + 1 && nums[j] == nums[j - 1]) {
                    // 跳过重复元素
                    continue;
                }

                int l = j + 1, r = nums.size() - 1;
                while (l < r) {
                    // int sum = nums[i] + nums[j] + nums[l] + nums[r]; // 会溢出
                    if (nums[i] + nums[j] > target - (nums[l] + nums[r])) {
                        --r;
                    } else if (nums[i] + nums[j] < target - (nums[l] + nums[r])) {
                        ++l;
                    } else {
                        result.push_back({nums[i], nums[j], nums[l], nums[r]});

                        for (++l; l < r && nums[l] == nums[l - 1]; ++l) {
                        }
                        for (--r; l < r && nums[r] == nums[r + 1]; --r) {
                        }
                    }
                }
            }
        }

        return result;
    }

    void quick_sort(std::vector<int>& nums, int left, int right) {
        if (nums.empty() || left >= right) {
            return;
        }

        int l = left, r = right;
        int pivot = nums[left];
        while (l < r) {
            for (; l < r && nums[r] >= pivot; --r) {
            }
            for (; l < r && nums[l] <= pivot; ++l) {
            }
            if (l < r) {
                std::swap(nums[l], nums[r]);
            }
        }
        std::swap(nums[left], nums[l]);

        quick_sort(nums, left, l - 1);
        quick_sort(nums, r + 1, right);
    }
};
