
// 2. 两数相加
//
// 给你两个 非空 的链表，表示两个非负的整数。
// 它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

#include "header.h"

class Solution {
public:
    // 由于两个链表是是逆序存储数字的，只需从头遍历两个链表相加即可
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *res_head = new ListNode(), *res = res_head;  // 总结果
        int carry = 0;                                         // 进位
        ListNode *node1 = l1, *node2 = l2;                     // 正在参与计算的数字

        while (node1 != nullptr || node2 != nullptr || carry != 0) {
            int num1 = node1 != nullptr ? node1->val : 0;
            int num2 = node2 != nullptr ? node2->val : 0;

            int sum = num1 + num2 + carry;       // 求和
            res->next = new ListNode(sum % 10);  //当前位得数
            res = res->next;
            carry = sum / 10;  //进位

            // 向后遍历链表
            node1 = node1 != nullptr ? node1->next : nullptr;
            node2 = node2 != nullptr ? node2->next : nullptr;
        }

        return res_head->next;
    }
};
