package strings

import "fmt"

// 官方题解方法二：做乘法
// 长度为m和n的两个数相乘，长度为m+n或者m+n-1
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := res[i+j+1] + n1*n2
			res[i+j+1] = sum % 10
			res[i+j] = res[i+j] + sum/10
		}
	}

	var ans string

	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}
		ans += fmt.Sprintf("%d", res[i])
	}
	return ans
}
