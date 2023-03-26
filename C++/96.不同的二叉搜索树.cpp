
// 96. 不同的二叉搜索树
//
// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的二叉搜索树有多少种？
// 返回满足题意的二叉搜索树的种数。

#include <vector>

class Solution {
public:
    // 动态规划
    int numTrees(int n) {
        // dp[i]表示总数为i个节点时所能的不同二叉排序树的总数
        std::vector<int> dp(n + 1);
        dp[0] = 1;
        for (int i = 1; i <= n; i++) {
            for (int j = 0; j < i; j++) {
                // 求总共i个节点, 选择第j个节点作为根结点时的不同二叉排序树的数量
                dp[i] += dp[j] * dp[i - j - 1];
            }
        }
        return dp[n];
    }
};
