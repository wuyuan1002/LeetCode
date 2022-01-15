package main

// 986. 区间列表的交集

// 给定两个由一些 闭区间 组成的列表，firstList 和 secondList ，
// 其中 firstList[i] = [starti, endi] 而 secondList[j] = [startj, endj] 。
// 每个区间列表都是成对 不相交 的，并且 已经排序 。

// 返回这 两个区间列表的交集 。
//
// 形式上，闭区间 [a, b]（其中 a <= b）表示实数 x 的集合，而 a <= x <= b 。
// 两个闭区间的 交集 是一组实数，要么为空集，要么为闭区间。例如，[1, 3] 和 [2, 4] 的交集为 [2, 3] 。

// func main() {
// 	f := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
// 	s := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
// 	fmt.Println(intervalIntersection(f, s))
// }

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if firstList == nil || secondList == nil {
		return nil
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	result := make([][]int, 0)
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {

		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])

		if lo <= hi {
			result = append(result, []int{lo, hi})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}
