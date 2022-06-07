
// 49. 字母异位词分组
//
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

class Solution {
public:
    // 1. 排序 -- 用map存排序后的字符串和各个字母异位词字符串
    // 2. 计数 -- 用map存计数结果和各个字母异位词字符串
    std::vector<std::vector<std::string>> groupAnagrams(std::vector<std::string>& strs)
    {
        std::unordered_map<std::string, std::vector<std::string>> mp;
        for (std::string& str : strs) {
            std::string s = str;
            std::sort(s.begin(), s.end());
            mp[s].emplace_back(str);
        }

        std::vector<std::vector<std::string>> res;
        for (auto& p : mp) {
            res.push_back(std::move(p.second));
        }

        return res;
    }
};
