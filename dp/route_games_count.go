package dp

import (
	"net/http"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type SimpleCount struct {
	Count int `json:"count" db:"count"`
}

// GetGamesCount returns the number of official league games in the database for a given team and season.
// Requires a valid auth record that belongs to the team queried.
func GetGamesCount(app core.App) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		// verify access status
		requireAuth := apis.RequireAuth()
		if err := requireAuth.Func(e); err != nil {
			return err
		}

		teamID := e.Request.PathValue("team")
		team, err := app.FindRecordById("teams", teamID)
		if err != nil {
			return e.BadRequestError("Failed to load teams record", err)
		}

		info, err := e.RequestInfo()
		if err != nil {
			return e.BadRequestError("Failed to retrieve request info", err)
		}

		canAccess, err := e.App.CanAccessRecord(team, info, team.Collection().ViewRule)
		if !canAccess {
			return e.ForbiddenError("Only members and admins can query team game count", err)
		}

		// do actual work
		season := e.Request.URL.Query().Get("season")

		if season == "" {
			return apis.NewBadRequestError("Query param for season is required", nil)
		}

		result, err := LoadCount(app, teamID, season)
		if err != nil {
			return apis.NewInternalServerError("Error querying event", err)
		}

		return e.JSON(http.StatusOK, result)
	}
}

func LoadCount(app core.App, teamID string, season string) (SimpleCount, error) {
	result := SimpleCount{}
	err := app.DB().
		Select("COUNT(events.id) AS count").
		From("events").
		Where(dbx.NewExp("bsm_id != 0")).
		AndWhere(dbx.NewExp("team = {:team}", dbx.Params{"team": teamID})).
		AndWhere(dbx.NewExp("strftime('%Y', events.starttime) = {:season}", dbx.Params{"season": season})).
		One(&result)

	if err != nil {
		return result, err
	}
	return result, nil
}
