package day9

import (
	"testing"
)

const input string = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestSolveA(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "15"

	got, err := problem.SolveA()

	if err != nil {
		t.Fatalf("Cannot solve Problem A: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}

func TestSolveB(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "61229"

	got, err := problem.SolveB()

	if err != nil {
		t.Fatalf("Cannot solve Problem B: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}
