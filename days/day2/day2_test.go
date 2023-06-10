package day2

import "testing"

const input string = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestSolveA(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "150"

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

	want := "900"

	got, err := problem.SolveB()

	if err != nil {
		t.Fatalf("Cannot solve Problem B: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}
