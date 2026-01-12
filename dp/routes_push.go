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

func deletePushSubscription(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		info, err := e.RequestInfo()
		if err != nil {
			return e.InternalServerError("", err)
		}
		endpoint := info.Body["endpoint"].(string)

		record, err := app.FindFirstRecordByFilter(
			PushSubscriptionsCollection,
			"endpoint = {:endpoint}",
			dbx.Params{"endpoint": endpoint},
		)
		if err != nil {
			app.Logger().Error("failed to find subscription to delete", "error", err, "endpoint", endpoint)
			return e.BadRequestError("", err)
		}
		var sub PushSubscription
		sub.SetProxyRecord(record)

		if sub.User() != e.Auth.Id {
			return e.ForbiddenError("", nil)
		}

		err = app.Delete(sub)
		if err != nil {
			app.Logger().Error("failed to delete subscription", "error", err, "sub", sub)
			return e.InternalServerError("", err)
		}

		return e.NoContent(http.StatusNoContent)
	}
}
