
// 32. 最长有效括号
//
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

#include <algorithm>
#include <string>

class Solution {
public:
    int longestValidParentheses(std::string s)
    {
        int left = 0, right = 0, max_len = 0;

        // 正序遍历字符串
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == '(') {
                left++;
            } else {
                right++;
            }

            if (left == right) {
                max_len = std::max(max_len, left + right);
            } else if (right > left) {
                left = right = 0;
            }
        }

        // 倒序遍历一次字符串, 统计"(()"这种左括号始终大于右括号的情况
        left = right = 0;
        for (int i = s.size() - 1; i >= 0; i--) {
            if (s[i] == '(') {
                left++;
            } else {
                right++;
            }

            if (left == right) {
                max_len = std::max(max_len, left + right);
            } else if (left > right) {
                left = right = 0;
            }
        }

        return max_len;
    }
};
