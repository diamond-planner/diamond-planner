package dp

import (
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

const (
	EventSeriesCollection = "eventseries"
)

var _ core.RecordProxy = (*EventSeries)(nil)

// EventSeries is a RecordProxy for the `eventseries` collection.
// It offers type-safe getters and setters for every field.
type EventSeries struct {
	core.BaseRecordProxy
}

// ID returns the record id.
func (s *EventSeries) ID() string {
	return s.Id
}

// Title returns the series title.
func (s *EventSeries) Title() string {
	return s.GetString("title")
}
func (s *EventSeries) SetTitle(title string) {
	s.Set("title", title)
}

// Team returns the id of the team the series belongs to.
func (s *EventSeries) Team() string {
	return s.GetString("team")
}
func (s *EventSeries) SetTeam(team string) {
	s.Set("team", team)
}

// Desc returns the optional description.
func (s *EventSeries) Desc() string {
	return s.GetString("desc")
}
func (s *EventSeries) SetDesc(desc string) {
	s.Set("desc", desc)
}

// Location returns the location string.
func (s *EventSeries) Location() string {
	return s.GetString("location")
}
func (s *EventSeries) SetLocation(location string) {
	s.Set("location", location)
}

// SeriesStart returns the first day/time of the series.
func (s *EventSeries) SeriesStart() types.DateTime {
	return s.GetDateTime("series_start")
}
func (s *EventSeries) SetSeriesStart(start types.DateTime) {
	s.Set("series_start", start)
}

// SeriesEnd returns the last day/time of the series.
func (s *EventSeries) SeriesEnd() types.DateTime {
	return s.GetDateTime("series_end")
}
func (s *EventSeries) SetSeriesEnd(end types.DateTime) {
	s.Set("series_end", end)
}

// Interval returns the number of days between occurrences.
func (s *EventSeries) Interval() int {
	return s.GetInt("interval")
}
func (s *EventSeries) SetInterval(interval int) {
	s.Set("interval", interval)
}

// Duration returns the number of days between occurrences.
func (s *EventSeries) Duration() int {
	return s.GetInt("duration")
}
func (s *EventSeries) SetDuration(duration int) {
	s.Set("duration", duration)
}
