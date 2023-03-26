
// 35. 搜索插入位置
//
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。

#include <vector>

class Solution {
public:
    // 704. 二分查找
    // 34. 在排序数组中查找元素的第一个和最后一个位置
    // offer 11、leetcode 33、35、153
    int searchInsert(std::vector<int>& nums, int target) {
        int l = 0, r = nums.size() - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] > target) {
                r = mid - 1;
            } else if (nums[mid] < target) {
                l = mid + 1;
            } else {
                return mid;
            }
        }
        return r + 1;
    }
};
