package main

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/diamond-planner/diamond-planner/dp"
	"github.com/pocketbase/pocketbase/tests"
)

type ServiceEntryData struct {
	ID           string
	ServiceDate  string
	Minutes      int
	Member       string
	Title        string
	Description  string
	AdminComment string
	BoardMember  string
	Club         string
}

var serviceEntryDataAliceForClubAAA = ServiceEntryData{
	ID:           "pqzhx6vjzhmaya0",
	ServiceDate:  "2025-12-22 12:00:00.000Z",
	Minutes:      120,
	Member:       "wvizeo6lht334n9", // Alice Anderson
	Title:        "Did things",
	Description:  "very many things",
	AdminComment: "LGTM",
	BoardMember:  "8qjhl2zp4rt625r", // Andy A Admin
	Club:         "vogf6j5vwmwrl6g", // Club AAA
}

func TestServiceEntryRules(t *testing.T) {
	aliceToken, err := generateToken(dp.UserCollection, alice.Email)
	if err != nil {
		t.Fatal(err)
	}

	bobToken, err := generateToken(dp.UserCollection, bob.Email)
	if err != nil {
		t.Fatal(err)
	}

	aClubAdminToken, err := generateToken(dp.UserCollection, aClubAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	bClubAdminToken, err := generateToken(dp.UserCollection, bClubAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	andreaToken, err := generateToken(dp.UserCollection, andreaAntelopeAdmin.Email)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		{
			Name:   "Alice creates service entry for herself in her club",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, alice.ID, clubA.ID)),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"member":"` + alice.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice cannot create service entry for herself in another club",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, alice.ID, clubB.ID)),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to create record."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice cannot create service entry for someone else",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, bob.ID, clubA.ID)),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to create record."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club Admin can create service entry for member in their club",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, alice.ID, clubA.ID)),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"member":"` + alice.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club Admin cannot create service entry for member in another club",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, bob.ID, clubB.ID)),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to create record."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice can read her own service entry",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"id":"` + serviceEntryDataAliceForClubAAA.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Bob cannot read Alice's service entry",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": bobToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice cannot read another's service entry",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/pqzhx6vjzhmaya1", // fake ID for another user
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club Admin can read member's service entry in their club",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"id":"` + serviceEntryDataAliceForClubAAA.ID + `"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club Admin cannot read service entry in another club",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": bClubAdminToken,
			},
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice can update her own service entry",
			Method: http.MethodPatch,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:            strings.NewReader(`{"title":"Updated title"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"title":"Updated title"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice cannot update another's service entry",
			Method: http.MethodPatch,
			URL:    "/api/collections/serviceentries/records/pqzhx6vjzhmaya1",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:            strings.NewReader(`{"title":"Updated title"}`),
			ExpectedStatus:  http.StatusNotFound,
			ExpectedContent: []string{"The requested resource wasn't found."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Club Admin can update member's service entry in their club",
			Method: http.MethodPatch,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"title":"Updated by admin"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"title":"Updated by admin"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Alice can delete her own service entry",
			Method: http.MethodDelete,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus: http.StatusNoContent,
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "Club Admin can delete member's service entry in their club",
			Method: http.MethodDelete,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			ExpectedStatus: http.StatusNoContent,
			TestAppFactory: setupTestApp,
		},
		{
			Name:   "Team admin (Andrea) cannot create entry for Alice",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": andreaToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s"
			}`, alice.ID, clubA.ID)),
			ExpectedStatus:  http.StatusBadRequest,
			ExpectedContent: []string{"Failed to create record."},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Regular user cannot set AdminComment when creating",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s",
				"admin_comment": "SHOULD BE IGNORED"
			}`, alice.ID, clubA.ID)),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":""`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Regular user cannot set AdminComment when updating",
			Method: http.MethodPatch,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body:            strings.NewReader(`{"admin_comment":"SHOULD BE IGNORED"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":"LGTM"`}, // Should remain original value from serviceEntryDataAliceForClubAAA
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Admin can set AdminComment when creating",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s",
				"admin_comment": "Admin set this"
			}`, alice.ID, clubA.ID)),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":"Admin set this"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Admin can set AdminComment when updating",
			Method: http.MethodPatch,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			Body:            strings.NewReader(`{"admin_comment":"Admin updated this"}`),
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":"Admin updated this"`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Regular user cannot see AdminComment in response",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":""`},
			TestAppFactory:  setupTestApp,
		},
		{
			Name:   "Admin can see AdminComment in response",
			Method: http.MethodGet,
			URL:    "/api/collections/serviceentries/records/" + serviceEntryDataAliceForClubAAA.ID,
			Headers: map[string]string{
				"Authorization": aClubAdminToken,
			},
			ExpectedStatus:  http.StatusOK,
			ExpectedContent: []string{`"admin_comment":"LGTM"`},
			TestAppFactory:  setupTestApp,
		},
		// TODO: fix rule, failing!
		{
			Name:   "Board member must be an admin of the club",
			Method: http.MethodPost,
			URL:    "/api/collections/serviceentries/records",
			Headers: map[string]string{
				"Authorization": aliceToken,
			},
			Body: strings.NewReader(fmt.Sprintf(`{
				"service_date": "2025-12-22 10:00:00.000Z",
				"minutes": 60,
				"member": "%s",
				"title": "Cleaned up",
				"club": "%s",
				"board_member": "%s"
			}`, alice.ID, clubA.ID, andreaAntelopeAdmin.ID)), // Andrea is not a Club admin
			ExpectedStatus: http.StatusBadRequest,
			TestAppFactory: setupTestApp,
		},
	}

	for _, s := range scenarios {
		s.Test(t)
	}
}
