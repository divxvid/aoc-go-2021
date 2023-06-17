package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/divxvid/aoc-go-2021/days/day9"
	"github.com/divxvid/aoc-go-2021/models"
	"github.com/divxvid/aoc-go-2021/utils"
)

const fileName = "inputs/day9.txt"

func main() {
	var problem models.Problem

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error in reading from File: %+v", err)
	}

	problem, err = day9.New(string(content))
	if err != nil {
		log.Fatalf("Error in creating new instance: %+v", err)
	}

	outputA, err := utils.TimeSolve(problem.SolveA)
	if err != nil {
		log.Fatalf("Error in SolveA: %+v", err)
	}
	fmt.Printf("Solve A:\n%s\n", outputA)

	outputB, err := utils.TimeSolve(problem.SolveB)
	if err != nil {
		log.Fatalf("Error in SolveB: %+v", err)
	}
	fmt.Printf("Solve B:\n%s\n", outputB)
}
