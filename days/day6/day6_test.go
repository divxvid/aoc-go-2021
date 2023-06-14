package day6

import "testing"

const input string = `3,4,3,1,2`

func TestSolveA(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "5934"

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

	want := "26984457539"

	got, err := problem.SolveB()

	if err != nil {
		t.Fatalf("Cannot solve Problem B: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}
