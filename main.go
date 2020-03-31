package main

import (
	"github.com/daiLlew/flexiC/cli"
	"github.com/daiLlew/flexiC/cli/out"
)

func main() {
	err := cli.Run()
	if err != nil {
		out.WriteWithPrefix(out.RED, err.Error())
	}
}