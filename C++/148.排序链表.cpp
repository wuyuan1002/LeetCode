
// 148. 排序链表
//
// 给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
//
// 进阶：
// 你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

#include "header.h"

class Solution {
public:
    // 类似Hot100 23, offer 51
    // 1.插入排序 O(n2)
    // 2.归并排序 O(nlogn) -- 自顶向下，或自底向上
    ListNode* sortList(ListNode* head)
    {
        return sort(head, nullptr);
    }

    ListNode* sort(ListNode* head, ListNode* tail)
    {
        // [head, tail)
        if (head == nullptr) {
            return nullptr;
        }

        if (head->next == tail) {
            head->next = nullptr;
            return head;
        }

        // 寻找中点
        ListNode *slow = head, *fast = head;
        while (fast != tail && fast->next != tail) {
            slow = slow->next;
            fast = fast->next->next;
        }
        ListNode* mid = slow;

        // 合并两个链表
        return mergeTwoLists(sort(head, mid), sort(mid, tail));
    }

    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2)
    {
        if (list1 == nullptr) {
            return list2;
        } else if (list2 == nullptr) {
            return list1;
        } else if (list1->val < list2->val) {
            list1->next = mergeTwoLists(list1->next, list2);
            return list1;
        } else {
            list2->next = mergeTwoLists(list1, list2->next);
            return list2;
        }
    }
};
