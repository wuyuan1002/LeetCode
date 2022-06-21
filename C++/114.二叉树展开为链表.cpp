
// 114. 二叉树展开为链表
//
// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，
// 而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。

#include "header.h"

class Solution {
public:
    // 1. 前序遍历构造链表
    // 2. 后序遍历构造链表
    TreeNode* pre = nullptr;
    void flatten(TreeNode* root)
    {
        if (root == nullptr) {
            return;
        }

        // 后序遍历二叉树, 用pre记录上一个遍历到的节点, 也就是链表的后一个节点
        flatten(root->right);
        flatten(root->left);

        // 将当前节点的后继节点记为上一个被遍历到的节点
        root->left = nullptr;
        root->right = pre;

        pre = root;
    }
};
