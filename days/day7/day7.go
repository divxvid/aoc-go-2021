package day7

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type Day7 struct {
	position []int
}

func New(data string) (*Day7, error) {
	data = strings.Trim(data, "\r\n")
	position := make([]int, 0)
	for _, v := range strings.Split(data, ",") {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		position = append(position, vInt)
	}
	sort.Slice(position, func(i, j int) bool {
		return position[i] <= position[j]
	})
	return &Day7{position: position}, nil
}

func (d *Day7) SolveA() (string, error) {
	ans := 10000001
	for i := 0; i < len(d.position); i++ {
		dist := calculateFuelLinear(d.position[i], d.position)
		if dist < ans {
			ans = dist
		}
	}
	return strconv.Itoa(ans), nil
}

func calculateFuelLinear(x int, position []int) int {
	result := 0
	for _, v := range position {
		dist := v - x
		if dist < 0 {
			dist *= -1
		}
		result += dist
	}
	return result
}

func (d *Day7) SolveB() (string, error) {
	ans := math.MaxInt
	limit := d.position[len(d.position)-1]
	for i := d.position[0]; i < limit; i++ {
		fuelBurnt := calculateFuelIncreasing(i, d.position)
		if fuelBurnt < ans {
			ans = fuelBurnt
		}
	}
	return strconv.Itoa(ans), nil
}

func calculateFuelIncreasing(x int, position []int) int {
	result := 0
	for _, v := range position {
		dist := v - x
		if dist < 0 {
			dist *= -1
		}
		result += calculateFuelBurn(dist)
	}
	return result
}

func calculateFuelBurn(dist int) int {
	return (dist * (dist + 1)) / 2
}
