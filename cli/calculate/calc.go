package calculate

import (
	"fmt"
	"time"

	"github.com/daiLlew/flexiC/cli/input"
)

func TotalDuration(times []input.TimeRange) error {
	fmt.Printf(">> %v\n", times)

	var totalDuration float64 = 0

	for _, t := range times {
		startT, err := t.StartTime()
		if err != nil {
			return err
		}
		fmt.Printf("start %+v", startT)

		endT, err := t.EndTime()
		if err != nil {
			return err
		}
		fmt.Printf("end %+v", endT)

		totalDuration += endT.Sub(startT).Minutes()
	}

	fmt.Printf("%+v", totalDuration)

	total := time.Minute * time.Duration(totalDuration)
	fmt.Printf("Total: %v", total)
	return nil
}
