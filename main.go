package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1 := a[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(a)
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	byteMap := make(map[byte]int, 0)
	for i := 0; i < len(s1); i++ {
		byteMap[s1[i]] = byteMap[s1[i]] + 1
	}

	for i := 0; i <= len(s2)-len(s1); i++ {
		var tmpMap = copyMap(byteMap)
		var fail bool
		for j := i; j < i+len(s1); j++ {
			if _, ok := tmpMap[s2[j]]; !ok {
				fail = true
				break
			}
			count := tmpMap[s2[j]]
			count--
			tmpMap[s2[j]] = count
			if count < 0 {
				fail = true
				break
			}
		}
		if !fail {
			return true
		}
	}
	return false
}

func copyMap(src map[byte]int) map[byte]int {
	dest := make(map[byte]int, len(src))
	for k, v := range src {
		dest[k] = v
	}
	return dest
}

type ListNode struct {
	Val  int
	Next *ListNode
}
