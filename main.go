package main

import (
	"os"

	"github.com/dngferreira/aws-nuke/v2/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		os.Exit(-1)
	}
}
