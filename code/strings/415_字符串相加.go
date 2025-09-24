package strings

import "fmt"

func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""

	i, j := len(num1)-1, len(num2)-1

	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = fmt.Sprintf("%d%s", result%10, ans)
		add = result / 10
	}
	return ans

}
