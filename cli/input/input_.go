package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	inputPattern string = "^[0-9]{4}\\s+[0-9]{4}$"
	hhMMFormat          = "1504"
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
	fmt.Println("Input time: hhMM hhMM")

	var input string
	for sc.Scan() {
		input = strings.TrimSpace(sc.Text())

		if len(input) == 0 {
			fmt.Println("please enter a time")
			continue
		}

		if input == "q" {
			fmt.Println("quitting")
			// todo fix this.
			return nil, nil
		}

		if input == "d" {
			break
		}

		if !inputRegex.Match([]byte(input)) {
			fmt.Println("invalid input")
			continue
		}

		args := strings.Split(input, " ")
		times = append(times, TimeRange{start: args[0], end: args[1]})
	}

	return times, nil
}
