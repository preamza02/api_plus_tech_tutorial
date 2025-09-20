package flags

import "github.com/urfave/cli/v3"

var (
	QueryFlag = &cli.StringSliceFlag{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "query parameters in key=value format",
	}
	HeaderFlag = &cli.StringSliceFlag{
		Name:    "header",
		Aliases: []string{"H"},
		Usage:   "custom headers in key=value format",
	}
	JsonFlag = &cli.StringFlag{
		Name:    "json",
		Aliases: []string{"j"},
		Usage:   "json body for Post/Put requests",
	}
)
