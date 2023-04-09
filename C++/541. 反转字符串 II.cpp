
// 541. 反转字符串 II
//
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。
//
// 如果剩余字符少于 k 个，则将剩余字符全部反转。
// 如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

#include <algorithm>
#include <string>

class Solution {
public:
    // 1. 每隔 2k 个字符的前 k 个字符进行反转
    // 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
    // 3. 剩余字符少于 k 个，则将剩余字符全部反转
    std::string reverseStr(std::string s, int k) {
        int n = s.length();
        for (int i = 0; i < n; i += 2 * k) {
            std::reverse(s.begin() + i, s.begin() + std::min(i + k, n));
        }
        return s;
    }
};
