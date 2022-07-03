
// 338. 比特位计数
//
// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，
// 计算其二进制数中的 1 的数目并将它们作为数组返回。

#include <vector>

class Solution {
public:
    // 同offer 15
    // 计算二进制中1的个数
    // 1. 使用1和数字做与运算，之后每次把1左移一位
    // 2. 每个数和它-1的数字做与运算，可以把该整数的最右面的1变成0，那么该整数有多少个1就可以做多少次这样的运算
    std::vector<int> countBits(int n)
    {
        std::vector<int> result(n + 1, 0);
        for (int i = 0; i <= n; i++) {
            int num = i, count = 0;
            while (num != 0) {
                count++;
                num &= (num - 1);
            }
            result[i] = count;
        }

        return result;
    }
};
