/**
 * 剑指 Offer 36. 二叉搜索树与双向链表
 *
 * 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。
 * 要求不能创建任何新的节点，只能调整树中节点指针的指向。
 *
 * 左指针表示双链表向前指，右指针表示双链表向后指
 */
public class offer36 {
    public static void main(String[] args) {

    }

    private Node pre; // 记录遍历的上一个节点
    private Node head; // 记录双向链表的头节点

    public Node treeToDoublyList(Node root) {
        if (root == null) {
            return null;
        }
        recur(root);
        head.left = pre;
        pre.right = head;
        return head;
    }

    public void recur(Node node) {
        if (node == null) {
            return;
        }
        recur(node.left);
        if (pre == null) {
            head = node;
        } else {
            pre.right = node;
            node.left = pre;
        }
        pre = node;
        recur(node.right);
    }
}

class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
}
