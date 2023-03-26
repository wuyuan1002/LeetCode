
// 94. 二叉树的中序遍历
//
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。

#include <vector>

#include "header.h"

class Solution {
public:
    std::vector<int> inorderTraversal(TreeNode* root) {
        std::vector<int> res;
        inorder(root, res);
        return res;
    }

    void inorder(TreeNode* root, std::vector<int>& res) {
        if (root == nullptr) {
            return;
        }

        inorder(root->left, res);
        res.emplace_back(root->val);
        inorder(root->right, res);
    }
};
