
// 24. 两两交换链表中的节点

// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

#include "header.h"

class Solution {
public:
    // 25. K个一组翻转链表
    // 递归
    ListNode* swapPairs(ListNode* head) {
        ListNode* next_group_head = head;
        for (int i = 0; i < 2; i++) {
            if (next_group_head == nullptr) {
                return head;
            }
            next_group_head = next_group_head->next;
        }

        // 反转当前组
        ListNode* new_head = reverse(head, next_group_head);
        // 反转剩余部分
        head->next = swapPairs(next_group_head);

        return new_head;
    }

    ListNode* reverse(ListNode* head, ListNode* next_group_head) {
        ListNode *pre = next_group_head, *node = head, *next = head;
        while (node != next_group_head) {
            next = node->next;
            node->next = pre;
            pre = node;
            node = next;
        }
        return pre;
    }
};
