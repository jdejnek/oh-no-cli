package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jdejnek/ono/commands"
	"github.com/urfave/cli/v2"
)

// commands: get, create, ls, delete
// trigger verb from commands
// format url from flags

func main() {
	app := &cli.App{
		Name:     "ono",
		Usage:    "query the oh-no-api",
		Commands: commands.Commands,
	}

	PrintColor := "\033[1;33m%s\033[0m %s"
	cli.AppHelpTemplate = fmt.Sprintf(PrintColor,
		`
   ____     __  __           _   __   ____           ______    __     ____
  / __ \   / / / /          / | / /  / __ \         / ____/   / /    /  _/
 / / / /  / /_/ /  ______  /  |/ /  / / / / ______ / /       / /     / /  
/ /_/ /  / __  /  /_____/ / /|  /  / /_/ / /_____// /___    / /___ _/ /   
\____/  /_/ /_/          /_/ |_/   \____/         \____/   /_____//___/   
                                                                          
`, cli.AppHelpTemplate)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
