package main

// this method will solve the sudoku
func (s *Sudoku) Solver() bool {
	// this is for row
	for i := 0; i < 9; i++ {
		// this is for col
		for j := 0; j < 9; j++ {
			if s.board[i][j] == 0 {

				for k := 1; k <= 9; k++ {
					if s.Ok(i, j, k) == true {
						s.board[i][j] = k

						if s.Solver() == true {
							return true
						} else {
							s.board[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}

