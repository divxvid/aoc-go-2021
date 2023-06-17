package day9

import (
	"sort"
	"strconv"

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
	result := 0
	for i, row := range d.grid {
		for j := range row {
			if d.isDepression(i, j) {
				result += (row[j] + 1)
			}
		}
	}
	return strconv.Itoa(result), nil
}

func (d *Day9) isDepression(i, j int) bool {
	if i-1 >= 0 && d.grid[i-1][j] <= d.grid[i][j] {
		return false
	}
	if i+1 < len(d.grid) && d.grid[i+1][j] <= d.grid[i][j] {
		return false
	}
	if j-1 >= 0 && d.grid[i][j-1] <= d.grid[i][j] {
		return false
	}
	if j+1 < len(d.grid[0]) && d.grid[i][j+1] <= d.grid[i][j] {
		return false
	}
	return true
}

func (d *Day9) SolveB() (string, error) {
	result := 1
	basinSizes := make([]int, 0)

	basinMask := make([][]bool, len(d.grid))
	for i := range d.grid {
		basinRow := make([]bool, len(d.grid[0]))
		basinMask[i] = basinRow
	}

	for i, row := range d.grid {
		for j := range row {
			if basinMask[i][j] || row[j] == 9 {
				continue
			}
			basinSize := d.makeBasin(i, j, basinMask)
			basinSizes = append(basinSizes, basinSize)
		}
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[j] <= basinSizes[i]
	})

	for _, x := range basinSizes[:3] {
		result *= x
	}

	return strconv.Itoa(result), nil
}

func (d *Day9) makeBasin(i int, j int, mask [][]bool) int {
	if i < 0 || j < 0 || i >= len(mask) || j >= len(mask[0]) {
		return 0
	}
	if d.grid[i][j] == 9 || mask[i][j] {
		return 0
	}
	mask[i][j] = true

	totalSize := 1
	totalSize += d.makeBasin(i+1, j, mask)
	totalSize += d.makeBasin(i-1, j, mask)
	totalSize += d.makeBasin(i, j+1, mask)
	totalSize += d.makeBasin(i, j-1, mask)

	return totalSize
}
