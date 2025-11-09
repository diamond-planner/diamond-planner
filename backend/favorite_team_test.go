package main

import (
	"net/http"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestFavoriteTeamEndpoint(t *testing.T) {
	scenarios := []tests.ApiScenario{
		{
			Name:            "try with different http method, e.g. POST",
			Method:          http.MethodPost,
			URL:             "/api/team/favorite",
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"wasn't found"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "missing team parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?season=2025",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Team ID not found in request"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "invalid team parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=invalid&season=2025",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Team ID must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "missing season parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Season not found in request"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "invalid season parameter",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219&season=invalid",
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Season must be an integer"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:            "valid request",
			Method:          http.MethodGet,
			URL:             "/api/team/favorite?team=23219&season=2025",
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"league_group", "playoff_series", "table", "table_row"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
