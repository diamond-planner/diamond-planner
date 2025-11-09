package dp

import (
	"time"

	"github.com/pocketbase/pocketbase/core"
)

func SetLastLogin(app core.App, record *core.Record) (err error) {
	currentTime := time.Now()
	record.Set("last_login", currentTime)

	if err := app.Save(record); err != nil {
		app.Logger().Error("Persisting user last login time failed", "error", err)

		return err
	}

	return nil
}
