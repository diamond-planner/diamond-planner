package dp

import "github.com/pocketbase/pocketbase/core"

func notifyWithTestPushMessage(s PushService) func(e *core.RequestEvent) error {
	return func(e *core.RequestEvent) error {
		// TODO: read from DB
		return nil
	}
}
