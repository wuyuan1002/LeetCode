import java.util.ArrayList;

/**
 * 查找树中的公共祖先 -- 书上的最后一题
 */
public class offer68_3_普通树的最近公共祖先 {
}

/**
 * 有多个孩子的树
 */
class TreeNodeC {
    int val;
    ArrayList<TreeNodeC> childrens = new ArrayList<>();

    TreeNodeC(int x) {
        val = x;
    }
}

// 分别查找并保存从root到两个节点的路径，之后遍历路径相当于求两个链表的第一个公共节点
class Solution2 {

    public TreeNodeC lowestCommonAncestor(TreeNodeC root, TreeNodeC p, TreeNodeC q) {
        if (root == null || p == null || q == null) {
            return null;
        }

        // 寻找p和q在书中的路径, 然后查找两个链表的第一个公共节点
        ArrayList<TreeNodeC> pathOfP = new ArrayList<>();
        ArrayList<TreeNodeC> pathOfQ = new ArrayList<>();

        if (findPath(root, p, pathOfP) && findPath(root, p, pathOfQ)) {
            // 用来记录上一个节点
            TreeNodeC last = null;
            for (int i = 0; i < pathOfP.size() && i < pathOfQ.size(); i++) {
                if (pathOfP.get(i) == pathOfP.get(i)) {
                    last = pathOfP.get(i);
                    continue;
                }
                break;
            }

            return last;
        }

        return null;
    }


    /**
     * 寻找从root到node的路径 -- 先序遍历root
     */
    private boolean findPath(TreeNodeC root, TreeNodeC node, ArrayList<TreeNodeC> path) {

        if (root == null) {
            return false;
        }

        // 是否已找到
        boolean found = false;

        // 把当前节点放进路径中
        path.add(root);
        if (root == node) {
            return true;
        }

        for (TreeNodeC n : root.childrens) {
            found = findPath(n, node, path);
            if (found) {
                break;
            }
        }

        // 若没找到，把当前节点从路径中移除
        if (!found) {
            path.remove(root);
        }

        return found;
    }
}