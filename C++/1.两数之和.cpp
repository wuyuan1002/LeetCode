// 1.两数之和

// 给定一个整数数组 nums和一个整数目标值 target，
// 请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
// 你可以按任意顺序返回答案。

#include <iostream>
#include <map>
#include <vector>

class Solution {
public:
    // 使用map存放已经遍历过的数字, 遍历之后的数字时在map中查找
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        std::map<int, int> hash;  // k: 数值, v: 下标
        for (int i = 0; i < nums.size(); i++) {
            std::map<int, int>::iterator iter = hash.find(target - nums[i]);
            if (iter != hash.end()) {
                return {iter->second, i};
            }
            hash.insert(std::make_pair(nums[i], i));
        }
        return {};
    }
};

int main() {
    std::vector<int> list = {1, 5, 2, 7, 34, 87, 23};
    auto res = Solution().twoSum(list, 36);
    for (auto i : res) {
        std::cout << i << std::endl;
    }
    return 0;
}
