package strings

import "strings"

// 官方题解方法一： 栈
func simplifyPath(path string) string {
	stack := make([]string, 0)
	var stackLen int
	arr := strings.Split(path, "/")
	for i := 0; i < len(arr); i++ {
		if arr[i] == "" || arr[i] == "." {
			continue
		}

		if arr[i] == ".." {
			if stackLen != 0 {
				stack = stack[:stackLen-1]
				stackLen -= 1
			}
		} else {
			stack = append(stack, arr[i])
			stackLen += 1
		}
	}

	return "/" + strings.Join(stack, "/")
}
