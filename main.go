package main

import (
	"fmt"
	"sort"
)

func main() {
	board := NewBoard(6)
	run(board)
	fmt.Println(board)
}

func run(board *Board) *Board {
	if board.Queens() == board.n {
		return board
	}
	for _, coordinate := range board.AvailableFields() {
		board.PlaceQueen(coordinate)
		next := run(board)
		if next != nil {
			return next
		}
		board.RemoveQueen(coordinate)
	}
	return nil
}

const (
	Empty Field = iota
	Queen
	Occupied
)

type Field int

func (f Field) String() string {
	switch f {
	case Empty:
		return "."
	case Queen:
		return "Q"
	case Occupied:
		return "X"
	default:
		return "?"
	}
}

type Board struct {
	Fields []Field
	n      int
}

func NewBoard(n int) *Board {
	if n < 4 {
		panic("Boardgröße ist mind. 4x4")
	}
	return &Board{
		Fields: make([]Field, n*n),
		n:      n,
	}
}

func (b Board) String() string {
	sb := ""
	for i, v := range b.Fields {
		if i%b.n == 0 {
			sb += "\n"
		}
		sb += fmt.Sprintf("%v", v)
	}
	return sb
}

func (b *Board) PlaceQueen(n int) {
	b.Fields[n] = Queen
	for _, v := range b.GetOccupiedFields(n) {
		if v == n {
			continue
		}
		b.Fields[v] = Occupied
	}
}

func (b *Board) RemoveQueen(n int) {
	b.Fields[n] = Empty
	for _, v := range b.GetOccupiedFields(n) {
		b.Fields[v] = Empty
	}
}

func (b Board) Queens() int {
	count := 0
	for _, v := range b.Fields {
		if v == Queen {
			count++
		}
	}
	return count
}

func (b Board) AvailableFields() []int {
	res := make([]int, 0)
	for i, v := range b.Fields {
		if v == Empty {
			res = append(res, i)
		}
	}
	return res
}

func (b Board) GetOccupiedFields(n int) []int {
	line := n / b.n
	column := n % b.n
	res := make([]int, 0)
	res = append(res, b.getLineNumbers(line)...)
	res = append(res, b.getColumnNumbers(column)...)
	res = append(res, b.getDiagonals(line, column)...)
	return sortAndRemoveDuplicates(res)
}

func (b Board) getLineNumbers(line int) []int {
	res := make([]int, 0)
	for i := line * b.n; i < line*b.n+b.n; i++ {
		res = append(res, i)
	}
	return sortAndRemoveDuplicates(res)
}

func (b Board) getColumnNumbers(col int) []int {
	res := make([]int, 0)
	for i := col; i < b.n*b.n; i += b.n {
		res = append(res, i)
	}
	return sortAndRemoveDuplicates(res)
}

func (b Board) getDiagonals(line, col int) []int {
	res := make([]int, 0)
	res = append(res, getTopRight(line, col, b.n)...)
	res = append(res, getTopLeft(line, col, b.n)...)
	res = append(res, getBottomRight(line, col, b.n)...)
	res = append(res, getBottomLeft(line, col, b.n)...)
	return sortAndRemoveDuplicates(res)
}

func getTopRight(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line-1, col-1
	for {
		if i < 0 || j < 0 {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i--
		j--
	}
	return res
}

func getTopLeft(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line-1, col+1
	for {
		if i < 0 || j >= limit {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i--
		j++
	}
	return res
}

func getBottomRight(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line+1, col-1
	for {
		if i >= limit || j < 0 {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i++
		j--
	}
	return res
}

func getBottomLeft(line, col, limit int) []int {
	res := make([]int, 0)
	i, j := line+1, col+1
	for {
		if i >= limit || j >= limit {
			break
		}
		currentField := i*limit + j
		res = append(res, currentField)
		i++
		j++
	}
	return res
}

func sortAndRemoveDuplicates(nums []int) []int {
	seen := make(map[int]bool, 0)
	res := make([]int, 0)

	for v := range nums {
		if seen[nums[v]] == false {
			seen[nums[v]] = true
			res = append(res, nums[v])
		}
	}
	sort.Ints(res)
	return res
}
