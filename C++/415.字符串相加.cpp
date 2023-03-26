
// 415. 字符串相加
//
// 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

#include <algorithm>
#include <string>

class Solution {
public:
    // 类似 43字符串相乘、67二进制求和
    // 大数相加
    std::string addStrings(std::string num1, std::string num2) {
        if (num1 == "" || num1 == "0") {
            return num2;
        } else if (num2 == "" || num2 == "0") {
            return num1;
        }

        std::string res;                               // 总结果
        int carry = 0;                                 // 进位
        int i = num1.size() - 1, j = num2.size() - 1;  // 两个指针分别指向两数末尾

        for (int n1 = 0, n2 = 0; i >= 0 || j >= 0 || carry > 0; --i, --j) {
            n1 = i >= 0 ? num1[i] - '0' : 0;
            n2 = j >= 0 ? num2[j] - '0' : 0;

            int sum = n1 + n2 + carry;      // 求和
            res.push_back(sum % 10 + '0');  // 当前位得数
            carry = sum / 10;               // 进位
        }

        // 翻转字符串
        std::reverse(res.begin(), res.end());

        return res;
    }
};
