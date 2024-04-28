/**
 *
 * 剑指 Offer 68 - I. 二叉搜索树的最近公共祖先
 *
 * 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉搜索树: root =[6,2,8,0,4,7,9,null,null,3,5]
 *
 */
public class offer68_1_二叉搜索树的最近公共祖先 {

}

// 二叉搜索树是排序的，从上到下第一个位于两个值中间的数字就是最近公共祖先
class Solution {

    /**
     * 循环解法
     */
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || p == null || q == null) {
            return null;
        }

        TreeNode min = min(p, q);
        TreeNode max = max(p, q);

        TreeNode node = root;
        while (node != null) {
            if (node.val > max.val) { // 如果当前节点的值比两个都大，则去当前节点的左面找
                node = node.left;
            } else if (node.val < min.val) { // 如果当前节点的值比两个都小，则去当前节点的右面找
                node = node.right;
            } else {
                return node;
            }
        }

        return null;
    }

    /**
     * 求小节点
     */
    private TreeNode min(TreeNode p, TreeNode q) {
        if (p.val < q.val) {
            return p;
        }
        return q;
    }

    /**
     * 求大节点
     */
    private TreeNode max(TreeNode p, TreeNode q) {
        if (p.val > q.val) {
            return p;
        }
        return q;
    }


    /**
     * 递归解法
     */
    public TreeNode lowestCommonAncestor1(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null) {
            return null;
        }

        TreeNode min = min(p, q);
        TreeNode max = max(p, q);

        if (root.val > max.val) { // 如果当前节点的值比两个都大，则去当前节点的左面找
            return lowestCommonAncestor1(root.left, p, q);
        } else if (root.val < min.val) { // 如果当前节点的值比两个都小，则去当前节点的右面找
            return lowestCommonAncestor1(root.right, p, q);
        } else {
            return root;
        }
    }
}
