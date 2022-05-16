// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
// 并返回其长度。如果不存在符合条件的子数组，返回 0 。

#include <vector>

class Solution {
public:
    int minSubArrayLen(int target, std::vector<int>& nums)
    {
        int min_len = INT32_MAX; // 最小的窗口长度
        int sum = 0;
        for (int l = 0, r = 0; r < nums.size(); ++r) {
            sum += nums[r];
            while (sum > target) {
                min_len = std::min(min_len, r - l + 1);
                sum -= nums[l++];
            }
        }

        return min_len != INT32_MAX ? min_len : 0;
    }
};

int main()
{
    std::vector<int> list = { 12, 43, 54, 23, 65, 23, 45, 76, 23, 65, 23 };
    printf("%d", Solution().minSubArrayLen(200, list));
    return 0;
}