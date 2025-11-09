package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/spf13/cast"
)

// OAuthUpdateUserData updates user record with information found in the OAuth2 response.
// This could also be used to validate the signup key, but does not do that at the moment
// (see generic record creation hook that also works for password auth)
func OAuthUpdateUserData(e *core.RecordAuthWithOAuth2RequestEvent) error {
	if e.Record == nil {
		return e.Next()
	}
	user := &User{}
	user.SetProxyRecord(e.Record)

	rawUser := e.OAuth2User.RawUser
	firstName, firstNameExists := rawUser["given_name"]
	lastName, lastNameExists := rawUser["family_name"]

	if firstNameExists {
		user.SetFirstName(cast.ToString(firstName))
	}
	if lastNameExists {
		user.SetLastName(cast.ToString(lastName))
	}

	err := e.App.Save(user)
	if err != nil {
		e.App.Logger().Warn("Could not persist new OAuth2 user data", "rawUser", rawUser, "record", user.ID())
		return e.Next()
	}
	return e.Next()
}
