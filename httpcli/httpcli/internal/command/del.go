package command

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/preamza02/httpcli/internal/argument"
	"github.com/preamza02/httpcli/internal/flags"
	"github.com/urfave/cli/v3"
)

func delAction(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 0 {
		fmt.Println(cmd.Args().Len(), cmd.Args().Slice())
		return fmt.Errorf("please provide exactly one URL")
	}
	url := cmd.StringArg("url")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	headers := cmd.StringSlice("header")
	for _, header := range headers {
		parts := strings.SplitN(header, "=", 2)
		if len(parts) == 2 {
			req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		} else {
			fmt.Println("Invalid header:", header)
		}
	}
	query := cmd.StringSlice("query")
	if query != nil {
		q := req.URL.Query()
		for _, param := range query {
			parts := strings.SplitN(param, "=", 2)
			if len(parts) == 2 {
				q.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			} else {
				fmt.Println("Invalid query param:", param)
			}
		}
		req.URL.RawQuery = q.Encode()

	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func initDelCommand() *cli.Command {
	return &cli.Command{
		Name:    "del",
		Aliases: []string{"d"},
		Usage:   "make a DELETE request",
		Flags: []cli.Flag{
			flags.QueryFlag,
			flags.HeaderFlag,
		},
		Arguments: []cli.Argument{
			argument.UrlArg,
		},
		Action: delAction,
	}
}
