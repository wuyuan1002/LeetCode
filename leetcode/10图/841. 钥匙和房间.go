package main

// 841. 钥匙和房间

// 有 n 个房间，房间按从 0 到 n - 1 编号。最初，除 0 号房间外的其余所有房间都被锁住。你的目标是进入所有的房间。然而，你不能在没有获得钥匙的时候进入锁住的房间。
//
// 当你进入一个房间，你可能会在里面找到一套不同的钥匙，每把钥匙上都有对应的房间号，即表示钥匙可以打开的房间。你可以拿上所有钥匙去解锁其他房间。
//
// 给你一个数组 rooms 其中 rooms[i] 是你进入 i 号房间可以获得的钥匙集合。如果能进入 所有 房间返回 true，否则返回 false。

// canVisitAllRooms .
// 当x号房间中有y号房间的钥匙时，表示可以从x号房间去往y号房间，
// 如果我们将这n个房间看成有向图中的n个节点，那么上述关系就可以看作是图中的x号点到y号点的一条有向边，
// 这样一来，问题就变成了给定一张有向图，询问从0号节点出发是否能够到达所有的节点。
//
// 使用深度优先搜索的方式遍历整张图，统计可以到达的节点个数，并利用数组visited标记当前节点是否访问过，防止节点被重复访问
func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	// 深度优先遍历，统计从0号节点出发能够到达的所有节点
	dfsCanVisitAllRooms(rooms, visited, 0)

	// 查看所有节点，若存在节点没有访问到则返回false，否则返回true
	for _, isVisited := range visited {
		if !isVisited {
			return false
		}
	}

	return true
}

// dfsCanVisitAllRooms .
// 深度优先遍历，从节点room出发访问所有能到达的所有节点
func dfsCanVisitAllRooms(rooms [][]int, visited []bool, room int) {
	// 若当前节点已被访问则直接返回
	if visited[room] {
		return
	}

	// 标记当前节点已访问
	visited[room] = true

	// 深度优先遍历，访问当前节点能访问到的其它节点
	for _, child := range rooms[room] {
		dfsCanVisitAllRooms(rooms, visited, child)
	}
}
