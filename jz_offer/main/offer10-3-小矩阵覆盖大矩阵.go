package main

// func main() {
// 	fmt.Println(coverWays())
// }

func coverWays() int {
	count := 0
	cover1 := 1 // 前一个
	cover2 := 1 // 前两个
	for i := 2; i <= 8; i++ {
		count = cover1 + cover2
		cover2 = cover1
		cover1 = count
	}
	return count
}
