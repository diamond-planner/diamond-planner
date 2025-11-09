package bsm

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestRealAPIClient_HTTPMethods(t *testing.T) {
	apiKey := "test-key"

	// Fixtures
	matches := []Match{{MatchID: "M1", HomeTeamName: "H", AwayTeamName: "A"}}
	table := Table{LeagueID: 42, LeagueName: "Test League", Season: 2025}
	team := Team{ID: 7, Name: "Skylarks", Season: 2025, ShortName: "SKY"}
	leagues := []LeagueGroup{{ID: 5, Season: 2025, Name: "LL", Acronym: "LL", League: League{ID: 9, Name: "L", Acronym: "L", Season: 2025}}}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// All endpoints must include api_key
		q := r.URL.Query()
		if q.Get("api_key") != apiKey {
			t.Fatalf("expected api_key=%s, got %s", apiKey, q.Get("api_key"))
		}

		switch {
		case r.URL.Path == "/matches.json":
			// expect some optional params are passed through unchanged
			_ = q.Get(SeasonFilter) // allowed to be empty in this test
			_ = q.Get(CompactFilter)
			_ = json.NewEncoder(w).Encode(matches)
			return

		case strings.HasPrefix(r.URL.Path, "/leagues/") && strings.HasSuffix(r.URL.Path, "/table.json"):
			_ = json.NewEncoder(w).Encode(table)
			return

		case strings.HasPrefix(r.URL.Path, "/teams/") && strings.HasSuffix(r.URL.Path, ".json"):
			_ = json.NewEncoder(w).Encode(team)
			return

		case r.URL.Path == "/league_groups.json":
			// check season filter present
			if q.Get(SeasonFilter) == "" {
				t.Fatalf("expected %s query param to be set", SeasonFilter)
			}
			_ = json.NewEncoder(w).Encode(leagues)
			return
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	client := &RealAPIClient{BaseAPIURL: server.URL}

	// Test LoadMatchesWithFilterParams
	params := map[string]string{SeasonFilter: strconv.Itoa(2025), CompactFilter: "true"}
	gotMatches, err := client.LoadMatchesWithFilterParams(params, apiKey)
	if err != nil {
		t.Fatalf("LoadMatchesWithFilterParams error: %v", err)
	}
	if len(gotMatches) != 1 || gotMatches[0].MatchID != "M1" {
		t.Fatalf("unexpected matches response: %+v", gotMatches)
	}

	// Test LoadSingleTable
	gotTable, err := client.LoadSingleTable(42, apiKey)
	if err != nil {
		t.Fatalf("LoadSingleTable error: %v", err)
	}
	if gotTable.LeagueID != 42 || gotTable.LeagueName != "Test League" {
		t.Fatalf("unexpected table: %+v", gotTable)
	}

	// Test LoadSingleTeamByID
	gotTeam, err := client.LoadSingleTeamByID(7, apiKey)
	if err != nil {
		t.Fatalf("LoadSingleTeamByID error: %v", err)
	}
	if gotTeam.ID != 7 || gotTeam.ShortName != "SKY" {
		t.Fatalf("unexpected team: %+v", gotTeam)
	}

	// Test FetchLeagueGroupsForSeason
	gotLeagues, err := client.FetchLeagueGroupsForSeason(apiKey, 2025)
	if err != nil {
		t.Fatalf("FetchLeagueGroupsForSeason error: %v", err)
	}
	if len(gotLeagues) != 1 || gotLeagues[0].ID != 5 {
		t.Fatalf("unexpected leagues: %+v", gotLeagues)
	}
}

func TestRealAPIClient_GetAPIURL_EncodesParams(t *testing.T) {
	c := &RealAPIClient{BaseAPIURL: "https://example.test"}
	params := map[string]string{"a": "1", "b": "two"}
	u := c.GetAPIURL("resource.json", params, "k")
	q := u.Query()
	if q.Get("api_key") != "k" || q.Get("a") != "1" || q.Get("b") != "two" {
		t.Fatalf("unexpected query: %s", u.RawQuery)
	}
	if u.Scheme != "https" || u.Host != "example.test" || u.Path != "/resource.json" {
		t.Fatalf("unexpected URL: %s", u.String())
	}
}
