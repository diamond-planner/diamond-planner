package bsm

// LeaderboardDataset BSM data type for a single leaderboard category
type LeaderboardDataset struct {
	Person      Person      `json:"person"`
	LeagueEntry LeagueEntry `json:"league_entry"`
	Rank        int         `json:"rank"`
	Value       interface{} `json:"value"` // BSM returns both strings and integers for different endpoints
}

// LeaderboardData BSM response for a single league in a stat category
type LeaderboardData struct {
	League        LeagueGroup          `json:"league"`
	StatsCategory string               `json:"stats_category"`
	Data          []LeaderboardDataset `json:"data"`
}

// LeaderboardSummary Custom implementation type,
// adds several sets of LeaderboardData together to be served in a single API response
type LeaderboardSummary struct {
	LeagueID  string            `json:"league_id"`
	StatsType string            `json:"stats_type"`
	Data      []LeaderboardData `json:"data"`
}
