package strings

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	length := len(s)
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				for l := 1; l < 4; l++ {
					if i+j+k+l == length {
						s1 := s[:i]
						s2 := s[i : i+j]
						s3 := s[i+j : i+j+k]
						s4 := s[i+j+k : i+j+k+l]
						if check(s1) && check(s2) && check(s3) && check(s4) {
							ans = append(ans, strings.Join([]string{s1, s2, s3, s4}, "."))
						}
					}
				}
			}
		}
	}
	return ans
}

func check(numStr string) bool {
	num, _ := strconv.Atoi(numStr)

	if num <= 255 {
		if numStr[0] != '0' || (numStr[0] == '0' && len(numStr) == 1) {
			return true
		}
	}

	return false
}
