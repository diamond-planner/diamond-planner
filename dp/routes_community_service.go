package dp

import (
	"errors"
	"net/http"
	"slices"
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

func GetUserCommunityService() func(event *core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		authUser := &User{}
		authUser.SetProxyRecord(event.Auth)

		userID := event.Request.PathValue("user")
		userRecord, err := event.App.FindRecordById(UserCollection, userID)
		if err != nil {
			return event.NotFoundError("invalid user provided", err)
		}
		user := &User{}
		user.SetProxyRecord(userRecord)

		requireUserAccess := RequireUserAccess(user.Record)
		if err := requireUserAccess.Func(event); err != nil {
			return event.ForbiddenError("no access", err)
		}

		var season int
		seasonParam := event.Request.PathValue("season")
		if seasonParam != "" {
			season, err = strconv.Atoi(seasonParam)
			if err != nil {
				event.App.Logger().Error("failed to parse season: %v", "error", err)
				return err
			}
		} else {
			return event.BadRequestError("season parameter is required", errors.New("season parameter is required"))
		}

		var readableClubs []Club
		clubRecords, err := event.App.FindRecordsByIds(ClubsCollection, user.Club())
		for _, clubRecord := range clubRecords {
			club := Club{}
			club.SetProxyRecord(clubRecord)
			if slices.Contains(club.Admins(), authUser.Id) || authUser.Id == user.Id {
				readableClubs = append(readableClubs, club)
			} else {
				event.App.Logger().Debug("user does not have access to club", "clubID", clubRecord.Id, "user", user.Id)
			}
		}

		serviceData, err := getCommunityServiceDataForUser(event.App, event.Auth.Id, *user, season, readableClubs)
		if err != nil {
			return event.InternalServerError("", err)
		}

		return event.JSON(http.StatusOK, serviceData)
	}
}
