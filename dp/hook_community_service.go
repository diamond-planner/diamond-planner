package dp

import (
	"slices"

	"github.com/pocketbase/pocketbase/core"
)

func HideCommunityServiceAdminFields(e *core.RecordEnrichEvent) error {
	if e.RequestInfo.Auth != nil && e.RequestInfo.Auth.Collection().Name == UserCollection {
		serviceEntry := &ServiceEntry{}
		serviceEntry.SetProxyRecord(e.Record)

		clubID := serviceEntry.Club()
		clubRecord, err := e.App.FindRecordById(ClubsCollection, clubID)
		if err != nil {
			e.App.Logger().Error("Failed to load club record", "clubID", clubID, "error", err, "file", "hook_community_service.go")
			return err
		}

		club := &Club{}
		club.SetProxyRecord(clubRecord)

		admins := club.Admins()
		if !slices.Contains(admins, e.RequestInfo.Auth.Id) {
			serviceEntry.HideAdminFields()
		}
	}

	return e.Next()
}
