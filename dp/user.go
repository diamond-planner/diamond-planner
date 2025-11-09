package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	UserCollection = "users"
)

var _ core.RecordProxy = (*User)(nil)

// User RecordProxy for collection `users`.
// Provides type-safe struct access to user records.
type User struct {
	core.BaseRecordProxy
}

func (u *User) ID() string {
	return u.Id
}

func (u *User) Username() string {
	return u.GetString("username")
}

func (u *User) SetUsername(username string) {
	u.Set("username", username)
}

func (u *User) Email() string {
	return u.GetString("email")
}

func (u *User) SetEmail(email string) {
	u.Set("email", email)
}

func (u *User) EmailVisibility() bool {
	return u.GetBool("emailVisibility")
}

func (u *User) SetEmailVisibility(visible bool) {
	u.Set("emailVisibility", visible)
}

func (u *User) DisplayOnWebsite() bool {
	return u.GetBool("display_on_website")
}

func (u *User) SetDisplayOnWebsite(display bool) {
	u.Set("display_on_website", display)
}

func (u *User) Verified() bool {
	return u.GetBool("verified")
}

func (u *User) SetVerified(v bool) {
	u.Set("verified", v)
}

func (u *User) FirstName() string {
	return u.GetString("first_name")
}

func (u *User) SetFirstName(first string) {
	u.Set("first_name", first)
}

func (u *User) LastName() string {
	return u.GetString("last_name")
}

func (u *User) SetLastName(last string) {
	u.Set("last_name", last)
}

func (u *User) Avatar() string {
	return u.GetString("avatar")
}

func (u *User) SetAvatar(filename string) {
	u.Set("avatar", filename)
}

func (u *User) Image() string {
	return u.GetString("image")
}

func (u *User) SetImage(filename string) {
	u.Set("image", filename)
}

func (u *User) Number() string {
	return u.GetString("number")
}

func (u *User) SetNumber(num string) {
	u.Set("number", num)
}

func (u *User) Teams() []string {
	return u.GetStringSlice("teams")
}

func (u *User) SetTeams(teamIDs []string) {
	u.Set("teams", teamIDs)
}

func (u *User) Club() []string {
	return u.GetStringSlice("club")
}

func (u *User) SetClub(clubIDs []string) {
	u.Set("club", clubIDs)
}

func (u *User) LastLogin() types.DateTime {
	return u.GetDateTime("last_login")
}

func (u *User) SetLastLogin(dt types.DateTime) {
	u.Set("last_login", dt)
}

func (u *User) Birthday() types.DateTime {
	return u.GetDateTime("birthday")
}

func (u *User) SetBirthday(dt types.DateTime) {
	u.Set("birthday", dt)
}

func (u *User) Position() []string {
	return u.GetStringSlice("position")
}

func (u *User) SetPosition(pos []string) {
	u.Set("position", pos)
}

func (u *User) Bats() string {
	return u.GetString("bats")
}

func (u *User) SetBats(bats string) {
	u.Set("bats", bats)
}

func (u *User) Throws() string {
	return u.GetString("throws")
}

func (u *User) SetThrows(throws string) {
	u.Set("throws", throws)
}

func (u *User) Umpire() string {
	return u.GetString("umpire")
}

func (u *User) SetUmpire(grade string) {
	u.Set("umpire", grade)
}

func (u *User) Scorer() string {
	return u.GetString("scorer")
}

func (u *User) SetScorer(grade string) {
	u.Set("scorer", grade)
}

func (u *User) BSMID() int {
	return u.GetInt("bsm_id")
}

func (u *User) SetBSMID(id int) {
	u.Set("bsm_id", id)
}

func (u *User) SignupKey() string {
	return u.GetString("signup_key")
}

func (u *User) SetSignupKey(key string) {
	u.Set("signup_key", key)
}
