package dp

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/diamond-planner/diamond-planner/bsm"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/spf13/cobra"
)

type DiamondPlanner struct {
	*pocketbase.PocketBase
}

func NewDiamondPlanner(client bsm.APIClient) *DiamondPlanner {
	dp := &DiamondPlanner{
		pocketbase.New(),
	}

	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	if isGoRun {
		log.Print("App was started with `go run`, automigrations are active.")
	}

	migratecmd.MustRegister(dp, dp.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Dashboard
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	BindDPHooks(dp, client)

	//------------------- Commands -------------------------//

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "import:games",
		Run: func(cmd *cobra.Command, args []string) {
			// app is a pointer here because commands only work after app.Start()
			gamesImportService := GameImportService{App: *dp}
			gamesImportService.ImportGames()
		},
	})

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "import:leagues",
		Run: func(cmd *cobra.Command, args []string) {
			err := ImportLeagueGroups(dp, client, nil, nil)
			if err != nil {
				log.Print("Error while running LeagueGroupImport: " + err.Error())
			}
		},
	})

	dp.RootCmd.AddCommand(&cobra.Command{
		Use: "push:generateVAPID",
		Run: func(cmd *cobra.Command, args []string) {
			privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
			if err != nil {
				log.Print("Error while generating VAPID keys: " + err.Error())
			}
			fmt.Printf("Private key: %s\nPublic key: %s\n", privateKey, publicKey)
		},
	})

	return dp
}

func BindDPHooks(app core.App, client bsm.APIClient) {

	//------------------- Hooks -------------------------//

	app.OnRecordAuthRequest(UserCollection).BindFunc(func(e *core.RecordAuthRequestEvent) error {
		err := SetLastLogin(e.App, e.Record)
		if err != nil {
			app.Logger().Error(
				"Error upon setting auth record last login field",
				"error", err, "user", e.Record.Id,
			)
		}
		return e.Next()
	})

	app.OnRecordAuthWithOAuth2Request(UserCollection).BindFunc(func(e *core.RecordAuthWithOAuth2RequestEvent) error {
		return OAuthUpdateUserData(e)
	})

	app.OnRecordCreateRequest(UserCollection).BindFunc(func(event *core.RecordRequestEvent) error {
		return ValidateSignupKey(event)
	})

	app.OnRecordsListRequest(LeagueGroupsCollection).BindFunc(func(event *core.RecordsListRequestEvent) error {
		return TriggerLeagueImport(event.App, client, event)
	})

	app.OnRecordCreateRequest(EventsCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return ValidateEventTimes(e)
	})

	app.OnRecordUpdateRequest(EventsCollection).BindFunc(func(e *core.RecordRequestEvent) error {
		return ValidateEventTimes(e)
	})

	app.OnRecordEnrich(EventsCollection).BindFunc(func(event *core.RecordEnrichEvent) error {
		return AddEventParticipationData(event.App, event)
	})

	app.OnRecordEnrich(ServiceEntryCollection).BindFunc(func(event *core.RecordEnrichEvent) error {
		return HideCommunityServiceAdminFields(event)
	})

	app.OnRecordAfterCreateSuccess(EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return CreateOrUpdateEventsForSeries(e)
	})

	app.OnRecordAfterUpdateSuccess(EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return CreateOrUpdateEventsForSeries(e)
	})

	app.OnRecordDelete(EventSeriesCollection).BindFunc(func(e *core.RecordEvent) error {
		return DeleteEventsForSeries(e)
	})

	//------------------- Serve static dir -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided public dir (if it exists)
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), true))

		return se.Next()
	})

	//------------------- Custom Routes -------------------------//

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/hello", func(e *core.RequestEvent) error {
			return e.String(http.StatusOK, "Hello ballplayers!")
		})
		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/gamecount/{team}", GetGamesCount(app))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/import/{club}/leagues", StartLeagueGroupsImport(app, client))

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/stats/{user}",
			GetUserStats()).
			Bind(apis.RequireAuth())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/communityservice/{user}/{season}",
			GetUserCommunityService()).
			Bind(apis.RequireAuth())

		return se.Next()
	})

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		clubAdminGroup := se.Router.Group("/api/clubs/{club}/admin")

		clubAdminGroup.Bind(RequireClubAdminAccess())
		clubAdminGroup.GET("/communityservice", GetClubCommunityService())
		return se.Next()
	})

	//------------------- Cronjobs -------------------------//

	if os.Getenv("APPLICATION_CONTEXT") != "Development" {
		app.Cron().MustAdd("LeagueGroupImport", "0 * * * *", func() {
			err := ImportLeagueGroups(app, client, nil, nil)
			if err != nil {
				app.Logger().Error("Error while running cronjob LeagueGroupImport: " + err.Error())
			}
		})

		app.Cron().MustAdd("GamesImport", "0 * * * *", func() {
			gamesImportService := GameImportService{App: app}
			gamesImportService.ImportGames()
		})
	}
}
