package dp

import (
	"fmt"
	"log/slog"
	"slices"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type UserServiceApp interface {
	FindRecordsByFilter(collectionModelOrIdentifier any, filter string, sort string, limit int, offset int, params ...dbx.Params) ([]*core.Record, error)
	Logger() *slog.Logger
	ExpandRecords(records []*core.Record, expands []string, optFetchFunc core.ExpandFetchFunc) map[string]error
}

type ClubCommunityServiceRow struct {
	ID           string `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Name         string `db:"name" json:"name"`
	Target       int    `db:"target" json:"target"`
	TotalMinutes int    `db:"total_minutes" json:"total_minutes"`
	TargetMet    bool   `db:"target_met" json:"target_met"`
}

func getCommunityServiceDataForClub(app core.App, clubID string, season string) ([]ClubCommunityServiceRow, error) {
	var result []ClubCommunityServiceRow
	query := app.DB().
		Select(
			"users.id",
			"users.email",
			"concat(users.first_name, ' ', users.last_name) AS name",
			"clubs.service_requirement AS target",
			"SUM(serviceentries.minutes) AS total_minutes",
			"CASE WHEN SUM(serviceentries.minutes) >= clubs.service_requirement THEN 1 ELSE 0 END AS target_met",
		).
		From(ServiceEntryCollection).
		InnerJoin("users", dbx.NewExp("users.id = serviceentries.member")).
		InnerJoin("clubs", dbx.NewExp("clubs.id = serviceentries.club")).
		Where(dbx.NewExp("strftime('%Y', serviceentries.service_date) = {:season}", dbx.Params{"season": season})).
		AndWhere(dbx.NewExp("serviceentries.club = {:clubID}", dbx.Params{"clubID": clubID})).
		GroupBy("users.id", "users.email")

	err := query.All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type UserCommunityServiceItem struct {
	Club           Club           `json:"club"`
	CurrentMinutes int            `json:"current_minutes"`
	Entries        []ServiceEntry `json:"entries"`
}

type UserCommunityServiceData struct {
	User   User                       `json:"user"`
	Season int                        `json:"season"`
	Items  []UserCommunityServiceItem `json:"items"`
}

func getCommunityServiceDataForUser(app UserServiceApp, authID string, user User, season int, clubs []Club) (UserCommunityServiceData, error) {
	serviceData := UserCommunityServiceData{
		User:   user,
		Season: season,
	}

	start, end, err := GetYearStartAndEnd(season)
	if err != nil {
		return serviceData, err
	}

	for _, club := range clubs {
		item := UserCommunityServiceItem{
			Club:           club,
			CurrentMinutes: 0,
		}
		records, err := app.FindRecordsByFilter(
			ServiceEntryCollection,
			"club = {:clubID} && member = {:userID} && service_date >= {:start} && service_date <= {:end}",
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
		errs := app.ExpandRecords(records, []string{"club"}, nil)
		if len(errs) > 0 {
			return serviceData, fmt.Errorf("failed to expand: %v", errs)
		}

		item.Entries = make([]ServiceEntry, len(records))
		for i, record := range records {
			entry := ServiceEntry{}
			entry.SetProxyRecord(record)
			if !slices.Contains(club.Admins(), authID) {
				entry.HideAdminFields()
			}
			item.CurrentMinutes += entry.Minutes()
			item.Entries[i] = entry
		}
		serviceData.Items = append(serviceData.Items, item)
	}

	return serviceData, nil
}
