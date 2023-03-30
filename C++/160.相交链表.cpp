
// 160. 相交链表

// 给你两个单链表的头节点 headA 和 headB ，
// 请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

#include <header.h>

class Solution {
public:
    // offer 52
    // 1.把两个链表分别入栈, 然后同时弹出元素, 直到第一个不相等的元素
    // 2.将第一个链表全部入map, 之后从头到尾遍历第二个链表, 判断map中是否存在节点
    // 3.双指针
    // 根据题目意思, 如果两个链表相交, 那么相交点之后的长度是相同的
    // 我们需要做的事情是, 让两个链表从同距离末尾同等距离的位置开始遍历。
    // 这个位置只能是较短链表的头结点位置。为此, 我们必须消除两个链表的长度差
    // 指针 pA 指向 A 链表, 指针 pB 指向 B 链表, 依次往后遍历
    // 如果 pA 到了末尾, 则 pA = headB 继续遍历
    // 如果 pB 到了末尾, 则 pB = headA 继续遍历
    // 比较长的链表指针指向较短链表head时, 长度差就消除了
    // 如此, 只需要将最短链表遍历两次即可找到位置
    // 遍历结束pA == pB时, 若有交点, 则指针同时指向交点, 若没有交点, 则两指针分别同时指向两链表末尾null
    ListNode* getIntersectionNode(ListNode* headA, ListNode* headB) {
        ListNode *p = headA, *q = headB;
        while (p != q) {
            p = p != nullptr ? p->next : headB;
            q = q != nullptr ? q->next : headA;
        }
        return p;
    }
};
