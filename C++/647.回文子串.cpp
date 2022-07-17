
// 647. 回文子串
// 
// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
// 回文字符串 是正着读和倒过来读一样的字符串。
// 子字符串 是字符串中的由连续字符组成的一个序列。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

#include <string>
#include <vector>

class Solution {
public:
    // 动态规划 -- 二维dp
    // dp[i][j] -- s[i]与s[j]之间的子串是否为回文串（[i, j]为左闭右闭）
    //
    // 当s[i]与s[j]不相等，dp[i][j]一定是false。
    // 当s[i]与s[j]相等时，这就复杂一些了，有如下三种情况
    //      情况一：下标i 与 j相同，同一个字符例如a，当然是回文子串
    //      情况二：下标i 与 j相差为1，例如aa，也是回文子串
    //      情况三：下标：i 与 j相差大于1的时候，例如cabac，此时s[i]与s[j]已经相同了，我们看i到j区间是不是回文子串就看aba是不是回文就可以了，那么aba的区间就是 i+1 与 j-1区间，这个区间是不是回文就看dp[i + 1][j - 1]是否为true。
    int countSubstrings(std::string s)
    {
        std::vector<std::vector<bool>> dp(s.size(), std::vector<bool>(s.size(), false)); // dp数组
        int result = 0;

        // 因为dp[i][j]依赖于dp[i + 1][j - 1], 因此dp数组应该是从左下往右上遍历 -- 字符串要倒着遍历
        for (int i = s.size() - 1; i >= 0; i--) {
            for (int j = i; j < s.size(); j++) {
                if (s[i] == s[j] && (j - i <= 1 || dp[i + 1][j - 1])) {
                    result++;
                    dp[i][j] = true;
                }
            }
        }

        return result;
    }
};
