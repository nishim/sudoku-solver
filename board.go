package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * Cell is Sudoku cell.
 */
type Cell struct {
	Num        int
	candidates map[int]bool
}

/**
 * Board is Sudoku board.
 */
type Board struct {
	Cells [9][9]Cell
}

/**
 * NewBoard generate Board.
 */
func NewBoard(cells string) *Board {
	splitted := strings.Split(cells, ",")
	if len(splitted) != 81 {
		panic("invalid argument")
	}

	b := &Board{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b.Cells[i][j] = newCell(splitted[(i*9)+j])
		}
	}

	return b
}

func newCell(s string) Cell {
	n, err := strconv.Atoi(s)
	if err != nil {
		c := make(map[int]bool)
		for i := 1; i <= 9; i++ {
			c[i] = true
		}
		return Cell{0, c}
	}

	return Cell{n, nil}
}

/**
 * Print board.
 */
func (board *Board) Print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board.Cells[i][j]
			var s string
			if c.Num == 0 {
				s = " "
			} else {
				s = strconv.Itoa(c.Num)
			}
			fmt.Printf("%s | ", s)
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

/**
 * Solved is
 */
func (board *Board) Solved() int {
	n := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board.Cells[i][j].Num != 0 {
				n++
			}
		}
	}

	return n
}

func (board *Board) Solve() {
	board.elim()
}

func (board *Board) emptyCell(i int, j int) bool {
	return board.cellnum(i, j) == 0
}

func (board *Board) cellnum(i int, j int) int {
	return board.Cells[i][j].Num
}

func (board *Board) elim() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !board.emptyCell(i, j) {
				continue
			}

			for y := 0; y < 9; y++ {
				if y == i || board.emptyCell(y, j) {
					continue
				}
				board.Cells[i][j].candidates[board.cellnum(y, j)] = false
			}

			for x := 0; x < 9; x++ {
				if x == j || board.emptyCell(i, x) {
					continue
				}
				board.Cells[i][j].candidates[board.cellnum(i, x)] = false
			}

			for _, y := range r(i) {
				for _, x := range r(j) {
					if y == i && x == j {
						continue
					}
					if board.emptyCell(y, x) {
						continue
					}

					board.Cells[i][j].candidates[board.cellnum(y, x)] = false
				}
			}
		}
	}

	board.update()
}

func (board *Board) update() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if !board.emptyCell(y, x) {
				continue
			}

			n := 0
			for k, v := range board.Cells[y][x].candidates {
				if !v {
					continue
				}
				if n == 0 {
					n = k
				} else {
					// 複数候補あり
					n = 0
					break
				}
			}

			board.Cells[y][x].Num = n
		}
	}
}

func r(i int) []int {
	if i >= 0 && i <= 2 {
		return []int{0, 1, 2}
	} else if i >= 3 && i <= 5 {
		return []int{3, 4, 5}
	} else {
		return []int{6, 7, 8}
	}
}
