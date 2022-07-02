
// 215. 数组中的第K个最大元素
//
// 在未排序的数组中找到第 k 个最大的元素。
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

#include <algorithm>
#include <vector>

class Solution {
public:
    int findKthLargest(std::vector<int>& nums, int k)
    {
        if (k <= 0 || k > nums.size()) {
            return 0;
        }

        int left = 0, right = nums.size() - 1;
        int index = partition(nums, left, right);
        while (index != k - 1) {
            index = index > k - 1 ? partition(nums, left, index - 1) : partition(nums, index + 1, right);
        }
        return nums[index];
    }

    int partition(std::vector<int>& nums, int left, int right)
    {
        int l = left, r = right;
        int povit = nums[left];

        while (l < r) {
            for (; l < r && nums[r] <= povit; r--) { }
            for (; l < r && nums[l] >= povit; l++) { }
            if (l < r) {
                std::swap(nums[l], nums[r]);
            }
        }
        std::swap(nums[left], nums[l]);

        return l;
    }
};
