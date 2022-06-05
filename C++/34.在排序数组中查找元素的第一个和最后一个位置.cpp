
// 34. 在排序数组中查找元素的第一个和最后一个位置
//
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回[-1, -1]。
// 进阶：
// 你可以设计并实现时间复杂度为O(log n)的算法解决此问题吗？

#include <vector>

class Solution {
public:
    // 二分法
    std::vector<int> searchRange(std::vector<int>& nums, int target)
    {
        return { get_first(nums, target), get_last(nums, target) };
    }

    int get_first(std::vector<int>& nums, int target)
    {
        int l = 0, r = nums.size() - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] > target) {
                r = mid - 1;
            } else if (nums[mid] < target) {
                l = mid + 1;
            } else {
                if (mid > 0 && nums[mid - 1] == target) {
                    r = mid - 1;
                } else {
                    return mid;
                }
            }
        }
        return -1;
    }

    int get_last(std::vector<int>& nums, int target)
    {
        int l = 0, r = nums.size() - 1;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] > target) {
                r = mid - 1;
            } else if (nums[mid] < target) {
                l = mid + 1;
            } else {
                if (mid < nums.size() - 1 && nums[mid + 1] == target) {
                    l = mid + 1;
                } else {
                    return mid;
                }
            }
        }
        return -1;
    }
};

int main()
{
    std::vector<int> nums = { 1 };
    Solution().searchRange(nums, 1);
}
