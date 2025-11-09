package main

import (
	"os"
	"testing"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
)

const (
	testDataDir          = "./test_pb_data"
	testTeamID           = "847u6a1af1loox5"
	testSeasonQueryParam = "2025"
	validUserID          = "jm866jti5piq9g9"
)

func generateToken(collectionNameOrId string, email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	record, err := app.FindAuthRecordByEmail(collectionNameOrId, email)
	if err != nil {
		return "", err
	}

	return record.NewAuthToken()
}

func setupTestApp(t testing.TB) *tests.TestApp {
	testApp, err := tests.NewTestApp(testDataDir)

	if err != nil {
		t.Fatal(err)
	}
	// real implementation used for integration tests
	mockClient := bsm.RealAPIClient{
		BaseAPIURL: os.Getenv("BSM_API_URL"),
	}
	if mockClient.BaseAPIURL == "" {
		t.Fatal("BSM_API_URL not set, aborting")
	}

	// no need to clean up since scenario.Test() will do that for us
	dp.BindDPHooks(testApp, mockClient)
	BindTiBHooks(testApp, mockClient)

	return testApp
}
