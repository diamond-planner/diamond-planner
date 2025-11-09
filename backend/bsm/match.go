package bsm

import (
	"strings"
)

type Match struct {
	ID              int         `json:"id"`
	MatchID         string      `json:"match_id"`
	Time            string      `json:"time"`
	LeagueID        int         `json:"league_id"`
	HomeRuns        *int        `json:"home_runs,omitempty"`
	AwayRuns        *int        `json:"away_runs,omitempty"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	State           string      `json:"state"`
	Season          int         `json:"season"`
	SecondLeagueID  *int        `json:"second_league_id,omitempty"`
	PlannedInnings  int         `json:"planned_innings"`
	Name            *string     `json:"name,omitempty"`
	HumanState      string      `json:"human_state"`
	ScoresheetURL   *string     `json:"scoresheet_url,omitempty"`
	HomeTeamName    string      `json:"home_team_name"`
	AwayTeamName    string      `json:"away_team_name"`
	HomeLeagueEntry LeagueEntry `json:"home_league_entry"`
	AwayLeagueEntry LeagueEntry `json:"away_league_entry"`
	League          League      `json:"league"`
	Field           Field       `json:"field"`
}

// IsDerby checks if the match is a derby game between teams with the same name
func (m *Match) IsDerby(teamName string) bool {
	return strings.Contains(m.HomeTeamName, teamName) &&
		strings.Contains(m.AwayTeamName, teamName)
}

// GetWinnerForMatch determines the winner of the match
func (m *Match) GetWinnerForMatch() GameWinner {
	if m.HomeRuns == nil || m.AwayRuns == nil {
		return GameWinnerNone
	}

	if *m.HomeRuns > *m.AwayRuns {
		return GameWinnerHome
	} else if *m.AwayRuns > *m.HomeRuns {
		return GameWinnerAway
	}

	return GameWinnerNone
}

// GetMatchState determines the current state of the match
func (m *Match) GetMatchState(teamName string) MatchState {
	if m.State == "planned" {
		return MatchStateNotYetPlayed
	}

	if m.State == "cancelled" || m.State == "canceled" {
		return MatchStateCancelled
	}

	if m.IsDerby(teamName) {
		return MatchStateDerby
	}

	winner := m.GetWinnerForMatch()

	isHomeTeam := strings.Contains(m.HomeTeamName, teamName)
	isAwayTeam := strings.Contains(m.AwayTeamName, teamName)

	if (winner == GameWinnerHome && isHomeTeam) ||
		(winner == GameWinnerAway && isAwayTeam) {
		return MatchStateWon
	} else if (winner == GameWinnerAway && isHomeTeam) ||
		(winner == GameWinnerHome && isAwayTeam) {
		return MatchStateLost
	}

	return MatchStateFinal
}

// IsPlayoffGame checks if the match is a playoff game
func (m *Match) IsPlayoffGame() bool {
	return strings.Contains(m.MatchID, "PO")
}
