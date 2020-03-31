package out

import (
	"fmt"

	"github.com/fatih/color"
)

var (
	blue   = color.New(color.FgHiBlue)
	red    = color.New(color.FgHiRed)
	green  = color.New(color.FgHiGreen)
	prefix = "[flexiC] "
)

type Colour int

const (
	RED Colour = iota + 1
	BLUE
	GREEN
)

func getColor(c Colour) *color.Color {
	switch c {
	case RED:
		return red
	case BLUE:
		return blue
	default:
		return green
	}
}

func Write(c Colour, msg string, args ...interface{}) {
	getColor(c).Printf(prefix+msg, args...)
}
func NewLine() {
	fmt.Println()
}
