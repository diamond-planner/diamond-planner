package migrations

import (
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		settings := app.Settings()

		settings.Meta.AppName = "DiamondPlanner"
		settings.Meta.AppURL = os.Getenv("PUBLIC_POCKETBASE_URL")

		settings.Logs.MaxDays = 5
		settings.Logs.LogAuthId = true
		settings.Logs.LogIP = true
		settings.Logs.MinLevel = 0

		settings.RateLimits.Enabled = true

		settings.Backups.Cron = "0 0 * * *"
		settings.Backups.CronMaxKeep = 3

		return app.Save(settings)
	}, func(app core.App) error {
		// add down queries...

		return nil
	})
}
