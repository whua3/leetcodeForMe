package strings

import "strings"

const (
	intMax = 1<<31 - 1
	intMin = -1 << 31
)

func myAtoi(s string) int {
	// 1. 跳过前导空格
	s = strings.TrimSpace(s)

	i, n := 0, len(s)

	if n == 0 {
		return 0
	}

	// 2. 处理正负号
	positive := true

	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		i++
		positive = false
	}

	// 3. 处理数字
	result := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// 检查结果是否溢出
		if result > intMax/10 || (result == intMax/10 && digit > intMax%10) {
			if positive {
				return intMax
			}
			return intMin
		}

		result = result*10 + digit
		i++
	}

	if !positive {
		result = -1 * result
	}
	return result
}

func MyAtoi2(s string) int {
	s = strings.TrimSpace(s)

	if len(s) == 0 {
		return 0
	}

	var positive bool = true

	if s[0] == '-' {
		positive = false
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	var zeroCnt int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeroCnt++
		} else {
			break
		}
	}

	s = s[zeroCnt:]

	var ans int
	for i := 0; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		digit := int(s[i] - '0')
		if ans > intMax/10 || (ans == intMax/10 && digit > intMax%10) {
			// 溢出
			if positive {
				return intMax
			}
			return intMin
		}
		ans = ans*10 + digit
	}

	if positive {
		return ans
	}

	return -1 * ans
}
