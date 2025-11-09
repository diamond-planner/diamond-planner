package dp

import (
	"fmt"
	"os"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

func CreateOrUpdateEventsForSeries(e *core.RecordEvent) error {
	events, err := generateSeriesEvents(e.App, e.Record)
	if err != nil {
		return err
	}

	err = e.App.RunInTransaction(func(txApp core.App) error {
		for _, record := range events {
			if err := txApp.Save(record); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return e.Next()
}

func DeleteEventsForSeries(e *core.RecordEvent) error {
	eventSeries := e.Record

	eventsToBeDeleted, err := e.App.FindRecordsByFilter(
		EventsCollection,
		"series = {:seriesID}",
		"",
		0,
		0,
		dbx.Params{"seriesID": eventSeries.Id},
	)
	if err != nil {
		return err
	}

	err = e.App.RunInTransaction(func(txApp core.App) error {
		for _, record := range eventsToBeDeleted {
			if err := txApp.Delete(record); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return e.Next()
}

func generateSeriesEvents(app core.App, record *core.Record) ([]*core.Record, error) {
	eventSeries := &EventSeries{}
	eventSeries.SetProxyRecord(record)

	startDate := eventSeries.SeriesStart()
	endDate := eventSeries.SeriesEnd()
	interval := eventSeries.Interval()
	duration := eventSeries.Duration()

	eventCollection, err := app.FindCollectionByNameOrId(EventsCollection)
	if err != nil {
		return nil, err
	}

	existingEvents, err := app.FindRecordsByFilter(
		EventsCollection,
		"series = {:seriesID}",
		"",
		0,
		0,
		dbx.Params{"seriesID": eventSeries.Id},
	)
	if err != nil {
		return nil, err
	}

	var events []*core.Record
	existingEventsMap := make(map[string]*Event)

	timezone := os.Getenv("TIME_ZONE")
	if timezone == "" {
		timezone = "Europe/Berlin"
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}

	// Create a map of existing events for easy lookup
	for _, event := range existingEvents {
		eventProxy := &Event{}
		eventProxy.SetProxyRecord(event)

		key := fmt.Sprintf("%s---%s", eventProxy.StartTime().Time().In(location).Format(time.RFC3339), eventProxy.EndTime().Time().In(location).Format(time.RFC3339))
		existingEventsMap[key] = eventProxy
	}

	// reads timezone information to ensure that the later call to `AddDate()` accounts for daylight savings time traversal
	currentDate, err := types.ParseDateTime(startDate.Time().In(location))
	if err != nil {
		return nil, err
	}

	for currentDate.Before(endDate) {
		// Create start and end times for this specific event
		eventStart := currentDate
		eventEnd := currentDate.Add(time.Duration(duration) * time.Minute)

		// Check if an event already exists for this time slot
		key := fmt.Sprintf("%s---%s", eventStart.Time().Format(time.RFC3339), eventEnd.Time().Format(time.RFC3339))
		event, exists := existingEventsMap[key]

		if !exists {
			// Not found - create a new event
			event = &Event{}
			event.SetProxyRecord(core.NewRecord(eventCollection))
		}

		event.SetStartTime(eventStart)
		event.SetEndTime(eventEnd)
		event.SetTitle(eventSeries.Title())
		event.SetTeam(eventSeries.Team())
		event.SetDesc(eventSeries.Desc())
		event.SetLocation(eventSeries.Location())
		event.SetSeries(eventSeries.Id)
		event.SetType(Practice.String())

		events = append(events, event.Record)

		// Mark this event as processed
		delete(existingEventsMap, key)

		currentDate = currentDate.AddDate(0, 0, interval)
	}

	// Delete any leftover events that are no longer part of the updated series
	for _, staleEvent := range existingEventsMap {
		err := app.Delete(staleEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to delete stale event: %w", err)
		}
	}

	return events, nil
}
