package day6

import (
	"strconv"
	"strings"
)

type Day6 struct {
	initialConfig []int64
}

func New(data string) (*Day6, error) {
	data = strings.Trim(data, "\r\n")
	initialConfig := make([]int64, 0)
	for _, v := range strings.Split(data, ",") {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		initialConfig = append(initialConfig, vInt)
	}
	return &Day6{initialConfig: initialConfig}, nil
}

func (d *Day6) SolveA() (string, error) {
	return "unimplemented", nil
}

func (d *Day6) SolveB() (string, error) {
	return "unimplemented", nil
}
