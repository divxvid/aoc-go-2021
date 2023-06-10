package utils

import "strings"

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
