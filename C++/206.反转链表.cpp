// 206. 反转链表

// 给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

#include "header.h"

class Solution {
public:
    // 三指针
    ListNode* reverseList(ListNode* head)
    {
        ListNode *pre = nullptr, *node = head, *next = head;
        while (node != nullptr) {
            next = node->next;
            node->next = pre;
            pre = node;
            node = next;
        }
        return pre;
    }
};