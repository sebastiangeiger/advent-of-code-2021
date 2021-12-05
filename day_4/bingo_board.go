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
