package dp

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type ParticipationStatsByPerson struct {
	ID         string `db:"id" json:"id"`
	LastName   string `db:"last_name" json:"lastName"`
	FirstName  string `db:"first_name" json:"firstName"`
	Type       string `db:"type" json:"type"`
	InCount    int    `db:"inCount" json:"inCount"`
	OutCount   int    `db:"outCount" json:"outCount"`
	MaybeCount int    `db:"maybeCount" json:"maybeCount"`
	TotalCount int    `db:"TotalCount" json:"totalCount"`
}

func GetParticipationStats(user *core.Record, app core.App, seasonParam string, eventTypeParam string, team string) ([]ParticipationStatsByPerson, error) {
	var participations []ParticipationStatsByPerson

	query := app.DB().
		Select(
			"users.id",
			"users.last_name",
			"users.first_name",
			"events.type",
			"COUNT(CASE WHEN participations.state = 'in' THEN participations.id END)    AS inCount",
			"COUNT(CASE WHEN participations.state = 'out' THEN participations.id END)   AS outCount",
			"COUNT(CASE WHEN participations.state = 'maybe' THEN participations.id END) AS maybeCount",
			"COUNT(participations.id)",
		).
		From("users").
		LeftJoin("participations", dbx.NewExp("participations.user = users.id")).
		InnerJoin("events", dbx.NewExp("participations.event = events.id")).
		Where(dbx.NewExp("events.starttime <= datetime('now')")). // participation stats make no sense for future events
		GroupBy("users.id", "users.first_name", "users.last_name", "events.type")

	if user.Id != "" {
		query.AndWhere(dbx.NewExp("users.id = {:id}", dbx.Params{"id": user.Id}))
	}

	if seasonParam != "" {
		query.AndWhere(dbx.NewExp("strftime('%Y', events.starttime) = {:season}", dbx.Params{"season": seasonParam}))
	}

	if eventTypeParam != "" {
		query.AndWhere(dbx.NewExp("events.type = {:eventType}", dbx.Params{"eventType": eventTypeParam}))
	}

	if team != "" {
		query.AndWhere(dbx.NewExp("events.team = {:team}", dbx.Params{"team": team}))
	}

	err := query.All(&participations)
	if err != nil {
		app.Logger().Error("failed to load participations: %v", "error", err)
		return participations, err
	}
	return participations, nil
}
