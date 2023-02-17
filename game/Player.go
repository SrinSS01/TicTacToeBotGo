package game

type PlayerType struct {
	Value string
}

var (
	CROSS  = PlayerType{Value: "x"}
	NOUGHT = PlayerType{Value: "o"}
)

type Player struct {
	Type         *PlayerType
	OpponentType *PlayerType
}
