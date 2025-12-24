package dp

import (
	"time"

	"github.com/pocketbase/pocketbase/tools/types"
)

func GetYearStartAndEnd(year int) (start, end types.DateTime, err error) {
	startTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC)

	err = start.Scan(startTime)
	if err != nil {
		return start, end, err
	}
	err = end.Scan(endTime)
	if err != nil {
		return start, end, err
	}

	return start, end, nil
}
