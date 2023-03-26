
// 剑指 Offer 51. 数组中的逆序对
//
// 在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。
// 输入一个数组，求出这个数组中的逆序对的总数。

#include <vector>

class Solution {
public:
    // 归并排序, 在归并的过程中求逆序对
    int reversePairs(std::vector<int>& nums) {
        std::vector<int> tmp(nums.size());
        return merge_sort(nums, 0, nums.size() - 1, tmp);
    }

    // 归并排序, 同时返回逆序对
    int merge_sort(std::vector<int>& nums, int l, int r, std::vector<int>& tmp) {
        if (l >= r) {
            return 0;
        }

        // 递归排序两个子数组
        int mid = l + (r - l) / 2;
        int reverse_count = merge_sort(nums, l, mid, tmp) + merge_sort(nums, mid + 1, r, tmp);

        // 两个子数组排序完成后, 合并两个子数组

        // 先将原数组中的数据暂存至tmp临时数组
        for (int i = l; i <= r; ++i) {
            tmp[i] = nums[i];
        }
        // 合并两个子数组, 同时统计逆序对
        int i = l, j = mid + 1;
        for (int k = l; k <= r; k++) {
            if (i == mid + 1) {
                nums[k] = tmp[j++];
            } else if (j == r + 1 || tmp[i] <= tmp[j]) {
                nums[k] = tmp[i++];
            } else {
                // 如果右侧元素先入结果集，则左侧剩余的所有元素都与它组成逆序对，相反如果左边元素先插入，则说明不存在逆序对
                nums[k] = tmp[j++];
                reverse_count += mid - i + 1;  // 统计逆序对
            }
        }
        return reverse_count;
    }
};

int main() {
    std::vector<int> nums = {7, 5, 6, 4};
    Solution().reversePairs(nums);
}