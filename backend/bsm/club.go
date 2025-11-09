package bsm

type Club struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Acronym   string `json:"acronym"`
	ShortName string `json:"short_name"`
	LogoURL   string `json:"logo_url"`
}
