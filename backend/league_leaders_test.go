package main

import (
	"net/http"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestGetLeagueLeadersEndpoint(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "Invalid league ID",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/abc",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"league ID must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Valid request with default statsType",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/5732",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_id", "stats_type", "data"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "Valid request with specified statsType",
			Method:          http.MethodGet,
			URL:             "/api/bsm/relay/top10/5732?statsType=pitching",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_id", "stats_type", "data"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
