
// 102. 二叉树的层序遍历
//
// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

#include <queue>
#include <utility>
#include <vector>

#include "header.h"

class Solution {
public:
    std::vector<std::vector<int>> levelOrder(TreeNode* root) {
        if (root == nullptr) {
            return {};
        }

        std::vector<std::vector<int>> result;  // 总结果
        std::queue<TreeNode*> queue;           // 队列存放遍历时的节点
        queue.push(root);

        while (!queue.empty()) {
            std::vector<int> res;                                 // 一层的结果
            for (int count = queue.size(); count > 0; count--) {  // 遍历某一层, count: 当前层节点个数
                TreeNode* node = queue.front();
                queue.pop();
                res.push_back(node->val);

                if (node->left != nullptr) {
                    queue.push(node->left);
                }
                if (node->right != nullptr) {
                    queue.push(node->right);
                }
            }
            result.push_back(std::move(res));
        }

        return result;
    }
};
