package array

// 官方题解方法一：深度优先搜索

// 用数组保存city有没有去过，实现一下dfs，如果联通没有去过的就去一下，把visited设置为true
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected)) // 城市数

	var dfs func(int)
	dfs = func(from int) {
		visited[from] = true
		for to, connected := range isConnected[from] {
			if connected == 1 && !visited[to] {
				// 如果连同，并且没有去过
				dfs(to)
			}
		}
	}

	var provinceNum int
	for city, isVisited := range visited {
		if !isVisited {
			provinceNum++
			dfs(city)
		}
	}
	return provinceNum
}
