package dp

import (
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

// LoadUserStats gets participation stats by event type and participation state
func LoadUserStats(user *core.Record, requestEvent *core.RequestEvent) (PersonalAttendanceStatsItem, error) {
	statsItem := PersonalAttendanceStatsItem{}
	seasonParam := requestEvent.Request.URL.Query().Get("season")
	eventTypeParam := requestEvent.Request.URL.Query().Get("eventType")
	teamParam := requestEvent.Request.URL.Query().Get("team")

	if teamParam != "" {
		team, err := requestEvent.App.FindRecordById("teams", teamParam, nil)
		if err != nil {
			return statsItem, err
		}
		statsItem.TeamName = team.GetString("name")
	} else {
		statsItem.TeamName = "All Teams"
	}

	if seasonParam != "" {
		season, err := strconv.Atoi(seasonParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse season: %v", "error", err)
			return statsItem, err
		}
		statsItem.Season = season
	}

	if eventTypeParam != "" {
		eventType, err := ParseEventType(eventTypeParam)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse eventType: %v", "error", err)
			return statsItem, err
		}
		statsItem.Type = eventType
	}

	eventCounts, err := GetEventCounts(requestEvent.App, user, seasonParam, teamParam, true)
	if err != nil {
		requestEvent.App.Logger().Error("failed to get eventCounts: %v", "error", err)
		return statsItem, err
	}
	statsItem.TotalPossibleEvents = eventCounts.AllEvents

	// data structure for frontend
	totalsIn := ParticipationTotal{
		Type:   In,
		Amount: 0,
	}
	totalsMaybe := ParticipationTotal{
		Type:   Maybe,
		Amount: 0,
	}
	totalsOut := ParticipationTotal{
		Type:   Out,
		Amount: 0,
	}

	totalsGames := AttendanceTotal{
		Type:     Game,
		Attended: 0,
		Total:    eventCounts.Games,
	}
	totalsPractice := AttendanceTotal{
		Type:     Practice,
		Attended: 0,
		Total:    eventCounts.Practices,
	}
	totalsMisc := AttendanceTotal{
		Type:     Misc,
		Attended: 0,
		Total:    eventCounts.Misc,
	}

	participations, err := GetParticipationStats(user, requestEvent.App, seasonParam, eventTypeParam, teamParam)
	if err != nil {
		return statsItem, err
	}
	// query can only get the totals for which participation data exists,
	// add raw possible totals from the previous query
	for i, participation := range participations {
		eventType, err := ParseEventType(participation.Type)
		if err != nil {
			requestEvent.App.Logger().Error("failed to parse eventType: %v", "error", err)
			return statsItem, err
		}
		switch eventType {
		case Game:
			participations[i].TotalCount = eventCounts.Games
			totalsGames.Attended = participation.InCount
		case Practice:
			participations[i].TotalCount = eventCounts.Practices
			totalsPractice.Attended = participation.InCount
		case Misc:
			participations[i].TotalCount = eventCounts.Misc
			totalsMisc.Attended = participation.InCount
		}

		totalsIn.Amount = totalsIn.Amount + participation.InCount
		totalsMaybe.Amount = totalsMaybe.Amount + participation.MaybeCount
		totalsOut.Amount = totalsOut.Amount + participation.OutCount
	}

	statsItem.ParticipationTotals = append(statsItem.ParticipationTotals, totalsIn, totalsMaybe, totalsOut)
	statsItem.AttendanceTotals = append(statsItem.AttendanceTotals, totalsPractice, totalsMisc, totalsGames)
	statsItem.Values = append(statsItem.Values, participations...)

	return statsItem, nil
}
