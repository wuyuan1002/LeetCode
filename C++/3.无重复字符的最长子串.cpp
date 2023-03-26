
// 3. 无重复字符的最长子串
//
// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

#include <algorithm>
#include <map>
#include <string>

class Solution {
public:
    // Offer 48
    // 双指针, 滑动窗口, 同时使用map存储字符出现过的下标
    int lengthOfLongestSubstring(std::string s) {
        std::map<char, int> hash;  // 用于存每个字符和它的下标
        int max_len = 0;           // 最大长度

        // 左右指针遍历字符串
        for (int l = 0, r = 0; r < s.size(); r++) {
            if (hash.find(s[r]) != hash.end()) {
                // 若该字符已经存在, 将左边界向前移动, 不能直接移动到index+1, 原因: abba
                l = std::max(l, hash[s[r]] + 1);
            }

            // 若map中存在则更新最新下标, 若不存在则将kv插入
            hash[s[r]] = r;

            // 计算最大长度
            max_len = std::max(max_len, r - l + 1);
        }

        return max_len;
    }
};
