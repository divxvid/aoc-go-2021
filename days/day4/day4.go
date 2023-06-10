package day4

import (
	"strconv"
	"strings"

	"github.com/divxvid/aoc-go-2021/utils"
)

type Day4 struct {
	drawingSequence []int
	boards          [][][]int
}

func (d *Day4) SolveA() (string, error) {
	numbersCut := make([][5][5]bool, len(d.boards))
	for i := 0; i < len(d.boards); i++ {
		numbersCut[i] = [5][5]bool{}
	}

	var winningBoardNumber int
	var lastDrawnNumber int

outer:
	for _, drawnNumber := range d.drawingSequence {
		for i, board := range d.boards {
			cutNumber(drawnNumber, board, &numbersCut[i])
		}
		for i := 0; i < len(d.boards); i++ {
			if hasWon(numbersCut[i]) {
				winningBoardNumber = i
				lastDrawnNumber = drawnNumber
				break outer
			}
		}
	}

	winningBoard := d.boards[winningBoardNumber]
	unmarkedNumbers := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if numbersCut[winningBoardNumber][i][j] {
				continue
			}
			unmarkedNumbers += winningBoard[i][j]
		}
	}

	result := unmarkedNumbers * lastDrawnNumber

	return strconv.Itoa(result), nil
}

func hasWon(b [5][5]bool) bool {
	for row := 0; row < 5; row++ {
		ok := true
		for col := 0; col < 5; col++ {
			ok = ok && b[row][col]
		}
		if ok {
			return true
		}
	}

	for col := 0; col < 5; col++ {
		ok := true
		for row := 0; row < 5; row++ {
			ok = ok && b[row][col]
		}
		if ok {
			return true
		}
	}
	return false
}

func (d *Day4) SolveB() (string, error) {
	numbersCut := make([][5][5]bool, len(d.boards))
	for i := 0; i < len(d.boards); i++ {
		numbersCut[i] = [5][5]bool{}
	}

	var winningBoardNumber int
	var lastDrawnNumber int
	boardsLeft := len(d.boards)
	skip := make([]bool, len(d.boards))

	for _, drawnNumber := range d.drawingSequence {
		for i, board := range d.boards {
			if skip[i] {
				continue
			}
			cutNumber(drawnNumber, board, &numbersCut[i])
		}
		for i := 0; i < len(d.boards); i++ {
			if skip[i] {
				continue
			}
			if hasWon(numbersCut[i]) {
				winningBoardNumber = i
				lastDrawnNumber = drawnNumber
				skip[i] = true
				boardsLeft--
			}
		}
		if boardsLeft == 0 {
			break
		}
	}

	winningBoard := d.boards[winningBoardNumber]
	unmarkedNumbers := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if numbersCut[winningBoardNumber][i][j] {
				continue
			}
			unmarkedNumbers += winningBoard[i][j]
		}
	}

	result := unmarkedNumbers * lastDrawnNumber

	return strconv.Itoa(result), nil
}

func New(data string) (*Day4, error) {
	lines := utils.SplitIntoLines(data)
	sequenceString := lines[0]

	drawingSequence := make([]int, 0)
	for _, value := range strings.Split(sequenceString, ",") {
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return nil, err
		}
		drawingSequence = append(drawingSequence, int(intValue))
	}

	boards := make([][][]int, 0)
	for start := 1; start < len(lines); start += 5 {
		board, err := makeBoard(lines, start)
		if err != nil {
			return nil, err
		}
		boards = append(boards, board)
	}

	return &Day4{
		drawingSequence: drawingSequence,
		boards:          boards,
	}, nil
}

func makeBoard(data []string, start int) ([][]int, error) {
	board := make([][]int, 5)
	for i, line := range data[start:(start + 5)] {
		boardRow := make([]int, 0)
		line = strings.Trim(line, " ")
		for _, value := range strings.Split(line, " ") {
			if value == "" {
				continue
			}
			intValue, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return nil, err
			}
			boardRow = append(boardRow, int(intValue))
		}
		board[i] = boardRow
	}
	return board, nil
}

func cutNumber(number int, board [][]int, cuts *[5][5]bool) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == number {
				cuts[i][j] = true
			}
		}
	}
}
