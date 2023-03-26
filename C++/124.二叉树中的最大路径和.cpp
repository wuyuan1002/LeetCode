
// 124. 二叉树中的最大路径和
//
// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
// 同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
// 路径和 是路径中各节点值的总和。
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。

#include <stdint.h>

#include <algorithm>

#include "header.h"

class Solution {
public:
    int maxPathSum(TreeNode* root) {
        int max_sum = INT32_MIN;
        path_sum(root, &max_sum);
        return max_sum;
    }

    int path_sum(TreeNode* root, int* max_sum) {
        if (root == nullptr) {
            return 0;
        }

        // 获取左右子树能够提供的最大值 -- 只有在最大贡献值大于0时, 才会选取对应子节点
        int left_sum = std::max(path_sum(root->left, max_sum), 0);
        int right_sum = std::max(path_sum(root->right, max_sum), 0);

        // 更新最大路径和
        *max_sum = std::max(*max_sum, root->val + left_sum + right_sum);

        // 返回当前节点能够为父节点提供的最大贡献值
        return root->val + std::max(left_sum, right_sum);
    }
};
