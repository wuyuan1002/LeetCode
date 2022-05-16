// 704. 二分查找

// 给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
// 写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。

#include <vector>

class Solution {
public:
    int search(std::vector<int>& nums, int target)
    {
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
        return -1;
    }
};

int main()
{
    std::vector<int> nums = { 2, 5, 7, 12, 34, 54, 56, 77, 89 };
    printf("%d", Solution().search(nums, 7));
    return 0;
}