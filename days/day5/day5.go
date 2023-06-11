package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/divxvid/aoc-go-2021/utils"
)

type Day5 struct {
	lines []Line
}

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (d *Day5) SolveA() (string, error) {
	return "unimplemented", nil
}

func (d *Day5) SolveB() (string, error) {
	return "unimplemented", nil
}

func New(data string) (*Day5, error) {
	lines := make([]Line, 0)

	for _, line := range utils.SplitIntoLines(data) {
		splits := strings.Split(line, "->")
		x1, y1, err := splitNumbers(splits[0])
		if err != nil {
			return nil, err
		}
		x2, y2, err := splitNumbers(splits[1])
		if err != nil {
			return nil, err
		}
		newLine := Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
		lines = append(lines, newLine)
	}

	fmt.Println(lines)
	return &Day5{lines: lines}, nil
}

func splitNumbers(data string) (int, int, error) {
	data = strings.Trim(data, " ")
	splits := strings.Split(data, ",")
	x, err := strconv.Atoi(splits[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(splits[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}
