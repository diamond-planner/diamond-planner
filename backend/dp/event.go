package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	EventsCollection = "events"
)

var _ core.RecordProxy = (*Event)(nil)

// Event RecordProxy for collection `events`.
// Provides type-safe struct access to event records.
type Event struct {
	core.BaseRecordProxy
}

func (e *Event) ID() string {
	return e.Id
}

func (e *Event) BSMID() int {
	return e.GetInt("bsm_id")
}

func (e *Event) SetBSMID(bsmID int) {
	e.Set("bsm_id", bsmID)
}

func (e *Event) Title() string {
	return e.GetString("title")
}

func (e *Event) SetTitle(title string) {
	e.Set("title", title)
}

func (e *Event) Desc() string {
	return e.GetString("desc")
}

func (e *Event) SetDesc(desc string) {
	e.Set("desc", desc)
}

func (e *Event) StartTime() types.DateTime {
	return e.GetDateTime("starttime")
}

func (e *Event) SetStartTime(startTime types.DateTime) {
	e.Set("starttime", startTime)
}

func (e *Event) EndTime() types.DateTime {
	return e.GetDateTime("endtime")
}

func (e *Event) SetEndTime(endTime types.DateTime) {
	e.Set("endtime", endTime)
}

func (e *Event) MeetingTime() types.DateTime {
	return e.GetDateTime("meetingtime")
}

func (e *Event) SetMeetingTime(meetingTime types.DateTime) {
	e.Set("meetingtime", meetingTime)
}

func (e *Event) Type() string {
	return e.GetString("type")
}

func (e *Event) SetType(eventType string) {
	e.Set("type", eventType)
}

func (e *Event) Team() string {
	return e.GetString("team")
}

func (e *Event) SetTeam(team string) {
	e.Set("team", team)
}

func (e *Event) Location() string {
	return e.GetString("location")
}

func (e *Event) SetLocation(location string) {
	e.Set("location", location)
}

func (e *Event) MatchJSON() string {
	return e.GetString("match_json")
}

func (e *Event) SetMatchJSON(matchJSON string) {
	e.Set("match_json", matchJSON)
}

func (e *Event) Guests() string {
	return e.GetString("guests")
}

func (e *Event) SetGuests(guests string) {
	e.Set("guests", guests)
}

func (e *Event) IsCancelled() bool {
	return e.GetBool("cancelled")
}

func (e *Event) SetCancelled(cancelled bool) {
	e.Set("cancelled", cancelled)
}

func (e *Event) Attire() string {
	return e.GetString("attire")
}

func (e *Event) SetAttire(attire string) {
	e.Set("attire", attire)
}

func (e *Event) Series() string {
	return e.GetString("series")
}

func (e *Event) SetSeries(series string) {
	e.Set("series", series)
}
