
// 3. 无重复字符的最长子串
//
// 给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。

#include <algorithm>
#include <string>
#include <unordered_map>

class Solution {
public:
    // 滑动窗口，同时使用map存储字符出现过的下标
    int lengthOfLongestSubstring(std::string s)
    {
        std::unordered_map<char, int> map;
        int max_len = 0;
        int l = 0, r = 0;
        for (; r < s.size(); ++r) {
            if (map.find(s[r]) != map.end()) {
                // 若该字符已经存在, 将左边界向前移动, 不能直接移动到index+1, 原因: abba
                l = std::max(l, map[s[r]] + 1);
            }
            // 若map中存在则更新最新下标, 若不存在则将kv插入
            map[s[r]] = r;

            // 更新最长长度
            max_len = std::max(max_len, r - l + 1);
        }

        return max_len;
    }
};
