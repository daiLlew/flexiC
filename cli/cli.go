package cli

import (
	"time"

	"github.com/daiLlew/flexiC/cli/calculate"
	"github.com/daiLlew/flexiC/cli/input"
	"github.com/daiLlew/flexiC/cli/out"
	"github.com/spf13/cobra"
)

const (
	resultFMT             = "Total: %02dh:%02dm\n"
	dailyTargetTimeInMins = 444 // 7 hrs 24 mins
)

func Run() error {
	c := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			times, err := input.GetTimes()
			if err != nil {
				if err == input.ExitErr {
					out.Write(out.BLUE, "Goodbye\n")
					return nil
				}
				return err
			}

			d, err := calculate.GetDuration(times)
			if err != nil {
				return err
			}

			d = d.Round(time.Minute) // round to nearest min.
			hrs := d / time.Hour
			mins := d - hrs*time.Hour
			mins = mins / time.Minute

			colour := out.GREEN
			if d.Minutes() < dailyTargetTimeInMins {
				colour = out.RED
			}

			out.Write(colour, resultFMT, hrs, mins)
			return nil
		},
	}
	return c.Execute()
}
