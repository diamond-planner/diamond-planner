package bsm

// LeagueGroup represents the BSM structure "Liga/Gruppe".
// It is always unique to single season.
// Caution: several LeagueGroup might have the same League.
type LeagueGroup struct {
	ID      int    `json:"id"`
	Season  int    `json:"season"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
	League  League `json:"league"`
}
