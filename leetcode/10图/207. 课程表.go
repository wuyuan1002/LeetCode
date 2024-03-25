package main

// 207. 课程表

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// canFinish 所有课程构成一个有向无环图
// 1. 广度优先搜索（BFS）
// 2. 深度优先搜索（DFS）
//    若A依赖B，B依赖C，则完成A之前必须先完成B，完成C之前必须完成C,而完成C不需要依赖其它节点，
//    此时认为C的入度为0，将C入队列并将其从图中删除，删除C后可以发现B的入度变为0，此时再将B加
//    入到队列并从图中删除，以此往复不断将入度为0的节点加入到队列直至图中没有节点为止。
//    此方法可以顺便检查图中是否有环，因为若图中有环，则环上节点的入度一定不会变为0。而BFS方法
//    则无法验证图中是否有环，若图中有环那么会造成无限递归或死循环。
func canFinish(numCourses int, prerequisites [][]int) bool {
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

	// 最终若结果列表长度等于总节点数，说明图中没有环，可以正常完成，否则说明无法完成
	return len(result) == numCourses
}
