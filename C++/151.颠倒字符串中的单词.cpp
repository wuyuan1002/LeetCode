
// 151. 翻转字符串里的单词

// 给你一个字符串 s ，逐个翻转字符串中的所有单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
//
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。

#include <algorithm>
#include <iostream>
#include <string>

class Solution {
public:
    // 先翻转整个字符串，再翻转每一个单词, 或者先翻转每个单词，再翻转整个字符串
    std::string reverseWords(std::string s)
    {
        if (s.empty()) {
            return s;
        }

        // 去掉字符串两端空格
        s.erase(0, s.find_first_not_of(' '));
        s.erase(s.find_last_not_of(' ') + 1);

        // 翻转整个字符串
        reverse(s, 0, s.size() - 1);

        // 翻转每个单词
        int start = 0; // 当前单词首字母下标
        for (int i = 1; i < s.size(); ++i) {
            if (s[i] == ' ') { // 翻转前一个单词
                reverse(s, start, i - 1);
                start = i + 1;
            } else if (i == s.size() - 1) { // 到达最后一个字母
                reverse(s, start, i);
            }
        }

        // 去掉字母间的多余空格
        for (int i = 0; i < s.size(); ++i) {
            if (s[i] == ' ' && s[i - 1] == ' ') {
                s.erase(i--, 1); // 删除i位置后把下标向前移动一位, 防止漏掉空格, 如: ab   cdefg
            }
        }

        return s;
    }

    // 反转给定string的指定区间
    void reverse(std::string& s, int l, int r)
    {
        while (l < r) {
            std::swap(s[l++], s[r--]);
        }
    }
};
