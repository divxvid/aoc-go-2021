package utils

import (
	"fmt"
	"strings"
	"time"
)

func SplitIntoLines(data string) []string {
	lines := make([]string, 0)

	for _, line := range strings.Split(data, "\n") {
		line = strings.Trim(line, "\r")
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

type SolveFuncType func() (string, error)

func TimeSolve(f SolveFuncType) (string, error) {
	startTime := time.Now()
	ret, err := f()
	timeTaken := time.Since(startTime)

	fmt.Printf("Time Taken for Solving: %d us\n", timeTaken.Microseconds())
	return ret, err
}
