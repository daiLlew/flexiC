package main

import (
	"log"

	"github.com/daiLlew/flexiC/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		log.Fatal(err)
	}
}