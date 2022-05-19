
// 142. 环形链表 II
//
// 给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

#include <header.h>

class Solution {
public:
    ListNode* detectCycle(ListNode* head)
    {
        ListNode *fast = head, *slow = head;
        while (fast != nullptr && fast->next != nullptr) {
            fast = fast->next->next;
            slow = slow->next;

            if (fast == slow) {
                slow = head;
                while (slow != fast) {
                    slow = slow->next;
                    fast = fast->next;
                }
                return slow;
            }
        }
        return nullptr;
    }
};