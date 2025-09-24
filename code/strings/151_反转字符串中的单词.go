package strings

func reverseWords(s string) string {
	var mark int
	var isMark bool = false
	var ans string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if !isMark {
				continue
			}
			isMark = false
			ans = s[mark:i] + " " + ans
		} else {
			if !isMark {
				isMark = true
				mark = i
			}
		}
	}
	if isMark {
		ans = s[mark:] + " " + ans
	}
	ans = ans[0 : len(ans)-1]
	return ans
}
