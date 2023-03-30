
// 142. 环形链表 II
//
// 给定一个链表的头节点 head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

#include <header.h>

class Solution {
public:
    // 141
    // 使用快慢指针确认链表有环后, 两指针一个从head开始一个从相遇位置开始,
    // 继续向前移动, 直到相遇时相遇点就是入环的起点
    ListNode* detectCycle(ListNode* head) {
        ListNode *fast = head, *slow = head;
        while (fast != nullptr && fast->next != nullptr) {
            fast = fast->next->next;
            slow = slow->next;

            if (fast == slow) {
                fast = head;
                while (fast != slow) {
                    fast = fast->next;
                    slow = slow->next;
                }
                return slow;
            }
        }
        return nullptr;
    }
};
