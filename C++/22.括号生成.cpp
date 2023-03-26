
// 22. 括号生成
//
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

#include <string>
#include <vector>

class Solution {
public:
    // 回溯法 -- 先放左括号, 再放右括号, 左右括号各数都等于n时即表示得到一个结果
    std::vector<std::string> generateParenthesis(int n) {
        std::string res;
        std::vector<std::string> result;
        dfs(0, 0, n, res, result);
        return result;
    }

    void dfs(int left_count, int right_count, int n, std::string& res, std::vector<std::string>& result) {
        // 左右括号各数都等于n时即表示得到一个结果
        if (left_count == n && right_count == n) {
            result.push_back(res);
            return;
        }

        // 优先放左括号
        if (left_count < n) {
            res.push_back('(');
            dfs(left_count + 1, right_count, n, res, result);
            res.pop_back();
        }

        // 其次放右括号 -- 为什么不用(right_count < n)这个条件呢？因为只有左括号数量大于右括号数量时, 放右括号才是合法的, 才会有正确的左括号匹配, 如已经是"()"了, 就没必要再放右括号了, 放了也是不合法的
        if (left_count > right_count) {
            res.push_back(')');
            dfs(left_count, right_count + 1, n, res, result);
            res.pop_back();
        }
    }
};
