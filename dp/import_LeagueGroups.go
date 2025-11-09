package dp

import (
	"sync"

	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/types"
)

// ImportLeagueGroups imports league groups concurrently, either for one given club or all clubs in the database.
func ImportLeagueGroups(app core.App, client bsm.LeagueGroupClient, clubID *string, season *int) (err error) {
	filter := "bsm_id != 0"
	params := dbx.Params{}

	if clubID != nil {
		filter = filter + " && id = {:clubID}"
		params["clubID"] = *clubID
	}

	clubs, err := app.FindRecordsByFilter("clubs", filter, "", 0, 0, params)
	if err != nil {
		return err
	}

	selectedSeason := types.NowDateTime().Time().Year()
	if season != nil {
		selectedSeason = *season
	}

	processedClubs := 0
	processedLeagueGroups := 0
	var wg sync.WaitGroup

	for _, club := range clubs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			leagueGroups, err := client.FetchLeagueGroupsForSeason(club.GetString("bsm_api_key"), selectedSeason)
			if err != nil {
				app.Logger().Error("Error fetching league groups", "error", err, "club", club.GetString("name"), "season", selectedSeason)
				return
			}
			err = createOrUpdateLeagueGroups(app, leagueGroups, club)
			if err != nil {
				// Logging happens in the called method where more detailed data is available.
				return
			}
			processedClubs++
			processedLeagueGroups += len(leagueGroups)
		}()
	}

	wg.Wait()
	app.Logger().Info("League Group Import successfully imported all league groups", "Number of clubs processed", processedClubs, "Number of league groups processed", processedLeagueGroups)
	return nil
}

func createOrUpdateLeagueGroups(app core.App, leagueGroups []bsm.LeagueGroup, club *core.Record) (err error) {
	for _, leagueGroup := range leagueGroups {
		record, err := app.FindFirstRecordByData(LeagueGroupsCollection, "bsm_id", leagueGroup.ID)

		// if not found, it throws an error, so create new record
		if err != nil {
			collection, err := app.FindCollectionByNameOrId(LeagueGroupsCollection)
			if err != nil {
				app.Logger().Error("Error getting collection", "error", err)
				return err
			}
			record = core.NewRecord(collection)
			setLeagueGroupRecordValues(record, leagueGroup, club)
		}
		// no error - update existing record
		setLeagueGroupRecordValues(record, leagueGroup, club)

		if err := app.Save(record); err != nil {
			app.Logger().Error("Persisting LeagueGroup record failed: ", "error", err, "leagueGroup", leagueGroup, "record", record)
			return err
		}
	}
	return
}

func setLeagueGroupRecordValues(record *core.Record, leagueGroup bsm.LeagueGroup, club *core.Record) {
	lg := &LeagueGroup{}
	lg.SetProxyRecord(record)

	lg.SetBSMID(leagueGroup.ID)
	lg.SetSeason(leagueGroup.Season)
	lg.SetName(leagueGroup.Name)
	lg.SetAcronym(leagueGroup.Acronym)
	// Use the raw record API for array append operations
	record.Set("clubs+", club.Id)

	return
}
