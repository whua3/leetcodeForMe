package other

import (
	"math/rand"
	"time"
)

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}
func rand10() int {
	for {
		// 第一次调用Rand7()，生成1-7，减1后变为0-6
		// 乘以7得到0,7,14,21,28,35,42
		a := rand7() - 1

		// 第二次调用Rand7()，生成1-7，减1后变为0-6
		b := rand7() - 1

		// 组合得到0-48的均匀分布随机数
		num := a*7 + b

		// 如果num在0-39范围内，直接映射到1-10
		if num < 40 {
			return num%10 + 1
		}

		// 否则重新生成，这样保证了均匀性
	}
}
