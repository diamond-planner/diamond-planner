package bsm

type Field struct {
	BSMID            int     `json:"id"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	ContactName      *string `json:"contact_name"` // Nullable field
	AddressAddon     string  `json:"address_addon"`
	Street           string  `json:"street"`
	PostalCode       string  `json:"postal_code"`
	City             string  `json:"city"`
	Country          string  `json:"country"`
	Longitude        float64 `json:"longitude"`
	Latitude         float64 `json:"latitude"`
	SpectatorTotal   int     `json:"spectator_total"`
	SpectatorSeats   int     `json:"spectator_seats"`
	OtherInformation string  `json:"other_information"`
	Groundrules      string  `json:"groundrules"`
	HumanCountry     string  `json:"human_country"`
	PhotoURL         string  `json:"photo_url"`
}
