// 27. 移除元素

// 给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

#include <vector>

class Solution {
public:
    // i指向的是第一个等于val的值, i之前的是不等于val的所有值
    int removeElement(std::vector<int>& nums, int val)
    {
        int i = 0;
        for (int j = 0; j < nums.size(); j++) {
            if (nums[j] != val) {
                nums[i++] = nums[j];
                // std::swap(nums[i++], nums[j]);
            }
        }
        return i;
    }
};

int main()
{
    std::vector<int> list = { 1, 6, 3, 2, 5, 3, 3, 7, 5, 3, 9 };
    printf("%d", Solution().removeElement(list, 3));
}