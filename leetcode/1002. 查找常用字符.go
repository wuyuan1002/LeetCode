package main

// 1002. 查找常用字符

// 给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。
// 例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。
//
// 你可以按任意顺序返回答案。

// func main() {

// }

// 1. 调一个字符串，遍历每个字符，再在剩下的两个字符串中查看是否存在
// 2. 使用map或数组记录每个字符出现的次数
func commonChars(words []string) []string {
	if words == nil || len(words) == 0 {
		return nil
	}

	hash := make([]int, 26)      // 存每个字符出现的最少次数
	for _, c := range words[0] { // 先使用第一个字符串初始化hash
		hash[c-'a']++
	}

	for i := 1; i < len(words); i++ {
		// 将一个字符串的字符出现次数先存入临时hash表中
		tmphash := make([]int, 26)
		for _, c := range words[i] {
			tmphash[c-'a']++
		}

		// 将hash和tmphash中的字符出现次数最少的更新在hash中
		for j := 0; j < 26; j++ {
			hash[j] = min(hash[j], tmphash[j])
		}
	}

	// 最终hash中剩余的字符就是所有字符串中共同出现的字符的出现次数
	result := make([]string, 0)
	for i, nums := range hash {
		for ; nums > 0; nums-- {
			result = append(result, string(byte(i+'a')))
		}
	}

	return result
}
