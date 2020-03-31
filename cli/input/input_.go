package input

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/daiLlew/flexiC/cli/out"
)

const (
	inputPattern string = "^[0-9]{4}\\s+[0-9]{4}$"
	hhMMFormat          = "1504"
)

var (
	ExitErr = errors.New("exit invoked")
)

type TimeRange struct {
	start string
	end   string
}

func (tr TimeRange) StartTime() (time.Time, error) {
	return time.Parse(hhMMFormat, tr.start)
}

func (tr TimeRange) EndTime() (time.Time, error) {
	return time.Parse(hhMMFormat, tr.end)
}

func GetTimes() ([]TimeRange, error) {
	inputRegex, err := regexp.Compile(inputPattern)
	if err != nil {
		return nil, err
	}

	times := make([]TimeRange, 0)

	sc := bufio.NewScanner(os.Stdin)
	out.Write(out.BLUE, "Input start end times using the format %q.\n", "hhMM hhMM")
	out.Write(out.BLUE, "Enter return after each pair to submit (repeat as necessary).\n")
	out.Write(out.BLUE, "Enter %q to complete and display the total duration.\n", "d")

	out.Write(out.BLUE, "")
	var input string
	for sc.Scan() {
		out.Write(out.BLUE, "")
		input = strings.TrimSpace(sc.Text())

		if len(input) == 0 {
			continue
		}

		if input == "q" {
			out.NewLine()
			return nil, ExitErr
		}

		if input == "d" {
			out.NewLine()
			break
		}

		if !inputRegex.Match([]byte(input)) {
			continue
		}

		args := strings.Split(input, " ")
		times = append(times, TimeRange{start: args[0], end: args[1]})
	}

	return times, nil
}
