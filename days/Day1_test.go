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

	day := NewDay1()
	day.ReadFromString(input)
	got, err := day.SolveA()
	if err != nil {
		t.Fatalf("Error in SolveA: %+v\n", err)
	}

	if want != got {
		t.Fatalf("Got: %v\nWant: %v\n", got, want)
	}
}
