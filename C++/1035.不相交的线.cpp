
// 1035. 不相交的线
//
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足满足：
//  nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
// 以这种方法绘制线条，并返回可以绘制的最大连线数。

#include <algorithm>
#include <vector>

class Solution {
public:
    // 其实就是求两个字符串的最长公共子序列的长度, 因为公共子序列数字的相对位置不变, 这样画出的线一定不相交
    //
    // 动态规划 -- 同最长公共子序列
    int maxUncrossedLines(std::vector<int>& nums1, std::vector<int>& nums2) {
        std::vector<std::vector<int>> dp(nums1.size() + 1, std::vector<int>(nums2.size() + 1, 0));

        for (int i = 1; i <= nums1.size(); i++) {
            for (int j = 1; j <= nums2.size(); j++) {
                if (nums1[i - 1] == nums2[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = std::max(dp[i][j - 1], dp[i - 1][j]);
                }
            }
        }

        return dp[nums1.size()][nums2.size()];
    }
};
