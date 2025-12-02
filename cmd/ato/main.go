package main

import (
	"os"

	"github.com/atopile/atopile/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
