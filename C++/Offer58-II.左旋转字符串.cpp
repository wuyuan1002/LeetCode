
// 剑指 Offer 58 - II. 左旋转字符串
//
// 字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
// 请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，
// 该函数将返回左旋转两位得到的结果"cdefgab"。

#include <string>

class Solution {
public:
    // leetcode 151、189、541
    // 先翻转整个字符串，再翻转每一部分, 或者先翻转每个部分，再翻转整个字符串
    std::string reverseLeftWords(std::string s, int n) {
        if (s.empty() || n >= s.length()) {
            return s;
        }

        reverse(s, 0, n - 1);
        reverse(s, n, s.length() - 1);
        reverse(s, 0, s.length() - 1);

        return s;
    }

    // 反转给定string的指定区间 []
    void reverse(std::string& s, int l, int r) {
        while (l < r) {
            std::swap(s[l++], s[r--]);
        }
    }
};
