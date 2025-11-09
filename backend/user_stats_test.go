package main

import (
	"net/http"
	"testing"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
)

func TestGetUserStatsEndpoint(t *testing.T) {
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
			Name:            "Unauthenticated request",
			Method:          http.MethodGet,
			URL:             "/api/stats/123",
			ExpectedStatus:  http.StatusUnauthorized,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Invalid token",
			Method: http.MethodGet,
			URL:    "/api/stats/" + validUserID,
			Headers: map[string]string{
				"Authorization": invalidToken,
			},
			ExpectedStatus:  http.StatusForbidden,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "User not found",
			Method: http.MethodGet,
			URL:    "/api/stats/doesnotexist",
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"message", "data", "status"},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Valid request",
			Method: http.MethodGet,
			URL:    "/api/stats/" + validUserID,
			Headers: map[string]string{
				"Authorization": validToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{"season", "totalPossibleEvents", "attendanceTotals", "participationTotals", "teamName"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
