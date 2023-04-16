
// 150. 逆波兰表达式求值
//
// 根据 逆波兰表示法，求表达式的值。
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

#include <stack>
#include <vector>

class Solution {
public:
    // 遇到数字则入栈；遇到算符则取出栈顶两个数字进行计算，并将结果压入栈中
    int evalRPN(std::vector<string>& tokens) {
        // 力扣修改了后台测试数据，需要用longlong
        std::stack<long long> st;
        for (std::string token : tokens) {
            if (token == "+" || token == "-" || token == "*" || token == "/") {
                long long num1 = st.top();
                st.pop();
                long long num2 = st.top();
                st.pop();

                if (token == "+") {
                    st.push(num2 + num1);
                } else if (token == "-") {
                    st.push(num2 - num1);
                } else if (token == "*") {
                    st.push(num2 * num1);
                } else if (token == "/") {
                    st.push(num2 / num1);
                }
            } else {
                st.push(stoll(token));
            }
        }

        int result = st.top();
        st.pop();  // 把栈里最后一个元素弹出（其实不弹出也没事）
        return result;
    }
};
