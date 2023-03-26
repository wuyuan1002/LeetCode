
// 101. 对称二叉树
//
// 给定一个二叉树，检查它是否是镜像对称的。
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

#include "header.h"

class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        return isSymme(root, root);
    }

    bool isSymme(TreeNode* left, TreeNode* right) {
        if (left == nullptr && right == nullptr) {
            return true;
        }
        if (left == nullptr || right == nullptr) {
            return false;
        }
        if (left->val != right->val) {
            return false;
        }
        return isSymme(left->left, right->right) && isSymme(left->right, right->left);
    }
};
