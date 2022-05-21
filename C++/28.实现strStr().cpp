
// 28. 实现 strStr()
//
// 给你两个字符串 haystack 和 needle,
// 请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。
// 如果不存在，则返回  -1 。

#include <string>

class Solution {
public:
    int strStr(std::string haystack, std::string needle)
    {
        int index = -1;
        for (int i = 0; i < haystack.size(); ++i) {
            int k = i, j = 0;
            index = i;
            for (; k < haystack.size() && j < needle.size(); ++k, ++j) {
                if (haystack[k] != needle[j]) {
                    break;
                }
            }
            if (j == needle.size()) {
                break;
            }
            index = -1;
        }

        return index;
    }
};
