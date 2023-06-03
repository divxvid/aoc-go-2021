package main

import (
	"fmt"
	"log"

	"github.com/divxvid/aoc-go-2021/days"
)

const fileName = "inputs/ay1/test.txt"

func main() {
	var problem Problem
	problem = days.NewDay1()

	err := problem.ReadFromFile(fileName)
	if err != nil {
		log.Fatalf("Error in ReadFromFile: %+v", err)
	}

	outputA, err := problem.SolveA()
	if err != nil {
		log.Fatalf("Error in SolveA: %+v", err)
	}
	fmt.Println("Solve A:\n", outputA)

	problem.ReadFromFile(fileName)
	outputB, err := problem.SolveB()
	if err != nil {
		log.Fatalf("Error in SolveB: %+v", err)
	}
	fmt.Println("Solve B:\n", outputB)
}
