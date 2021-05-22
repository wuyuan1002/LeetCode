
/**
 * 剑指 Offer 37. 序列化二叉树
 *
 * 请实现两个函数，分别用来序列化和反序列化二叉树。
 */
public class offer37_序列化二叉树 {
    private int index = 0;

    // 编码
    public String serialize(TreeNode root) {
        if (root == null) {
            return "null";
        }
        StringBuilder builder = new StringBuilder();
        serialize1(root, builder);
        return builder.substring(0, builder.length() - 1);
    }

    private void serialize1(TreeNode node, StringBuilder builder) {
        if (node == null) {
            builder.append("null,");
            return;
        }
        builder.append(node.val).append(",");
        serialize1(node.left, builder);
        serialize1(node.right, builder);
    }

    // 解码
    public TreeNode deserialize(String data) {
        if (data == null || "".equals(data)) {
            return null;
        }
        String[] arr = data.split(",");
        return deserialize1(arr);
    }

    private TreeNode deserialize1(String[] arr) {
        if ("null".equals(arr[index])) {
            index++;
            return null;
        }
        TreeNode node = new TreeNode(Integer.parseInt(arr[index++]));
        node.left = deserialize1(arr);
        node.right = deserialize1(arr);
        return node;
    }
}

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}