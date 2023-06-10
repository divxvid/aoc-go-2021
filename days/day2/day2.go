package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type direction int

const (
	forward direction = iota
	up
	down
)

type Day2 struct {
	moves []Move
}

type Move struct {
	dir    direction
	amount int
}

func New(data string) (*Day2, error) {
	lines := strings.Split(data, "\n")
	moves := make([]Move, 0)

	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		lineContent := strings.Split(line, " ")
		if len(lineContent) != 2 {
			return nil, fmt.Errorf("Line does not have direction and amount. Line: %s\n", line)
		}

		var dir direction
		switch lineContent[0] {
		case "forward":
			dir = forward
		case "up":
			dir = up
		case "down":
			dir = down
		default:
			return nil, fmt.Errorf("Invalid Direction: %s\n", lineContent[0])
		}

		amount, err := strconv.ParseInt(lineContent[1], 10, 64)
		if err != nil {
			return nil, err
		}

		moves = append(moves, Move{dir: dir, amount: int(amount)})
	}
	ret := &Day2{
		moves: moves,
	}
	return ret, nil
}

func (d *Day2) SolveA() (string, error) {
	horizontalPosition := 0
	depth := 0

	for _, move := range d.moves {
		switch move.dir {
		case forward:
			horizontalPosition += move.amount
		case up:
			depth -= move.amount
		case down:
			depth += move.amount
		}
	}

	result := horizontalPosition * depth
	return strconv.Itoa(result), nil
}

func (d *Day2) SolveB() (string, error) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, move := range d.moves {
		switch move.dir {
		case forward:
			horizontalPosition += move.amount
			depth += aim * move.amount
		case up:
			aim -= move.amount
		case down:
			aim += move.amount
		}
	}

	result := horizontalPosition * depth
	return strconv.Itoa(result), nil
}
