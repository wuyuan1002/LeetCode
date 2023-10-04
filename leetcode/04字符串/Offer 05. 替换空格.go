package main

// Offer 05. 替换空格

// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

// replaceSpace .
// 统计s中空格的个数，扩充字符串预留替换所需的空间，使用两个指针分别从尾部遍历字符串和数组，进行字符复制或空格替换
func replaceSpace(s string) string {
	// 统计空格个数
	spaceNum := 0
	for _, c := range s {
		if c == ' ' {
			spaceNum++
		}
	}
	if spaceNum == 0 {
		return s
	}

	// 扩充字符串，预留用于空格替换的空间 -- 长度正好等于替换完后的字符串长度
	sb := append([]byte(s), make([]byte, spaceNum*2)...)

	// 双指针遍历 -- 分别指向字符串最后一个字母和数组末尾
	for i, j := len(s)-1, len(sb)-1; i >= 0; i, j = i-1, j-1 {
		if sb[i] != ' ' {
			sb[j] = sb[i]
		} else {
			sb[j] = '0'
			j--
			sb[j] = '2'
			j--
			sb[j] = '%'
		}
	}
	return string(sb)
}
