package day8

import (
	"strings"

	"github.com/divxvid/aoc-go-2021/utils"
)

type Query struct {
	jumbledSegments []string
	digits          []string
}

type Day8 struct {
	queries []Query
}

func New(data string) (*Day8, error) {
	queries := make([]Query, 0)

	for _, line := range utils.SplitIntoLines(data) {
		parts := strings.Split(line, "|")
		segments := strings.Split(parts[0], " ")
		digits := strings.Split(parts[1], " ")
		for i := range segments {
			segments[i] = strings.Trim(segments[i], " ")
		}
		for i := range digits {
			digits[i] = strings.Trim(digits[i], " ")
		}

		queries = append(queries, Query{jumbledSegments: segments, digits: digits})
	}

	return &Day8{
		queries: queries,
	}, nil
}

func (d *Day8) SolveA() (string, error) {
	return "unimplemented", nil
}

func (d *Day8) SolveB() (string, error) {
	return "unimplemented", nil
}
