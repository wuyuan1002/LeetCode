
// 141. 环形链表
//
// 给你一个链表的头节点 head ，判断链表中是否有环。

#include <header.h>

class Solution {
public:
    ListNode* hasCycle(ListNode* head) {
        ListNode *fast = head, *slow = head;
        while (fast != nullptr && fast->next != nullptr) {
            fast = fast->next->next;
            slow = slow->next;

            if (fast == slow) {
                return fast;
            }
        }
        return nullptr;
    }
};