package main

// 1. 两数之和

func main() {

}

// 可以在查找时将数字存入map中，这样，下次使用时可以直接获取
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	hash := make(map[int]int)
	for i, n := range nums {
		if index, ok := hash[target-n]; ok {
			return []int{i, index}
		}
		hash[n] = i
	}
	return nil
}
