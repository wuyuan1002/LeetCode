
// 105. 从前序与中序遍历序列构造二叉树
//
// 根据一棵树的前序遍历与中序遍历构造二叉树。
// 注意:
// 你可以假设树中没有重复的元素。

#include <iostream>
#include <vector>

#include "header.h"

class Solution {
public:
    TreeNode* buildTree(std::vector<int>& preorder, std::vector<int>& inorder) {
        return buildTree(preorder, 0, preorder.size(), inorder, 0, inorder.size());
    }

    TreeNode* buildTree(std::vector<int>& preorder, int preorder_start, int preorder_end, std::vector<int>& inorder, int inorder_start, int inorder_end) {
        // [preorder_start, preorder_end) 左闭右开
        // [inorder_start, inorder_end) 左闭右开
        std::cout << preorder_start << " " << preorder_end << " " << inorder_start << " " << inorder_end << " " << std::endl;
        if (preorder_start >= preorder_end || inorder_start >= inorder_end) {
            std::cout << "跳出" << std::endl;

            return nullptr;
        }

        int root_val = preorder[preorder_start];  // 当前跟节点值
        std::cout << root_val << std::endl;

        int root_index = inorder_start;  // 根结点在中序遍历中的下标
        for (; root_index < inorder_end; root_index++) {
            if (inorder[root_index] == root_val) {
                break;
            }
        }
        int left_length = root_index + 1;
        int right_length = inorder_end - left_length;

        TreeNode* root = new TreeNode(root_val);
        root->left = buildTree(preorder, preorder_start + 1, root_index + 1, inorder, inorder_start, root_index);
        root->right = buildTree(preorder, root_index + 1, preorder_end, inorder, root_index + 1, inorder_end);

        return root;
    }
};

int main() {
    // std::vector<int> preorder = { 3, 9, 20, 15, 7 };
    // std::vector<int> inorder = { 9, 3, 15, 20, 7 };
    std::vector<int> preorder = {1, 2, 3};
    std::vector<int> inorder = {3, 2, 1};
    Solution().buildTree(preorder, inorder);
}
