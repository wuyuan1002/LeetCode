// 19. 删除链表的倒数第 N 个结点
//
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
// 进阶：你能尝试使用一趟扫描实现吗？

#include "header.h"

class Solution {
public:
    // 双指针
    // 先让右指针向后移动n, 之后左右指针同时向后移动, 当右指针到末尾时, 要删除的元素正好时左指针的下一个节点
    ListNode* removeNthFromEnd(ListNode* head, int n) {
        ListNode *l = head, *r = head;

        // 右指针先向后移动n
        for (int i = 0; i < n; i++) {
            if (r == nullptr) {
                return head;
            }
            r = r->next;
        }

        // 双指针同时向后移动
        // 这里如果正好删除的是head, 则 r == nullptr, 不会进入这个循环
        while (r != nullptr && r->next != nullptr) {
            l = l->next;
            r = r->next;
        }

        if (r == nullptr) {
            // 要删除的元素正好的head
            return head->next;
        } else {
            l->next = l->next->next;
            return head;
        }
    }
};
