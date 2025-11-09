package dp

import (
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/hook"
)

const UserAccessMiddlewareId = "requireUserAccess"

// RequireUserAccess verifies that the current auth can access a record in a request param. In DiamondPlanner
// context, that is most commonly used for `users` and means one (or more) of 3 things:
//
//   - the user is the auth record itself
//   - the auth record is an admin of one of the requested user's teams
//   - the auth record is an admin of one of the requested user's clubs
func RequireUserAccess(record *core.Record) *hook.Handler[*core.RequestEvent] {
	return &hook.Handler[*core.RequestEvent]{
		Id:   UserAccessMiddlewareId,
		Func: requireUserAccess(record),
	}
}

func requireUserAccess(record *core.Record) func(*core.RequestEvent) error {
	return func(event *core.RequestEvent) error {
		if event.Auth.Id == record.Id {
			return event.Next()
		}

		info, err := event.RequestInfo()
		if err != nil {
			return err
		}

		canAccess, err := event.App.CanAccessRecord(record, info, record.Collection().ListRule)
		if !canAccess {
			return apis.NewForbiddenError("", err)
		}

		return event.Next()
	}
}
