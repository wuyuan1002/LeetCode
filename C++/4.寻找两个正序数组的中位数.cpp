
// 4. 寻找两个正序数组的中位数
//
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
// 请你找出并返回这两个正序数组的中位数 。

#include <vector>

class Solution {
public:
    // 1. 合并两个数组，求中位数 -- 合并过程类似归并排序的一部分，见offer 51
    // 2. 不需要合并两个数组，由于两个数组长度已知，因此最终的中位数是第几个也是已知的，使用两个指针同时遍历两个数组即可得到中位数
    double findMedianSortedArrays(std::vector<int>& nums1, std::vector<int>& nums2) {
        int len = (nums1.size() + nums2.size()), mid = len / 2;  // 总长度、中位数是总数字中的第几个
        int middle = 0, next = 0;                                // 中位数、中位数的下一个数字（若总数为偶数时使用）

        if (len % 2 != 0) {
            mid += 1;
        }

        for (int i = 0, j = 0; mid >= 0; --mid) {
            if (len % 2 != 0 && mid == 0) {
                // 若总长度为奇数且中位数已找到, 就没必要继续寻找next, 直接break
                break;
            }

            // 两指针向后移动, 若有一个已到达末尾, 则指移动另一个
            if (i == nums1.size()) {
                if (mid == 1) {
                    middle = nums2[j];  // 找到
                } else if (mid == 0) {
                    next = nums2[j];  // 找中位数的下一个数字
                }
                ++j;
            } else if (j == nums2.size()) {
                if (mid == 1) {
                    middle = nums1[i];  // 找到
                } else if (mid == 0) {
                    next = nums1[i];  // 找中位数的下一个数字
                }
                ++i;
            } else {
                if (nums1[i] <= nums2[j]) {
                    if (mid == 1) {
                        middle = nums1[i];  // 找到
                    } else if (mid == 0) {
                        next = nums1[i];  // 找中位数的下一个数字
                    }
                    ++i;
                } else {
                    if (mid == 1) {
                        middle = nums2[j];  // 找到
                    } else if (mid == 0) {
                        next = nums2[j];  // 找中位数的下一个数字
                    }
                    ++j;
                }
            }
        }

        if (len % 2 == 0) {
            return (middle + next) / 2.0;
        } else {
            return middle;
        }
    }
};

int main() {
    std::vector<int> nums1 = {};
    std::vector<int> nums2 = {1};
    Solution().findMedianSortedArrays(nums1, nums2);
}