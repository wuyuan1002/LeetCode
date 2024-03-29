// 25. K个一组翻转链表

// 给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
// k是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

#include "header.h"

class Solution {
public:
    // 24. 两两交换链表中的节点
    // 递归
    ListNode* reverseKGroup(ListNode* head, int k) {
        ListNode* next_group_head = head;
        for (int i = 0; i < k; i++) {
            if (next_group_head == nullptr) {
                return head;
            }
            next_group_head = next_group_head->next;
        }

        // 反转当前组
        ListNode* new_head = reverse(head, next_group_head);
        // 反转剩余部分
        head->next = reverseKGroup(next_group_head, k);

        return new_head;
    }

    // 反转链表
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
