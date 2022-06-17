
// 102. 二叉树的层序遍历
//
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

#include "header.h"
#include <queue>
#include <vector>

class Solution {
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root)
    {
        if (root == nullptr) {
            return {};
        }

        std::vector<std::vector<int>> result; // 总结果
        std::queue<TreeNode*> queue; // 队列存放遍历时的节点
        queue.push(root);

        while (!queue.empty()) {
            int current_count = queue.size(); // 当前行总数
            std::vector<int> res; // 一层的结果
            for (; current_count > 0; current_count--) { // 遍历某一层
                TreeNode* node = queue.front();
                queue.pop();
                res.emplace_back(node->val);

                if (node->left != nullptr) {
                    queue.push(node->left);
                }
                if (node->right != nullptr) {
                    queue.push(node->right);
                }
            }
            result.emplace_back(res);
        }

        return result;
    }
};
