package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/divxvid/aoc-go-2021/days/day1"
)

const fileName = "inputs/day1.txt"

type Problem interface {
	SolveA() (string, error)
	SolveB() (string, error)
}

func main() {
	var problem Problem

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Error in reading from File: %+v", err)
	}

	problem, err = day1.New(string(content))
	if err != nil {
		log.Fatalf("Error in creating new instance: %+v", err)
	}

	outputA, err := problem.SolveA()
	if err != nil {
		log.Fatalf("Error in SolveA: %+v", err)
	}
	fmt.Printf("Solve A:\n%s\n", outputA)

	outputB, err := problem.SolveB()
	if err != nil {
		log.Fatalf("Error in SolveB: %+v", err)
	}
	fmt.Printf("Solve B:\n%s\n", outputB)
}
