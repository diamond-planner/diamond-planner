package bsm

// Table represents a league table with standings information
type Table struct {
	LeagueID   int    `json:"league_id"`
	LeagueName string `json:"league_name"`
	Season     int    `json:"season"`
	Rows       []Row  `json:"rows"`
}

// Row represents a single row in the league table standings
type Row struct {
	Rank          string  `json:"rank"`
	LeagueEntryID int     `json:"league_entry_id"`
	TeamName      string  `json:"team_name"`
	ShortTeamName string  `json:"short_team_name"`
	MatchCount    int     `json:"match_count"`
	WinsCount     float64 `json:"wins_count"` // Using float64 for half-wins (T-Ball league sometimes has ties)
	LossesCount   float64 `json:"losses_count"`
	Quota         string  `json:"quota"`
	GamesBehind   string  `json:"games_behind"`
	Streak        string  `json:"streak"`
}

// DetermineTableRow determines the table row for a team
func DetermineTableRow(team Team, table *Table) *Row {
	if table == nil || len(table.Rows) == 0 {
		return nil
	}

	for _, row := range table.Rows {
		if row.TeamName == team.Name || row.ShortTeamName == team.ShortName {
			return &row
		}
	}

	return nil
}
