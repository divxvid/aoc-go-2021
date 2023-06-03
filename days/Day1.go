package days

import (
	"bufio"
	"fmt"
	"os"
)

type Day1 struct {
}

// ReadFromFile implements main.Problem
func (day *Day1) ReadFromFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// ReadFromFile implements main.Problem
func (day *Day1) ReadFromString(data string) error {

	return nil
}

// SolveA implements main.Problem
func (*Day1) SolveA() (string, error) {
	return "Unimplemented", nil
}

// SolveB implements main.Problem
func (*Day1) SolveB() (string, error) {
	return "Unimplemented", nil
}

func NewDay1() *Day1 {
	return &Day1{}
}
