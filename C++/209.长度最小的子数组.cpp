
// 209. 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，
// 并返回其长度。如果不存在符合条件的子数组，返回 0 。

#include <vector>

class Solution {
public:
    // 双指针, 滑动窗口
    // 右指针不断向前移动把数字加到sum, 当发现 sum >= target 时, 开始向前移动左指针, 并不断记录窗口大小
    int minSubArrayLen(int target, std::vector<int>& nums) {
        int min_len = INT32_MAX;  // 最小窗口长度
        int sum = 0;
        for (int l = 0, r = 0; r < nums.size(); r++) {
            sum += nums[r];
            while (sum >= target) {
                min_len = std::min(min_len, r - l + 1);
                sum -= nums[l++];
            }
        }
        return min_len == INT32_MAX ? 0 : min_len;
    }
};

int main() {
    std::vector<int> list = {12, 43, 54, 23, 65, 23, 45, 76, 23, 65, 23};
    printf("%d", Solution().minSubArrayLen(200, list));
    return 0;
}