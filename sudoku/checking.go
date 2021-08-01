package main

type Sudoku struct {
	board [][]int
}

// intial the board for sudoko
func (s *Sudoku) SudokuInit (unSolved [][]int) {
	//s.board = make([][]int, 9)
	//for ints := range s.board {
	//	s.board[ints] = make([]int, 9)
	//}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			s.board[i][j] = unSolved[i][j]
		}
	}
}

// check if the number is exist in a row twice
func (s *Sudoku) InRow(row int, num int) bool {
	for i := 0; i < 9; i++ {
		if s.board[row][i] == num {
			return true
		}
	}
	return false
}

// check if the number is exist in a col twice
func (s *Sudoku) InCol(col int, num int) bool {
	for i := 0; i < 9; i++ {
		if s.board[i][col] == num {
			return true
		}
	}
	return false
}

// check if the number is exist in a box twice
func (s *Sudoku) InBox(row int, col int, num int) bool {
	rowStart := row - (row % 3)
	colStart := col - (col % 3)

	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if s.board[i][j] == num {
				return true
			}
		}
	}
	return  false
}

// mixing all of the check method
func (s *Sudoku) Ok (row int, col int, num int) bool {
	return !s.InRow(row, num) && !s.InCol(col, num) && !s.InBox(row, col, num)
}
