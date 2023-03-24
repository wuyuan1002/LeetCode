
// 442. 数组中重复的数据
//
// 给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。
// 找到所有出现两次的元素。
//
// 你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

#include <iostream>
#include <math.h>
#include <vector>

class Solution {
public:
    // 1. 遍历数组, 存入map的同时检查出现次数
    // 2. offer 03

    // 因为每个数字都在[1, n]的范围内, 所以可以使用数组下标作为原地hash的key(省去创建map的空间复杂度),
    // 使用下标位置数字的正负值来表示对应数字是否已出现过, 遍历时若下标处数字为正, 说明现在数字是第一次出现,
    // 将下标处数字置为负数, 若下标处数字为负数, 说明当前数字已经出现过, 为重复数字, 写入结果并将下标处数字
    // 重新置为正数(为了将数组恢复原样)
    std::vector<int> findDuplicates(std::vector<int>& nums)
    {
        std::vector<int> result;

        for (int n : nums) {
            // 获取当前数字的绝对值 -- 因为当前位置数字可能为负数, 所以要使用绝对值
            int num = std::abs(n);

            // 若用于标记当前数字是否存在过的下标处为负数, 说明当前数字已经出现过, 写入结果集
            if (nums[num - 1] < 0) {
                result.emplace_back(num);
            }

            // 将用于标记的下标处数字置为负数或正数,
            // 若原来是正数, 则置为负数, 说明当前数字没出现过
            // 若原来是负数, 则置为正数, 说明当前数字之前出现过, 已记入结果集, 将数组元素复原
            nums[num - 1] = -nums[num - 1];
        }

        return result;
    }
};

int main()
{
    std::vector<int> list = { 4, 3, 2, 7, 8, 2, 3, 1 };
    auto res = Solution().findDuplicates(list);
    for (auto i : res) {
        std::cout << i << std::endl;
    }
    return 0;
}
