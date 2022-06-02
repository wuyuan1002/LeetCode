
// 23. 合并K个升序链表
//
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

#include "header.h"
#include <vector>

class Solution {
public:
    // leetcode 21
    // 1. 分治法 -- 类似归并排序，先解决前半部分链表合并，后解决后半部分链表合并，再合并两部分已合并好的链表；其实就是分割成两个链表合并问题
    ListNode* mergeKLists(std::vector<ListNode*>& lists)
    {
        if (lists.empty()) {
            return nullptr;
        }

        return merge(lists, 0, lists.size() - 1);
    }

    ListNode* merge(std::vector<ListNode*>& lists, int l, int r)
    {
        if (l >= r) {
            return lists[l];
        }

        int mid = l + (r - l) / 2;
        ListNode* left = merge(lists, l, mid);
        ListNode* right = merge(lists, mid + 1, r);

        return mergeTwoLists(left, right);
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

    // ---------
    // 2. 先合并数组中第1个和第2个链表，再将合并好的链表和第3个合并，再将合并好的链表和第4个合并...直到合并完数组中的所有链表
    ListNode* mergeKLists(std::vector<ListNode*>& lists)
    {
        if (lists.empty()) {
            return nullptr;
        }

        ListNode* merge_list = lists[0];
        for (int i = 1; i < lists.size(); ++i) {
            merge_list = mergeTwoLists(merge_list, lists[i]);
        }
        return merge_list;
    }
};
