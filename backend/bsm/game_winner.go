package bsm

// GameWinner represents the winner of a game
type GameWinner int

const (
	GameWinnerNone GameWinner = iota
	GameWinnerHome
	GameWinnerAway
)
