package dp

import "github.com/pocketbase/pocketbase/core"

var _ core.RecordProxy = (*Location)(nil)

// Location RecordProxy for collection `locations`.
// Testing type-safe struct access.
type Location struct {
	core.BaseRecordProxy
}

func (a *Location) ID() string {
	return a.Id
}

func (a *Location) BSMID() int {
	return a.GetInt("bsm_id")
}

func (a *Location) SetBSMID(bsmID int) {
	a.Set("bsm_id", bsmID)
}

// Name contains mostly useless data like "Baseball"
// there is also `internal_name` in the record model to have a field that automatic imports do not overwrite
func (a *Location) Name() string {
	return a.GetString("name")
}

func (a *Location) SetName(name string) {
	a.Set("name", name)
}

func (a *Location) Description() string {
	return a.GetString("description")
}

func (a *Location) SetDescription(description string) {
	a.Set("description", description)
}

// AddressAddon for many BSM datasets contains the actual name of the field
func (a *Location) AddressAddon() string {
	return a.GetString("address_addon")
}

func (a *Location) SetAddressAddon(addressAddon string) {
	a.Set("address_addon", addressAddon)
}

func (a *Location) Street() string {
	return a.GetString("street")
}

func (a *Location) SetStreet(street string) {
	a.Set("street", street)
}

func (a *Location) PostalCode() string {
	return a.GetString("postal_code")
}

func (a *Location) SetPostalCode(postalCode string) {
	a.Set("postal_code", postalCode)
}

func (a *Location) Country() string {
	return a.GetString("country")
}

func (a *Location) SetCountry(country string) {
	a.Set("country", country)
}

func (a *Location) City() string {
	return a.GetString("city")
}

func (a *Location) SetCity(city string) {
	a.Set("city", city)
}

func (a *Location) Longitude() float64 {
	return a.GetFloat("longitude")
}

func (a *Location) SetLongitude(longitude float64) {
	a.Set("longitude", longitude)
}

func (a *Location) Latitude() float64 {
	return a.GetFloat("latitude")
}

func (a *Location) SetLatitude(latitude float64) {
	a.Set("latitude", latitude)
}

func (a *Location) SpectatorTotal() int {
	return a.GetInt("spectator_total")
}

func (a *Location) SetSpectatorTotal(spectatorTotal int) {
	a.Set("spectator_total", spectatorTotal)
}

func (a *Location) SpectatorSeats() int {
	return a.GetInt("spectator_seats")
}

func (a *Location) SetSpectatorSeats(spectatorSeats int) {
	a.Set("spectator_seats", spectatorSeats)
}

func (a *Location) OtherInformation() string {
	return a.GetString("other_information")
}

func (a *Location) SetOtherInformation(otherInformation string) {
	a.Set("other_information", otherInformation)
}

func (a *Location) GroundRules() string {
	return a.GetString("groundrules")
}

func (a *Location) SetGroundRules(groundRules string) {
	a.Set("groundrules", groundRules)
}

func (a *Location) HumanCountry() string {
	return a.GetString("human_country")
}

func (a *Location) SetHumanCountry(humanCountry string) {
	a.Set("human_country", humanCountry)
}

func (a *Location) PhotoURL() string {
	return a.GetString("photo_url")
}

func (a *Location) SetPhotoURL(photoURL string) {
	a.Set("photo_url", photoURL)
}

func (a *Location) Club() string {
	return a.GetString("club")
}

func (a *Location) SetClub(club string) {
	a.Set("club", club)
}
