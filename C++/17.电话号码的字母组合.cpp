
// 17. 电话号码的字母组合
//
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

#include <algorithm>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <vector>

class Solution {
private:
    const std::string phone[10] = {
        "", // 0
        "", // 1
        "abc", // 2
        "def", // 3
        "ghi", // 4
        "jkl", // 5
        "mno", // 6
        "pqrs", // 7
        "tuv", // 8
        "wxyz", // 9
    };

public:
    std::vector<std::string> letterCombinations(std::string digits)
    {
        if (digits == "") {
            return {};
        }
        std::vector<std::string> result;
        std::string res;
        dfs(digits, 0, res, result);
        return result;
    }

    void dfs(const std::string& digits, int index, std::string& res, std::vector<std::string>& result)
    {
        if (index == digits.size()) {
            result.push_back(res);
            return;
        }

        for (char c : phone[digits[index] - '0']) {
            res.push_back(c);
            dfs(digits, index + 1, res, result);
            res.pop_back();
        }
    }
};
