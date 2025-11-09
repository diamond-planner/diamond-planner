package bsm

import "testing"

func intPtr(v int) *int { return &v }

func TestMatch_IsDerby(t *testing.T) {
	m := Match{HomeTeamName: "Berlin Skylarks I", AwayTeamName: "Berlin Skylarks II"}
	if !m.IsDerby("Berlin Skylarks") {
		t.Fatalf("expected derby to be true")
	}

	nm := Match{HomeTeamName: "Berlin Skylarks I", AwayTeamName: "Potsdam Porcupines"}
	if nm.IsDerby("Berlin Skylarks") {
		t.Fatalf("expected derby to be false")
	}
}

func TestMatch_GetWinnerForMatch(t *testing.T) {
	// nil runs => none
	m := Match{HomeRuns: nil, AwayRuns: nil}
	if got := m.GetWinnerForMatch(); got != GameWinnerNone {
		t.Fatalf("expected GameWinnerNone, got %v", got)
	}

	// home win
	m = Match{HomeRuns: intPtr(5), AwayRuns: intPtr(2)}
	if got := m.GetWinnerForMatch(); got != GameWinnerHome {
		t.Fatalf("expected GameWinnerHome, got %v", got)
	}

	// away win
	m = Match{HomeRuns: intPtr(2), AwayRuns: intPtr(3)}
	if got := m.GetWinnerForMatch(); got != GameWinnerAway {
		t.Fatalf("expected GameWinnerAway, got %v", got)
	}

	// tie
	m = Match{HomeRuns: intPtr(3), AwayRuns: intPtr(3)}
	if got := m.GetWinnerForMatch(); got != GameWinnerNone {
		t.Fatalf("expected GameWinnerNone for tie, got %v", got)
	}
}

func TestMatch_GetMatchState(t *testing.T) {
	team := "Skylarks"

	// planned
	m := Match{State: "planned"}
	if got := m.GetMatchState(team); got != MatchStateNotYetPlayed {
		t.Fatalf("expected NotYetPlayed, got %v", got)
	}

	// cancelled (british)
	m = Match{State: "cancelled"}
	if got := m.GetMatchState(team); got != MatchStateCancelled {
		t.Fatalf("expected Cancelled, got %v", got)
	}

	// canceled (american)
	m = Match{State: "canceled"}
	if got := m.GetMatchState(team); got != MatchStateCancelled {
		t.Fatalf("expected Cancelled, got %v", got)
	}

	// derby has priority over winner/loser
	m = Match{State: "finished", HomeTeamName: "Berlin Skylarks I", AwayTeamName: "Berlin Skylarks II"}
	if got := m.GetMatchState("Berlin Skylarks"); got != MatchStateDerby {
		t.Fatalf("expected Derby, got %v", got)
	}

	// won as home team
	m = Match{State: "finished", HomeTeamName: "Skylarks", AwayTeamName: "Opponents", HomeRuns: intPtr(4), AwayRuns: intPtr(1)}
	if got := m.GetMatchState(team); got != MatchStateWon {
		t.Fatalf("expected Won, got %v", got)
	}

	// lost as home team
	m = Match{State: "finished", HomeTeamName: "Skylarks", AwayTeamName: "Opponents", HomeRuns: intPtr(1), AwayRuns: intPtr(2)}
	if got := m.GetMatchState(team); got != MatchStateLost {
		t.Fatalf("expected Lost, got %v", got)
	}

	// won as away team
	m = Match{State: "finished", HomeTeamName: "Opponents", AwayTeamName: "Skylarks", HomeRuns: intPtr(1), AwayRuns: intPtr(2)}
	if got := m.GetMatchState(team); got != MatchStateWon {
		t.Fatalf("expected Won (away), got %v", got)
	}

	// lost as away team
	m = Match{State: "finished", HomeTeamName: "Opponents", AwayTeamName: "Skylarks", HomeRuns: intPtr(3), AwayRuns: intPtr(2)}
	if got := m.GetMatchState(team); got != MatchStateLost {
		t.Fatalf("expected Lost (away), got %v", got)
	}

	// final when tie or unrelated team
	m = Match{State: "finished", HomeTeamName: "Others A", AwayTeamName: "Others B", HomeRuns: intPtr(3), AwayRuns: intPtr(3)}
	if got := m.GetMatchState(team); got != MatchStateFinal {
		t.Fatalf("expected Final for tie/unrelated, got %v", got)
	}
}

func TestMatch_IsPlayoffGame(t *testing.T) {
	m := Match{MatchID: "2025-PO-123"}
	if !m.IsPlayoffGame() {
		t.Fatalf("expected playoff game to be true")
	}

	nm := Match{MatchID: "2025-RS-123"}
	if nm.IsPlayoffGame() {
		t.Fatalf("expected playoff game to be false")
	}
}
