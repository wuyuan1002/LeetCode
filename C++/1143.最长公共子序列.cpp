
// 1143. 最长公共子序列
//
// 给定两个字符串text1 和text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    // 同1035. 不相交的线
    // 类似于 392. 判断子序列、718. 最长重复子数组、115. 不同的子序列、礼物的最大价值、零钱兑换、不同路径
    // 动态规划，二维dp -- dp[i][j]表示以下标i-1结尾的text1和下标j-1结尾的text2的数组的字符串的公共子序列长度
    int longestCommonSubsequence(std::string text1, std::string text2)
    {
        if (text1.empty() || text2.empty()) {
            return 0;
        }

        // dp[i][j]表示以下标i-1结尾的text1和下标j-1结尾的text2的数组的字符串的公共子序列长度
        // dp[i][j] =
        // 若 A[i-1] == B[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
        // 若 A[i-1] != B[j-1], 则 dp[i][j] = max(dp[i-1][j], dp[i][j-1]) --> 公共子序列可以不连续, 因此不用像718那样使用0
        std::vector<std::vector<int>> dp(text1.size() + 1, std::vector<int>(text2.size() + 1, 0));

        for (int i = 1; i <= text1.size(); i++) {
            for (int j = 1; j <= text2.size(); j++) {
                if (text1[i - 1] == text2[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = std::max(dp[i][j - 1], dp[i - 1][j]);
                }
            }
        }

        return dp[text1.size()][text2.size()];
    }

    // 滚动数组 -- 一维dp
    int longestCommonSubsequence(std::string text1, std::string text2)
    {
        if (text1.empty() || text2.empty()) {
            return 0;
        }

        std::vector<int> dp(text2.size() + 1, 0);
        dp[0] = 0;

        for (int i = 1; i <= text1.size(); i++) {
            int up_left = 0; // 记录左上角的值, 防止被覆盖
            for (int j = 1; j <= text2.size(); j++) {
                int tmp = dp[j];
                if (text1[i - 1] == text2[j - 1]) {
                    dp[j] = up_left + 1;
                } else {
                    dp[j] = std::max(dp[j - 1], dp[j]);
                }

                up_left = tmp; // 记录下一个字母的左上角的值
            }
        }

        return dp[text2.size()];
    }
};
