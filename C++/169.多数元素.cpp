
// 169. 多数元素

// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于⌊ n/2 ⌋的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

#include <vector>

class Solution {
public:
    // 摩尔投票
    // 求水王数
    // 相当于一次删除两个不同的数字，有水王数时最后剩下的就是水王数，
    // 但若没有水王数时，也可能剩下来，此时需要再遍历一次数组查看剩下来数字的出现次数，比如1,2,3,4,5会剩下5
    int majorityElement(std::vector<int>& nums) {
        if (nums.size() == 0) {
            return 0;
        }

        int veto = 0, num = nums[0];
        for (int n : nums) {
            if (veto == 0) {
                num = n;
            }

            if (n == num) {
                veto++;
            } else {
                veto--;
            }
        }

        return num;
    }
};
