
// 392. 判断子序列
//
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
//
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
// （例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
//
// 进阶：
//
// 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，
// 你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

#include <string>
#include <vector>

class Solution {
public:
    // 1. 双指针
    bool isSubsequence(std::string s, std::string t)
    {
        int i = 0, j = 0;
        while (i < s.size() && j < t.size()) {
            if (s[i] == t[j]) {
                i++;
            }
            j++;
        }
        return i == s.size();
    }

    // 2. 动态规划
    // 718. 最长重复子数组、1143. 最长公共子序列、115. 不同的子序列
    bool isSubsequence(std::string s, std::string t)
    {
        // 1. 可以用最长公共子序列的dp求法求出最长公共子序列, 然后判断最长公共子序列长度是否等于s的长度
        // 2. 用下面这种dp -- 类似编辑距离
        //     dp[i][j]表示以下标i-1为结尾的字符串s, 和以下标j-1为结尾的字符串t, 相同子序列的长度为dp[i][j]
        // dp[i][j] =
        // 若 A[i-1] == B[j-1], 则 dp[i][j] = dp[i-1][j-1] + 1
        // 若 A[i-1] != B[j-1], 则 dp[i][j] = dp[i][j-1]
        std::vector<std::vector<int>> dp(s.size() + 1, std::vector<int>(t.size() + 1, 0));

        for (int i = 1; i <= s.size(); i++) {
            for (int j = 1; j <= t.size(); j++) {
                if (s[i - 1] == t[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = dp[i][j - 1];
                }
            }
        }

        return dp[s.size()][t.size()] == s.size();
    }
};
