package p0304rangesumquery2dimmutable

type NumMatrix struct {
	n, m   int
	bit    []int
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	nm := NumMatrix{n, m, make([]int, n*m+1), matrix}
	for i := range matrix {
		copy(nm.bit[i*n+1:], matrix[i])
	}
	for i := 1; i < n*m; i++ {
		parent := i + (i & -i)
		if parent <= n*m {
			nm.bit[parent] += nm.bit[i]
		}
	}
	return nm
}

func (this *NumMatrix) Update(row int, col int, val int) {
	i := row*this.n + col + 1
	d := this.matrix[row][col] - val
	this.matrix[row][col] = val
	for i < len(this.bit) {
		this.bit[i] -= d
		i += i & -i
	}
}

func (this *NumMatrix) Sum(i int) int {
	var res int
	for i > 0 {
		res += this.bit[i]
		i -= i & -i
	}
	return res
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var res int
	for i := row1; i <= row2; i++ {
		start := i*this.n + col1
		end := i*this.n + col2 + 1
		res += this.Sum(end) - this.Sum(start)
	}
	return res
}
