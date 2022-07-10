
// 718. 最长重复子数组

// 给两个整数数组 nums1 和 nums2 ，返回 两个数组中 公共的 、长度最长的子数组的长度 。

// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
// 输出：3
// 解释：长度最长的公共子数组是 [3,2,1] 。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 1. 动态规划
    // 2. 滑动窗口

    // 类似于 1143. 最长公共子序列
    // 1. 二维dp
    int findLength(std::vector<int>& nums1, std::vector<int>& nums2)
    {
        // dp[i][j]表示以下标i-1结尾的nums1和下标j-1结尾的nums2的数组的重复子数组长度
        // dp[i][j] =
        // 若 A[i-1] == B[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
        // 若 A[i-1] != B[j-1], 则 dp[i][j] = 0 --> 重复子数组需要连续, 因此不能像1143那样使用max(dp[i-1][j], dp[i][j-1])

        if (nums1.empty() || nums2.empty()) {
            return 0;
        }

        int max_len = 0;

        std::vector<std::vector<int>> dp(nums1.size() + 1, std::vector<int>(nums2.size() + 1, 0));
        dp[0][0] = 0;

        for (int i = 1; i <= nums1.size(); i++) {
            for (int j = 1; j <= nums2.size(); j++) {
                if (nums1[i - 1] == nums2[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = 0;
                }

                max_len = std::max(max_len, dp[i][j]);
            }
        }

        return max_len;
    }

    // 一维dp
    int findLength(std::vector<int>& nums1, std::vector<int>& nums2)
    {
        if (nums1.empty() || nums2.empty()) {
            return 0;
        }

        int max_len = 0;

        std::vector<int> dp(nums2.size() + 1, 0);
        dp[0] = 0;

        for (int i = 1; i <= nums1.size(); i++) {
            for (int j = nums2.size(); j >= 1; j--) {
                if (nums1[i - 1] == nums2[j - 1]) {
                    dp[j] = dp[j - 1] + 1;
                } else {
                    dp[j] = 0;
                }

                max_len = std::max(max_len, dp[j]);
            }
        }

        return max_len;
    }
};
