
// 516. 最长回文子序列
// 给你一个字符串 s ，找出其中最长的回文子序列，并返回该序列的长度。
// 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    // 类似5、647 -- 回文子串是要连续的，回文子序列可不是连续的
    // 动态规划
    // dp[i][j]表示字符串s在[i, j]范围内最长的回文子序列的长度
    //      若s[i]与s[j]相等: dp[i][j] = dp[i+1][j-1]+2（相当于把s[i]和s[j]处两个字符加入到最长子序列中）
    //      若s[i]与s[j]不相等: dp[i][j] = max(dp[i+1][j], dp[i][j-1])（两个字符不相等, 查看把两个分别加入能构成的最长长度）
    //
    // 从递推公式：dp[i][j] = dp[i + 1][j - 1] + 2; 可以看出 递推公式是计算不到 i 和j相同时候的情况，
    // 所以需要手动初始化一下，当i与j相同，那么dp[i][j]一定是等于1的，即：一个字符的回文子序列长度就是1
    int longestPalindromeSubseq(std::string s) {
        // 创建并初始化dp
        std::vector<std::vector<int>> dp(s.size(), std::vector<int>(s.size(), 0));
        for (int i = 0; i < s.size(); i++) {
            dp[i][i] = 1;
        }

        // 开始dp -- 因为dp[i][j]依赖于dp[i + 1][j - 1], 因此dp数组应该是从左下往右上遍历 -- 字符串要倒着遍历
        for (int i = s.size() - 1; i >= 0; i--) {
            for (int j = i + 1; j < s.size(); j++) {
                if (s[i] == s[j]) {
                    dp[i][j] = dp[i + 1][j - 1] + 2;
                } else {
                    dp[i][j] = std::max(dp[i + 1][j], dp[i][j - 1]);
                }
            }
        }

        return dp[0][s.size() - 1];
    }
};
