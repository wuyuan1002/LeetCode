/**
 * 剑指 Offer 68 - II. 二叉树的最近公共祖先
 *
 * 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
 *
 * 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
 *
 * 例如，给定如下二叉树: root =[3,5,1,6,2,0,8,null,null,7,4]
 */
public class offer68_2_二叉树的最近公共祖先 {

}

class Solution1 {

    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || p == root || q == root) {
            return root;
        }

        // 后序遍历二叉树，从底向上寻找第一个p和q在它两侧的root节点

        TreeNode left = lowestCommonAncestor(root.left, p, q); // 从左子树中寻找p或q
        TreeNode right = lowestCommonAncestor(root.right, p, q); // 从右子树中寻找p或q

        // 若p和q分别在root的两侧，则root为最近公共祖先
        if (left == null) {
            return right;
        } else if (right == null) {
            return left;
        } else {
            return root;
        }
    }
}