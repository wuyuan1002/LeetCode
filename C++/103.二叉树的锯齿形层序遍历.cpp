
// 103. 二叉树的锯齿形层序遍历
//
// 给定一个二叉树，返回其节点值的锯齿形层序遍历。
// （即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

#include <deque>
#include <utility>
#include <vector>

#include "header.h"

class Solution {
public:
    std::vector<std::vector<int>> zigzagLevelOrder(TreeNode* root) {
        if (root == nullptr) {
            return {};
        }

        std::vector<std::vector<int>> result;  // 总结果
        std::deque<TreeNode*> deque;           // 双端队列 -- 存放遍历时的节点
        deque.push_back(root);

        int num = 0;  // 当前是第几行
        while (!deque.empty()) {
            std::vector<int> res;                                 // 一层总结果
            for (int count = deque.size(); count > 0; count--) {  // 遍历某一层, count: 当前层节点个数
                if (num % 2 == 0) {                               // 偶数行, 从队头读, 左右放队尾
                    TreeNode* node = deque.front();
                    deque.pop_front();

                    res.push_back(node->val);

                    if (node->left != nullptr) {
                        deque.push_back(node->left);
                    }
                    if (node->right != nullptr) {
                        deque.push_back(node->right);
                    }
                } else {  // 奇数行, 从队尾读, 右左放队头
                    TreeNode* node = deque.back();
                    deque.pop_back();

                    res.push_back(node->val);

                    if (node->right != nullptr) {
                        deque.push_front(node->right);
                    }
                    if (node->left != nullptr) {
                        deque.push_front(node->left);
                    }
                }
            }

            result.push_back(std::move(res));
            num++;
        }

        return result;
    }
};
