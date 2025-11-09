package main

import (
	"net/http"
	"testing"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
)

func TestGamesCountEndpoint(t *testing.T) {
	validToken, err := generateToken(dp.UserCollection, "karltestmichel@example.com")
	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := generateToken(dp.UserCollection, "otheruser@example.com")
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "Unauthenticated",
			Method:          http.MethodGet,
			URL:             "/api/gamecount/" + testTeamID,
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"data\":{},\"message\":\"The request requires valid record authorization token."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Invalid token",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID,
			Headers: map[string]string{
				"Authorization": invalidToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"Only members and admins can query team game count"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Bogus Team",
			Method: http.MethodGet,
			URL:    "/api/gamecount/doesnotexist",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to load teams record"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid user and team without query",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Query param for season is required"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid user and team with query",
			Method: http.MethodGet,
			URL:    "/api/gamecount/" + testTeamID + "?season=" + testSeasonQueryParam,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"count\":0"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
