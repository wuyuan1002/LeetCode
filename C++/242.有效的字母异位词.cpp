
// 242. 有效的字母异位词

// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
// 注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。

#include <string>
#include <unordered_map>

class Solution {
public:
    // 使用数组记录
    bool isAnagram(std::string s, std::string t)
    {
        int record[26] = { 0 };
        for (int i = 0; i < s.size(); ++i) {
            ++record[s[i] - 'a'];
        }
        for (int i = 0; i < t.size(); ++i) {
            --record[t[i] - 'a'];
        }

        for (int i = 0; i < 26; ++i) {
            if (record[i] != 0) {
                return false;
            }
        }
        return true;
    }

    // 使用map记录
    bool isAnagram1(std::string s, std::string t)
    {
        std::unordered_map<char, int> record;

        for (int i = 0; i < s.size(); ++i) {
            ++record[s[i]];
        }
        for (int i = 0; i < t.size(); ++i) {
            --record[t[i]];
        }

        for (auto item : record) {
            if (item.second != 0) {
                return false;
            }
        }
        return true;
    }
};