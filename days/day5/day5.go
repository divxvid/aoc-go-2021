package day5

import (
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
	lines := make([]Line, 0)
	for _, line := range d.lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			lines = append(lines, line)
		}
	}

	return solveGeneric(lines)
}

func (d *Day5) SolveB() (string, error) {
	return solveGeneric(d.lines)
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

func solveGeneric(lines []Line) (string, error) {
	xMax := 0
	yMax := 0

	for _, line := range lines {
		if line.x1 > xMax {
			xMax = line.x1
		}
		if line.x2 > xMax {
			xMax = line.x2
		}
		if line.y1 > yMax {
			yMax = line.y1
		}
		if line.y2 > yMax {
			yMax = line.y2
		}
	}

	grid := make([][]int, xMax+1)
	for i := 0; i <= xMax; i++ {
		grid[i] = make([]int, yMax+1)
	}

	for _, line := range lines {
		fillGridWithLine(grid, line)
	}

	ans := 0
	for i := 0; i <= xMax; i++ {
		for j := 0; j <= yMax; j++ {
			if grid[i][j] > 1 {
				ans++
			}
		}
	}
	return strconv.Itoa(ans), nil
}

func fillGridWithLine(grid [][]int, line Line) {
	if line.x1 == line.x2 {
		y1 := line.y1
		y2 := line.y2
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			grid[line.x1][y]++
		}
	} else if line.y1 == line.y2 {
		x1 := line.x1
		x2 := line.x2
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		for x := x1; x <= x2; x++ {
			grid[x][line.y1]++
		}
	} else {
		handleDiagonalLine(grid, line)
	}
}

func handleDiagonalLine(grid [][]int, line Line) {
	deltaX := line.x2 - line.x1
	deltaY := line.y2 - line.y1

	stepX := 1
	if deltaX < 0 {
		stepX = -1
	}
	stepY := 1
	if deltaY < 0 {
		stepY = -1
	}

	totalSteps := deltaX
	if totalSteps < 0 {
		totalSteps *= -1
	}

	x := line.x1
	y := line.y1
	for i := 0; i <= totalSteps; i++ {
		grid[x][y]++
		x += stepX
		y += stepY
	}
}
