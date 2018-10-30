package sudoku

import "fmt"

type cell struct {
	num        int
	candidates map[int]bool
}

var b [9][9]cell

func main() {
	board := [][]int{
		{0, 0, 2, 0, 0, 3, 5, 9, 1},
		{5, 1, 0, 0, 0, 7, 2, 3, 0},
		{9, 0, 0, 0, 5, 1, 0, 7, 0},
		{6, 9, 0, 5, 1, 0, 0, 0, 3},
		{1, 0, 3, 0, 0, 0, 4, 0, 2},
		{0, 0, 0, 0, 6, 4, 9, 0, 5},
		{0, 6, 0, 7, 0, 0, 0, 2, 0},
		{2, 0, 0, 8, 0, 0, 3, 5, 0},
		{0, 8, 9, 1, 0, 0, 6, 0, 0},
		// -------------------------------------
		// | 7 | 4 | 2 | 6 | 8 | 3 | 5 | 9 | 1 |
		// -------------------------------------
		// | 5 | 1 | 8 | 4 | 9 | 7 | 2 | 3 | 6 |
		// -------------------------------------
		// | 9 | 3 | 6 | 2 | 5 | 1 | 8 | 7 | 4 |
		// -------------------------------------
		// | 6 | 9 | 4 | 5 | 1 | 2 | 7 | 8 | 3 |
		// -------------------------------------
		// | 1 | 5 | 3 | 9 | 7 | 8 | 4 | 6 | 2 |
		// -------------------------------------
		// | 8 | 2 | 7 | 3 | 6 | 4 | 9 | 1 | 5 |
		// -------------------------------------
		// | 4 | 6 | 5 | 7 | 3 | 9 | 1 | 2 | 8 |
		// -------------------------------------
		// | 2 | 7 | 1 | 8 | 4 | 6 | 3 | 5 | 9 |
		// -------------------------------------
		// | 3 | 8 | 9 | 1 | 2 | 5 | 6 | 4 | 7 |
		// -------------------------------------
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != 0 {
				b[i][j] = cell{board[i][j], nil}
			} else {
				c := make(map[int]bool)
				for k := 1; k <= 9; k++ {
					c[k] = true
				}
				b[i][j] = cell{board[i][j], c}
			}
		}
	}

	p()
	elim()
	elim()
	elim()
	elim()
	elim()
	elim()
	p()
}

/*
 * 消去法
 */
func elim() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j].num != 0 {
				continue
			}

			for x := 0; x < 9; x++ {
				if x == j || b[i][x].num == 0 {
					continue
				}
				b[i][j].candidates[b[i][x].num] = false
			}
			for y := 0; y < 9; y++ {
				if y == i || b[y][j].num == 0 {
					continue
				}
				b[i][j].candidates[b[y][j].num] = false
			}

			// for y := range r(i) {
			// 	for x := range r(j) {
			// 		if y == i && x == j {
			// 			continue
			// 		}
			// 		if b[y][x].num == 0 {
			// 			continue
			// 		}
			// 		b[i][j].candidates[b[y][x].num] = false
			// 	}
			// }

			b[i][j].num = c(b[i][j].candidates)
		}
	}
}

/*
 * 候補から
 */
func c(c map[int]bool) int {
	n := 0
	for k, v := range c {
		if !v {
			// 候補じゃない
			continue
		}

		if n == 0 {
			n = k
		} else {
			// 候補が複数
			return 0
		}
	}

	return n
}

func p() {
	a := 0
	fmt.Print("\n")
	for i := 0; i < 9; i++ {
		fmt.Print("\n-------------------------------------\n| ")
		for j := 0; j < 9; j++ {
			c := b[i][j]
			if c.num == 0 {
				fmt.Print("  | ")
			} else {
				fmt.Printf("%d | ", c.num)
				a++
			}
		}
	}
	fmt.Println("\n-------------------------------------")
	fmt.Printf("             %2d / 81\n", a)
}
