package main

import "fmt"

func main() {
	fmt.Println("write a sudoku to solve")

	s := new(Sudoku)
	s.board = make([][]int, 9)
	for ints := range s.board {
		s.board[ints] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Scanf("%d", &s.board[i][j])
		}
	}
	s.SudokuInit(s.board)
	s.Print()

	if s.Solver() {
		s.Print()
		fmt.Println("solved successfully")
	} else {
		fmt.Println("unsolvable")
	}
}
