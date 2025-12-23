package dp

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func GetUserStats() func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		userID := event.Request.PathValue("user")
		user, err := event.App.FindRecordById(UserCollection, userID)
		if err != nil {
			return event.NotFoundError("invalid user provided", err)
		}

		requireUserAccess := RequireUserAccess(user)
		if err := requireUserAccess.Func(event); err != nil {
			return event.ForbiddenError("no access", err)
		}

		statsItem, err := LoadUserStats(user, event)
		if err != nil {
			return event.InternalServerError("failed to load user stat", err)
		}

		return event.JSON(http.StatusOK, statsItem)
	}
}
