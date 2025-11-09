package bsm

// MatchState represents the state of a match
type MatchState int

const (
	MatchStateNotYetPlayed MatchState = iota
	MatchStateCancelled
	MatchStateDerby
	MatchStateWon
	MatchStateLost
	MatchStateFinal
)
