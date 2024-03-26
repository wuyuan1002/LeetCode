package main

// 210. 课程表 II

// 现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
// 给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，
// 表示在选修课程 ai 前 必须 先选修 bi 。
//
// 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
// 返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。
// 如果不可能完成所有课程，返回 一个空数组 。

// findOrder 所有课程构成一个有向无环图
// 1. 深度优先搜索（DFS）
// 2. 广度优先搜索（BFS）
//    若A依赖B，B依赖C，则完成A之前必须先完成B，完成C之前必须完成C,而完成C不需要依赖其它节点，
//    此时认为C的入度为0，将C入队列并将其从图中删除，删除C后可以发现B的入度变为0，此时再将B加
//    入到队列并从图中删除，以此往复不断将入度为0的节点加入到队列直至图中没有节点为止。
//    此方法可以顺便检查图中是否有环，因为若图中有环，则环上节点的入度一定不会变为0。而BFS方法
//    则无法验证图中是否有环，若图中有环那么会造成无限递归或死循环。
func findOrder(numCourses int, prerequisites [][]int) []int {
	result := make([]int, 0, numCourses) // 总结果
	edges := make([][]int, numCourses)   // 存每个节点及依赖该节点的节点
	indeg := make([]int, numCourses)     // 存每个节点及其入度（当前节点依赖的节点数）

	// 遍历课程依赖关系，初始化edges和indeg
	for _, info := range prerequisites {
		// 将依赖当前节点的节点加入到当前节点的依赖队列
		edges[info[1]] = append(edges[info[1]], info[0])
		// 递增依赖当前节点的节点的入度
		indeg[info[0]]++
	}

	// 用于存储当前入度为0的节点
	emptyIn := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			emptyIn = append(emptyIn, i)
		}
	}

	// 不断将入度为0的节点加入到总结果并将依赖该节点的节点入度减一并将新的入度变为0的节点加入到入度为0的队列中
	for len(emptyIn) > 0 {
		// 获取当前入度为0的一个节点，并将其加入到结果队列
		num := emptyIn[0]
		emptyIn = emptyIn[1:]
		result = append(result, num)

		// 遍历依赖当前入度为0的节点的节点，将其入度减一，若其入度变为0则将其加入到入度为0的队列中
		for _, n := range edges[num] {
			indeg[n]--
			if indeg[n] == 0 {
				emptyIn = append(emptyIn, n)
			}
		}
	}

	// 最终若结果列表长度等于总节点数，说明图中没有环，可以正常完成，返回总结果，否则返回空数组
	if len(result) == numCourses {
		return result
	}
	return []int{}
}
