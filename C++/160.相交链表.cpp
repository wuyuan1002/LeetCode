
// 160. 相交链表

// 给你两个单链表的头节点 headA 和 headB ，
// 请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

#include <header.h>

class Solution {
public:
    ListNode* getIntersectionNode(ListNode* headA, ListNode* headB) {
        ListNode *p = headA, *q = headB;
        while (p != q) {
            if (p != nullptr) {
                p = p->next;
            } else {
                p = headB;
            }

            if (q != nullptr) {
                q = q->next;
            } else {
                q = headA;
            }
        }
        return p;
    }
};