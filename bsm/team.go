package bsm

type Team struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Season        int           `json:"season"`
	ShortName     string        `json:"short_name"`
	Clubs         []Club        `json:"clubs"`
	LeagueEntries []LeagueEntry `json:"league_entries"`
}
