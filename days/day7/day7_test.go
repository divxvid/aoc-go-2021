package day7

import "testing"

const input string = `16,1,2,0,4,2,7,1,2,14`

func TestSolveA(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "37"

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

	want := "168"

	got, err := problem.SolveB()

	if err != nil {
		t.Fatalf("Cannot solve Problem B: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}
