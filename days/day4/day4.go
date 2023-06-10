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
	return "unsolved", nil
}

func (d *Day4) SolveB() (string, error) {
	return "unsolved", nil
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
