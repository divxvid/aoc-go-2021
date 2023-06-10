package day3

import (
	"fmt"
	"strconv"

	"github.com/divxvid/aoc-go-2021/utils"
)

type Day3 struct {
	values []string
}

func New(data string) (*Day3, error) {
	return &Day3{values: utils.SplitIntoLines(data)}, nil
}

func (d *Day3) SolveA() (string, error) {
	length := len(d.values[0])
	totalEntries := len(d.values)
	numOfOnes := countNumberOfOnes(d.values)

	gammaRate := 0
	epsilonRate := 0

	for i := 0; i < length; i++ {
		gammaRate <<= 1
		epsilonRate <<= 1

		numOfZeros := totalEntries - numOfOnes[i]
		if numOfOnes[i] >= numOfZeros {
			gammaRate += 1
		} else {
			epsilonRate += 1
		}
	}

	result := gammaRate * epsilonRate

	return strconv.Itoa(result), nil
}

func (d *Day3) SolveB() (string, error) {
	totalEntries := len(d.values)
	oxygenRatingData := make([]string, totalEntries)
	co2RatingData := make([]string, totalEntries)

	copy(oxygenRatingData, d.values)
	copy(co2RatingData, d.values)

	oxygenRating := calculateRating(oxygenRatingData, true)
	co2Rating := calculateRating(co2RatingData, false)

	result := oxygenRating * co2Rating

	return strconv.Itoa(result), nil
}

func calculateRating(data []string, isOxy bool) int {
	length := len(data[0])
	for i := 0; len(data) > 1 && i < length; i++ {
		numOnes := 0
		for _, value := range data {
			if value[i] == '1' {
				numOnes++
			}
		}
		numZeros := len(data) - numOnes
		var dominant byte

		if numOnes >= numZeros {
			dominant = '1'
		} else {
			dominant = '0'
		}

		newData := make([]string, 0)
		for _, value := range data {
			if value[i] == dominant && isOxy {
				newData = append(newData, value)
			}
			if value[i] != dominant && !isOxy {
				newData = append(newData, value)
			}
		}
		data = newData
	}

	rating, err := strconv.ParseInt(data[0], 2, 32)
	if err != nil {
		fmt.Printf("Cannot convert string to int: %+v\n", err)
	}

	return int(rating)
}

func countNumberOfOnes(data []string) []int {
	length := len(data[0])
	numOfOnes := make([]int, length)

	for _, value := range data {
		for i := 0; i < length; i++ {
			if value[i] == '1' {
				numOfOnes[i]++
			}
		}
	}

	return numOfOnes
}
