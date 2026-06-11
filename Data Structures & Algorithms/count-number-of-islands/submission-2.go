func numIslands(grid [][]byte) int {
	count := 0
	rows, cols := len(grid), len(grid[0])

	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				count += 1
				dfs(grid, r, c, rows, cols)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c, rows, cols int) {
	if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r+1, c, rows, cols)
	dfs(grid, r-1, c, rows, cols)
	dfs(grid, r, c+1, rows, cols)
	dfs(grid, r, c-1, rows, cols)
}
