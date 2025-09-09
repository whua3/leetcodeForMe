package strings

func LengthOfLongestSubstring(s string) int {
	// 记录字符是否出现过
	byteMap := make(map[byte]int, 0)
	length := len(s)
	p, ans := 0, 0

	for i := 0; i < length; i++ {
		if i != 0 {
			// 移除第一个字符
			delete(byteMap, s[i-1])
		}
		for p < length && byteMap[s[p]] == 0 {
			byteMap[s[p]]++
			p++
		}
		if p-i > ans {
			ans = p - i
		}
	}
	return ans
}
