package day6

import (
	"fmt"
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
	return solveGeneric(80, d.initialConfig)
}

func (d *Day6) SolveB() (string, error) {
	return solveGeneric(256, d.initialConfig)
}

func solveGeneric(numDays int64, nums []int64) (string, error) {
	var result int64 = int64(len(nums))
	dp := make([]int64, numDays+1)
	for i := 0; i <= int(numDays); i++ {
		dp[i] = -1
	}

	for _, x := range nums {
		result += f(numDays-x-1, dp)
	}
	return fmt.Sprintf("%d", result), nil
}

func f(x int64, dp []int64) int64 {
	if x < 0 {
		return 0
	}
	if dp[x] != -1 {
		return dp[x]
	}
	result := 1 + f(x-7, dp) + f(x-9, dp)
	dp[x] = result
	return result
}
