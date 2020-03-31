package calculate

import (
	"time"

	"github.com/daiLlew/flexiC/cli/input"
)

func GetDuration(times []input.TimeRange) (time.Duration, error) {
	var totalDuration float64 = 0

	for _, t := range times {
		startT, err := t.StartTime()
		if err != nil {
			return 0, err
		}

		endT, err := t.EndTime()
		if err != nil {
			return 0, err
		}

		totalDuration += endT.Sub(startT).Minutes()
	}

	return time.Minute * time.Duration(totalDuration), nil
}
