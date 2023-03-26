// 203. 移除链表元素

// 给你一个链表的头节点 head 和一个整数 val ，
// 请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

#include "header.h"

class Solution {
public:
    ListNode* removeElements(ListNode* head, int val) {
        ListNode* dummy = new ListNode(0, head);
        for (ListNode* node = dummy; node != nullptr && node->next != nullptr;) {
            if (node->next->val == val) {
                node->next = node->next->next;
            } else {
                node = node->next;
            }
        }
        head = dummy->next;
        delete dummy;
        return head;
    }
};
