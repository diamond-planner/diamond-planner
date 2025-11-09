package dp

import (
	"testing"
	"time"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tools/types"
)

// helper that must never fail inside the test body
func mustParseDateTime(t *testing.T, tm time.Time) types.DateTime {
	t.Helper()

	parsed, err := types.ParseDateTime(tm)
	if err != nil {
		t.Fatalf("failed to parse time %v: %v", tm, err)
	}

	return parsed
}

// registers an empty `events` collection so that PocketBase queries don't fail
func prepareEventsCollection(t *testing.T, app *tests.TestApp) {
	t.Helper()

	coll := core.NewBaseCollection(EventsCollection)

	coll.Fields.Add(&core.TextField{Name: "title"})
	coll.Fields.Add(&core.TextField{Name: "desc"})
	coll.Fields.Add(&core.TextField{Name: "type"})
	coll.Fields.Add(&core.DateField{Name: "starttime"})
	coll.Fields.Add(&core.DateField{Name: "endtime"})

	// not text fields but a database query should get them anyway
	coll.Fields.Add(&core.TextField{Name: "team"})
	coll.Fields.Add(&core.TextField{Name: "location"})
	coll.Fields.Add(&core.TextField{Name: "series"})

	if err := app.Save(coll); err != nil {
		t.Fatalf("failed to register events collection: %v", err)
	}
}

// returns a fresh in-memory `EventSeries` record already populated with the
// arguments provided.
func buildSeriesRecord(
	t *testing.T,
	title string,
	start, end time.Time,
	intervalDays int,
	durationMinutes int,
) *EventSeries {
	t.Helper()

	seriesColl := core.NewBaseCollection(EventSeriesCollection)
	record := core.NewRecord(seriesColl)
	series := &EventSeries{}
	series.SetProxyRecord(record)

	series.SetTitle(title)
	series.SetSeriesStart(mustParseDateTime(t, start))
	series.SetSeriesEnd(mustParseDateTime(t, end))
	series.SetInterval(intervalDays)
	series.SetDuration(durationMinutes)
	series.SetTeam("dummy-team")
	series.SetLocation("dummy-location")

	return series
}

// ---------------------------------------------------------------------------
// TESTS
// ---------------------------------------------------------------------------

func TestGenerateSeriesEvents_WeeklyAcrossSpringDST(t *testing.T) {
	app, err := tests.NewTestApp()
	if err != nil {
		t.Fatal(err)
	}
	defer app.Cleanup()

	prepareEventsCollection(t, app)

	loc, _ := time.LoadLocation("Europe/Berlin")

	start := time.Date(2024, time.March, 25, 18, 0, 0, 0, loc) // one week before DST start
	end := time.Date(2024, time.April, 15, 18, 0, 0, 0, loc)   // three-week window

	series := buildSeriesRecord(t, "Spring practice", start, end, 7, 90)

	events, genErr := generateSeriesEvents(app, series.Record)
	if genErr != nil {
		t.Fatalf("generateSeriesEvents failed: %v", genErr)
	}

	expectedCount := 3 // 25-Mar, 01-Apr, 08-Apr (loop stops before 15-Apr)
	if len(events) != expectedCount {
		t.Fatalf("expected %d events, got %d", expectedCount, len(events))
	}

	for i, ev := range events {
		st := ev.GetDateTime("starttime").Time().In(loc)

		if st.Hour() != 18 || st.Minute() != 0 {
			t.Errorf("event %d has wrong local time: %v (want 18:00)", i, st)
		}
	}
}

func TestGenerateSeriesEvents_WeeklyAcrossAutumnDST(t *testing.T) {
	app, err := tests.NewTestApp()
	if err != nil {
		t.Fatal(err)
	}
	defer app.Cleanup()

	prepareEventsCollection(t, app)

	loc, _ := time.LoadLocation("Europe/Berlin")

	start := time.Date(2024, time.October, 14, 18, 0, 0, 0, loc) // before the end-of-DST change
	end := time.Date(2024, time.November, 4, 18, 0, 0, 0, loc)

	series := buildSeriesRecord(t, "Autumn practice", start, end, 7, 60)

	events, genErr := generateSeriesEvents(app, series.Record)
	if genErr != nil {
		t.Fatalf("generateSeriesEvents failed: %v", genErr)
	}

	expectedCount := 3 // 14-Oct, 21-Oct, 28-Oct
	if len(events) != expectedCount {
		t.Fatalf("expected %d events, got %d", expectedCount, len(events))
	}

	for i, ev := range events {
		st := ev.GetDateTime("starttime").Time().In(loc)
		if st.Hour() != 18 || st.Minute() != 0 {
			t.Errorf("event %d has wrong local time: %v (want 18:00)", i, st)
		}
	}
}

func TestGenerateSeriesEvents_DailyShortInterval(t *testing.T) {
	app, err := tests.NewTestApp()
	if err != nil {
		t.Fatal(err)
	}
	defer app.Cleanup()

	prepareEventsCollection(t, app)

	loc, _ := time.LoadLocation("Europe/Berlin")

	start := time.Date(2024, time.July, 1, 10, 0, 0, 0, loc)
	end := time.Date(2024, time.July, 5, 10, 0, 0, 0, loc) // 4-day window

	series := buildSeriesRecord(t, "Intensive camp", start, end, 1, 120)

	events, genErr := generateSeriesEvents(app, series.Record)
	if genErr != nil {
		t.Fatalf("generateSeriesEvents failed: %v", genErr)
	}

	expectedCount := 4 // 01-Jul, 02-Jul, 03-Jul, 04-Jul
	if len(events) != expectedCount {
		t.Fatalf("expected %d events, got %d", expectedCount, len(events))
	}
}
