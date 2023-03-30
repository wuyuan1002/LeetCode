
// 141. 环形链表
//
// 给你一个链表的头节点 head ，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。

#include <header.h>

class Solution {
public:
    // 142
    // 1. 遍历所有节点, 使用map记录已经访问过的节点
    // 2. 快慢指针
    // 分别定义 fast 和 slow 指针, 从头结点出发, fast指针每次移动两个节点,
    // slow指针每次移动一个节点, 如果 fast 和 slow指针在途中相遇, 说明这个链表有环
    bool hasCycle(ListNode *head) {
        ListNode *fast = head, *slow = head;
        while (fast != nullptr && fast->next != nullptr) {
            fast = fast->next->next;
            slow = slow->next;

            if (fast == slow) {
                return true;
            }
        }
        return false;
    }
};
