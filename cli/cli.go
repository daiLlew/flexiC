package cli

import (
	"fmt"

	"github.com/daiLlew/flexiC/cli/calculate"
	"github.com/daiLlew/flexiC/cli/input"
	"github.com/spf13/cobra"
)

func Run() error {
	c := &cobra.Command{RunE: calculateFlexi}
	return c.Execute()
}

func calculateFlexi(cmd *cobra.Command, args []string) error {
	fmt.Println("flexiC")
	times, err := input.GetTimes()
	if err != nil {
		return err
	}

	return calculate.TotalDuration(times)
}
