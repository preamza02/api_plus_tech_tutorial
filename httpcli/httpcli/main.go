package main

import (
	"context"
	"log"
	"os"

	"github.com/preamza02/httpcli/internal/command"
)

func main() {
	cmd := command.InitCommand()
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
