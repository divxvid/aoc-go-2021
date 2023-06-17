package day9

import (
	"github.com/divxvid/aoc-go-2021/utils"
)

type Day9 struct {
	grid [][]int
}

func New(data string) (*Day9, error) {
	grid := make([][]int, 0)
	for _, line := range utils.SplitIntoLines(data) {
		row := make([]int, 0)
		for _, ch := range line {
			chValue := int(ch) - int('0')
			row = append(row, chValue)
		}
		grid = append(grid, row)
	}
	return &Day9{grid: grid}, nil
}

func (d *Day9) SolveA() (string, error) {
	return "unimplemented", nil
}

func (d *Day9) SolveB() (string, error) {
	return "unimplemented", nil
}
