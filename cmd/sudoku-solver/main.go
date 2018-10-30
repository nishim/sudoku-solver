package main

import (
	"fmt"

	"github.com/nishim/sudoku-solver"
)

func main() {
	b := sudoku.NewBoard("9,6,,,3,,7,,,,,,9,,1,5,8,,8,1,,5,,,9,3,,1,,,3,6,,,,8,,,4,2,8,,,6,9,,8,7,,9,,2,,,7,3,1,,,9,,,5,,,,,,3,,9,1,,9,,8,,2,,,7")
	b.Print()
	fmt.Println(b.Solved())
	b.Solve()
	b.Print()
	fmt.Println(b.Solved())
	b.Solve()
	b.Print()
	fmt.Println(b.Solved())
	b.Solve()
	b.Print()
	fmt.Println(b.Solved())
	b.Solve()
	b.Print()
	fmt.Println(b.Solved())
}
