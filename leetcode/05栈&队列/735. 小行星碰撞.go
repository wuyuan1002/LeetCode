package main

// 735. 小行星碰撞

// 给定一个整数数组 asteroids，表示在同一行的小行星。
//
// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
// 找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。

// asteroidCollision .
// 由于碰撞抵消总是从相邻行星之间发生，我们可以使用「栈」来模拟该过程，
// 从前往后处理所有的asteroids[i]，使用栈存储当前未被抵消的行星，
//
// 遇到向右移动的行星直接入栈，遇到向左移动的行星则将其挨个与栈顶元素进行碰撞操作，
// 向左移动的行星只有两种结局，要么栈内向右移动的行星比它大被爆炸，
// 要么它比栈内向右移动的行星都大，将栈内向右移动的行星都被它爆炸了，它左侧没有能使其爆炸的向右移动的行星，它永远不会被爆炸
func asteroidCollision(asteroids []int) []int {
	// 构造一个栈，栈内存当前没有被爆炸的行星
	stack := make([]int, 0)

	for _, aster := range asteroids {
		alive := true // 用来表示当前行星有没有被栈内行星爆炸掉

		// 若当前行星向左且栈顶行星当前向右，则将两行星进行碰撞
		for aster < 0 && alive && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = abs(aster) > abs(stack[len(stack)-1]) // 当前行星是否会被爆炸
			if abs(stack[len(stack)-1]) <= abs(aster) {   // 栈顶行星是否会被爆炸
				stack = stack[:len(stack)-1]
			}
		}

		// 若最终当前行星还存在，则将其加入栈中
		if alive {
			stack = append(stack, aster)
		}
	}

	// 栈内行星就是最终会剩余的行星
	return stack
}

// 求绝对值
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
