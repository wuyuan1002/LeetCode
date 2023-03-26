
// 131. 分割回文串
//
// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 回文串 是正着读和反着读都一样的字符串。

#include <algorithm>
#include <string>
#include <vector>

class Solution {
public:
    std::vector<std::vector<std::string>> partition(std::string s) {
        std::vector<std::string> res;
        std::vector<std::vector<std::string>> result;
        dfs(s, 0, res, result);
        return result;
    }

    void dfs(std::string& s, int start, std::vector<std::string>& res, std::vector<std::vector<std::string>>& result) {
        if (start == s.size()) {
            result.push_back(res);
            return;
        }

        for (int i = start; i < s.size(); ++i) {
            // 从start位置开始寻找回文串，找到了就放入集合进入下一层递归
            if (is_partion(s, start, i)) {
                res.push_back(s.substr(start, i - start + 1));
                dfs(s, i + 1, res, result);
                res.pop_back();
            }
        }
    }

    // 判断是否为回文串
    bool is_partion(std::string& s, int l, int r) {
        for (; l < r; ++l, --r) {
            if (s[l] != s[r]) {
                return false;
            }
        }
        return true;
    }
};
