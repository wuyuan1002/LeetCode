package main

// 28. 实现 strStr()

func main() {

}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	index := -1
	for i := 0; i < len(haystack); i++ {
		k, j := i, 0
		index = i
		for k < len(haystack) && j < len(needle) && haystack[k] == needle[j] {
			k++
			j++
		}

		if j == len(needle) {
			break
		}
		index = -1
	}

	return index
}
