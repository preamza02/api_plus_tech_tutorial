package command

import (
	"github.com/urfave/cli/v3"
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:   "http client",
		Usage:  "make a http request",
		Action: getAction,
		Commands: []*cli.Command{
			initGetCommand(),
			initPostCommand(),
			initPutCommand(),
			initDelCommand(),
		},
	}
}
