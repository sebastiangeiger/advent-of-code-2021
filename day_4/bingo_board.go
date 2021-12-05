package day_4

type BingoBoard struct {
	board  [][]int
	marked [][]bool
}

func (bingoBoard *BingoBoard) Mark(number int) {
	dx, dy := bingoBoard.Dim()
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if bingoBoard.board[x][y] == number {
				bingoBoard.marked[x][y] = true
			}
		}
	}
}

func (bingoBoard *BingoBoard) Dim() (int, int) {
	return len(bingoBoard.board), len(bingoBoard.board[0])
}

func (bingoBoard *BingoBoard) HasWon() bool {
	dx, dy := bingoBoard.Dim()
	for x := 0; x < dx; x++ {
		rowHasWon := true
		for y := 0; y < dy; y++ {
			if bingoBoard.marked[x][y] == false {
				rowHasWon = false
				break
			}
		}
		if rowHasWon {
			return true
		}
	}
	for y := 0; y < dy; y++ {
		columnHasWon := true
		for x := 0; x < dx; x++ {
			if bingoBoard.marked[x][y] == false {
				columnHasWon = false
				break
			}
		}
		if columnHasWon {
			return true
		}
	}
	return false
}
