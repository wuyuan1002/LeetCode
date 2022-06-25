
// 128. 最长连续序列
//
// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 进阶：你可以设计并实现时间复杂度为O(n) 的解决方案吗？

#include <algorithm>
#include <unordered_set>
#include <vector>

class Solution {
public:
    // 类似 Hot100 300、674
    // 1.使用插入排序，在排序过程中计算最长连续序列
    // 2.依次遍历数组中每一个数字x，不断尝试寻找x+1,x+2...是否存在，这个寻找的过程，可以使用map或set来做到O(1)的时间复杂度
    // 3.不可以也像Hot100 300那样用动态规划
    int longestConsecutive(std::vector<int>& nums)
    {
        if (nums.size() == 0) {
            return 0;
        }

        // 将nums元素添加到set去重, 同时提供了O(1)的查询时间复杂度
        std::unordered_set<int> num_set(nums.begin(), nums.end());

        int longest = 0;
        for (int num : num_set) {
            if (num_set.find(num - 1) != num_set.end()) {
                // 若存在比当前元素小的元素, 直接跳过, 因为可以从更小的元素计算连续序列
                continue;
            }

            // 查找x+1,x+2,x+3,x+4...
            int i = 1;
            for (; num_set.find(num + i) != num_set.end(); i++) { }

            // 更新最长连续序列
            longest = std::max(longest, i);
        }

        return longest;
    }
};
