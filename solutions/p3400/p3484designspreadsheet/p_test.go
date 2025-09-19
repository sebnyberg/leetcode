package p3484designspreadsheet

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	cells [][27]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		cells: make([][27]int, rows),
	}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	i, j := idx(cell)
	this.cells[i][j] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	i, j := idx(cell)
	this.cells[i][j] = 0
}

func (this *Spreadsheet) getCellValue(cell string) int {
	i, j := idx(cell)
	return this.cells[i][j]
}

func (this *Spreadsheet) GetValue(formula string) int {
	exprs := strings.Split(formula[1:], "+")
	a := exprs[0]
	b := exprs[1]
	var res int
	if isnum(a) {
		res += mustint(a)
	} else {
		res += this.getCellValue(a)
	}
	if isnum(b) {
		res += mustint(b)
	} else {
		res += this.getCellValue(b)
	}
	return res
}

func idx(cellName string) (int, int) {
	cellIdx := int(cellName[0] - 'A')
	rowIdx, _ := strconv.Atoi(cellName[1:])
	return rowIdx - 1, cellIdx
}

func mustint(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}

func isnum(x string) bool {
	return x[0] < 'A' || x[0] > 'Z'
}
