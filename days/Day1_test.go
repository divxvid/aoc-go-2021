package days

import "testing"

const input = `199
200
208
210
200
207
240
269
260
263`

func TestSolveA(t *testing.T) {
	want := "7"

	day, err := NewDay1(input)
	if err != nil {
		t.Fatalf("Error in creating new Instance: %+v", err)
	}

	got, err := day.SolveA()
	if err != nil {
		t.Fatalf("Error in SolveA: %+v\n", err)
	}

	if want != got {
		t.Fatalf("Got: %v\nWant: %v\n", got, want)
	}
}

func TestSolveB(t *testing.T) {
	want := "5"

	day, err := NewDay1(input)
	if err != nil {
		t.Fatalf("Error in creating new Instance: %+v", err)
	}

	got, err := day.SolveB()
	if err != nil {
		t.Fatalf("Error in SolveA: %+v\n", err)
	}

	if want != got {
		t.Fatalf("Got: %v\nWant: %v\n", got, want)
	}
}
