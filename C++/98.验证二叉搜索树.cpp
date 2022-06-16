
// 98. 验证二叉搜索树
//
// 给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征：
//
// 节点的左子树只包含小于当前节点的数。
// 节点的右子树只包含大于当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。

#include "header.h"
#include <stdint.h>

class Solution {
public:
    // 1. 二叉搜索树中序遍历后是排序的，因此，在中序遍历时校验当前值是否比前一个值大即可
    // 2. 前序遍历，通过最大最小值限定，校验节点是否符合要求
    bool isValidBST(TreeNode* root)
    {
        return isValidTree(root, INT64_MIN, INT64_MAX);
    }
    bool isValidTree(TreeNode* root, int64_t min, int64_t max)
    {
        if (root == nullptr) {
            return true;
        }
        if (root->val <= min || root->val >= max) {
            return false;
        }

        return isValidTree(root->left, min, root->val) && isValidTree(root->right, root->val, max);
    }
};
