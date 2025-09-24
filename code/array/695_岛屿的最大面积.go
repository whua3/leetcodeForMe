package array

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			value := DFS(grid, i, j)
			if value > ans {
				ans = value
			}
		}
	}
	return ans
}

func DFS(grid [][]int, cur_i, cur_j int) int {
	if cur_i < 0 || cur_j < 0 || cur_i == len(grid) || cur_j == len(grid[0]) || grid[cur_i][cur_j] == 0 {
		return 0
	}

	ans := 1
	grid[cur_i][cur_j] = 0
	di := []int{-1, 1, 0, 0}
	dj := []int{0, 0, -1, 1}

	for i := 0; i < 4; i++ {
		next_i, next_j := cur_i+di[i], cur_j+dj[i]
		ans += DFS(grid, next_i, next_j)
	}
	return ans
}
