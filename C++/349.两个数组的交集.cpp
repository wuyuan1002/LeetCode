
// 349. 两个数组的交集
//
// 给定两个数组，编写一个函数来计算它们的交集。

#include <unordered_set>
#include <vector>

class Solution {
public:
    std::vector<int> intersection(std::vector<int>& nums1, std::vector<int>& nums2)
    {
        std::unordered_set<int> result;
        std::unordered_set<int> nums1_set(nums1.begin(), nums1.end());

        for (int num : nums2) {
            if (nums1_set.find(num) != nums1_set.end()) {
                result.insert(num);
            }
        }

        return std::vector<int>(result.begin(), result.end());
    }
};