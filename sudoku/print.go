package main

import "fmt"

func (s *Sudoku) Print() {
	fmt.Println("------------------------------")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 0; i < 3; i++ {
		for j := 3; j < 6; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 0; i < 3; i++ {
		for j := 6; j < 9; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")
	fmt.Println("------------------------------")

	for i := 3; i < 6; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 3; i < 6; i++ {
		for j := 3; j < 6; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 3; i < 6; i++ {
		for j := 6; j < 9; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")
	fmt.Println("------------------------------")

	for i := 6; i < 9; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 6; i <9; i++ {
		for j := 3; j < 6; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")

	for i := 6; i < 9; i++ {
		for j := 6; j < 9; j++ {
			fmt.Printf(" %d ", s.board[i][j])
		}
		fmt.Printf("|")
	}
	fmt.Printf("\n")
	fmt.Println("------------------------------\n\n")
}
