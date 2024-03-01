package commands

import (
	"fmt"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:  "get",
		Usage: "Fetch resources from Onomondo's database",
		Action: func(cCtx *cli.Context) error {
			url := fmt.Sprintf("http://localhost:8443/%v ", cCtx.Args().First())
			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println(url)
			fmt.Println(resp)

			return nil
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
		Usage: "Delete resrouces",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Resource deleted.", cCtx.Args())
			return nil
		},
	},
}
