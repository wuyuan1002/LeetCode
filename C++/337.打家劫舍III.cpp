
// 337. 打家劫舍 III
//
// 小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。
// 除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，
// 聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
// 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
// 给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。

#include "header.h"
#include <algorithm>

class Solution {
public:
    // 偷了一个节点后，不可以再偷它的子节点
    // 1. 先按层级广度优先遍历二叉树，算出每一层的总和，存在一个数组中，再和之前一个打家劫舍一样不能偷取相邻的即可
    // 2. 动态规划 -- 每个节点都可以选择偷或不偷
    // 		当前节点选择不偷: 当前节点能偷到的最大钱数 = 左孩子能偷到的钱 + 右孩子能偷到的钱
    // 		当前节点选择偷: 当前节点能偷到的最大钱数 = 左孩子选择自己不偷时能得到的钱 + 右孩子选择不偷时能得到的钱 + 当前节点的钱数
    int rob(TreeNode* root)
    {
        std::pair<int, int> price = dfs(root);
        return std::max(price.first, price.second);
    }

    // 后序遍历二叉树 -- 求每个节点的最大价值
    std::pair<int, int> dfs(TreeNode* node)
    {
        if (node == nullptr) {
            return std::pair<int, int>(0, 0);
        }

        std::pair<int, int> l_price = dfs(node->left); // 左节点分别选择偷和不偷时能得到的最大钱数
        std::pair<int, int> r_price = dfs(node->right); // 右节点分别选择偷和不偷时能得到的最大钱数

        int rob = node->val + l_price.second + r_price.second;
        int not_rob = std::max(l_price.first, l_price.second) + std::max(r_price.first, r_price.second);

        return std::pair<int, int>(rob, not_rob);
    }
};
