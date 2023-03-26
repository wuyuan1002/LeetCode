
// 115. 不同的子序列
//
// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。
// （例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
// 题目数据保证答案符合 32 位带符号整数范围。

#include <stdint.h>

#include <string>
#include <vector>

class Solution {
public:
    // 动态规划 -- 392
    // dp[i][j]表示s的前i-1个字符中，t的前j-1个字符出现的次数
    // dp[i][j] =
    // 若 s[i-1] == t[j-1], 则 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
    // 若 s[i-1] != t[j-1], 则 dp[i][j] = dp[i-1][j]
    int numDistinct(std::string s, std::string t) {
        std::vector<std::vector<uint64_t>> dp(s.size() + 1, std::vector<uint64_t>(t.size() + 1, 0));
        // 初始化dp, 表示s[i-1]出现空串的个数为1
        for (int i = 0; i < s.size() + 1; i++) {
            dp[i][0] = 1;
        }

        for (int i = 1; i <= s.size(); i++) {
            for (int j = 1; j <= t.size(); j++) {
                if (s[i - 1] == t[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + dp[i - 1][j];
                } else {
                    dp[i][j] = dp[i - 1][j];
                }
            }
        }

        return dp[s.size()][t.size()];
    }
};
