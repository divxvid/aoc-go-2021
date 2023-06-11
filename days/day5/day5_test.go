package day5

import "testing"

const input string = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestSolveA(t *testing.T) {
	problem, err := New(input)

	if err != nil {
		t.Fatalf("Cannot create new instance: %+v", err)
	}

	want := "5"

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

	want := "12"

	got, err := problem.SolveB()

	if err != nil {
		t.Fatalf("Cannot solve Problem B: %+v", err)
	}

	if got != want {
		t.Fatalf("Got: %s\nWant: %s\n", got, want)
	}
}
