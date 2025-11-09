package dp

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type EventCounts struct {
	Games     int `db:"games"`
	Practices int `db:"practices"`
	Misc      int `db:"misc"`
	AllEvents int `db:"all_events"`
}

func GetEventCounts(app core.App, user *core.Record, season string, team string, excludeFutureEvents bool) (EventCounts, error) {
	result := EventCounts{}
	query := app.DB().
		Select(
			"COUNT(CASE WHEN events.type = 'game' THEN events.id END) AS games",
			"COUNT(CASE WHEN events.type = 'practice' THEN events.id END) AS practices",
			"COUNT(CASE WHEN events.type = 'misc' THEN events.id END) AS misc",
			"COUNT(events.id) AS all_events",
		).
		From("events")

	if user != nil {
		teams := user.GetStringSlice("teams")

		if len(teams) > 0 {
			var expressions []dbx.Expression
			for _, team := range teams {
				expressions = append(expressions, dbx.HashExp{"team": team})
			}
			query.AndWhere(dbx.Or(expressions...))
		} else {
			// We have received a user as a parameter, but this user does not have any teams.
			// That should not happen in normal usage, but catch this edge case nonetheless.
			// The result of the function should be zero events,
			// since a user without teams does not have any events they could possibly attend.
			// So: add impossible restriction.
			query.AndWhere(dbx.NewExp("events.team = {:team}", dbx.Params{"team": "TOTALLY_NONEXISTENT_TEAM"}))
		}
	}

	if season != "" {
		query.AndWhere(dbx.NewExp("strftime('%Y', events.starttime) = {:season}", dbx.Params{"season": season}))
	}

	if team != "" {
		query.AndWhere(dbx.NewExp("events.team = {:team}", dbx.Params{"team": team}))
	}

	// i.e., count only events that have not yet started for attendance stats
	if excludeFutureEvents {
		query.AndWhere(dbx.NewExp("events.starttime <= datetime('now')"))
	}

	err := query.One(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
