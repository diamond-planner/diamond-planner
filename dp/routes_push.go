package dp

import (
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func notifyWithTestPushMessage(ps PushService, app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		userID := e.Request.PathValue("user")

		if userID != e.Auth.Id {
			return e.ForbiddenError("", nil)
		}

		subRecords, err := app.FindRecordsByFilter(
			PushSubscriptionsCollection,
			"user.id = {:userID}",
			"",
			0,
			0,
			dbx.Params{"userID": userID},
		)
		if err != nil {
			return err
		}

		for _, subRecord := range subRecords {
			sub := &PushSubscription{}
			sub.SetProxyRecord(subRecord)
			ws := sub.ToWebPushSubscription()

			err := ps.handleTestPush(&ws)
			if err != nil {
				app.Logger().Warn("failed to send push message", "error", err, "sub", sub)
				continue
			}
		}

		return e.JSON(http.StatusOK, "Test push sent")
	}
}
