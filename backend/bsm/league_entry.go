package bsm

type LeagueEntry struct {
	ID                        int     `json:"id"`
	CreatedAt                 string  `json:"created_at"`
	UpdatedAt                 string  `json:"updated_at"`
	GameDay                   *string `json:"game_day"`
	NotCompeting              bool    `json:"not_competing"`
	State                     string  `json:"state"`
	HumanGameDay              *string `json:"human_game_day,omitempty"`
	HumanLicenseCriteriaState *string `json:"human_license_criteria_state,omitempty"`
	Team                      *Team   `json:"team,omitempty"`
	League                    League  `json:"league"`
	//Field                       *Field   `json:"field,omitempty"`
}
