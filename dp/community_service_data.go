package dp

import (
	"log/slog"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type UserServiceApp interface {
	FindRecordsByFilter(collectionModelOrIdentifier any, filter string, sort string, limit int, offset int, params ...dbx.Params) ([]*core.Record, error)
	Logger() *slog.Logger
}

type CommunityServiceItem struct {
	Club    Club           `json:"club"`
	Entries []ServiceEntry `json:"entries"`
}

type CommunityServiceData struct {
	User   User                   `json:"user"`
	Season int                    `json:"season"`
	Items  []CommunityServiceItem `json:"items"`
}

func getCommunityServiceData(app UserServiceApp, user User, season int, clubs []Club) (CommunityServiceData, error) {
	serviceData := CommunityServiceData{
		User:   user,
		Season: season,
	}

	start, end, err := GetYearStartAndEnd(season)
	if err != nil {
		return serviceData, err
	}

	for _, club := range clubs {
		item := CommunityServiceItem{Club: club}
		records, err := app.FindRecordsByFilter(
			ServiceEntryCollection,
			"club = {:clubID} && user = {:userID} && service_date >= {:start} && service_date <= {:end}",
			"-service_date",
			0,
			0,
			dbx.Params{
				"clubID": club.Id,
				"userID": user.Id,
				"start":  start,
				"end":    end},
		)
		if err != nil {
			app.Logger().Error("Failed to load service entries", "clubID", club.Id, "error", err)
			return serviceData, err
		}
		item.Entries = make([]ServiceEntry, len(records))
		for i, record := range records {
			entry := ServiceEntry{}
			entry.SetProxyRecord(record)
			item.Entries[i] = entry
		}
		serviceData.Items = append(serviceData.Items, item)
	}

	return serviceData, nil
}
