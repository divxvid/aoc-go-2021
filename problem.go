package main

type Problem interface {
	ReadFromFile(string) error
	ReadFromString(string) error
	SolveA() (string, error)
	SolveB() (string, error)
}
