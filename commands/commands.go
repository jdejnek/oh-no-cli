package commands

import (
	"errors"
	"fmt"

	"github.com/jdejnek/ono/http_client"
	"github.com/urfave/cli/v2"
)

var (
	limit     string
	offset    string
	next_page string
	id        string

	Commands = []*cli.Command{
		{
			Name:  "get",
			Usage: "Fetch resources from Onomondo's database",
			Subcommands: []*cli.Command{
				{
					Name:  "sims",
					Usage: "Fetch sims",
					Action: func(cCtx *cli.Context) error {
						queryParams := http_client.QueryParams{
							Params: map[string]interface{}{
								"limit":  limit,
								"offset": offset,
							},
						}
						defer http_client.CallApiWithParams("GET", "/sims", queryParams)
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "limit",
							Value:       "1000",
							Aliases:     []string{"l"},
							Usage:       "Limit the number of items fetched",
							Destination: &limit,
						},
						&cli.StringFlag{
							Name:        "offset",
							Value:       "0",
							Aliases:     []string{"o"},
							Usage:       "Skip the first n items",
							Destination: &offset,
						},
						&cli.StringFlag{
							Name:        "next_page",
							Value:       "",
							Aliases:     []string{"n"},
							Usage:       "View next page of results",
							Destination: &next_page,
						},
					},
				},
				{
					Name:  "sim",
					Usage: "Fetch sim by id or label",
					Action: func(cCtx *cli.Context) error {
						if id == "" {
							return errors.New("missing flag: -id")
						}
						http_client.CallApiWithPath("GET", "/sims/", id)
						return nil
					},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:        "id",
							Aliases:     []string{"i"},
							Value:       "",
							Usage:       "Provide sim id",
							Destination: &id,
						},
					},
				},
				{
					Name:  "find",
					Usage: "Search for sims matching flag criteria",
					Action: func(cCtx *cli.Context) error {
						return nil
					},
				},
				{
					Name:  "connectors",
					Usage: "List connectors",
					Action: func(cCtx *cli.Context) error {
						http_client.CallApiWithPath("GET", "/connectors", "")
						return nil
					},
				},
			},
		},
		{
			Name:  "create",
			Usage: "Create resources",
			Action: func(cCtx *cli.Context) error {
				fmt.Println("Resource created.", cCtx.Args())
				return nil
			},
		},
		{
			Name:  "delete",
			Usage: "Delete resources",
			Action: func(cCtx *cli.Context) error {
				fmt.Println("Resource deleted.", cCtx.Args())
				return nil
			},
		},
	}
)
