package dp

import "github.com/pocketbase/pocketbase/core"

func ValidateEventTimes(e *core.RecordRequestEvent) error {
	startTime := e.Record.GetDateTime("starttime")
	endTime := e.Record.GetDateTime("endtime")
	meetingTime := e.Record.GetDateTime("meetingtime")

	if !endTime.IsZero() && startTime.After(endTime) {
		return e.BadRequestError("Event start time cannot be after end time", nil)
	}

	if !endTime.IsZero() && meetingTime.After(endTime) {
		return e.BadRequestError("Event meeting time cannot be after end time", nil)
	}

	return e.Next()
}
