package datastructure

func sunday(haystack, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if needle == "" {
		return 0
	}
	//计算偏移表
	shiftMap := map[byte]int{}
	for i := len(needle) - 1; i > -1; i-- {
		if _, ok := shiftMap[needle[i]]; !ok {
			shiftMap[needle[i]] = len(needle) - i
		}
	}
	ans := 0
	for ans+len(needle) <= len(haystack) {
		// 判断是否匹配
		if haystack[ans:ans+len(needle)] == needle {
			return ans
		}
		// 不匹配情况下，根据下一个字符的偏移，移动ans
		if step, ok := shiftMap[haystack[ans+len(needle)]]; ok {
			ans += step
		} else {
			ans += len(needle) + 1
		}
	}
	return -1
}
