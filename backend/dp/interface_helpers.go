package dp

import "log/slog"

// LogOnlyApp is a PocketBase app that only needs to log errors.
type LogOnlyApp interface {
	Logger() *slog.Logger
}
