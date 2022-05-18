// 19. 删除链表的倒数第 N 个结点
//
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

#include "header.h"

class Solution {
public:
    // 双指针
    ListNode* removeNthFromEnd(ListNode* head, int n)
    {
        ListNode *l = head, *r = head;
        for (int i = 0; i < n; ++i) {
            if (r == nullptr) {
                return head;
            }
            r = r->next;
        }

        while (r != nullptr && r->next != nullptr) {
            r = r->next;
            l = l->next;
        }

        if (r == nullptr) {
            return head->next;
        } else {
            l->next = l->next->next;
            return head;
        }
    }
};