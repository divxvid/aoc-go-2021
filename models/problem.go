package models

type Problem interface {
	SolveA() (string, error)
	SolveB() (string, error)
}
