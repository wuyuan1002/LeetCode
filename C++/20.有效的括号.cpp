
// 20. 有效的括号
//
// 给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。

#include <stack>
#include <string>
#include <unordered_map>

class Solution {
public:
    std::unordered_map<char, char> brackets = {
        { '(', ')' },
        { '[', ']' },
        { '{', '}' }
    };

    // 定义一个栈，遇到左括号则入栈，遇到右括号则和栈顶元素比较并弹出栈顶元素，若与栈顶元素不同，则说明不合法
    bool isValid(std::string s)
    {
        std::stack<char> stack;
        for (char c : s) {
            if (c == '(' || c == '[' || c == '{') {
                stack.push(c);
            } else if (!stack.empty() && (c == ')' || c == ']' || c == '}')) {
                if (c != brackets[stack.top()]) {
                    // 与栈中最后一个左括号对应的合法右括号比较，若不同则返回false
                    return false;
                }
                // 若当前右括号合法，则将栈顶左括号出栈，继续检查下一个右括号
                stack.pop();
            } else {
                // 遇到了其他不合法的字符或栈中没有左括号却遇到了右括号
                return false;
            }
        }

        return stack.empty();
    }
};
