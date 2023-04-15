
// 1047. 删除字符串中的所有相邻重复项
//
// 给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。
// 在 S 上反复执行重复项删除操作，直到无法继续删除。
// 在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

#include <stack>
#include <string>

class Solution {
public:
    // 使用栈保存遍历过程中的字符, 遍历下一个字符时与栈顶元素比较, 若相同则出栈, 最终栈内元素就是删除后的结果
    // 也可以使用veckor存遍历的字符, 这样在构建结果时不用进行反转操作
    std::string removeDuplicates(std::string s) {
        std::stack<char> stack;
        for (char c : s) {
            if (stack.empty() || c != stack.top()) {
                stack.push(c);
            } else {
                stack.pop();
            }
        }

        // 构造结果
        std::string result = "";
        while (!stack.empty()) {
            result += stack.top();
            stack.pop();
        }

        // 栈内元素是倒序的, 这里需要进行反转
        std::reverse(result.begin(), result.end());

        return result;
    }
};
