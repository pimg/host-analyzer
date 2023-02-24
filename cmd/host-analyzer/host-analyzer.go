package main

import (
	"log"
	"os"

	"github.com/pimg/host-analyzer/cmd/host-analyzer/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "hostname analyzer cli"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name servers"

	app.Commands = commands.GetRegistry().ImplementedCommands
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
