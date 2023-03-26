
// 104. 二叉树的最大深度
//
// 给定一个二叉树，找出其最大深度。
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
// 说明: 叶子节点是指没有子节点的节点。

#include <algorithm>

#include "header.h"

class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (root == nullptr) {
            return 0;
        }

        return std::max(maxDepth(root->left), maxDepth(root->right)) + 1;
    }
};
