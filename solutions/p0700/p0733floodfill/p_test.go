package p0733floodfill

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	m := len(image)
	n := len(image[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	originalColor := image[sr][sc]
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	seen[sr][sc] = true
	image[sr][sc] = newColor
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	curr := [][]int{{sr, sc}}
	next := [][]int{}
	for len(curr) > 0 {
		next = next[:0]
		for _, pos := range curr {
			for _, dir := range dirs {
				r := pos[0] + dir[0]
				c := pos[1] + dir[1]
				if !ok(r, c) || seen[r][c] || image[r][c] != originalColor {
					continue
				}
				image[r][c] = newColor
				seen[r][c] = true
				next = append(next, []int{r, c})
			}
		}
		curr, next = next, curr
	}
	return image
}
