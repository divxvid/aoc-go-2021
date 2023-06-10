package day1

import (
	"strconv"
	"strings"
)

type Day1 struct {
	depths []int64
}

// SolveA implements main.Problem
func (d *Day1) SolveA() (string, error) {
	length := len(d.depths)
	ans := 0
	for i := 1; i < length; i += 1 {
		if d.depths[i] > d.depths[i-1] {
			ans += 1
		}
	}
	return strconv.Itoa(ans), nil
}

// SolveB implements main.Problem
func (d *Day1) SolveB() (string, error) {
	length := len(d.depths)
	runningSum := d.depths[0] + d.depths[1] + d.depths[2]
	ans := 0
	for i := 3; i < length; i += 1 {
		newRunningSum := runningSum - d.depths[i-3] + d.depths[i]
		if newRunningSum > runningSum {
			ans += 1
		}
		runningSum = newRunningSum
	}
	return strconv.Itoa(ans), nil
}

func New(data string) (*Day1, error) {
	lines := strings.Split(data, "\n")
	depths := make([]int64, 0)
	for _, line := range lines {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		v, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		depths = append(depths, v)
	}
	return &Day1{depths: depths}, nil
}
