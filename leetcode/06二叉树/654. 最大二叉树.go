package main

// 654. 最大二叉树

// 给定一个不重复的整数数组 nums 。 最大二叉树 可以用下面的算法从 nums 递归地构建:
//
// 创建一个根节点，其值为 nums 中的最大值。
// 递归地在最大值 左边 的 子数组前缀上 构建左子树。
// 递归地在最大值 右边 的 子数组后缀上 构建右子树。
// 返回 nums 构建的 最大二叉树 。

// constructMaximumBinaryTree .
// 构造树一般采用的是前序遍历，因为先构造中间节点，然后递归构造左子树和右子树
// 每次都遍历nums获取最大值作为当前节点的值，然后左半部分为左子树右半部分为右子树，之后递归构建左右子树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	// 获取数组最大值的下标
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	// 构建当前节点并递归构建其左右子树
	node := &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}

	return node
}
